package main

import (
	"crypto/sha256"

	"github.com/lazyledger/smt"
)

type SMT struct {
	kv   KV
	tree *smt.SparseMerkleTree
}

type MapStoreWrapper struct {
	kv KV
}

func (m *MapStoreWrapper) Get(key []byte) ([]byte, error) {
	return m.kv.Get(key), nil
}

func (m *MapStoreWrapper) Set(key []byte, value []byte) error {
	m.kv.Set(key, value)
	return nil
}

func (m *MapStoreWrapper) Delete(key []byte) error {
	m.kv.Remove(key)
	return nil
}

var _ KV = &SMT{}

func newSMT(underlying KV) *SMT {
	return &SMT{
		kv:   underlying,
		tree: smt.NewSparseMerkleTree(&MapStoreWrapper{underlying}, sha256.New()),
	}
}

func (s *SMT) Get(key []byte) []byte {
	val, _ := s.tree.Get(key)
	return val
}

func (s *SMT) Set(key []byte, val []byte) {
	s.tree.Update(key, val)
}

func (s *SMT) Remove(key []byte) {
	s.tree.Delete(key)
}

func (s *SMT) Open(path string, cache uint64) error {
	return s.kv.Open(path, cache)
}

func (s *SMT) Close() error {
	return s.kv.Close()
}

func (s *SMT) Compact() {
	s.kv.Compact()
}

func (s *SMT) CommitVersion(v uint64) error {
	return s.kv.CommitVersion(v)
}

func (s *SMT) RemoveVersion(v uint64) error {
	return s.kv.RemoveVersion(v)
}

func (s *SMT) GetAt(v uint64, key []byte) []byte {
	panic("not implemented") // TODO: Implement
}

// helpers
func (s *SMT) Name() string {
	return "smt_" + s.kv.Name()
}
