```Go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var m sync.Map
        m.Store("key", "value")
        v, ok := m.Load("key")
        fmt.Println(v, ok)

        // Correct way to check if value is string
        if value, ok := v.(string); ok {
                if value == "value" {
                        fmt.Println("Value is equal")
                }
        }
}
```