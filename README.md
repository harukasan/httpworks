# http server sandboxes

To build:

```
bash ./build_all.sh
```

To run:

```
./local/bin/h2o -m master -c conf/h2o.conf &

curl --cacert conf/self.crt https://localhost:8443/ # serve a static file
curl --cacert conf/self.crt https://localhost:8444/ # proxy to h2o
curl --cacert conf/self.crt https://localhost:8444/ # proxy to nginx
```

```
./local/bin/nginx -c $(pwd)/conf/nginx.conf -p $(pwd) &

curl --cacert conf/self.crt https://localhost:9443/ # serve a static file
curl --cacert conf/self.crt https://localhost:9444/ # proxy to h2o
curl --cacert conf/self.crt https://localhost:9445/ # proxy to nginx
```

```
./local/bin/wrk -c 10 https://localhost:8443/
```

## Components:

- [nginx](https://nginx.org/) (with nginx-build)
- [pcre](http://www.pcre.org)
- [h2o](https://h2o.examp1e.net/)
- [wrk](https://github.com/wg/wrk)

## Copyright

All components are provided under the respective licenses.

The rest is licensed under MIT License.
