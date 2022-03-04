package main

import (
	"bytes"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"io"
	"net/http"
	"strings"
	"time"
)

func Get(url string) string {

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
	return result.String()
}

// UserData 用户数据项结构体
type UserData struct {
	userId      string
	password    string
	service     string
	queryString string
}

func Post(url string, data UserData) {
	req := HttpRequest.NewRequest()
	//	格式化cookie
	Cookie := fmt.Sprintf("EPORTAL_USER_GROUP=null; EPORTAL_COOKIE_OPERATORPWD=; EPORTAL_COOKIE_DOMAIN=false; EPORTAL_COOKIE_SAVEPASSWORD=true; EPORTAL_COOKIE_USERNAME=%s; EPORTAL_COOKIE_PASSWORD=%s; EPORTAL_COOKIE_SERVER=%s; EPORTAL_COOKIE_SERVER_NAME=%s; EPORTAL_AUTO_LAND=true; JSESSIONID=9F5310BE6D56FAF15D3A98772E2F0769", data.userId, data.password, data.service)
	req.SetHeaders(map[string]string{
		"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8", //这也是HttpRequest包的默认设置
		"Accept":       "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
		"Connection":   "keep-alive",
		"User-Agent":   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36",
		"Cookie":       Cookie,
	})
	postData := map[string]interface{}{
		"userId":          data.userId,
		"password":        data.password,
		"service":         data.service,
		"queryString":     data.queryString,
		"operatorPwd":     "",
		"validcode":       "",
		"operatorUserId":  "",
		"passwordEncrypt": "false",
	}
	resp, _ := req.Post(url, postData)
	body, _ := resp.Body()
	fmt.Println(string(body))

}
func main() {
	resString := Get("http://www.google.cn/generate_204")
	loginpageUrl := strings.Split(resString, "'")[1]
	loginUrl := strings.ReplaceAll(strings.Split(loginpageUrl, "?")[0], "index.jsp", "InterFace.do?method=login")
	queryString := strings.Split(loginpageUrl, "?")[1]
	queryString = strings.ReplaceAll(queryString, "&", "%2526")
	queryString = strings.ReplaceAll(queryString, "=", "%253D")
	Post(loginUrl, UserData{
		userId:      "3210706055",
		password:    "2441086385dlf",
		service:     "%E7%A7%BB%E5%8A%A8",
		queryString: queryString,
	})

}
