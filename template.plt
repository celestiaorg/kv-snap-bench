set datafile separator ','

set terminal pngcairo size 800,450 enhanced font 'DejaVu Sans,10'
set output 'PLOT_TITLE.png'

set title "PLOT_TITLE time"
set xlabel "Block number"
set ylabel "Time [ns]"

plot 'stats/RocksDB_PLOT_TITLE.csv' u 2:3 w lp t 'RocksDB'

