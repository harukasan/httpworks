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
(cd src/go-proxy && go build .)
src/go-proxy/go-proxy -addr=:10443 -tls -cert=conf/self.crt -ca-cert=conf/self.crt -key=conf/self.key -url=https://localhost:8443/

curl --cacert conf/self.crt https://localhost:9444/ # proxy to h2o
```

```
./local/bin/wrk -c 10 https://localhost:8443/
```

## Components:

- [nginx](https://nginx.org/) (with nginx-build)
- [pcre](http://www.pcre.org)
- [h2o](https://h2o.examp1e.net/)
- [wrk](https://github.com/wg/wrk)
- [nghttp2](https://github.com/nghttp2/nghttp2)

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

#### Go proxy (go-proxy -> h2o)

```
$ ./local/bin/wrk -c 10 https://localhost:10443/1k
Running 10s test @ https://localhost:10443/1k
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   590.55us    0.94ms  21.70ms   92.95%
    Req/Sec    12.45k   784.12    13.90k    78.50%
  247960 requests in 10.01s, 296.77MB read
Requests/sec:  24763.61
Transfer/sec:     29.64MB
```

### HTTP/2

#### h2o 2.2.2 (static)

```
$ ./local/bin/h2load -n100000 -t2 -c10 https://localhost:8443/
starting benchmark...
spawning thread #0: 5 total client(s). 50000 total requests
spawning thread #1: 5 total client(s). 50000 total requests
TLS Protocol: TLSv1.2
Cipher: ECDHE-RSA-AES256-GCM-SHA384
Server Temp Key: ECDH P-256 256 bits
Application protocol: h2

finished in 973.81ms, 102689.96 req/s, 5.19MB/s
requests: 100000 total, 100000 started, 100000 done, 100000 succeeded, 0 failed, 0 errored, 0 timeout
status codes: 100000 2xx, 0 3xx, 0 4xx, 0 5xx
traffic: 5.06MB (5301350) total, 1.15MB (1201050) headers (space savings 93.18%), 2.19MB (2300000) data
                     min         max         mean         sd        +/- sd
time for request:       25us      7.32ms        90us       116us    99.21%
time for connect:     3.92ms     13.39ms      8.62ms      2.61ms    70.00%
time to 1st byte:     4.04ms     13.52ms      9.69ms      3.16ms    50.00%
req/s           :   10271.87    10999.61    10477.20      222.19    80.00%
```

#### h2o 2.2.2 proxy (h2o -> h2o)

```
$ ./local/bin/h2load -n100000 -t2 -c10 https://localhost:8444/
starting benchmark...
spawning thread #0: 5 total client(s). 50000 total requests
spawning thread #1: 5 total client(s). 50000 total requests
TLS Protocol: TLSv1.2
Cipher: ECDHE-RSA-AES256-GCM-SHA384
Server Temp Key: ECDH P-256 256 bits
Application protocol: h2

finished in 2.08s, 48058.32 req/s, 2.43MB/s
requests: 100000 total, 100000 started, 100000 done, 100000 succeeded, 0 failed, 0 errored, 0 timeout
status codes: 100000 2xx, 0 3xx, 0 4xx, 0 5xx
traffic: 5.06MB (5301465) total, 1.15MB (1201165) headers (space savings 93.18%), 2.19MB (2300000) data
                     min         max         mean         sd        +/- sd
time for request:       62us      9.02ms       188us       224us    98.03%
time for connect:     5.31ms     13.11ms     10.35ms      2.17ms    70.00%
time to 1st byte:     5.87ms     13.82ms     11.66ms      2.30ms    90.00%
req/s           :    4806.71     5843.60     5147.21      334.31    70.00%
```

#### h2o 2.2.2 http_request (h2o -> h2o)

```
$ ./local/bin/h2load -n100000 -t1 -c10 https://localhost:8544/
starting benchmark...
spawning thread #0: 10 total client(s). 100000 total requests
TLS Protocol: TLSv1.2
Cipher: ECDHE-RSA-AES256-GCM-SHA384
Server Temp Key: ECDH P-256 256 bits
Application protocol: h2

finished in 12.23s, 8177.91 req/s, 431.58KB/s
requests: 100000 total, 100000 started, 100000 done, 100000 succeeded, 0 failed, 0 errored, 0 timeout
status codes: 100000 2xx, 0 3xx, 0 4xx, 0 5xx
traffic: 5.15MB (5404110) total, 1.24MB (1303810) headers (space savings 93.31%), 2.19MB (2300000) data
                     min         max         mean         sd        +/- sd
time for request:      349us     10.47ms      1.21ms       480us    90.40%
time for connect:     9.49ms     12.98ms     11.17ms      1.07ms    70.00%
time to 1st byte:    12.35ms     15.38ms     13.59ms      1.11ms    50.00%
req/s           :     817.82      822.94      820.53        1.69    50.00%
```

#### Go proxy (go-proxy -> h2o)

```
$ ./local/bin/h2load -n100000 -t2 -c10 https://localhost:10443/
starting benchmark...
spawning thread #0: 5 total client(s). 50000 total requests
spawning thread #1: 5 total client(s). 50000 total requests
TLS Protocol: TLSv1.2
Cipher: ECDHE-RSA-AES128-GCM-SHA256
Server Temp Key: X25519 253 bits
Application protocol: h2

finished in 5.55s, 18010.32 req/s, 862.24KB/s
requests: 100000 total, 100000 started, 100000 done, 100000 succeeded, 0 failed, 0 errored, 0 timeout
status codes: 100000 2xx, 0 3xx, 0 4xx, 0 5xx
traffic: 4.68MB (4902360) total, 783.20KB (802000) headers (space savings 95.44%), 2.19MB (2300000) data
                     min         max         mean         sd        +/- sd
time for request:      157us     10.24ms       545us       557us    93.84%
time for connect:     7.20ms     13.15ms      9.14ms      1.96ms    70.00%
time to 1st byte:     8.29ms     13.64ms     10.18ms      1.79ms    60.00%
req/s           :    1801.08     1821.26     1806.99        6.22    80.00%
```
