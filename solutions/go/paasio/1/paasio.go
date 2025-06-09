package paasio

import (
    "io"
    "sync/atomic"
)

type readCounter struct {
    reader	io.Reader
    count	int32
    read	int64
}

type writeCounter struct {
    writer	io.Writer
    count	int32
    written	int64
}

type readWriteCounter struct {
    readCounter
    writeCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{
        writer:	writer,
    }
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{
        reader:	reader,
    }
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
    atomic.AddInt32(&rc.count, 1)
    atomic.AddInt64(&rc.read, int64(n))
    return
}

func (rc *readCounter) ReadCount() (int64, int) {
	return atomic.LoadInt64(&rc.read), int(atomic.LoadInt32(&rc.count))
}

func (wc *writeCounter) Write(p []byte) (n int, err error) {
	n, err = wc.writer.Write(p)
    atomic.AddInt32(&wc.count, 1)
    atomic.AddInt64(&wc.written, int64(n))
    return
}

func (wc *writeCounter) WriteCount() (int64, int) {
	return atomic.LoadInt64(&wc.written), int(atomic.LoadInt32(&wc.count))
}
