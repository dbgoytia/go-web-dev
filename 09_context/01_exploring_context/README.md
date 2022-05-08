# Context package

Let's review what the godoc says:

`
Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context. The chain of function calls between them must propagate the Context, optionally replacing it with a derived Context created using WithCancel, WithDeadline, WithTimeout, or WithValue. When a Context is canceled, all Contexts derived from it are also canceled.
`

What this is saying is that we can set various limitations and boundareis
to our APIs by using a context. Your server is actually asking: "What is the context of what I'm doing here? And this context might contain various session variables, for example, your session id.

Limitations are set to a Context using WithDeadline, WithTimeout and WithValue, with which you can set healthy deadlines for incoming requests, and increase security around your APIs. For example:
* You have 5 seconds to process this request, otherwise we drop the connection.

Another very useful thing is that this Context is shared accross all spawned subprocesses. But don't fall in the trap of using it as a place to share variables around, that's not the intended way to use a Context. A healthy usage might be to keep for example track of a session within 
your system. Think of them as **configuration**.


# What's available right out of the gate?

If we print the current context, we would see something like this:
``` 
context.Background.WithValue(type *http.contextKey, val <not Stringer>)
    .WithValue(type *http.contextKey, val 127.0.0.1:8080)
    .WithCancel
    .WithCancel
```

I think this is the first time we encounter a **Stringer** in Go. A stringer is simply a type that can describe itself as a string in Go.

The Stringer interface looks like this:

```
type Stringer interface {
    String() string
}
```

Quick mini-example:
```
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
```

In this case, the String() method implements the Stringer interface.

Context also seems to be storing a reference to the Context key that the Server received, along with the configuration values.

```
.WithValue(type *http.contextKey, val 127.0.0.1:8080)
```

# Useful links:
* godoc: https://pkg.go.dev/context
