# Exercise 3
* Serve files in the "starting-files" folder

## Requirements

* To ser your images, use only:

```
func StripPrefix(prefix string, h Handler) Handler
func FileServer(root FileSystem) Handler
```

Constraint: you are not allowed to change the route being used for images in the template file