package gencode

import (
	"io"
	"time"
	"unsafe"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
	_ = time.Now()
)

type Args struct {
	A int64
	B int64
}

func (d *Args) Size() (s uint64) {

	s += 16
	return
}
func (d *Args) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{

		buf[0+0] = byte(d.A >> 0)

		buf[1+0] = byte(d.A >> 8)

		buf[2+0] = byte(d.A >> 16)

		buf[3+0] = byte(d.A >> 24)

		buf[4+0] = byte(d.A >> 32)

		buf[5+0] = byte(d.A >> 40)

		buf[6+0] = byte(d.A >> 48)

		buf[7+0] = byte(d.A >> 56)

	}
	{

		buf[0+8] = byte(d.B >> 0)

		buf[1+8] = byte(d.B >> 8)

		buf[2+8] = byte(d.B >> 16)

		buf[3+8] = byte(d.B >> 24)

		buf[4+8] = byte(d.B >> 32)

		buf[5+8] = byte(d.B >> 40)

		buf[6+8] = byte(d.B >> 48)

		buf[7+8] = byte(d.B >> 56)

	}
	return buf[:i+16], nil
}

func (d *Args) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.A = 0 | (int64(buf[0+0]) << 0) | (int64(buf[1+0]) << 8) | (int64(buf[2+0]) << 16) | (int64(buf[3+0]) << 24) | (int64(buf[4+0]) << 32) | (int64(buf[5+0]) << 40) | (int64(buf[6+0]) << 48) | (int64(buf[7+0]) << 56)

	}
	{

		d.B = 0 | (int64(buf[0+8]) << 0) | (int64(buf[1+8]) << 8) | (int64(buf[2+8]) << 16) | (int64(buf[3+8]) << 24) | (int64(buf[4+8]) << 32) | (int64(buf[5+8]) << 40) | (int64(buf[6+8]) << 48) | (int64(buf[7+8]) << 56)

	}
	return i + 16, nil
}

type Reply struct {
	C int64
}

func (d *Reply) Size() (s uint64) {

	s += 8
	return
}
func (d *Reply) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{

		buf[0+0] = byte(d.C >> 0)

		buf[1+0] = byte(d.C >> 8)

		buf[2+0] = byte(d.C >> 16)

		buf[3+0] = byte(d.C >> 24)

		buf[4+0] = byte(d.C >> 32)

		buf[5+0] = byte(d.C >> 40)

		buf[6+0] = byte(d.C >> 48)

		buf[7+0] = byte(d.C >> 56)

	}
	return buf[:i+8], nil
}

func (d *Reply) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.C = 0 | (int64(buf[0+0]) << 0) | (int64(buf[1+0]) << 8) | (int64(buf[2+0]) << 16) | (int64(buf[3+0]) << 24) | (int64(buf[4+0]) << 32) | (int64(buf[5+0]) << 40) | (int64(buf[6+0]) << 48) | (int64(buf[7+0]) << 56)

	}
	return i + 8, nil
}
