# Persisting data to file

Idea is to let the user upload a file, and then write some code to save it on the server.

# Note

Please read the notes on the previous lesson (form uplaod read contents) since it will contain
most of the same types for this implementation. The difference will be on how we're handling
the uploaded asset on the server side.

# Implementing a storing mechanism

Basically the magic here is this:

```
// store on server
dst, err := os.Create(filepath.Join("./user/", h.Filename))
HandleError(w, err)
defer dst.Close()

_, err = dst.Write(bs)
HandleError(w, err)
```

All we're doing is creating a new file using *os.Create*, and taking the filename from the
headers multipart.FileHeader. We're able to do dst.Write because the *os.Write* method basically
receives the array of bytes read by the *ioutil.ReadAll(f)* call.

```
func (f *File) Write(b []byte) (n int, err error) {
```