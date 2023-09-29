# benchmarks
Benchmarks between defferent packages or languages 

## http `hello world` Golang and C++ userver
Golang code:
```go
package main

import (
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8090", nil)
}

```

C++ userver [code](https://github.com/userver-framework/userver/blob/develop/samples/hello_service/hello_service.cpp)
```C++
#include <userver/utest/using_namespace_userver.hpp>

/// [Hello service sample - component]
#include <userver/components/minimal_server_component_list.hpp>
#include <userver/server/handlers/http_handler_base.hpp>
#include <userver/utils/daemon_run.hpp>

namespace samples::hello {

class Hello final : public server::handlers::HttpHandlerBase {
 public:
  // `kName` is used as the component name in static config
  static constexpr std::string_view kName = "handler-hello-sample";

  // Component is valid after construction and is able to accept requests
  using HttpHandlerBase::HttpHandlerBase;

  std::string HandleRequestThrow(
      const server::http::HttpRequest&,
      server::request::RequestContext&) const override {
    return "Hello world!\n";
  }
};

}  // namespace samples::hello
/// [Hello service sample - component]

/// [Hello service sample - main]
int main(int argc, char* argv[]) {
  const auto component_list =
      components::MinimalServerComponentList().Append<samples::hello::Hello>();
  return utils::DaemonMain(argc, argv, component_list);
}
/// [Hello service sample - main]
```

benchmark command
```bash
ab -c 1000 -n 100000 -k localhost:8080/hello
```

### Results

```
Golang
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

#############################################
#############################################

C++ userver
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
