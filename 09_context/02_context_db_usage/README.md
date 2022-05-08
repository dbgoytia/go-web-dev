# Modyfing Context

As we have seen in the previous example, we can modify the Context with the idea that it should **only** hold data that is direclty connected to a particular request. For example, session Id, user Id or similar values are good examples. Every per-request variable could be a good example.

In this case, we modified the context using a new value, "userID" that then we can share across the entire system. The new context looks like this:

```
&{0xc000110aa0 0xc000156200 {} 0x107ea20 false false false false 0 {0 0} 0xc00002a380 {0xc0001402a0 map[] false false} map[] false 0 -1 0 false false [] 0 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] [0 0 0] 0xc0000261c0 0}
    context.Background
        .WithValue(type *http.contextKey, val <not Stringer>)
        .WithValue(type *http.contextKey, val [::1]:8080)
        .WithCancel
        .WithCancel
        .WithValue(type string, val <not Stringer>)
        .WithValue(type string, val Bond)
```


As you can see it's storing a value that can't be printed as String (userID is an integer), and Bond that is the fname and can actually be printed as string.P