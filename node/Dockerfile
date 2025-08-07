ARG NODE_VERSION
FROM node:${NODE_VERSION}

WORKDIR /app

RUN apt-get update && apt-get install -y openssl ca-certificates
RUN apt-get dist-upgrade -y
RUN npm install -g pnpm

# RUN pnpm install

EXPOSE 3000

CMD ["pnpm", "dev"]
