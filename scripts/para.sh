
begin_time=`date +%s.%N`

#./parallel -j 5 --xapply ./task1.sh ::: `seq 1 5` ::: `seq 11 15`

echo "parallel..."

parallel -j 100 ./api.sh ::: `seq 1 1000`

end_time=`date +%s.%N`

start_s=$(echo $begin_time | cut -d '.' -f 1)
start_ns=$(echo $begin_time | cut -d '.' -f 2)
end_s=$(echo $end_time | cut -d '.' -f 1)
end_ns=$(echo $end_time | cut -d '.' -f 2)

time=$(( ( 10#$end_s - 10#$start_s ) * 1000 + ( 10#$end_ns / 1000000 - 10#$start_ns / 1000000 ) ))

#echo "begin_time:" $begin_time
#echo "end_time:" $end_time
echo "duration:" $time "ms"
