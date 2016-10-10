set terminal png size 1920,1440 crop enhanced font "/usr/share/fonts/truetype/times.ttf,30" dashlength 2;
set termoption linewidth 3;

set output "cache.png";
set title "";
set ylabel "milliseconds"
set xlabel "step size";
set logscale x 2;
plot "c.out" title "";
