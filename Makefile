plots=snapshot.png block.png compaction.png 

all: kv-snap-bench 

clean:
	rm -rf tmp/* stats/* *.png kv-snap-bench

run: stats/RocksDB stats/badger $(plots)

plot: $(plots)

kv-snap-bench: *.go go.mod go.sum
	go build

stats/RocksDB_%.csv: stats/RocksDB
	grep "$*" stats/RocksDB > $@

stats/badger_%.csv: stats/badger
	grep "$*" stats/badger > $@

stats/smt_RocksDB_%.csv: stats/smt_RocksDB
	grep "$*" stats/smt_RocksDB > $@

stats/smt_badger_%.csv: stats/smt_badger
	grep "$*" stats/smt_badger > $@

stats/iavl_%.csv: stats/IAVL
	grep "$*" stats/IAVL > $@


%.plt: template.plt
	sed -e 's/PLOT_TITLE/$*/g' template.plt > $*.plt

%.png: %.plt stats/RocksDB_%.csv stats/badger_%.csv stats/smt_RocksDB_%.csv stats/smt_badger_%.csv stats/iavl_%.csv
	gnuplot $*.plt
