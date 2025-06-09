package circular

import "errors"

type Buffer struct {
    buffer []Node
    readCursor, writeCursor int
}

type Node *byte

var (
    ErrEmptyBuffer = errors.New("Buffer is empty")
    ErrFullBuffer = errors.New("Buffer is full")
)

func NewBuffer(size int) *Buffer {
	return &Buffer{
        buffer:	make([]Node, size),
    }
}

func (b *Buffer) ReadByte() (byte, error) {
	switch b.buffer[b.readCursor] {
    case nil:
        return 0, ErrEmptyBuffer
    default:
        c := *b.buffer[b.readCursor]
        b.buffer[b.readCursor] = nil
        b.readCursor++
        b.readCursor %= len(b.buffer)
        return c, nil
    }
}

func (b *Buffer) WriteByte(c byte) error {
	switch b.buffer[b.writeCursor] {
    case nil:
        b.buffer[b.writeCursor] = &c
        b.writeCursor++
        b.writeCursor %= len(b.buffer)
        return nil
    default:
        return ErrFullBuffer
    }
}

func (b *Buffer) Overwrite(c byte) {
	switch b.buffer[b.writeCursor] {
    case nil:
        b.buffer[b.writeCursor] = &c
        b.writeCursor++
        b.writeCursor %= len(b.buffer)
    default:
        b.buffer[b.writeCursor] = &c
        b.readCursor++
        b.readCursor %= len(b.buffer)
        b.writeCursor++
        b.writeCursor %= len(b.buffer)
    }
}

func (b *Buffer) Reset() {
	for i := range b.buffer {
        b.buffer[i] = nil
    }
    b.readCursor, b.writeCursor = 0, 0
}