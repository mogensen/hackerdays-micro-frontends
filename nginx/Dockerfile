FROM nginx:1-alpine

RUN mkdir -p /usr/share/nginx/html/common

COPY nginx.conf /etc/nginx/conf.d/nginx.conf
COPY common /usr/share/nginx/html/common
