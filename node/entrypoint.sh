#!/bin/bash

# 安裝依賴
echo "Installing dependencies..."
pnpm install

# 啟動開發伺服器
echo "Starting pnpm dev..."
exec pnpm dev
