#!/bin/bash

# =============================
# V1.2
# add configuration file, logging, remove shell parameters
# 
# 1. in ./ruijie.conf, you can set userId, password, timeInterval
# 
# 2. in ./ruijie.conf, you can set save time in each log, and delete 
#    time 
# 
# 3. in ./ruijie.conf, you can select network from '移动','联通','电信','校园网'
# =============================
# V1.0
# save configuration in .cache
# 
# =============================


# header
source ./utils/init.sh
source ./utils/log4shell.sh
source ./ruijie.conf

LOG_LEVEL_ALL

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
    
    echo 'All parameters in ./ruijie.conf'

    exit 0 # shell shutdown
}

# help infos
case "$1" in
    -h|--help|?) help;;
esac

# None parameter
if [ -z "$userId" ]; then
    ERROR "User account is None!" >> $log_path
    exit 0
fi

# None parameter
if [ -z "$password" ]; then
    ERROR "User password is None!" >> $log_path
    exit 0
fi

# let it relax
if [ "$timeInterval" -lt 600 ]; then
    WARN "WARNING : relogin time is too short!" >> $log_path
    timeInterval=600
fi

# get network HEX encode
case "$network" in
    '移动') service="%E7%A7%BB%E5%8A%A8t";;
    '联通') service="%E8%81%94%E9%80%9A";;
    '电信') service="%E7%94%B5%E4%BF%A1";;
    '校园网'|'教育网') service="%E6%A0%A1%E5%9B%AD%E7%BD%91";;
    ?) ERROR "unknown network" >> $log_path && exit 0;;
esac

while true;
  do
    log4shell()  # logging manage

    captiveReturnCode=`curl -s -I -m 10 -o /dev/null -s -w %{http_code} http://www.google.cn/generate_204`
    if [ "${captiveReturnCode}" = "204" ]; then
      INFO "Already online!"
      sleep $timeInterval
      continue
    fi
    loginPageURL=`curl -s "http://www.google.cn/generate_204" | awk -F \' '{print $2}'`

    #Structure loginURL
    loginURL=`echo ${loginPageURL} | awk -F \? '{print $1}'`
    loginURL="${loginURL/index.jsp/InterFace.do?method=login}"
    queryString=`echo ${loginPageURL} | cut -d "?" -f 2`
    # service="%E6%A0%A1%E5%9B%AD%E7%BD%91"
    queryString="${queryString//&/%2526}"
    queryString="${queryString//=/%253D}"

    #Send Ruijie eportal auth request and output result
    if [ -n "${loginURL}" ]; then
      strA=`curl -s -A "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.91 Safari/537.36" -e "${loginPageURL}" -b "EPORTAL_COOKIE_USERNAME=; EPORTAL_COOKIE_PASSWORD=; EPORTAL_COOKIE_SERVER=; EPORTAL_COOKIE_SERVER_NAME=; EPORTAL_AUTO_LAND=; EPORTAL_USER_GROUP=; EPORTAL_COOKIE_OPERATORPWD=;" -d "userId=${userId}&password=${password}&service=${service}&queryString=${queryString}&operatorPwd=&operatorUserId=&validcode=&passwordEncrypt=false" -H "Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8" -H "Content-Type: application/x-www-form-urlencoded; charset=UTF-8" "${loginURL}"`
      INFO -e "Get: $strA" >> $log_path

      # if success 
      # Login successfully! \n 
    fi
  done
