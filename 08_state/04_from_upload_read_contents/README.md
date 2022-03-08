# Persisting data to file

Idea is to let the user upload a file, and then write some code to read it and return
the contents of the data in the template. Along the way, we will onderstand a bit about
the multipart.File and multipart.FileHeader types, along with a bit of the io package.


## The multipart.File type
One cool thing to notice in this implementation is the multipart.File that we receive we call FormFile from the net/http package:

```
func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error) {
```

You can see that you get access to:
* the multipart.File
* a pointer to the multipart.FileHeader
* and an error

Multipart requests combine one or more sets of data into a single body,
separated by boundaries. You typically use this requests for file uploads and for transferring data of several types in a single request (for example, a file along with a JSON object). 

If you dig a big deeper into the *multipart.File*, you'll see that it is an implementation of the File interface, that contains this properties:

```
type File interface {
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Closer
}
```

* io.Reader is an interface that wraps the basic Read method. It is used to determine the number of bytes to read and any error encountered

* io.ReadAt is the interface that wraps the basic ReadAt method. It 
is used to determine the number of bytes to read and any error encountered, but starting at offsett off in the underlying input source.

* io.Seeker is the interface that helps us read the file from beginning to end. It sets the offset for the next Read or Write to offset.

* io.Closer interface is the wrapper for the basic Close method
and any specific implementation should be documented since the basic
implementation is not defined.

## The multipart.Fileheader type

It is a struct that describes a file part of a multipart request

```
type FileHeader struct {
	Filename string
	Header   textproto.MIMEHeader
	Size     int64

	content []byte
	tmpfile string
}
```

This is the first time we encounter a *Google Protocol Buffer* which we will cover up later on! Basically the *textproto* package implements generic support for text-based request/response protocols in the style of HTTP, NNTP, and SMTP. 

For example, the multipart.FileHeader of our test file are:

```
header: &{testfile.txt map[Content-Disposition:[form-data; name="q"; filename="testfile.txt"] Content-Type:[text/plain]] 187 [66 101 32 99 114 117 109 98 108 101 100 46 32 83 111 32 119 105 108 100 32 102 108 111 119 101 114 115 32 119 105 108 108 32 99 111 109 101 32 117 112 32 119 104 101 114 101 32 121 111 117 32 97 114 101 46 10 89 111 117 32 104 97 118 101 32 98 101 101 110 32 115 116 111 110 121 32 102 111 114 32 116 111 32 109 97 110 121 32 121 101 97 114 115 46 10 84 114 121 32 115 111 109 101 116 104 105 110 103 32 100 105 102 102 101 114 101 110 116 46 10 83 117 114 114 101 110 100 101 114 46 10 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 45 32 82 117 109 105 46] }
```