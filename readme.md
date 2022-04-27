# 流媒体网站

## Api 设计
api设计
### User相关
| desc | url | method |
| :----: | :--- | :------: |
|用户注册| /user|  POST |
|用户登录| /user/:username| POST |
|获取用户信息 |/user/:username | GET |
|用户注销 | /user/:username| DELETE |
 

 ## Powershell 

 ### curl
 ```sh
curl -uri 'http://localhost:8000/user/123' -Method post

$body = @{
    user_name = "ado";
    pwd = "12345";
}
$JSON = $body | convertto-json
curl -uri 'http://localhost:8000/user' -Method post -Body $JSON
 ```

 ## Mongo

 ### 连接问题
 Robo 3T可以连接，但是代码连接不了，需要重启mongo并update用户
 [Unauthorized 故障排查手记](https://www.jianshu.com/p/293263675fdc)

 ## Session 与 Cookie
 Session是服务端保持的状态，服务器重启的时候需要去加载所有的session到cache中
 Cookie是前端保存的状态，Cookie里面可以存session key，前端登录过之后重登可以用sessionkey来登录

 ## 用到的库
 * httproute: http server
 * google/uuid: 生成uuid
 * sony/sonyflake: 生成id
 * go-mongo-driver: 采用mongo数据库

 ## TODO
 * 通用的，业务相关性不大的组件可以拆出成库
 * jwt用来生成session
 * 视频和评论分页