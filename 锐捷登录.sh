#!/bin/sh
while true;
  do
    captiveReturnCode=`curl -s -I -m 10 -o /dev/null -s -w %{http_code} http://www.google.cn/generate_204`
    if [ "${captiveReturnCode}" = "204" ]; then
      echo "You are already online!"
      exit 0
    fi

    userId=你的学号
    password='你的密码'
    loginPageURL=`curl -s "http://www.google.cn/generate_204" | awk -F \' '{print $2}'`
    #Structure loginURL
    loginURL=`echo ${loginPageURL} | awk -F \? '{print $1}'`
    loginURL="${loginURL/index.jsp/InterFace.do?method=login}"
    queryString=`echo ${loginPageURL} | cut -d "?" -f 2`
    service="%E7%A7%BB%E5%8A%A8"
    queryString="${queryString//&/%2526}"
    queryString="${queryString//=/%253D}"
    #Send Ruijie eportal auth request and output result
    if [ -n "${loginURL}" ]; then
      strA=`curl -s -A "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.91 Safari/537.36" -e "${loginPageURL}" -b "EPORTAL_COOKIE_USERNAME=; EPORTAL_COOKIE_PASSWORD=; EPORTAL_COOKIE_SERVER=; EPORTAL_COOKIE_SERVER_NAME=; EPORTAL_AUTO_LAND=; EPORTAL_USER_GROUP=; EPORTAL_COOKIE_OPERATORPWD=;" -d "userId=${userId}&password=${password}&service=${service}&queryString=${queryString}&operatorPwd=&operatorUserId=&validcode=&passwordEncrypt=false" -H "Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8" -H "Content-Type: application/x-www-form-urlencoded; charset=UTF-8" "${loginURL}"`
      echo $strA
    fi
    echo "检测了一次"
  done