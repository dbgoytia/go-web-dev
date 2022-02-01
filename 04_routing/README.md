ServeMux routing
======================

The idea of a mux to serve certain paths. For example, you can build RESTful applications
using ServeMux.

Old school, the path would point to a file, in currnet web
development practices, they just point to certain functions.

Remember
========

There are a lot of synonims for webserver, on of them being multiplexer, mux, webserver, servemux, etc.


HandlerInterface
===================

```
ServeHTTP(http.ResponseWriter, req *http.Request)
```

A Handler is an interface that knows how to "handle" http requests. 
It has a method called *ServeHTTP*, that receives an *HTTPReponseWriter* and a pointer to a *Request* struct.
In other words, anything that implements a method called *ServeHTTP* with that signature, is a Handler.


ServeMux architecture
=====================

Review handler arch for important information.


ServeMux
==============
For creating a new ServeMux, we can use the *NewServeMux* method, that returns a pointer to a ServeMux.
```
func NewServeMux() *ServeMux
```
The ServeMux implements the ServeHTTP method, meaning that ServeMux implements the Handler interface.
The handler is typically nil, in which the DefaultServeMux is used.
Because *ListenAndServe* receives a Handler, we can then pass in the *ServeMux* to ListenAndServe.


HandleFunc
==================

```
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

Basically takes a handler, and a pattern to attach it to the DefaultServeMux.


ListenAndServe
====================

```
func ListenAndServe(addr string, handler Handler) error
```

Listen and Serve takes a Handler as an argument.
Listens on the TCP network address and then calls Serve with handler to handle the incoming connections.
The handler is typically nil, in which the DefaultServeMux is used.


DefaultServeMux
================

```
var DefaultServeMux = &defaultServeMux
```

It is the default ServeMux used by Serve.


Disambiguation between func(ResponseWriter, *Request) vs HnalderFunc
======================================================================
