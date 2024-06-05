package main

import "os"

type Log struct {
	dir      string
	segments []*Segment

	offset int64
	buffer []*Data
	size   int64

	opts *LogOpts
}

type LogOpts struct {
	MaxBufferSize int64
}

type Data struct {
	Key   int64
	Value []byte
}

func NewLog(dir string) (*Log, error) {
	_, err := os.Stat(dir)
	if err != nil {
		err = os.MkdirAll("./"+dir, 0777)
		if err != nil {
			return nil, err
		}
	} else {
	}

	l := &Log{
		dir: dir,
	}

	return l, nil
}

func (l *Log) Append(data []byte) (int64, error) {
	size := 8 + len(data)

	key := l.offset

	l.buffer = append(l.buffer, &Data{
		Key:   key,
		Value: data,
	})

	l.offset += 1
	l.size += int64(size)

	if l.size > 512*1024*1024 {
		segment, err := NewSegment(l.dir, int64(len(l.segments)))
		if err != nil {
			return 0, err
		}

		segment.Write(l.buffer)
		l.segments = append(l.segments, segment)
		l.buffer = nil
		l.offset = 0
	}

	return key, nil
}

func (l *Log) Close() error {
	segment, err := NewSegment(l.dir, int64(len(l.segments)))
	if err != nil {
		return err
	}

	segment.Write(l.buffer)

	l.segments = append(l.segments, segment)

	segment.Close()

	return nil
}
