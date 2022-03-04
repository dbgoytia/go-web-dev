# State

State of a compute program is a term for all stored information, at a given instant, that a program
has access to (out of Wikipedia). The idea is basically 'who is this user' is the question we want
to be able to know in our servers.

There's a couple of things we first need to understand previous to actually implementing cookies.

First of them:
* How are we going to pass that data to the server?
* How are we going to redirect to the appropiate server?
* What cookies are and how we can manage them?

Stay tuned for this labs demonstrating this implementation.

## Passing data

Basically we have two ways of passing information, either through the URL with a Get method.
Or we can POST the data through a form.

```
post --> body
get --> url
```