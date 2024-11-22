ARG GO_VERSION
FROM golang:${GO_VERSION}

WORKDIR /app

COPY ./entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# 安裝 air
# https://github.com/air-verse/air
RUN go install github.com/air-verse/air@latest

# 暴露應用運行的端口
EXPOSE 8080

# 運行應用
ENTRYPOINT ["/entrypoint.sh"]
