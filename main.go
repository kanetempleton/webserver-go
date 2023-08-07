// main.go
package main

import (
    "fmt"
    "user/user"
)

func main() {
    u := user.NewUser(1, "john_doe", "john@example.com")
    fmt.Println(u)
}