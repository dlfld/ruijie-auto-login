# Ruijie Auto Login
```
    ____        _   _ _      ___    __ 
   / __ \__  __(_) (_|_)__  /   |  / / 
  / /_/ / / / / / / / / _ \/ /| | / /  
 / _, _/ /_/ / / / / /  __/ ___ |/ /___
/_/ |_|\__,_/_/_/ /_/\___/_/  |_/_____/
             /___/                     
```


## 1 简介
苦于校园网每次连接都要弹出网页，在网页上登录认证，遂想出此方法一键连接校园网（在不打开网页的情况下）

在这个程序的基础之上可以加一些触发条件比如把这个加到开机自启动里面，这样以后开机就自动联网，不用手动联网了。
<hr>


## 2 当前进度

✅ **Shell**版本实现日志管理  
✅ **GO**版本实现日志管理

❌ 校园网最近新增了一个选项：此设备无感认证，配置支持此功能

<hr/>

## 3 使用方式

### 3.1Shell
- 配置`shell/ruijie.conf`配置即可，（账号密码、运营商必填）
- 运行shell脚本，`nohup ruijie-connect-server1.2.sh &`

## 3.2 GO
- configuration.yaml中配置相应项即可，（账号密码、运营商必填）
    - 配置文件说明：
    ```
    UserId:             用户名
    Password:           用户密码
    Server:             网络服务商，可选参数：
                            移动: 1 | 移动 | yd
                            联通: 2 | 联通 | lt
                            移动: 3 | 电信 | dx
                            校园网: 0 | 教育网 | 校园网 | edu
    TimeInterval: 600   时间间隔（s,秒）
    LogPath:            日志文件相对路径（以可执行文件位置为基准,默认为'ruijie.log'）
    LogClearDay: 7      日志清除时间（day）
    ```
- 在对应平台编译后使用

## 4 求求帮我点一个Star✨
