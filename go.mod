module github.com/lazyledger/kv-snap-bench

go 1.15

require (
	github.com/cosmos/cosmos-sdk v0.40.0-rc5
	github.com/cosmos/iavl v0.15.3
	github.com/dgraph-io/badger/v2 v2.2007.2
	github.com/facebookgo/ensure v0.0.0-20200202191622-63f1cf65ac4c // indirect
	github.com/facebookgo/subset v0.0.0-20200203212716-c811ad88dec4 // indirect
	github.com/lazyledger/smt v0.1.1
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/tecbot/gorocksdb v0.0.0-20191217155057-f0fad39f321c
	github.com/tendermint/tm-db v0.6.3
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.1.0 // indirect
)

replace (
	github.com/cosmos/cosmos-sdk v0.40.0-rc5 => github.com/lazyledger/cosmos-sdk v0.40.0-rc5.0.20210121152417-3addd7f65d1c
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
)
