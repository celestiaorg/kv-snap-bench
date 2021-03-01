package main

type KV interface {
	Get(key []byte) []byte
	Set(key, val []byte)
	Remove(key []byte)
	// Iterator(start, end []byte)

	Open(path string, cache uint64) error
	Close() error
	Compact()

	CommitVersion(v uint64) error
	RemoveVersion(v uint64) error
	GetAt(v uint64, key []byte) []byte
	// IteratorAt(v uint64, start, end []byte)

	// helpers
	Name() string
}
