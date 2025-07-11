// cmd/vuecryptoengineframework/main.go
package main

import (
"flag"
"log"
"os"

"vuecryptoengineframework/internal/vuecryptoengineframework"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := vuecryptoengineframework.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
