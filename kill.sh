process=`ps -ef|grep $1 |grep -v grep`;
if [ "$process" == "" ];then
	echo "NO";
	break;
else 
	echo "Yes";
    ps -ef|grep $1 |grep -v grep;
	ps -ef|grep $1 |grep -v grep|cut -c 9-15|xargs kill ;
fi