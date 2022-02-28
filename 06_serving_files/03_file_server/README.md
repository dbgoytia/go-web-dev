# Serving files using a FileServer

The FileServer function takes a root FileSystem to convert the FilseSystem to a Handler, so we're able to use it inside our *Mux*.
It takes the root FileSystem to know where it will start serving files from.

```
func FileServer(root FileSystem) Handler
```

## The FileSystem interface

First, let's examine the [FileSystem interface]((https://pkg.go.dev/net/http#FileSystem)):

```
type FileSystem interface {
	Open(name string) (File, error)
}
```

The implementation inside the net/http package is the [type http.Dir](https://pkg.go.dev/net/http#Dir), simply because
the type Dir implements all the methods of the FileSystem interface.

```
 type Dir
    func (d Dir) Open(name string) (File, error)
```

Which means that we can use the http.Dir type to give *FileServer* a root FileSystem to serve files from. 

# Conclussion

This I think is the best approach. We can implement an *http.FileServer* to serve files on a specific directory. This will actually pick up all the files in the directory and be able to serve them.

This is opposed to the first scenario were we did a demo of the io.WriteString trying to read a file in the server.

So actually if you go to the root (http://localhost:8080), you'll be able to see all of the files in this exercise, as well as the image bien served on any browser!

```
curl localhost:8080
<pre>
<a href="README.md">README.md</a>
<a href="main.go">main.go</a>
<a href="toby.png">toby.png</a>
</pre>
```