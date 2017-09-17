// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package msgpack

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/m3db/m3cluster/placement"
	"github.com/m3db/m3metrics/protocol/msgpack"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	yaml "gopkg.in/yaml.v2"
)

var (
	testPlacementInstance = placement.NewInstance().
		SetID("testInstanceID").
		SetEndpoint("testInstanceAddress")
)

func TestInstanceQueueEnqueueClosed(t *testing.T) {
	opts := testServerOptions()
	queue := newInstanceQueue(testPlacementInstance, opts).(*queue)
	queue.writeFn = func([]byte) error { return nil }
	queue.closed = true

	require.Equal(t, errInstanceQueueClosed, queue.Enqueue(msgpack.NewBufferedEncoder()))
}

func TestInstanceQueueEnqueueQueueFullDropCurrent(t *testing.T) {
	opts := testServerOptions().SetInstanceQueueSize(1)
	queue := newInstanceQueue(testPlacementInstance, opts).(*queue)
	queue.dropType = DropCurrent

	// Fill up the queue and park the draining goroutine so the queue remains full.
	queue.writeFn = func([]byte) error {
		select {}
	}
	queue.bufCh <- msgpack.NewBufferedEncoder()
	queue.bufCh <- msgpack.NewBufferedEncoder()
	require.Equal(t, errWriterQueueFull, queue.Enqueue(msgpack.NewBufferedEncoder()))
}

func TestInstanceQueueEnqueueQueueFullDropOldest(t *testing.T) {
	opts := testServerOptions().SetInstanceQueueSize(1)
	queue := newInstanceQueue(testPlacementInstance, opts).(*queue)

	// Fill up the queue and park the draining goroutine so the queue remains full
	// until the enqueueing goroutine pulls a buffer off the channel.
	queue.writeFn = func([]byte) error {
		select {}
	}
	queue.bufCh <- msgpack.NewBufferedEncoder()
	queue.bufCh <- msgpack.NewBufferedEncoder()
	require.NoError(t, queue.Enqueue(msgpack.NewBufferedEncoder()))
}

func TestInstanceQueueEnqueueSuccessDrainSuccess(t *testing.T) {
	opts := testServerOptions()
	queue := newInstanceQueue(testPlacementInstance, opts).(*queue)
	var (
		res     []byte
		resLock sync.Mutex
	)
	queue.writeFn = func(data []byte) error {
		resLock.Lock()
		res = data
		resLock.Unlock()
		return nil
	}

	data := []byte("foobar")
	buffer := msgpack.NewBufferedEncoder()
	buffer.Buffer().Write(data)
	require.NoError(t, queue.Enqueue(buffer))

	// Wait for the queue to be drained.
	for {
		resLock.Lock()
		if res != nil {
			resLock.Unlock()
			break
		}
		resLock.Unlock()
		time.Sleep(100 * time.Millisecond)
	}

	require.Equal(t, data, res)
}

func TestInstanceQueueEnqueueSuccessDrainError(t *testing.T) {
	opts := testServerOptions()
	queue := newInstanceQueue(testPlacementInstance, opts).(*queue)
	var drained int32
	queue.writeFn = func(data []byte) error {
		atomic.StoreInt32(&drained, 1)
		return errTestWrite
	}

	buffer := msgpack.NewBufferedEncoder()
	require.NoError(t, queue.Enqueue(buffer))

	// Wait for the queue to be drained.
	for {
		if atomic.LoadInt32(&drained) == 1 {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func TestInstanceQueueEnqueueSuccessWriteError(t *testing.T) {
	opts := testServerOptions()
	queue := newInstanceQueue(testPlacementInstance, opts).(*queue)
	var done int32
	queue.writeFn = func(data []byte) error {
		err := queue.conn.Write(data)
		atomic.StoreInt32(&done, 1)
		return err
	}

	buffer := msgpack.NewBufferedEncoder()
	_, err := buffer.Buffer().Write([]byte{0x1, 0x2})
	require.NoError(t, err)
	require.NoError(t, queue.Enqueue(buffer))

	// Wait for the queue to be drained.
	for {
		if atomic.LoadInt32(&done) == 1 {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func TestInstanceQueueCloseAlreadyClosed(t *testing.T) {
	opts := testServerOptions()
	queue := newInstanceQueue(testPlacementInstance, opts).(*queue)
	queue.closed = true

	require.Equal(t, errInstanceQueueClosed, queue.Close())
}

func TestInstanceQueueCloseSuccess(t *testing.T) {
	opts := testServerOptions()
	queue := newInstanceQueue(testPlacementInstance, opts).(*queue)
	require.NoError(t, queue.Close())
	require.True(t, queue.closed)
	_, ok := <-queue.bufCh
	require.False(t, ok)
}

func TestDropTypeUnmarshalYAML(t *testing.T) {
	type S struct {
		A DropType
	}

	tests := []struct {
		input    []byte
		expected DropType
	}{
		{
			input:    []byte("a: oldest\n"),
			expected: DropOldest,
		},
		{
			input:    []byte("a: current\n"),
			expected: DropCurrent,
		},
	}

	for _, test := range tests {
		var s S
		err := yaml.Unmarshal(test.input, &s)
		require.NoError(t, err)
		assert.Equal(t, test.expected, s.A)
	}
}
