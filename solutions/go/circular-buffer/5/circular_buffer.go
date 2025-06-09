package circular

import "errors"

type Buffer struct {
    buffer []*byte
    readCursor, writeCursor int
}

var (
    ErrEmptyBuffer = errors.New("Buffer is empty")
    ErrFullBuffer = errors.New("Buffer is full")
)

func NewBuffer(size int) *Buffer {
	return &Buffer{buffer: make([]*byte, size)}
}

func (b *Buffer) ReadByte() (byte, error) {
    if c := b.buffer[b.readCursor]; c == nil {
        return 0, ErrEmptyBuffer
    } else {
        b.buffer[b.readCursor], b.readCursor = nil, (b.readCursor + 1) % len(b.buffer)
        return *c, nil
    }
}

func (b *Buffer) WriteByte(c byte) error {
    if b.buffer[b.writeCursor] != nil {
        return ErrFullBuffer
    }
    b.buffer[b.writeCursor], b.writeCursor = &c, (b.writeCursor + 1) % len(b.buffer)
    return nil
}

func (b *Buffer) Overwrite(c byte) {
	if b.buffer[b.writeCursor] != nil {
    	b.readCursor = (b.readCursor + 1) % len(b.buffer)
    }
    b.buffer[b.writeCursor], b.writeCursor = &c, (b.writeCursor + 1) % len(b.buffer)
}

func (b *Buffer) Reset() {
	for i := range b.buffer {
        b.buffer[i] = nil
    }
    b.readCursor, b.writeCursor = 0, 0
}