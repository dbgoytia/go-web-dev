# Serving files using http.ServeContent and http.ServeFile

Please visit the previous lab to understand the basics of why this works inside the "serving files" section.

## Using http.ServeContent

The key takeway in our code is that we can use ServeContent method to have a little bit of more control on the actual content that's being served.

```
// Get the file Stat information
fi, err := f.Stat()
if err != nil {
    http.Error(w, "file not found", 404)
    return
}

http.ServeContent(w, req, fi.Name(), fi.ModTime(), f)
```

This specially useful when you want to control cache operations, for example using the [etag](https://developer.mozilla.org/es/docs/Web/HTTP/Headers/ETag) header. This header controls the specific version of a resource, and lets caches be more efficient and save bandwith.

## Using http.ServeFile

This implementation is actually a lot simpler, it simply asks for the Writer, a Request, and a file to serve when hitting that so the Handler can return it to the client.