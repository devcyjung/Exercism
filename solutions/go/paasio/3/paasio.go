package paasio

import (
    "io"
    "sync/atomic"
    "sync"
)

type readCounter struct {
    reader	io.Reader
    count	atomic.Int32
    read	atomic.Int64
    lock	sync.RWMutex
}

type writeCounter struct {
    writer	io.Writer
    count	atomic.Int32
    written	atomic.Int64
    lock	sync.RWMutex
}

type readWriteCounter struct {
    readCounter
    writeCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	wc := writeCounter{
        writer:	writer,
    }
    return &wc
}

func NewReadCounter(reader io.Reader) ReadCounter {
	rc := readCounter{
        reader:	reader,
    }
    return &rc
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	rwc := readWriteCounter{
        readCounter:	readCounter{
            reader:		readwriter,
        },
        writeCounter:	writeCounter{
        	writer:		readwriter,
        },
    }
    return &rwc
}

func (rc *readCounter) Read(p []byte) (n int, err error) {
	rc.lock.Lock()
    n, err = rc.reader.Read(p)
    if len(p) == n && err == nil {
    	rc.count.Add(1)   
    }
    rc.read.Add(int64(n))
    rc.lock.Unlock()
    return
}

func (rc *readCounter) ReadCount() (int64, int) {
    rc.lock.RLock()
    defer rc.lock.RUnlock()
	return rc.read.Load(), int(rc.count.Load())
}

func (wc *writeCounter) Write(p []byte) (n int, err error) {
	wc.lock.Lock()
    n, err = wc.writer.Write(p)
    if len(p) == n && err == nil {
    	wc.count.Add(1)
    }
    wc.written.Add(int64(n))
    wc.lock.Unlock()
    return
}

func (wc *writeCounter) WriteCount() (int64, int) {
    wc.lock.RLock()
    defer wc.lock.RUnlock()
	return wc.written.Load(), int(wc.count.Load())
}
