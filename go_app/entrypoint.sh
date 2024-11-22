#!/bin/bash

# 初始化 go mod
if [ ! -f /app/go.mod ]; then
    echo 'go.mod not found, initializing module'
    go mod init erp
else
    echo 'go.mod already exists'
fi

# 安裝gin套件
go get -u github.com/gin-gonic/gin

air