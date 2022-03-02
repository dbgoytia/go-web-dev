# Exercise 3
* Serve files in the "starting-files" folder

## Requirements

* Parse Glob in an init function
* Use HandleFunc for each of the routes
* Combine apply & applyProcess into one func called "apply", inside this function use the following code
to create logic to respond differently to a POST and a GET method

```
if req.Method == http.MethodPost {
    // code here
    return
}
```
