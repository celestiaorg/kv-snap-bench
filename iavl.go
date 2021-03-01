package main

import (
	iavls "github.com/cosmos/cosmos-sdk/store/iavl"
	"github.com/cosmos/iavl"
	dbm "github.com/tendermint/tm-db"
)

type IAVL struct {
	db   dbm.DB
	tree *iavls.Store
}

func (i *IAVL) Get(key []byte) []byte {
	return i.tree.Get(key)
}

func (i *IAVL) Set(key []byte, val []byte) {
	i.tree.Set(key, val)
}

func (i *IAVL) Remove(key []byte) {
	i.tree.Delete(key)
}

func (i *IAVL) Open(path string, cache uint64) error {
	db, err := dbm.NewGoLevelDB("goleveldb", path)
	if err != nil {
		return err
	}
	t, err := iavl.NewMutableTreeWithOpts(db, int(cache/1024), &iavl.Options{InitialVersion: 0})
	if err != nil {
		return err
	}
	i.tree = iavls.UnsafeNewStore(t)
	i.db = db
	return nil
}

func (i *IAVL) Close() error {
	return i.db.Close()
}

func (i *IAVL) Compact() {
}

func (i *IAVL) CommitVersion(v uint64) error {
	i.tree.Commit()
	return nil
}

func (i *IAVL) RemoveVersion(v uint64) error {
	i.tree.DeleteVersions(int64(v))
	return nil
}

func (i *IAVL) GetAt(v uint64, key []byte) []byte {
	panic("not implemented") // TODO: Implement
}

// helpers
func (i *IAVL) Name() string {
	return "IAVL"
}

var _ KV = &IAVL{}

func newSMTLevelDB(dir string, cache int) *IAVL {
	return &IAVL{}
}
