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
> ./aliyun_ecs_api -s <mysql|redis|Linux>
