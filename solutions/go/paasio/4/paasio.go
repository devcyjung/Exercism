package paasio

import (
    "io"
    "sync/atomic"
)

type counterStats struct {
    bytes int64
    count int
}

type counter struct {
    atomic.Value
}

func (c *counter) initialize() {
    c.Store(counterStats{})
}

func (c *counter) loadState() (int64, int) {
    state := c.Load().(counterStats)
    return state.bytes, state.count
}

func (c *counter) addBytes(n int) {
    var oldState, newState counterStats
    for {
        oldState = c.Load().(counterStats)
        newState.bytes, newState.count = oldState.bytes + int64(n), oldState.count + 1
        if c.CompareAndSwap(oldState, newState) {
            return
        }
    }
}

type readCounter struct {
    r	io.Reader
    *counter
}

type writeCounter struct {
    w	io.Writer
    *counter
}

type readWriteCounter struct {
    readCounter
    writeCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
    c := &counter{}
    c.initialize()
	return &writeCounter{
        w:			writer,
        counter:	c,
    }
}

func NewReadCounter(reader io.Reader) ReadCounter {
    c := &counter{}
    c.initialize()
	return &readCounter{
        r:			reader,
        counter:	c,
    }
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
    c := &counter{}
    c.initialize()
	return &readWriteCounter{
        readCounter{
            r:			readwriter,
            counter:	c,
        },
        writeCounter{
            w:			readwriter,
            counter:	c,
        },
    }
}

func (rc *readCounter) Read(p []byte) (n int, err error) {
    n, err = rc.r.Read(p)
    if n > 0 {
        rc.addBytes(n)
    }
    return
}

func (rc *readCounter) ReadCount() (int64, int) {
    return rc.loadState()
}

func (wc *writeCounter) Write(p []byte) (n int, err error) {
    n, err = wc.w.Write(p)
    if n > 0 {
        wc.addBytes(n)
    }
    return
}

func (wc *writeCounter) WriteCount() (int64, int) {
    return wc.loadState()
}