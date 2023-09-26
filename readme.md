
## 项目目录说明
```shell
simpleTikTok
├── Dockerfile
├── apiServer // API 服务器代码
│   ├── biz
│   │   ├── handler
│   │   │   ├── ApiServer // 各个服务处理 HTTP 请求代码
│   │   │   │   ├── relation_service.go
│   │   │   │   ├── social_service.go
│   │   │   │   ├── user_service.go
│   │   │   │   └── video_service.go
│   │   │   ├── ping.go
│   │   │   └── response // 各个服务的构建响应代码
│   │   │       ├── relation.go
│   │   │       ├── social.go
│   │   │       ├── user.go
│   │   │       └── video.go
│   │   ├── model // API 服务 idl 生成代码，由 hz 生成
│   │   │   ├── ApiServer
│   │   │   │   └── apiServer.pb.go
│   │   │   └── api
│   │   │       └── api.pb.go
│   │   └── router // 路由代码，由 hz 生成
│   │       ├── ApiServer
│   │       │   ├── apiServer.go
│   │       │   └── middleware.go // API 服务中间件设定
│   │       └── register.go
│   ├── build.sh
│   ├── cmd
│   │   └── init.go
│   ├── main.go // API 服务启动运行入口
│   ├── router.go
│   ├── router_gen.go
│   ├── rpc // API 服务器相关 rpc 调用
│   │   ├── init.go
│   │   ├── relation.go
│   │   ├── social.go
│   │   ├── user.go
│   │   └── video.go
│   └── script
│       └── bootstrap.sh
├── config // 服务配置项
│   ├── config.yaml
│   └── docker 启动.md
├── data // minio服务视频存储文件夹
│   ├── images
│   └── videos
├── docker-compose.yml
├── go.mod
├── go.sum
├── idl // idl 文件夹
│   ├── BaseResp.proto
│   ├── RelationServer.proto
│   ├── SocialServer.proto
│   ├── UserServer.proto
│   ├── VideoServer.proto
│   ├── api.proto
│   ├── apiServer.proto
│   └── 代码生成.md
├── init.sql
├── kitex_gen // 由 kitex 生成的代码
│   ├── BaseResponse
│   ├── RelationService
│   ├── SocialService
│   ├── UserService
│   └── VideoService
├── package-lock.json
├── package.json
├── pkg // 公用包
│   ├── config // 配置读取
│   │   └── init.go
│   ├── db // mysql 数据库初始化
│   │   ├── init.go
│   │   └── 数据库实例使用说明.md
│   ├── encryption // 密码加密，用于登录和注册
│   │   └── encrypter.go
│   ├── errno // 错误代码以及错误工具包
│   │   ├── code.go
│   │   └── errno.go
│   ├── log // 日志包，基于 zap 封装
│   │   ├── log.go
│   │   ├── options.go
│   │   └── 日志使用说明.md
│   └── middleware // 中间件文件
│       └── jwt.go
├── script
│   └── bootstrap.sh
├── service // 微服务模块
│   ├── dtm // dtm 服务
│   │   ├── init.go 
│   │   ├── main.go // dtm 服务启动入口
│   │   └── request //用于处理需要使用 dtm 服务的请求
│   │       ├── commentRequest.go
│   │       ├── favoriteRequest.go
│   │       └── init.go
│   ├── message // 聊天服务（未实现）
│   │   ├── api
│   │   └── cmd
│   ├── relation // 关系服务
│   │   ├── build.sh
│   │   ├── dal // 数据库访问层
│   │   │   ├── init.go 
│   │   │   ├── model // gorm 模型
│   │   │   │   └── relation.go
│   │   │   ├── redis 
│   │   │   │   ├── init.go
│   │   │   │   └── relation.go // redis 数据访问
│   │   │   └── relation.go // mysql 数据访问
│   │   ├── handler.go // 关系服务 rpc 调用请求
│   │   ├── init.go // 关系服务初始化函数
│   │   ├── kitex_info.yaml
│   │   ├── main.go // 启动关系服务
│   │   ├── mq // 关系服务消息队列
│   │   │   ├── consumer.go
│   │   │   ├── init.go
│   │   │   └── producer.go
│   │   ├── response // 关系服务响应封装
│   │   │   └── resp.go
│   │   ├── rpc // 关系服务所调用的 rpc 服务
│   │   │   ├── init.go
│   │   │   └── user.go
│   │   └── service // 关系服务所提供服务
│   │       ├── getFollowList.go
│   │       ├── getFollowerList.go
│   │       ├── queryRelation.go
│   │       ├── queryUserInfosWithRelation.go
│   │       └── relationAction.go
│   ├── social // 结构与上面类似，不同服务根据所需有细微不同
│   │   ├── build.sh
│   │   ├── dal
│   │   │   ├── comment.go
│   │   │   ├── favorite.go
│   │   │   ├── init.go
│   │   │   └── model
│   │   │       ├── comment.go
│   │   │       └── favorite.go
│   │   ├── handler.go
│   │   ├── init.go
│   │   ├── kitex_info.yaml
│   │   ├── main.go
│   │   ├── response
│   │   │   └── resp.go
│   │   ├── rpc
│   │   │   └── init.go
│   │   └── service
│   │       ├── commentAction.go
│   │       ├── commentList.go
│   │       ├── favoriteAction.go
│   │       └── favoriteList.go
│   ├── user
│   │   ├── build.sh
│   │   ├── dal
│   │   │   ├── init.go
│   │   │   ├── model
│   │   │   │   ├── dynamicUserFields.go
│   │   │   │   ├── stableUserFields.go
│   │   │   │   ├── user.go
│   │   │   │   └── userInfo.go
│   │   │   ├── redis
│   │   │   │   ├── init.go
│   │   │   │   └── user.go
│   │   │   └── user.go
│   │   ├── handler.go
│   │   ├── init.go
│   │   ├── kitex_info.yaml
│   │   ├── main.go
│   │   ├── mq
│   │   │   ├── consumer.go
│   │   │   ├── init.go
│   │   │   └── producer.go
│   │   ├── response
│   │   │   └── resp.go
│   │   ├── rpc
│   │   │   └── init.go
│   │   ├── service
│   │   │   ├── changeUserFollowCount.go
│   │   │   ├── getUserInfo.go
│   │   │   ├── login.go
│   │   │   ├── mGetUserInfo.go
│   │   │   ├── register.go
│   │   │   └── updateUserInfo.go
│   │   ├── test
│   │   │   └── testRedis.go
│   │   └── user
│   └── video
│       ├── dal
│       │   ├── init.go
│       │   ├── model
│       │   │   └── video.go
│       │   └── video.go
│       ├── handler.go
│       ├── init.go
│       ├── kitex_info.yaml
│       ├── main.go
│       ├── minio // 视频服务 minio 相关函数
│       │   ├── docker 启动说明.md
│       │   ├── init.go
│       │   └── service.go
│       ├── mq
│       │   ├── consumer.go
│       │   ├── init.go
│       │   ├── producer.go
│       │   ├── updateInfoConsumer.go
│       │   └── updateInfoProducer.go
│       ├── response
│       │   └── resp.go
│       ├── rpc
│       │   └── init.go
│       └── service
│           ├── getVideoInfo.go
│           ├── getVideoListByIds.go
│           ├── getVideoPublishList.go
│           ├── publishAction.go
│           ├── updateVideoInfo.go
│           └── videoFeed.go
├── test
│   └── main.go
└── tmp // 日志文件存放
    └── log
        ├── common
        │   └── log.log
        ├── dtm
        │   └── dtm.log
        ├── relation
        │   └── relation.log
        ├── social
        │   └── social.log
        └── user
            └── user.log

```
