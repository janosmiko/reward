FROM nginx:1.17-alpine

COPY --from=janosmiko/envsubst /usr/bin/envsubst /usr/local/bin/envsubst

RUN set -eux \
  && apk add --no-cache bash shadow \
  && apk add --no-cache --virtual .build-deps curl \
#  && curl -fsSLo /usr/local/bin/envsubst \
#     "https://github.com/a8m/envsubst/releases/download/v1.2.0/envsubst-`uname -s`-`uname -m`" \
#  && chmod +x /usr/local/bin/envsubst \
  && apk del --no-cache .build-deps

ENV NGINX_UPSTREAM_HOST           php-fpm
ENV NGINX_UPSTREAM_PORT           9000
ENV NGINX_UPSTREAM_DEBUG_HOST     php-debug
ENV NGINX_UPSTREAM_DEBUG_PORT     9000
ENV NGINX_UPSTREAM_BLACKFIRE_HOST php-blackfire
ENV NGINX_UPSTREAM_BLACKFIRE_PORT 9000
ENV NGINX_ROOT                    /var/www/html
ENV NGINX_PUBLIC                  ''
ENV NGINX_TEMPLATE                application.conf
ENV XDEBUG_CONNECT_BACK_HOST      '""'

COPY ./etc/. /etc/

CMD find /etc/nginx -name '*.template' -exec sh -c 'envsubst <${1} >${1%.*}' sh {} \; \
  && nginx -g "daemon off;"

WORKDIR /var/www/html
