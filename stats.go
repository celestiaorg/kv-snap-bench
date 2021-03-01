package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode"
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
	startSize, _ := du(path)
	startTime := time.Now().UnixNano()
	fn()
	endTime := time.Now().UnixNano()
	endSize, _ := du(path)

	w := s.w[kv]
	w.Write([]string{operation, strconv.Itoa(n),
		strconv.FormatInt(endTime-startTime, 10),
		strconv.FormatFloat(float64(startSize)/1024.0/1024.0, 'f', 9, 64),
		strconv.FormatFloat(float64(endSize)/1024.0/1024.0, 'f', 9, 64),
	})
}

func (s *stats) flush() {
	for _, w := range s.w {
		w.Flush()
	}
}

func du(path string) (uint64, error) {
	cmd := exec.Command("du", "-ks", path)
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}
	str := string(output)

	return strconv.ParseUint(str[:strings.IndexFunc(str, unicode.IsSpace)], 10, 64)
}
