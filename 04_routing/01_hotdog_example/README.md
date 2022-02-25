# Most basic idea

I think the first idea we can get after creating a *Mux*, would be to 
use some sort of case based selection on each of the respones we would like to
give to our users. Let me walk you through this example. But first some recap of what we know so far.

# Recap 

## The Handler interface

Remember that the Handler interface allows us to respond to an HTTP request.

```
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

## The ServeHTTP method

We know that the ServeHTTP for the type *ServeMux* implements the Handler interace

```
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
```

## The default server

We know that the ListenAndServe method will listen on the TCP network address and call *Serve* to handle all requests to incoming connections. By default, this will all get assigned to the default server.

```
func (srv *Server) ListenAndServe() error
```

## The io.WriteString method for experimenting

Because we know that teh io.WriteString is expecting to write to a type that implemetns the *Writer* interface, and a string, we should be able to respond some arbitrary stuff to it just to test things out.

```
func WriteString(w Writer, s string) (n int, err error)
```

# Case based solution

We might be tempted to  write something like this to serve multipe routes
inside our server. This is not a great implementation, but it works.

```
func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(res, "doggy doggy doggy")
	case "/cat":
		io.WriteString(res, "kitty kitty kitty")
	}
}
```

## Where does this fall short?

* Your ServeHTTP method is now not performing a single task. This means that each time you want to update it to do something slightly different, you'll have to update them both.
* The route in which this served is controlled on the Handler, and not on the server.
* You're going to be unable to add this handles independently in multiple routers or on custom Multiplexers.



