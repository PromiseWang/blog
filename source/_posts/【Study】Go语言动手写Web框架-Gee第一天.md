---
title: 【Study】Go语言动手写Web框架-Gee第一天
tags:
  - Study
  - Go
abbrlink: 9ac00b9a
date: 2024-01-14 23:04:59
---

# 【Study】Go语言动手写Web框架-Gee第一天

项目来源: [7天用Go动手写/从零实现系列](https://github.com/geektutu/7days-golang)



## 第一天

### 1. net/http的简易demo

#### 代码和一些说明

``` go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))

}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

```

使用Go语言内置的`net/http`库, 简单写个demo

> 第10、11行：设置两个路由，分别是`/`和`/hello`，并且分别绑定了两个函数（第16、20行）`indexHandler`和`helloHandler`
>
> 处理`indexHandler`，返回的是URL；处理`helloHandler`，返回的是请求头内的信息
>
> 第12行：设置端口为9999，表示在9999端口进行监听。第二个参数为处理所有http请求实例，`nil`代表使用标准库中的实例处理。
>
> `http.ListenAndServe`底层源码：
>
> ```go
> // ListenAndServe listens on the TCP network address addr and then calls
> // Serve with handler to handle requests on incoming connections.
> // Accepted connections are configured to enable TCP keep-alives.
> //
> // The handler is typically nil, in which case the DefaultServeMux is used.
> //
> // ListenAndServe always returns a non-nil error.
> 
> func ListenAndServe(addr string, handler Handler) error {
> 	server := &Server{Addr: addr, Handler: handler}
> 	return server.ListenAndServe()
> }
> ```
>
> `Handler`底层源码：
>
> ```go
> // A Handler responds to an HTTP request.
> //
> // ServeHTTP should write reply headers and data to the ResponseWriter
> // and then return. Returning signals that the request is finished; it
> // is not valid to use the ResponseWriter or read from the
> // Request.Body after or concurrently with the completion of the
> // ServeHTTP call.
> //
> // Depending on the HTTP client software, HTTP protocol version, and
> // any intermediaries between the client and the Go server, it may not
> // be possible to read from the Request.Body after writing to the
> // ResponseWriter. Cautious handlers should read the Request.Body
> // first, and then reply.
> //
> // Except for reading the body, handlers should not modify the
> // provided Request.
> //
> // If ServeHTTP panics, the server (the caller of ServeHTTP) assumes
> // that the effect of the panic was isolated to the active request.
> // It recovers the panic, logs a stack trace to the server error log,
> // and either closes the network connection or sends an HTTP/2
> // RST_STREAM, depending on the HTTP protocol. To abort a handler so
> // the client sees an interrupted response but the server doesn't log
> // an error, panic with the value ErrAbortHandler.
> type Handler interface {
>     ServeHTTP(ResponseWriter, *Request)
> }
> ```
>
> `http.ListenAndServe()`的第二个参数类型可以看到是一个接口，需要实现`ServeHTTP`方法，也就是说所有HTTP请求都交给它来处理（第二节实现）
>
> `log.Fatal()`: 进入源码我们可以看到这样的注释`Fatal is equivalent to Print() followed by a call to os.Exit(1).`
>
> 可以得知，当程序报错之后打印日志并会立即退出，defer函数也不会执行，而不同于`panic()`
>
> `panic()`: 
>
> 1. 函数立刻停止执行 (注意是函数本身，不是应用程序停止)
> 2. defer函数被执行
> 3. 返回给调用者(caller)
> 4. 调用者函数假装也收到了一个panic函数，从而
>     4.1 立即停止执行当前函数
>     4.2 它defer函数被执行
>     4.3 返回给它的调用者(caller)
> 5. ...(递归重复上述步骤，直到最上层函数)
>     应用程序停止。

#### 一些运行结果

<img src="..\images\Gee1\url_path.png" style="zoom:50%;" />

<img src="..\images\Gee1\helloHandler.png" style="zoom:50%;" />

当然也可以使用curl工具

<img src="..\images\Gee1\curl.png" style="zoom:50%;" />

### 2. 实现http.Handler接口

需要构造一个对象，并且实现`ServeHTTP`方法

``` go
// 所有请求的统一实现
type Engine struct {}


func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
    engine := new(Engine)
    log.Fatal(http.ListenAndServe(":9999", engine))
}
```



...(未完待续)
