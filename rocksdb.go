package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/tecbot/gorocksdb"
)

type RocksKV struct {
	db        *gorocksdb.DB
	path      string
	ro        *gorocksdb.ReadOptions
	wo        *gorocksdb.WriteOptions
	snapshots map[uint64]*gorocksdb.Snapshot
}

var _ KV = &RocksKV{}

func newRocksKV() *RocksKV {
	return &RocksKV{
		snapshots: make(map[uint64]*gorocksdb.Snapshot),
	}
}

func (r *RocksKV) Open(path string, cache uint64) error {
	r.path = path
	bbto := gorocksdb.NewDefaultBlockBasedTableOptions()
	bbto.SetBlockCache(gorocksdb.NewLRUCache(cache))
	opts := gorocksdb.NewDefaultOptions()
	opts.SetBlockBasedTableFactory(bbto)
	opts.SetCreateIfMissing(true)
	var err error
	r.db, err = gorocksdb.OpenDb(opts, path)
	r.ro = gorocksdb.NewDefaultReadOptions()
	r.wo = gorocksdb.NewDefaultWriteOptions()
	return err
}

func (r *RocksKV) Close() error {
	r.db.Close()
	return nil
}

func (r *RocksKV) Get(key []byte) []byte {
	data, err := r.db.GetBytes(r.ro, key)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}

func (r *RocksKV) Set(key []byte, val []byte) {
	err := r.db.Put(r.wo, key, val)
	if err != nil {
		log.Fatalln(err)
	}
}

func (r *RocksKV) Remove(key []byte) {
	err := r.db.Delete(r.wo, key)
	if err != nil {
		log.Fatalln(err)
	}
}

func (r *RocksKV) Compact() {
	// TODO: check how to set the range correctly
	r.db.CompactRange(gorocksdb.Range{})
}

func (r *RocksKV) CommitVersion(v uint64) error {
	cp, err := r.db.NewCheckpoint()
	if err != nil {
		return err
	}
	return cp.CreateCheckpoint(r.getCheckpointDir(v), 0)
}

func (r *RocksKV) RemoveVersion(v uint64) error {
	os.RemoveAll(r.getCheckpointDir(v))
	return nil
}

func (r *RocksKV) getCheckpointDir(v uint64) string {
	return filepath.Join(r.path, "checkpoint"+strconv.FormatUint(v, 10))
}

func (r *RocksKV) GetAt(v uint64, key []byte) []byte {
	ro := gorocksdb.NewDefaultReadOptions()
	ro.SetSnapshot(r.snapshots[v])
	data, err := r.db.GetBytes(ro, key)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}

func (r *RocksKV) Name() string {
	return "RocksDB"
}
