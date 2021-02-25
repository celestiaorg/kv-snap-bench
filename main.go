package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

var (
	// TODO(tdybal): configuration
	cache       = uint64(128 * 1024 * 1024) // smaller value ensures that disk is used more
	nRounds     = 5
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
		&BadgerKV{},
		newSMT(newRocksKV()),
		newSMT(&BadgerKV{}),
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
		stats.measure(kv.Name(), "initial_insert", dir, 0, func() {
			insertInitialData(kv, maxKey, data)
		})

		block := 1
		for r := 1; r <= nRounds; r++ {
			log.Printf("simulation round: %d/%d, number of blocks: %d\n", r, nRounds, nBlocks)
			for b := 1; b <= nBlocks; b++ {
				log.Println("simulating block", block)
				data = randomValues(2 * opsPerBlock)
				stats.measure(kv.Name(), "block", dir, block, func() {
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

				if b%100 == 0 {
					log.Println("creating snapshot")
					stats.measure(kv.Name(), "snapshot", dir, block, func() {
						kv.CommitVersion(uint64(block))
					})
				}
				block++
			}
			log.Println("compacting database")
			stats.measure(kv.Name(), "compaction", dir, r, func() {
				kv.Compact()
			})
		}

		log.Println("closing store")
		kv.Close()

		log.Println("cleaning up")
		stats.flush()
		//os.RemoveAll(dir)
		log.Println("all done for", kv.Name())
	}

}
