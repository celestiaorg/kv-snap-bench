plots=snapshot.png compaction.png block.png

all: kv-snap-bench

run: stats/RocksDB stats/badger $(plots)

stats/RocksDB stats/badger:
	./kv-snap-bench

kv-snap-bench: *.go go.mod go.sum
	go build

stats/RocksDB_%.csv: stats/RocksDB
	grep "$*" stats/RocksDB > $@

stats/badger_%.csv: stats/badger
	grep "$*" stats/badger > $@

%.plt: template.plt
	sed -e 's/PLOT_TITLE/$*/g' template.plt > $*.plt

%.png: %.plt stats/RocksDB_%.csv stats/badger_%.csv
	gnuplot $*.plt
