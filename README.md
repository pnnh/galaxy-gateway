go示例项目

## 本地监听文件服务

```shell
cd data
http-server -a 0.0.0.0 -p 3700 --cors
```

## mod相关操作

```shell
# 更新packages
go get -u
# 整理packages
go mod tidy
```

### 编译构建

这里target目录定为输出目录

```shell
# 服务端构建
go build -o target/bin/ github.com/pnnh/galaxy-gateway
# 前端资源构建
cd browser
npm run build
```

### 构建Docker镜像

```bash
# 构建docker镜像
cd target
sudo docker build -f Dockerfile -t galaxy-gateway:latest .
# 测试执行构建的镜像
sudo docker run -p 8080:8080 galaxy-gateway
# 仅在本地测试时使用，将aws凭证文件挂载到docker容器
sudo docker run -p 8080:8080 -v $HOME/.aws/credentials:/root/.aws/credentials:ro galaxy-gateway
```
