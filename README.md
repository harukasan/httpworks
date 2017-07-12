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

## Benchmarks

### HTTP/1.1

#### h2o 2.2.2 (static)

```
$ ./local/bin/wrk -c 10 https://localhost:8443/1k
Running 10s test @ https://localhost:8443/1k
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   141.91us  443.74us  13.83ms   96.87%
    Req/Sec    58.03k     4.54k   68.11k    63.00%
  1154016 requests in 10.00s, 1.37GB read
Requests/sec: 115387.75
Transfer/sec:    140.30MB
```

#### h2o 2.2.2 proxy (h2o -> h2o)

```
$ ./local/bin/wrk -c 10 https://localhost:8444/1k
Running 10s test @ https://localhost:8444/1k
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   341.08us  674.51us  11.06ms   93.58%
    Req/Sec    24.97k     1.24k   28.04k    72.00%
  496617 requests in 10.00s, 603.85MB read
Requests/sec:  49656.78
Transfer/sec:     60.38MB
```

#### h2o 2.2.2 http_request (h2o -> h2o)

```
$ ./local/bin/wrk -c 10 https://localhost:8544/1k
Running 10s test @ https://localhost:8544/1k
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.42ms    1.00ms  11.56ms   91.69%
    Req/Sec     2.75k   142.68     3.07k    70.50%
  54804 requests in 10.00s, 67.37MB read
Requests/sec:   5477.99
Transfer/sec:      6.73MB
```

#### nginx (static)

```
$ ./local/bin/wrk -c 10 https://localhost:9443/1k
Running 10s test @ https://localhost:9443/1k
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   149.85us  382.93us   9.64ms   96.98%
    Req/Sec    45.17k     3.20k   53.12k    67.50%
  898430 requests in 10.00s, 1.06GB read
Requests/sec:  89814.81
Transfer/sec:    108.26MB
```

#### nginx proxy (nginx -> nginx)

```
$ ./local/bin/wrk -c 10 https://localhost:9445/1k
Running 10s test @ https://localhost:9445/1k
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     0.98ms  811.26us  16.73ms   92.99%
    Req/Sec     5.74k   363.71     6.40k    70.50%
  114222 requests in 10.00s, 137.68MB read
Requests/sec:  11418.67
Transfer/sec:     13.76MB
```

#### Go proxy

```
$ ./local/bin/wrk -c 10 https://localhost:10443/1k
Running 10s test @ https://localhost:10443/1k
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.01ms    3.61ms  30.90ms   76.06%
    Req/Sec     1.46k   157.82     1.92k    67.50%
  29092 requests in 10.01s, 34.82MB read
Requests/sec:   2905.23
Transfer/sec:      3.48MB
```

### HTTP/2

(TODO)
