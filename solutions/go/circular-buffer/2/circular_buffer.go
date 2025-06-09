package circular

import "errors"

type Buffer struct{
    buf							[]byte
    lenth, capac, start, end	int
}

func NewBuffer(size int) *Buffer {
    return &Buffer{
        buf:	make([]byte, size),
        capac:	size,
    }
}

func (b *Buffer) ReadByte() (r byte, e error) {
    if b.lenth == 0 {
        e = errors.New("Read from empty buffer!")
        return
    }
    r, b.start, b.lenth = b.buf[b.start], (b.start+1) % b.capac, b.lenth-1
    return
}

func (b *Buffer) WriteByte(c byte) (e error) {
    if b.lenth == b.capac {
        e = errors.New("Write on full buffer!")
        return
    }
    b.buf[b.end], b.end, b.lenth = c, (b.end+1) % b.capac, b.lenth+1
    return
}

func (b *Buffer) Overwrite(c byte) {
    if b.lenth == b.capac {
        b.ReadByte()
    }
    b.WriteByte(c)
    return
}

func (b *Buffer) Reset() {
    b.start, b.end, b.lenth = 0, 0, 0
    return
}