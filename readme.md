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
 ```

 ## Mongo

 ### 连接问题
 Robo 3T可以连接，但是代码连接不了，需要重启mongo并update用户
 [Unauthorized 故障排查手记](https://www.jianshu.com/p/293263675fdc)