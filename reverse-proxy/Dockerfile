FROM nginx:1.17-alpine

RUN apk add --no-cache gettext

COPY default.conf.template .
COPY docker-nginx-entry.sh .

ENTRYPOINT [ "sh", "/docker-nginx-entry.sh" ]   