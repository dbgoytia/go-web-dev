
# Cleaning up our FileServer


So far, our FileServer implementation basically will serve anything under the 
*root FileSystem* that we give to the *http.FileServer* function. This is not desirable because most likely, things like our go code are also hosted at the root.
So how do we order everything up?

# The http.StripPrefix function

So far, we know that in order to be able to *Handle* things inside a *Mux*, we have to give it a *Handler*. 

```
func Handle(pattern string, handler Handler) { DefaultServeMux.Handle(pattern, handler) }
```

And we know that the Handler interface is responsible for letting us create new handlers. Let's take this opportunity to recall the  Handler interface

```
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

So if we take a look at the *http.StripPrefix* function, we see that in fact it returns a Handler that we can use to give it to our Handle function:

```
func StripPrefix(prefix string, h Handler) Handler
```

## Implementation

So the idea is to have a folder, in which we can store assets in a sepparate location than those of the rest of our FileServer actual logic. For this I've created an *assets* folder, and we're going to be serving it under the prefix "/resources/".

```
http.StripPrefix(
    "/resources", 
    http.FileServer(http.Dir("./assets")))
```

If you can see, basically what you're saying is serve files from "./assets", and strip the "/resources" name in the prefix. So that when we're looking for a resource like toby.png:

```
<img src=/resources/toby.png>
```

our actual code is going to be looking for something like:

```
<img src=/assets/toby.png>
```

But you'll be consuming it through localhost:8080/resources/toby.png)


## Conclusion

I think this is going to be particularly useful once we start implementing templates
into our FileServer, and adding more logic to it so we can strip all of those up and just come up with a single file location.