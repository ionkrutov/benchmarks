//
//  hello.cc
//
//  Copyright (c) 2019 Yuji Hirose. All rights reserved.
//  MIT License
//

#include <httplib.h>

using namespace httplib;

int main(void) {
  Server svr;

  svr.Get("/hello", [](const Request & /*req*/, Response &res) {
    res.set_content("Hello World!\n", "text/plain");
  });

  svr.listen("127.0.0.1", 8080);
}