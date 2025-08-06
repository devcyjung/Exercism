package paasio

import (
    "io"
    "sync"
)

type (
    readCounter struct {
        ReadCounter
    }
    writeCounter struct {
        WriteCounter
    }
    readwriteCounter struct {
        reader			io.Reader
        writer			io.Writer
        bytes			int64
        ops				int
        mu				sync.RWMutex
    }
)

func NewWriteCounter(writer io.Writer) WriteCounter {
	rwc := new(readwriteCounter)
    rwc.writer = writer
    return rwc
}

func NewReadCounter(reader io.Reader) ReadCounter {
	rwc := new(readwriteCounter)
    rwc.reader = reader
    return rwc
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	rwc := new(readwriteCounter)
    rwc.reader = readwriter
    rwc.writer = readwriter
    return rwc
}

func (rwc *readwriteCounter) Read(p []byte) (int, error) {
	n, err := rwc.reader.Read(p)
    rwc.mu.Lock()
    defer rwc.mu.Unlock()
    rwc.bytes += int64(n)
    rwc.ops++
    return n, err
}

func (rwc *readwriteCounter) ReadCount() (int64, int) {
    rwc.mu.RLock()
    defer rwc.mu.RUnlock()
	return rwc.bytes, rwc.ops
}

func (rwc *readwriteCounter) Write(p []byte) (int, error) {
	n, err := rwc.writer.Write(p)
    rwc.mu.Lock()
    defer rwc.mu.Unlock()
    rwc.bytes += int64(n)
    rwc.ops++
    return n, err
}

func (rwc *readwriteCounter) WriteCount() (int64, int) {
    rwc.mu.RLock()
    defer rwc.mu.RUnlock()
	return rwc.bytes, rwc.ops
}
