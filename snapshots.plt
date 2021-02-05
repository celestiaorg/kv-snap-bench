set datafile separator ','

set terminal pngcairo size 800,450 enhanced font 'DejaVu Sans,10'
set output 'snapshots.png'

set title "Snapshot time"
set xlabel "Block number"
set ylabel "Time [ns]"

plot 'stats/RocksDB_snapshot.csv' u 2:3 w lp t 'RocksDB'

