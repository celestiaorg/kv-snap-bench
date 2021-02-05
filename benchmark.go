package main

import (
	"bufio"
	"encoding/csv"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type KV interface {
	Get(key []byte) []byte
	Set(key, val []byte)
	Remove(key []byte)
	// Iterator(start, end []byte)

	Open(path string, cache uint64) error
	Close() error
	Compact()

	CommitVersion(v uint64)
	RemoveVersion(v uint64)
	GetAt(v uint64, key []byte) []byte
	// IteratorAt(v uint64, start, end []byte)

	Name() string
}

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

func (s *stats) measure(kv, operation string, n int, fn func()) {
	start := time.Now().UnixNano()
	fn()
	end := time.Now().UnixNano()

	w := s.w[kv]
	w.Write([]string{operation, strconv.Itoa(n), strconv.FormatInt(end-start, 10)})
}

func (s *stats) flush() {
	for _, w := range s.w {
		w.Flush()
	}
}

var (
	// TODO(tdybal): configuration
	cache       = uint64(128 * 1024 * 1024) // smaller value ensures that disk is used more
	nRounds     = 10
	nBlocks     = 1000
	opsPerBlock = 2000
	maxKey      = 1_000_000
	minValSize  = 8
	maxValSize  = 1024
	dataDir     = "./tmp"
	statsDir    = "./stats"
)

func main() {
	log.Println("Starting KV-store versioning benchmark")

	stores := []KV{
		newRocksKV(),
		//&BadgerKV{},
	}

	stats := newStats(stores)

	for _, kv := range stores {
		log.Println("benchmarking", kv.Name())

		// TODO(tzdybal): add data directory option
		os.MkdirAll(dataDir, 0755)
		dir, err := ioutil.TempDir(dataDir, kv.Name())
		if err != nil {
			log.Fatalln("error creating temporary data directory:", err)
		}
		log.Println("initializing store in", dir)
		kv.Open(dir, cache)

		data := randomValues(maxKey)
		stats.measure(kv.Name(), "initial_insert", 0, func() {
			insertInitialData(kv, maxKey, data)
		})

		for b := 1; b <= nBlocks; b++ {
			log.Println("simulating block", b)
			data = randomValues(2 * opsPerBlock)
			stats.measure(kv.Name(), "block", b, func() {
				// update existing values
				for i := 0; i < opsPerBlock; i++ {
					key := rand.Intn(maxKey)
					kv.Set(hash(key), data[i])
				}
				// insert new values
				for i := 0; i < opsPerBlock; i++ {
					key := maxKey + i
					kv.Set(hash(key), data[opsPerBlock+i])
				}
			})

			log.Println("creating snapshot")
			stats.measure(kv.Name(), "snapshot", b, func() {
				kv.CommitVersion(uint64(b))
			})
		}

		log.Println("closing store")
		kv.Close()

		log.Println("cleaning up")
		stats.flush()
		os.RemoveAll(dir)
		log.Println("all done for", kv.Name())
	}

}
