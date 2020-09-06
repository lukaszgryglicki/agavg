#!/bin/bash
if [ $# -lt 2 ]
then
  echo "Usage: $0 x y"
  exit 1
fi
x="${1}"
y="${2}"
i=0
while true
do
  nx=`awk "BEGIN {printf \"%.13f\n\", ($x+$y)/2.0}"`
  ny=`awk "BEGIN {printf \"%.13f\n\", sqrt($x*$y)}"`
  if ( [ "$x" = "$nx" ] && [ "$y" = "$ny" ] && [ "$nx" = "$ny" ] )
  then
    break
  fi
  x=$nx
  y=$ny
  i=$((i+1))
  if [ "$i" = "100" ]
  then
    break
  fi
done
echo "agavg($1, $2) --> $x (converge in $i steps)"
