#!/usr/bin/env bash


tempFunc(){

	CPU_TEMP=$(cat /sys/class/thermal/thermal_zone0/temp)
	#echo "$(date) @ $(hostname)"
	#echo "-------------------------------------------"
	#echo "GPU => $(/opt/vc/bin/vcgencmd measure_temp)"
	#echo "CPU => $((CPU_TEMP/1000))'C"
	echo "CPU TEMP $CPU_TEMP"  
	echo "cpu_temp $((CPU_TEMP/1000))" | curl --data-binary @- http://127.0.0.1:9091/metrics/job/test_job & 
	#echo "CPU => $((CPU_TEMP/1000))'C"
}

while :
do 
	sleep 10
	tempFunc
done



