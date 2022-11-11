# MyChatRoom


### 功能实现

   分为用户和管理员两部分。用户模块实现了，登录，注册，邮箱验证，重置密码，跨域Token，群聊，私聊，添加和删除好友，查询，存储及修改资料。消息模块使用了websocket协议编程，实现消息的收发及存储查询。聊天室模块实现了存储及增删改查。
   内部实现了跨域，Token鉴权，mysql，mongodb，redis处理数据，文件配置及实时监听，日志持久化，实时，循环分割及软连接。swagger，postman接口文档，详细注释说明。错误处理，代码复用，Docker部署等功能。

### 技术应用

前端：暂无前端，有兴趣的可以自己做做
后端:：mysql、mongodb、redis数据库。gin、gorm框架。Viper配置管理，logurs、rotatelogs日志管理，swaggo接口文档，cors跨域，jwt跨域鉴权，md5加密,email邮箱验证、websocket协议编程，air热部署等

### 安装教程

#### docker下的安装
需要安装docker，
并在docker中拉取启动mysql:8.0.28:并创建数据库mychatroom；
拉取mongodb：5.0.8 并创建数据库mychatroom，集合message，user_room；
拉取redis:5.0.14 

修改config下的config.yaml文件
进入文件所在的文件夹，执行docker build -f Dockerfile -t myblog .
执行 docker run -p 9001:9001 -d myblog


#### 直接安装访问
需要配置go环境，并安装mysql：8.0.28，创建数据库mychatroom
安装mongodb：5.0.8 并创建数据库mychatroom，集合message，user_room
安装redis:5.0.14 

修改config.yaml文件
执行命令go mod tidy下载go依赖包
go run main.go执行


### 使用说明

需要在windows下或者docker下安装mysql,mongodb,redis
如果主机名不是localhost，端口号想修改，需要修改config/config.yaml

swagger导入依赖后，swag init初始化 进入路由界面即可http://localhost:9001/swagger/index.html
授权token格式 Bearer+空格+登录返回reponse数据中token中的字符串

### 参与贡献

Fork 本仓库
新建 Feat_xxx 分支
提交代码
新建 Pull Request

### 项目实例

swagger:
![swagger](https://github.com/shisanxiaobai/MyChatRoom/blob/main/image/Snipaste_2022-11-11_22-55-46.png)
