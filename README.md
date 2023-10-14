# Go Modules Sample

This tutorial we will learn how to use go modules and structure project as multiple packages.

What will you learn?

- Structure go project as a module outside GOPATH use go.mod
- 

## Question & Answer

1. Does project name must be defined in go.mod file? - YES
2. Does project directory have to have the same name with module name in go.mod file?
    - No, we can create project root directory name difference with module name in go.mod file


## Project Structure

```
myproject/
├── cmd/
│   └── main.go
├── pkg/
│   ├── api/
│   │   ├── api.go
│   │   └── handlers.go
│   ├── db/
│   │   ├── db.go
│   │   └── models.go
│   └── utils/
│       └── utils.go
└── go.mod
```

main.go

```go
package main

import (
    "fmt"
    "myproject/pkg/api"
    "myproject/pkg/db"
    "myproject/pkg/utils"
)

func main() {
    fmt.Println("Starting server...")
    api.HandleRequests()
    db.Connect()
    utils.Log("Server started successfully!")
}
```

api.go

```go
package api

import (
    "fmt"
    "net/http"
)

func HandleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/about", aboutPage)
    http.HandleFunc("/contact", contactPage)
    http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the home page!")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is the about page.")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Contact us at contact@example.com")
}
```

**Run application**

go run cmd/main.go

### Resolve Issues

1/ Issues 1

> package myproject/pkg/api is not in std (/usr/local/go/src/myproject/pkg/api)

That mean go compiler can not find you module named `myproject`. 
Make sure the module `myproject` is defined in go.mod for example:

**go.mod**

```go
module myproject

go 1.21.3
```# techlead
# techlead
