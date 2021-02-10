package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
)

type stats struct {
	w map[string]*csv.Writer
}

func newStats(stores []KV) *stats {
	s := &stats{w: make(map[string]*csv.Writer, len(stores))}

	os.Mkdir(statsDir, 0755)

	for _, store := range stores {
		file, err := os.Create(filepath.Join(statsDir, store.Name()))
		if err != nil {
			log.Fatalln("Cannot create output stat file:", err)
		}
		s.w[store.Name()] = csv.NewWriter(bufio.NewWriter(file))
		s.w[store.Name()].Write([]string{"operation", "n", "duration"})
	}

	return s
}

func (s *stats) measure(kv, operation, path string, n int, fn func()) {
	startSize, _ := getSize(path)
	startTime := time.Now().UnixNano()
	fn()
	endTime := time.Now().UnixNano()
	endSize, _ := getSize(path)

	w := s.w[kv]
	w.Write([]string{operation, strconv.Itoa(n),
		strconv.FormatInt(endTime-startTime, 10),
		strconv.FormatFloat(float64(startSize)/1024.0/1024.0/1024.0, 'f', 9, 64),
		strconv.FormatFloat(float64(endSize)/1024.0/1024.0/1024.0, 'f', 9, 64),
	})
}

func (s *stats) flush() {
	for _, w := range s.w {
		w.Flush()
	}
}

// sizeAggregator collects size per inode, so it's able to size of hardlinked files correctly
type sizeAggregator struct {
	files map[uint64]uint64
}

func (a *sizeAggregator) sum() uint64 {
	var sum uint64
	for _, v := range a.files {
		sum += v
	}
	return sum
}

func (a *sizeAggregator) collect(path string, info os.FileInfo, err error) error {
	if info == nil {
		return errors.New("info is nil")
	}
	sys := info.Sys()
	if sys == nil {
		return errors.New("Sys() is nil")
	}
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return errors.New("Not a syscall.Stat_t")
	}

	if info.IsDir() {
		a.files[stat.Ino] = uint64(info.Size())
	}
	return nil
}

func getSize(path string) (uint64, error) {
	a := sizeAggregator{files: make(map[uint64]uint64)}
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		return a.collect(path, info, err)
	})
	return a.sum(), err
}
