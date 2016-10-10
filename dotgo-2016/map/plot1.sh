set terminal png size 1920,1440 crop enhanced font "/usr/share/fonts/truetype/times.ttf,30" dashlength 2;
set termoption linewidth 3;

set output "list-vs-map1.png";
set ylabel "milliseconds"
set xlabel "elements"
plot "linear.out" title "linear" with lines, "map-small.out" title "map" with lines;
