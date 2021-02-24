FROM dockerhub.piggy.xiaozhu.com/build/go:v1.0.0 AS builder

WORKDIR /src/
COPY . /src/
RUN mkdir /src/build/config && \
    cp -rf config/* /src/build/config/ && \
    cp startup.sh /src/build/ && \
    go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct && \
    go env -w GOPRIVATE=gitlab.idc.xiaozhu.com && \
    go build -o /src/build/app /src/cmd/apiserver/main.go && \
    go build -o /src/build/cron /src/cmd/cron/main.go

FROM dockerhub.piggy.xiaozhu.com/runtime/golang:v1.0.1
WORKDIR /home/www/
COPY --from=builder /src/build/ /home/www/
RUN mkdir /home/www/log && chmod +x /home/www/startup.sh
EXPOSE 80
ENTRYPOINT ["/home/www/startup.sh"]
