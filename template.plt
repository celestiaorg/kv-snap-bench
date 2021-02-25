set datafile separator ','

set terminal pngcairo size 800,450 enhanced font 'DejaVu Sans,10'
set output 'PLOT_TITLE.png'

set title "PLOT_TITLE time"
set xlabel "Block number"
# set logscale y
set ylabel "Time [ns]"

plot 'stats/smt_RocksDB_PLOT_TITLE.csv' u 2:3 w l t 'SMT on RocksDB' smooth bezier, \
 'stats/smt_badger_PLOT_TITLE.csv' u 2:3 w l t 'SMT on badger' smooth bezier,
# 'stats/RocksDB_PLOT_TITLE.csv' u 2:3 w l t 'RocksDB' smooth bezier, \
# 'stats/badger_PLOT_TITLE.csv' u 2:3 w l t 'badger' smooth bezier, \

set output 'PLOT_TITLE_size.png'

set title "PLOT_TITLE size"
set xlabel "Block number"
unset logscale y
set ylabel "Size [GiB]"
set autoscale

plot 'stats/smt_RocksDB_PLOT_TITLE.csv' u 2:5 w l t 'SMT on RocksDB' smooth bezier, \
 'stats/smt_badger_PLOT_TITLE.csv' u 2:5 w l t 'SMT on badger' smooth bezier,
# 'stats/RocksDB_PLOT_TITLE.csv' u 2:5 w l t 'RocksDB' smooth bezier, \
# 'stats/badger_PLOT_TITLE.csv' u 2:5 w l t 'badger' smooth bezier, \
