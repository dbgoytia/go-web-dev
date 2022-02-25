# Serving files using io.Copy

In the previous sessions we learned how to serve templates using the text/template and the html/template modules in Go. In this lesson, we're going to see how to serve files to clients using the io module (io.Copy).

# First approach

Based on what we've done previously in the past, it could seem that we could use the io.WriteString to return the file, using something like this:

```
func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<!--image doesn't serve -->
	<img src="/toby.jpg">
	`)
}
```

## Why does this doesn't work?

I think the issue here is that in spite we're writing the correct html code for opening up the file, we're actually looking it in the client side, and we should open it up on the server side, so the return of this code is a broken link on the client.

## Why did I think this could be a good approach?

Because we know the that the *io.WriteString* method takes in an object that implements the writer interface:

```
func WriteString(w Writer, s string) (n int, err error)
```

We know that the *http.ResponseWriter* implements the Writer interface. We know this because the *http.ResponseWriter* interface has a Write method that accepts a *[]bytes(str)*, so we should be able to write to that *Writer* (I know..  bare with me) using the io.WriteString method.

```

```

# Correct implementation

## The File type

There's two particular things I'd like to focus on the File type inside the *os* module to further explain.

By looking at the available methods, I can see that we're able to do two major things with Files, reading them and Writing to them.

---

* Reading:

```
func (f *File) Read(b []byte) (n int, err error)
```


It works with almost all markdown flavours (the below blank line matters).


* Writing:
```
func (f *File) Write(b []byte) (n int, err error)
```

---

This is particularly interesting because if we see the signature of those functions, we can understand that the Read method implements the Reader interface:

```
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

And the Writer method implements the Writer interface:

```
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

So this simply means that we're either able to *Read* or *Write* to files, which of course makes a lot of sense.

## The http.ResponseWriter type

We know that the ResponseWriter also implements the Writer interface, meaning, we can send an arbitrary array of bytes into it for writing

```
type ResponseWriter interface {
	Header() Header
	Write([]byte) (int, error)
	WriteHeader(statusCode int)
}
```

So in fact, the *http.ResponseWriter* type implements the *Writer* interface.

## Implementation

Ok then, based on the previous explanation we're able to *Read* and *Write* to files. Then comes the next piece of the puzzle. This would be how do we get the *File, that implements the *Read* interface, written as a reponse?

This is where the *io.Copy* function comes in. Based on the documentation, I can see that the Copy function, takes an object that implements the *Writer* interface, and an Object that implements the *Reader* interface.

```
func Copy(dst Writer, src Reader) (written int64, err error)
```

Therefore we should be able to write our read object, to our *http.ResponseWriter* object, using our **os.File* object.


```
func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("puppy.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	io.Copy(w, f)
}
```

Then you can tie this handler to your *Mux*.

--- 
I wonder if this is the best implementation?
--- 