
set terminal png size 1920,1440 crop enhanced font "/usr/share/fonts/truetype/times.ttf,30" dashlength 2;
set termoption linewidth 3;
set output "insert.png";

set ylabel "milliseconds"
set xlabel "elements"

set logscale x;

plot "bench.out" using 1:2 title "slice" with lines, "bench.out" using 1:3 title "list" with lines;
