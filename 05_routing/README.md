ServeMux routing
======================

The idea of a mux to serve certain paths. For example, you can build RESTful applications
using ServeMux.

Old school, the path would point to a file, in currnet web
development practices, they just point to certain functions.

Remember
========

* There are a lot of synonims for webserver, on of them being multiplexer, mux, webserver, etc.

* A Handler is an interface that knows how to "handle" http requests. That interface has a method named
*ServeHTTP* with two parameters: an *HTTPResponseWriter* interface, and a pointer to a *Request* struct.
In other words, anything that has a method called *ServeHTTP* with that signature is a Handler:

```
ServeHTTP(http.ResponseWriter, req *http.Request)
```

ServeMux architecture
=====================

![Alt text](05_routing/handler_arch.jpg?raw=true "servemux_arch")
