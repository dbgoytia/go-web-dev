# Creating a Static Website server

What we've learned so far is that http.FileServer is able to serve files from a particular root FileSystem. One interesting thing that happens, is that when you serve an index.html to it, you'll automatically start serving that static website using your Mux!