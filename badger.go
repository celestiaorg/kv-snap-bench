package main

type BadgerKV struct {
}

var _ KV = &BadgerKV{}

func (b *BadgerKV) Open(path string, cache uint64) error {
	panic("not implemented") // TODO: Implement
}

func (b *BadgerKV) Close() error {
	panic("not implemented") // TODO: Implement
}

func (b *BadgerKV) Get(key []byte) []byte {
	panic("not implemented") // TODO: Implement
}

func (b *BadgerKV) Set(key []byte, val []byte) {
	panic("not implemented") // TODO: Implement
}

func (b *BadgerKV) Remove(key []byte) {
	panic("not implemented") // TODO: Implement
}

func (b *BadgerKV) Compact() {
	panic("not implemented") // TODO: Implement
}

func (b *BadgerKV) CommitVersion(v uint64) {
	panic("not implemented") // TODO: Implement
}

func (b *BadgerKV) RemoveVersion(v uint64) {
	panic("not implemented") // TODO: Implement
}

func (b *BadgerKV) GetAt(v uint64, key []byte) []byte {
	panic("not implemented") // TODO: Implement
}

func (b *BadgerKV) Name() string {
	return "Badger"
}
