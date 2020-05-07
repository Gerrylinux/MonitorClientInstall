# expoter-install 脚本
## Quick Start
- 服务部署

```bash
> go get -v -u github.com/hashicorp/consul/api
> cd cmd
> go build -o expoter-install main.go
```
- 服务运行
```sh
> ./expoter-install -s <mysql|redis|Linux|kafka>
```

- 有其他服务安装需求，在继续增加。
- 依赖服务版本调整在install/install.go

