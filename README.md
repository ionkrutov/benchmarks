# benchmarks
Benchmarks between defferent packages or languages 

## http `hello world`

Here are shown the results of benchmarks of a simple http server wich returns only `hello world` to the request.
All tests was made on one machine. For detailed information see [Appendix](#appendix).

| Framework                                                        | RPC   | Remark                                                               |
|------------------------------------------------------------------|-------|----------------------------------------------------------------------|
| Golang net/http                                                  | 61692 |                                                                      |
| C++ userver [repo](https://github.com/userver-framework/userver) | 53483 | From Yandex                                                          |
| C++ httplib [repo](https://github.com/yhirose/cpp-httplib)       | 406   | **Very simple library.** **This library uses 'blocking' socket I/O** |


### Appendix

Benchmark command
```bash
ab -c 1000 -n 100000 -k localhost:8080/hello
```

```bash
# Golang net/http
#############################################

Server Software:        
Server Hostname:        localhost
Server Port:            8090

Document Path:          /hello
Document Length:        12 bytes

Concurrency Level:      1000
Time taken for tests:   1.621 seconds
Complete requests:      100000
Failed requests:        0
Keep-Alive requests:    100000
Total transferred:      15300000 bytes
HTML transferred:       1200000 bytes
Requests per second:    61692.03 [#/sec] (mean)
Time per request:       16.210 [ms] (mean)
Time per request:       0.016 [ms] (mean, across all concurrent requests)
Transfer rate:          9217.66 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   3.1      0      49
Processing:     7   16   0.8     16      49
Waiting:        0   16   0.8     16      19
Total:          7   16   3.3     16      64

Percentage of the requests served within a certain time (ms)
  50%     16
  66%     16
  75%     16
  80%     16
  90%     16
  95%     16
  98%     17
  99%     35
 100%     64 (longest request)
```



```bash
# C++ userver
#############################################
Server Software:        userver/1.0.0
Server Hostname:        localhost
Server Port:            8080

Document Path:          /hello
Document Length:        13 bytes

Concurrency Level:      1000
Time taken for tests:   1.870 seconds
Complete requests:      100000
Failed requests:        0
Keep-Alive requests:    100000
Total transferred:      36600000 bytes
HTML transferred:       1300000 bytes
Requests per second:    53483.03 [#/sec] (mean)
Time per request:       18.698 [ms] (mean)
Time per request:       0.019 [ms] (mean, across all concurrent requests)
Transfer rate:          19116.00 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   3.3      0      55
Processing:     6   18   1.4     18      55
Waiting:        1   18   1.4     18      40
Total:          6   18   4.0     18      91

Percentage of the requests served within a certain time (ms)
  50%     18
  66%     18
  75%     18
  80%     18
  90%     18
  95%     18
  98%     19
  99%     29
 100%     91 (longest request)
```

C++ show enough good results, but for a `simple` http server golang seems is a little bit faster.

```bash
# C++ httplib
#############################################
Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /hello
Document Length:        13 bytes

Concurrency Level:      1000
Time taken for tests:   30.007 seconds
Complete requests:      12197
Failed requests:        0
Keep-Alive requests:    9761
Total transferred:      1291430 bytes
HTML transferred:       158561 bytes
Requests per second:    406.47 [#/sec] (mean)
Time per request:       2460.233 [ms] (mean)
Time per request:       2.460 [ms] (mean, across all concurrent requests)
Transfer rate:          42.03 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    3   9.0      0      52
Processing:     0 1964 4203.1     44   12208
Waiting:        0 1937 4215.3      0   12208
Total:          0 1966 4204.8     44   12208

Percentage of the requests served within a certain time (ms)
  50%     44
  66%     47
  75%     48
  80%     65
  90%  12116
  95%  12147
  98%  12168
  99%  12176
 100%  12208 (longest request)
```
