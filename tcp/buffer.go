package tcp

import (
	"errors"
	"io"
)

//读缓冲区 供 解码器使用
type buffer struct {
	reader io.Reader
	buf    []byte
	start  int
	end    int
}

func newBuffer(reader io.Reader, len int) buffer {
	buf := make([]byte, len)
	return buffer{reader, buf, 0, 0}
}

func (b *buffer) len() int {
	return b.end - b.start
}

// grow 将有用的字节前移
func (b *buffer) grow() {
	if b.start == 0 {
		return
	}
	copy(b.buf, b.buf[b.start:b.end])
	b.end -= b.start
	b.start = 0
}

// readFromReader 从reader里面读取数据，如果reader阻塞，会发生阻塞
func (b *buffer) readFromReader() (int, error) {
	b.grow()
	n, err := b.reader.Read(b.buf[b.end:])
	if err != nil {
		return n, err
	}
	b.end += n
	return n, nil
}

// seek 返回n个字节，而不产生移位，如果没有足够字节，返回错误
func (b *buffer) seek(start, end int) ([]byte, error) {
	if start < end && b.end-b.start >= end {
		buf := b.buf[b.start+start : b.start+end]
		return buf, nil
	}
	return nil, errors.New("not enough")
}

// read 舍弃offset个字段，读取n个字段,如果没有足够的字节，返回错误
func (b *buffer) read(offset, limit int) ([]byte, error) {
	if b.len() < offset+limit {
		return nil, errors.New("not enough")
	}
	b.start += offset
	buf := b.buf[b.start : b.start+limit]
	b.start += limit
	return buf, nil
}

// read all
func (b *buffer) readAll() []byte {
	return b.buf[:b.end]
}
