```Go
func main() {
  var m sync.Map
  m.Store("key", "value")
  v, ok := m.Load("key")
  fmt.Println(v, ok)
  // Incorrect way to check if value is string
  if v == "value" {
    fmt.Println("Value is equal")
  }
}
```