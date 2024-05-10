# 使用最新的golang基础镜像
FROM golang:latest as builder

# 设置工作目录
WORKDIR /app

# 复制go mod和sum文件
COPY go.mod go.sum ./

# 下载依赖项
RUN go mod download

# 复制全部的源文件到容器内
COPY . .

# 编译Go程序为二进制文件，假设输出的文件名为app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# 使用scratch为最终镜像，只包含我们需要的程序和资源
FROM scratch

# 从builder阶段复制编译好的程序和conf目录到当前目录
COPY --from=builder /app/app .
COPY --from=builder /app/conf .

# 暴露必要的端口号
EXPOSE ${PORT}

# 启动Go应用程序
ENTRYPOINT ["./app"]
