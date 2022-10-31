package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"auto_login/pojo"

	"github.com/kirinlabs/HttpRequest"
)

// Get 发送GET请求
func Get(url string) (string, int) {

	// 超时时间：5秒
	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String(), resp.StatusCode
}

// Post 发送登陆的Post请求
func Post(url string, data *pojo.UserData) *HttpRequest.Response {
	req := HttpRequest.NewRequest()
	//	格式化cookie
	Cookie := fmt.Sprintf("EPORTAL_USER_GROUP=null; EPORTAL_COOKIE_OPERATORPWD=; EPORTAL_COOKIE_DOMAIN=false; EPORTAL_COOKIE_SAVEPASSWORD=true; EPORTAL_COOKIE_USERNAME=%s; EPORTAL_COOKIE_PASSWORD=%s; EPORTAL_COOKIE_SERVER=%s; EPORTAL_COOKIE_SERVER_NAME=%s; EPORTAL_AUTO_LAND=true; JSESSIONID=9F5310BE6D56FAF15D3A98772E2F0769", data.UserId, data.Password, data.Server, data.Server)
	req.SetHeaders(map[string]string{
		"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8", //这也是HttpRequest包的默认设置
		"Accept":       "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
		"Connection":   "keep-alive",
		"User-Agent":   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36",
		"Cookie":       Cookie,
	})
	postData := map[string]interface{}{
		"userId":          data.UserId,
		"password":        data.Password,
		"service":         data.Server,
		"queryString":     data.QueryString,
		"operatorPwd":     "",
		"validcode":       "",
		"operatorUserId":  "",
		"passwordEncrypt": "false",
	}
	resp, _ := req.Post(url, postData)
	//body, _ := resp.Body()
	//fmt.Println(string(body))
	return resp
}
