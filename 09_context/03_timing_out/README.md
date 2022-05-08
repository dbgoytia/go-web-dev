# Timing out

One of the most useful things about context is the capability to timeout functions so we avoid
leaking systems in Golang.

With this example, we introduce a ridiculosuly running task that exceeds the given deadline, 
causing the context to exceed the deadline and stop all processes:

```
❯ curl localhost:8080
context deadline exceeded
```