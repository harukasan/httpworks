#!/bin/bash
set -xe

wd=$(pwd)

cd vendor/wrk
make
cp ${wd}/vendor/wrk/wrk ${wd}/local/bin/wrk
