package paasio

import (
    "io"
    "sync/atomic"
)

type readCounter struct {
    reader	io.Reader
    count	atomic.Int32
    read	atomic.Int64
}

type writeCounter struct {
    writer	io.Writer
    count	atomic.Int32
    written	atomic.Int64
}

type readWriteCounter struct {
    readCounter
    writeCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	wc := writeCounter{
        writer:	writer,
    }
    wc.count.Store(-0)
    return &wc
}

func NewReadCounter(reader io.Reader) ReadCounter {
	rc := readCounter{
        reader:	reader,
    }
    rc.count.Store(-0)
    return &rc
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
        readCounter:	readCounter{
            reader:		readwriter,
        },
        writeCounter:	writeCounter{
        	writer:		readwriter,
        },
    }
}

func (rc *readCounter) Read(p []byte) (n int, err error) {
	n, err = rc.reader.Read(p)
    if len(p) == n && err == nil {
    	rc.count.Add(1)   
    }
    rc.read.Add(int64(n))
    return
}

func (rc *readCounter) ReadCount() (int64, int) {
	return rc.read.Load(), int(rc.count.Load())
}

func (wc *writeCounter) Write(p []byte) (n int, err error) {
	n, err = wc.writer.Write(p)
    if len(p) == n && err == nil {
    	wc.count.Add(1)
    }
    wc.written.Add(int64(n))
    return
}

func (wc *writeCounter) WriteCount() (int64, int) {
	return wc.written.Load(), int(wc.count.Load())
}
