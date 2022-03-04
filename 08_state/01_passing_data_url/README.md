# Passing data through the URL

For the URL remember that the anatomy of a URl is basically conformed of:

* the protocol
* sudbomain
* domain
* top level domain (TLD)
* path
* a query string

## The query string

A Query string is basically a 'key=value' store that we can send using regular GET methods,
and use in our backend retrieving the value out of that query string.

Example:

```
domain.com/shoes?type=sneakers&sort=price_ascending
```

The query string is asking for shoes of type sneaker, sorted by ascending price. Those options
of course are completely tied to the backend implementation of such methods.