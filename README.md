# Echo router builder

A builder design pattern for the echo framework router. This lib made mvc projects easier to create and simpler to maintain. See the example bellow:

### Example

controller file: ``controllers/access.go``:

```go

package controllers

import (
  "github.com/labstack/echo/v4"
  r "github.com/Mth-Ryan/echo_router_builder"
)

AccessController = r.NewController("/access").
  View("/login", "login.tmpl.html", nil).
  Post("/login", loginAction))
  
func loginAction(c echo.Context) error {
  ...
}
```

main file: ``main.go``

```go
  package main
  
  import (
    m "github.com/labstack/echo/v4/middleware"
    r "github.com/Mth-Ryan/echo_router_builder"
    "example.com/project/controllers"
  )
  
  builder := r.NewBuilder(m.Logger(), m.Recovery()).
    RegisterStatic("/public", "assets").
    RegisterViews("views", ".tmpl.html").
    Register(AccessController)
    
  e := builder.Build
  e.Logger.Fatal(e.Start(":8080"))
```
