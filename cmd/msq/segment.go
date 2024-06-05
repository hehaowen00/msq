package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

const (
	OP_APPEND = 1
	OP_DELETE = 2
)

type Segment struct {
	file   *os.File
	offset int64
	size   int64

	index    map[int64]location
	metadata *Metadata
}

type location struct {
	key    int64
	offset int64
	size   int64
}

func NewSegment(dir string, num int64) (*Segment, error) {
	var f *os.File
	f, err := os.OpenFile(path.Join(dir, fmt.Sprintf("%d.data", num)), os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return nil, err
	}

	metadata, err := NewMetadata(dir, num)
	if err != nil {
		return nil, err
	}

	segment := &Segment{
		file:  f,
		index: map[int64]location{},

		metadata: metadata,
	}

	return segment, nil
}

func (s *Segment) Write(data []*Data) error {
	for i, d := range data {
		log.Println(i)
		r, err := s.file.Seek(0, io.SeekEnd)
		if err != nil {
			return err
		}

		log.Println("pos", r)

		length := []uint8{}
		length = binary.LittleEndian.AppendUint64(length, uint64(d.Key))
		length = binary.LittleEndian.AppendUint64(length, uint64(len(d.Value)))
		log.Println(length)
		log.Println(s.file.Write(length))
		log.Println(s.file.Write(d.Value))

		l := location{
			key:    d.Key,
			offset: r,
			size:   int64(len(d.Value)),
		}

		s.index[int64(d.Key)] = l

		err = s.metadata.Append(l)
		if err != nil {
			return err
		}

		log.Println(s.index)
		log.Println(s.metadata)
	}

	return nil
}

func (s *Segment) Delete(l location) {}

func (s *Segment) Close() {
	s.file.Close()
	s.metadata.Close()
}
