# Exercise 3
* Serve files in the "starting-files" folder

## Requirements

* To ser your images, use only:

```
fs := http.FileServer(http.Dir("public"))
```

Hint: look to see what type FileServer returns, then think it through.