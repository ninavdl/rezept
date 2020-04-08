# Copy resources and build html/js/css bundle
FROM node:13.12-alpine AS builder

WORKDIR /app

COPY package.json package-lock.json ./
COPY src ./src

RUN npm install
RUN npm run build

# Replace environment variables in index.html and serve 
FROM nginx:1.17-alpine AS server

RUN apk add --no-cache gettext

COPY docker-nginx-entry.sh /
COPY default.conf /etc/nginx/conf.d

WORKDIR /usr/share/nginx/html

COPY --from=builder /app/dist .
COPY --from=builder /app/dist/index.html index.html.template

ENTRYPOINT [ "sh", "/docker-nginx-entry.sh" ]