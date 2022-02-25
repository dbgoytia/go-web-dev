# Serving files using a FileServer

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