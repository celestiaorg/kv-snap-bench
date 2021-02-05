.PHONY: run

plots=snapshot.png compaction.png block.png

all: kv-snap-bench

run: stats/RocksDB $(plots)

stats/RocksDB:
	./kv-snap-bench

kv-snap-bench: *.go go.mod go.sum
	go build

stats/RocksDB_%.csv: stats/RocksDB
	grep "$*" stats/RocksDB > $@

%.plt: template.plt
	sed -e 's/PLOT_TITLE/$*/g' template.plt > $*.plt

%.png: %.plt stats/RocksDB_%.csv
	gnuplot $*.plt


