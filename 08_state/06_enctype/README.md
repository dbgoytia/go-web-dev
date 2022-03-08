# Enctype

When you make a POST request, you have to encode the data that forms the
body of hte request in some way. HTML forms provide three methods of encoding:

* application/x-www-form-urlencoded (default)
* multipart/form-data
* text/plain

Although, the specifics of the formats don't matter to most developers. The
important points are: When you are writing client-side code, all you need to
know is use multipart/form-data when your forms include any of the input file
elements:

HTML:
```
<input type="file">
```
