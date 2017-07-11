#!/bin/bash
set -xe

numproc=$(cat /proc/cpuinfo | grep processor | wc -l)
wd=$(pwd)

echo ${numproc}

mkdir -p src/h2o
cd src/h2o
cmake -DWITH_BUNDLED_SSL=on -DWITH_MRUBY=on -DCMAKE_INSTALL_PREFIX=${wd}/local ${wd}/vendor/h2o
make install -j ${numproc}
