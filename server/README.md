## 代码结构
> ps: 本项目不按照三层架构拆分了，感觉有点乱。
```lua
server
├── api -- api文件夹
│   ├── v1 -- 主要api
│   │   ├── 
├── config -- 配置文件
├── core -- 主要结构
├── initialize -- 初始化文件
├── logs -- 日志文件
├── global -- 公共调用
├── model -- 主要模型
├── router -- 路由
├── service -- 主要业务处理
├── utils -- 工具
├── .air.conf -- air 配置文件
├── config.yaml -- 项目配置文件
├── Dockerfile -- docker 打包文件
├── go.mod -- go的mod文件
├── main.go -- 主文件
├── Makefile -- 
└── README.md  -- 
```