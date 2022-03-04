# ruijie-auto-login
用于登录网页认证版本的锐捷客户端
### 苦于校园网每次连接都要弹出网页，在网页上登录认证，遂想出此方法一键连接校园网（在不打开网页的情况下）
在这个脚本的基础之上可以加一些触发条件比如把这个加到开机自启动里面，这样以后开机就自动联网，不用手动联网了。

目前各个版本均实现配置文件式配置，自动检测是否掉线，如果掉线自动连接

## 当前进度

✅ Shell版本实现日志管理

<hr/>

# 用法：
首先得手动登录一次获取一些必要的信息字段保存到脚本里面
## ☞ 在断网情况下打开浏览器（推荐Chrome）
![Image](https://gitee.com/dlfdd/readme-image-folder/raw/master/ruijie-step-1.png)
## ☞ 点击再次连接or连接  ① 这个时候右边会现实很多请求 找到后缀为login的请求并点击，这个时候右边会出来一堆参数② 往下拉 在Form Data这一项中找到 service这个参数 复制他的值
![Image](https://gitee.com/dlfdd/readme-image-folder/raw/master/ruijie-step-2.png)
![Image](https://gitee.com/dlfdd/readme-image-folder/raw/master/ruijie-step-3.png)
## ☞ 把这个值粘贴到脚本中的service这一项中并填上自己的学号和密码就ok了
![Image](https://gitee.com/dlfdd/readme-image-folder/raw/master/ruijie-step-4.png)

​                        

