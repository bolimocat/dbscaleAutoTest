#/bin/bash
source /etc/profile
#echo 'kill -15 pid'
KILL_TYPE=$1
PID_PORT=$2
REBORNID=`ps -ef|grep "$PID_PORT"|grep 'mysqld_safe'`
#echo $REBORNID
REBORN=(`echo $REBORNID|cut -d \  -f 1` `echo $REBORNID|cut -d \  -f 2` `echo $REBORNID|cut -d \  -f 3` `echo $REBORNID|cut -d \  -f 4` `echo $REBORNID|cut -d \  -f 5` `echo $REBORNID|cut -d \  -f 6` `echo $REBORNID|cut -d \  -f 7` `echo $REBORNID|cut -d \  -f 8` `echo $REBORNID|cut -d \  -f 9` `echo $REBORNID|cut -d \  -f 10` `echo $REBORNID|cut -d \  -f 11` `echo $REBORNID|cut -d \  -f 12`)
#echo 'REBORN : '${REBORN[8]}' '${REBORN[9]}' '${REBORN[10]}' '${REBORN[11]}' '${REBORN[12]}

TESTPID=`ps -ef|grep "$PID_PORT"|grep 'mysqld '`
#echo $TESTPID
IP=(`echo $TESTPID|cut -d \  -f 1` `echo $TESTPID|cut -d \  -f 2` `echo $TESTPID|cut -d \  -f 3`)
#echo 'kill -15 '${IP[1]}' '${IP[2]}
kill -15 ${IP[1]} ${IP[2]} 
sleep 15
#echo 'REBORN : '${REBORN[8]}' '${REBORN[9]}' '${REBORN[10]}' '${REBORN[11]}' '${REBORN[12]}
echo 'REBORN : '${REBORN[8]}' '${REBORN[9]}' '${REBORN[10]}' '${REBORN[11]}' '${REBORN[12]} >> /home/greatdb/kill_pid.result
nohup ${REBORN[8]} ${REBORN[9]} ${REBORN[10]} ${REBORN[11]} ${REBORN[12]} > /dev/null 2>&1 &
echo '------'  >> /home/greatdb/kill_pid.result
sleep 5
