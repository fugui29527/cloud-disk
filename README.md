# cloud-disk go-zero
创建core服务
goctl api new core

启动服务
go run core.go -f etc/core-api.yaml

#生成代码命令
goctl api go -api core.api -dir . -style gozero