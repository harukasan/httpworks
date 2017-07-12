#!/bin/bash
set -xe

numproc=$(cat /proc/cpuinfo | grep processor | wc -l)
wd=$(pwd)

echo ${numproc}

cd vendor/nghttp2
autoreconf -i
autoconf
automake
./configure --prefix ${wd}/local --disable-python-bindings
make install -j ${numproc}
