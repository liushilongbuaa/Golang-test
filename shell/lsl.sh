#!bin/bash
#报表需先转utf-8编码
printf "start: \t\t\t\t" && date +%S:%N
cat $1 | awk -F "," '$11>0{print$1,$2,$7,$11}' > nozero_cover.csv
printf "drop 0: \t\t\t" && date +%S:%N
res1=`cat nozero_cover.csv | grep -a "4/27" | sort -u -k 2.1n | wc -l`
res2=`cat nozero_cover.csv | grep -a "4/28" | sort -u -k 2.1n | wc -l`
res3=`cat nozero_cover.csv | grep -a "4/29" | sort -u -k 2.1n | wc -l`
res4=`cat nozero_cover.csv | grep -a "4/30" | sort -u -k 2.1n | wc -l`
res5=`cat nozero_cover.csv | grep -a "5/1" | sort -u -k 2.1n | wc -l`
res6=`cat nozero_cover.csv | grep -a "5/2" | sort -u -k 2.1n | wc -l`
res7=`cat nozero_cover.csv | grep -a "5/3" | sort -u -k 2.1n | wc -l`
echo $res1 $res2 $res3 $res4 $res5 $res6 $res7 | awk '{printf "%-20s%.2f\n","active/day",($1+$2+$3+$4+$5+$6+$7)/7}'
printf "active/day: \t\t\t" && date +%S:%N
total=`cat nozero_cover.csv | awk -F" " 'BEGIN{sum=0}{sum+=$4}END{printf "%15.2f",sum}'`
echo $total | awk '{printf "%-20s%.2f\n","consume/day",$total/7}'
printf "consume/day: \t\t\t" && date +%S:%N
ka=`cat nozero_cover.csv | grep "KA客户" | awk -F" " 'BEGIN{sum=0}{sum+=$4}END{printf "%15.2f",sum}'`
echo $ka | awk '{printf "%-20s%.2f\n","ka-consume/day",$ka/7}'
printf "ka-consume/day: \t\t\t" && date +%S:%N
zhx=`cat nozero_cover.csv | egrep "渠道代理商|中小客户" | awk -F" " 'BEGIN{sum=0}{sum+=$4}END{printf "%15.2f",sum}'`
echo $zhx | awk '{printf "%-20s%.2f\n","zhx-consume/day",$zhx/7}'
printf "zhx-consume/day: \t\t\t" && date +%S:%N
zhxcount=`cat nozero_cover.csv | egrep "渠道代理商|中小客户" | awk -F" " '{print$1,$2}' | sort -u | wc -l`

echo $zhx $zhxcount | awk '{printf "%-20s%.2f\n" ,"zhx-consume/solo",$1/$2}'

kacount=`cat nozero_cover.csv | grep "KA客户" | awk -F" " '{print$1,$2}' | sort -u | wc -l`
echo $ka $kacount | awk '{printf "%-20s%.2f\n" ,"ka",$1/$2}'
printf "ka: \t\t\t" && date +%S:%N