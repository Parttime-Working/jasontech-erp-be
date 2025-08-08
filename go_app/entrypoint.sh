#!/bin/bash

# 檢查 go.mod 是否存在
if [ ! -f /app/go.mod ]; then
    echo '❌ go.mod not found, initializing module'
    go mod init erp
else
    echo '✅ go.mod found'
fi

# 自動下載所有依賴 (類似 npm install)
echo '📦 下載依賴中...'
go mod download

# 確保依賴完整性 (類似 npm ci)
echo '🔍 驗證依賴完整性...'
go mod verify

# 清理未使用的依賴
echo '🧹 清理依賴...'
go mod tidy

echo '🚀 啟動應用程式...'
air
