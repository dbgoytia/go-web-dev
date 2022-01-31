Handler
=======

The Handler type is an *interface* that is the logic to actually "handle" connections to HTTP servers. This interface, of course can be implemented by any other type, making them a value of type **Handler** too.

[http.Handler](https://godoc.org/net/http#Handler)

```Go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

A Handler type is any type's method, that has this signature:

```Go
ServeHTTP(ResponseWriter, *Request)
```


Basically will handle a request, and a client to respond to.

***

Request
===

The request struct:


``` Go
type Request struct {
    Method string
    URL *url.URL
	//	Header = map[string][]string{
	//		"Accept-Encoding": {"gzip, deflate"},
	//		"Accept-Language": {"en-us"},
	//		"Foo": {"Bar", "two"},
	//	}
    Header Header
    Body io.ReadCloser
    ContentLength int64
    Host string
    // This field is only available after ParseForm is called.
    Form url.Values
    // This field is only available after ParseForm is called.
    PostForm url.Values
    MultipartForm *multipart.Form
    // RemoteAddr allows HTTP servers and other software to record
	// the network address that sent the request, usually for
	// logging. 
    RemoteAddr string
}
```


***

