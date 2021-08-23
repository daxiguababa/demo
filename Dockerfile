FROM builder
# 按需安装依赖包

# 设置Go编译参数
COPY . /app
RUN cd /app;\
    go env -w GOPROXY=https://goproxy.cn,direct; \

RUN GOOS=linux go build -o ./main ./main.go; \

COPY --from=builder /app/main /usr/local/bin
ENTRYPOINT [ "main" ]
