ARG NODE_VERSION
FROM node:${NODE_VERSION}

WORKDIR /app

RUN apt-get update && apt-get install -y openssl
RUN apt-get dist-upgrade -y
RUN corepack enable && corepack prepare pnpm@latest --activate

# RUN pnpm install

EXPOSE 3000

CMD ["pnpm", "dev"]
