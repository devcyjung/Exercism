package circular

import "errors"

type Buffer struct{
    buf		[]byte
    len		int
    cap		int
    start	int
    end		int
}

func NewBuffer(size int) *Buffer {
    return &Buffer{
        buf:	make([]byte, size),
        len:	0,
        cap:	size,
        start:	0,
        end:	0,
    }
}

func (b *Buffer) ReadByte() (r byte, e error) {
    if b.len == 0 {
        e = errors.New("Read from empty buffer!")
        return
    }
    r = b.buf[b.start]
    b.start = (b.start+1) % b.cap
    b.len--
    return
}

func (b *Buffer) WriteByte(c byte) (e error) {
    if b.len == b.cap {
        e = errors.New("Write on full buffer!")
        return
    }
    b.buf[b.end] = c
    b.end = (b.end+1) % b.cap
    b.len++
    return
}

func (b *Buffer) Overwrite(c byte) {
    if b.len == b.cap {
        b.ReadByte()
    }
    b.WriteByte(c)
    return
}

func (b *Buffer) Reset() {
    b.start = 0
    b.end = 0
    b.len = 0
    return
}