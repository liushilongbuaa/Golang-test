#!bash
if [[ $1 == 6 ]] 
then
	cccomm=service_linkV3_ipv6
	ccser=ipv6-tomerge_servicelink_build
	lbcomm=ipv6_test_servicelink_build
	lbser=ipv6_tomerge_servicelink_build
else
    cccomm=service_linkV3
    ccser=service_linkV3
    lbcomm=service_linkV3
    lbser=service_linkV3
fi

check_cc_common(){
    echo "################## cc-common ##################"
    cd e:/go/src/jd.com/cc/jstack-cc-common/
    git checkout $cccomm
    git pull
}

check_cc_server(){
    echo "################## cc-server ##################"
    cd e:/go/src/jd.com/cc/jstack-cc-server/
    git checkout $ccser
    git pull
}

check_lb_common(){
    echo "################## lb-common ##################"
    cd e:/go/src/jd.com/lb/jstack-lb-common/
    git checkout $lbcomm
    git pull
}

check_lb_server(){
    echo "################## lb-server ##################"
    cd e:/go/src/jd.com/lb/jstack-lb-server/
    git checkout $lbser
    git pull
}

check_external(){
    echo "################## external-sdk ##################"
    cd e:/go/src/jd.com/jstack-external-sdk/
    git checkout service_link
	# git checkout .
    git pull
}

main(){
    check_cc_common
    check_cc_server
    check_lb_common
    check_lb_server
    check_external
}

#main
#exit

for i in $(seq 1 5 )
do
	echo ${i}Hahaha
done

# 脚本名称叫test.sh 入参三个: 1 2 3 
# $* 为"1 2 3"（一起被引号包住）
# $@ 为"1" "2" "3"（分别被包住）
# $# 为3（参数数量）
# $0 当前脚本的文件名
# $n 第一个参数是$1，第二个参数是$2。
# $? 上个命令的返回值
# $$ 当前shell进程ID
if (( $1 == 6 ))
then
    echo HHHHH
fi
# [] 和 test 一样
# [[]] shell内置命令，更强大，每项左右都要加空格
# let 和 (()) 用于算术计算，可以直接使用变量名如var而不需要$var这样的形式