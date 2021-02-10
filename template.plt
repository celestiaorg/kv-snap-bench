set datafile separator ','

set terminal pngcairo size 800,450 enhanced font 'DejaVu Sans,10'
set output 'PLOT_TITLE.png'

set title "PLOT_TITLE time"
set xlabel "Block number"
set ylabel "Time [ns]"

plot 'stats/RocksDB_PLOT_TITLE.csv' u 2:3 w l t 'RocksDB', \
 'stats/badger_PLOT_TITLE.csv' u 2:3 w l t 'badger'

set output 'PLOT_TITLE_SIZE.png'

set title "PLOT_TITLE size"
set xlabel "Block number"
set ylabel "Size [GiB]"

plot 'stats/RocksDB_PLOT_TITLE.csv' u 2:5 w l t 'RocksDB', \
 'stats/badger_PLOT_TITLE.csv' u 2:5 w l t 'badger'
