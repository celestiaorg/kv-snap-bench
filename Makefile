plots=snapshots.png

all: kv-snap-bench $(plots)

kv-snap-bench: *.go go.mod go.sum
	go build

stats/RocksDB_snapshot.csv: stats/RocksDB
	grep "snapshot" stats/RocksDB > stats/RocksDB_snapshot.csv

snapshots.png: snapshots.plt stats/RocksDB_snapshot.csv
	gnuplot snapshots.plt
