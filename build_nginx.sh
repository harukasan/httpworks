#!/bin/bash
set -xe

wd=$(pwd)
nginx_version=1.13.1

bindir=${wd}/local/bin

mkdir -p ${bindir}

tar xf vendor/pcre/pcre-8.41.tar.bz2 -C ${wd}/src

if [ "$(uname)" == "Linux" ]; then
  tar xf vendor/nginx-build/nginx-build-linux-amd64-0.10.0.tar.gz -C ${bindir}
elif  [ "$(uname)" == "Darwin" ]; then
  tar xf vendor/nginx-build/nginx-build-darwin-amd64-0.10.0.tar.gz -C ${bindir}
fi

local/bin/nginx-build -v ${nginx_version} -d ${wd}/src \
  -libressl -libresslversion=2.5.4 \
  --prefix=${wd}/local/nginx --with-pcre=${wd}/src/pcre-8.41 --sbin-path=${wd}/local/bin/nginx

cd ${wd}/src/nginx/${nginx_version}/nginx-${nginx_version}
make install

