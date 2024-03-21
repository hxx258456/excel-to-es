package reader

import "golang.org/x/exp/mmap"

type Reader struct {
	*mmap.ReaderAt
	index int64
}

func NewReader(readAt *mmap.ReaderAt) *Reader {
	return &Reader{
		ReaderAt: readAt,
	}
}

func (r *Reader) Read(buf []byte) (n int, err error) {
	n, err = r.ReadAt(buf, r.index)
	r.index += int64(n)
	return
}
