package main

import (
	"github.com/dgraph-io/badger/v2"
)

type BadgerKV struct {
	db *badger.DB
}

var _ KV = &BadgerKV{}

func (b *BadgerKV) Open(path string, cache uint64) error {
	opts := badger.DefaultOptions(path).WithBlockCacheSize(int64(cache)).WithSyncWrites(false)
	var err error
	b.db, err = badger.Open(opts)
	return err
}

func (b *BadgerKV) Close() error {
	return b.db.Close()
}

func (b *BadgerKV) Get(key []byte) []byte {
	var val []byte
	txn := b.db.NewTransaction(false)
	i, err := txn.Get(key)
	if err == nil {
		val, _ = i.ValueCopy(nil)
	}
	txn.Discard()

	return val
}

func (b *BadgerKV) Set(key []byte, val []byte) {
	b.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, val)
	})
}

func (b *BadgerKV) Remove(key []byte) {
	b.db.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
}

func (b *BadgerKV) Compact() {
	b.db.RunValueLogGC(0.5) // according to godoc, 0.5 is recommended value
}

func (b *BadgerKV) CommitVersion(v uint64) error {
	//panic("not implemented") // TODO: Implement
	b.db.Sync()
	return nil
}

func (b *BadgerKV) RemoveVersion(v uint64) error {
	//panic("not implemented") // TODO: Implement
	return nil
}

func (b *BadgerKV) GetAt(v uint64, key []byte) []byte {
	panic("not implemented") // TODO: Implement
}

func (b *BadgerKV) Name() string {
	return "badger"
}
