FROM centos:8

ARG RPM_SCRIPT=https://packagecloud.io/install/repositories/varnishcache/varnish64/script.rpm.sh

COPY --from=janosmiko/envsubst /usr/bin/envsubst /usr/local/bin/envsubst

RUN set -eux \
  && curl -fsSLo- "${RPM_SCRIPT}" | bash \
  && dnf install -y epel-release \
  && dnf install -y varnish \
  && dnf clean all \
  && rm -rf /var/cache/dnf

ENV VCL_CONFIG      /etc/varnish/default.vcl
ENV CACHE_SIZE      256m
ENV VARNISHD_PARAMS -p default_ttl=3600 -p default_grace=3600 \
    -p feature=+esi_ignore_https -p feature=+esi_disable_xml_check

COPY default.vcl /etc/varnish/default.vcl.template

ENV BACKEND_HOST    nginx
ENV BACKEND_PORT    80
ENV ACL_PURGE_HOST  0.0.0.0/0

EXPOSE 	80
CMD envsubst < /etc/varnish/default.vcl.template > /etc/varnish/default.vcl \
    && varnishd -F -f $VCL_CONFIG -s malloc,$CACHE_SIZE $VARNISHD_PARAMS
