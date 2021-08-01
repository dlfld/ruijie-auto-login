# !/usr/bin/env python
# -*-coding:utf-8-*-
import requests

USERNAME = '2017081162'
PASSWORD = 'HDD050267a'
SERVICE = r'%E7%A7%BB%E5%8A%A8'
CAPTIVE_SERVER = r'http://www.google.cn/generate_204'


def get_captive_server_response():
    return requests.get(CAPTIVE_SERVER)


def login(response):
    response_text = response.text
    # login_page_url = response_text.split('\'')[1]
    login_page_url = response_text.split('\'')[1]
    login_url = login_page_url.split('?')[0].replace('index.jsp', 'InterFace.do?method=login')
    query_string = login_page_url.split('?')[1]
    query_string = query_string.replace('&', '%2526')
    query_string = query_string.replace('=', '%253D')
    headers = {
        'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8',
        'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8',
        'Connection': 'keep-alive',
        "Cookie": "EPORTAL_USER_GROUP=null; EPORTAL_COOKIE_OPERATORPWD=; EPORTAL_COOKIE_DOMAIN=false; EPORTAL_COOKIE_SAVEPASSWORD=true; EPORTAL_COOKIE_USERNAME=2017081162; EPORTAL_COOKIE_PASSWORD=HDD050267a; EPORTAL_COOKIE_SERVER=%E7%A7%BB%E5%8A%A8; EPORTAL_COOKIE_SERVER_NAME=%E7%A7%BB%E5%8A%A8; EPORTAL_AUTO_LAND=true; JSESSIONID=9F5310BE6D56FAF15D3A98772E2F0769",
        'User-Agent': 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36'
    }
    login_post_data = 'userId={}&password={}&service={}&queryString={}&operatorPwd=&operatorUserId=&validcode=&passwordEncrypt=false'.format(
        USERNAME, PASSWORD, SERVICE, query_string)
    login_result = requests.post(
        url=login_url,
        data=login_post_data,
        headers=headers
    )

    res = login_result.content.decode('utf-8')
    print(res)
    return res


if __name__ == '__main__':
    captive_server_response = get_captive_server_response()
    while captive_server_response.status_code != 204:
        # Login when user is offline
        login(captive_server_response)
        captive_server_response = get_captive_server_response()
    print('You are already online.')
