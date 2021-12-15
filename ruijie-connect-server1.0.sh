#!/bin/bash

#Exit the script when is already online, use www.google.cn/generate_204 to check the online status
# captiveReturnCode=`curl -s -I -m 10 -o /dev/null -s -w %{http_code} http://www.google.cn/generate_204`
# if [ "${captiveReturnCode}" = "204" ]; then
#   echo "You are already online!"
#   exit 0
# fi

#If not online, begin Ruijie Auth
#Get Ruijie login page URL

help(){
	echo '========================================'	
	echo '         ____          _   _  _         '
	echo '        / __ \ __  __ (_) (_)(_)___     '
    echo '       / /_/ // / / // / / // // _ \    '
	echo '      / _, _// /_/ // / / // //  __/    '
	echo '     /_/ |_| \____//_/_/ //_/ \___/     '
	echo '                    /___/               '
	echo '========================================'	
	
	echo 'This a shell script for automaticly con-'
	echo 'nect Ruijie Authorized Network in CUIT.'
	echo ''
	
	echo 'Please input bellow parameters:'
    echo "-u, account     user account."
    echo "-p, password    password."
    echo "-s, save        if .cache is None, save login informations in ./.cache, else update infors."

    exit 0 # shell shutdown
}

# help infos
case "$1" in
    -h|--help|?)
	help
;;
esac


check(){
    if [ -z "$userId" ]; then
        echo "ERROR -u: user ID is None!"
        exit 0
    # else
    #     echo "Current userId is $userId."
    fi

    if [ -z "$password" ]; then
        echo "ERROR -p: Password is None!"
        exit 0
    # else
    #     echo "Current password is $password."
    fi
}


getdir(){
    local workdir=$(cd $(dirname $1); pwd)
    echo $workdir/.cache

    return $?
}

savecache(){
    # $1 save path
    # $2 account 
    # $3 password
    check

    # if file not exists, create it
    if [ ! -e $1 ]; then
        touch $1
    fi

    # overwrite
    echo "userId=$2" > $1
    echo "password=$3" >> $1
}

loadcache(){
    if [ -e $1 ]; then
        source $1
    else
        userId=''
        password=''
    fi
}


# read cache file
cachepath=$(getdir $0)

# if input parameters is None, go to read cache file
if [ -z "$1" ]; then
	loadcache $cachepath
fi


while getopts "u:p:s" opt; do
	case $opt in
		u) 
        userId=${OPTARG}
        # echo "read userId $userId"
        ;;
		p) 
        password=${OPTARG}
        # echo "read password $password"
        ;;
        s) 
        # echo "save cache file"
        savecache $cachepath $userId $password
        ;;
	esac
done

# check userId and password
check


while true;
  do
    SDATE=`date '+%Y-%m-%d %H:%M:%S'`  # datetime string at now

    captiveReturnCode=`curl -s -I -m 10 -o /dev/null -s -w %{http_code} http://www.google.cn/generate_204`
    if [ "${captiveReturnCode}" = "204" ]; then
      echo "${SDATE} Already online!"
      sleep 600
      continue

    fi
    loginPageURL=`curl -s "http://www.google.cn/generate_204" | awk -F \' '{print $2}'`

    #Structure loginURL
    loginURL=`echo ${loginPageURL} | awk -F \? '{print $1}'`
    loginURL="${loginURL/index.jsp/InterFace.do?method=login}"
    queryString=`echo ${loginPageURL} | cut -d "?" -f 2`
    service="%E6%A0%A1%E5%9B%AD%E7%BD%91"
    queryString="${queryString//&/%2526}"
    queryString="${queryString//=/%253D}"

    #Send Ruijie eportal auth request and output result
    if [ -n "${loginURL}" ]; then
      strA=`curl -s -A "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.91 Safari/537.36" -e "${loginPageURL}" -b "EPORTAL_COOKIE_USERNAME=; EPORTAL_COOKIE_PASSWORD=; EPORTAL_COOKIE_SERVER=; EPORTAL_COOKIE_SERVER_NAME=; EPORTAL_AUTO_LAND=; EPORTAL_USER_GROUP=; EPORTAL_COOKIE_OPERATORPWD=;" -d "userId=${userId}&password=${password}&service=${service}&queryString=${queryString}&operatorPwd=&operatorUserId=&validcode=&passwordEncrypt=false" -H "Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8" -H "Content-Type: application/x-www-form-urlencoded; charset=UTF-8" "${loginURL}"`
      echo -e "${SDATE} login successfully! \n Get: $strA"
    fi
  done
