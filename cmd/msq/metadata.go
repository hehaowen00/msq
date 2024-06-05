package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"path"
)

type Metadata struct {
	file    *os.File
	entries []location
}

func NewMetadata(dir string, id int64) (*Metadata, error) {
	f, err := os.OpenFile(path.Join(dir, fmt.Sprintf("%d.meta", id)), os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return nil, err
	}

	m := Metadata{
		file: f,
	}

	return &m, nil
}

func (m *Metadata) Append(l location) error {
	m.entries = append(m.entries, l)

	var data []byte
	binary.LittleEndian.AppendUint64(data, uint64(l.key))
	binary.LittleEndian.AppendUint64(data, uint64(l.offset))
	binary.LittleEndian.AppendUint64(data, uint64(l.size))

	_, err := m.file.Write(data)

	return err
}

func (m *Metadata) Close() {
	m.file.Close()
}
