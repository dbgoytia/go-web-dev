## Exercises

Here a couple of exercises that I encourage you to do on your own. You can go to each of this and see the instructions in the **README.md** file located inside each folder. This are numbered, because for exercise 2, you'll need the answer from exercise 1, and so on.

## Topics covered

* You'll get a deeper understanding of how Webservers work
* [Understand the template type from the html/template package](https://pkg.go.dev/html/template)
* [Understand the use of the multiplexer and the net/http package](https://pkg.go.dev/net/http)
* Question when to use an http.Handle, and when to use an http.HandleFunc



## Implementation notes

Remember this core differences between *http.Handle* and *http.HandleFunc*

Problem: I want to create an object (type) that responds to HTTP requests.

Solution: use http.Handle and the object you create should implement ServeHTTP interface from the HTTP package.

Problem: I want a function to respond to my HTTP request.

Solution: Use http.HandleFunc for that and register your function with the HTTP server.