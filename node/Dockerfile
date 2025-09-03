ARG NODE_VERSION
FROM node:${NODE_VERSION}

WORKDIR /app

RUN apt-get update && apt-get install -y openssl ca-certificates
RUN apt-get dist-upgrade -y
RUN npm install -g pnpm

# 複製並設定 entrypoint 腳本
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

EXPOSE 3000

ENTRYPOINT ["/entrypoint.sh"]
