package dotenv-mustache

import (
   "github.com/joho/godotenv"
   "github.com/alexkappa/mustache"
   "os"
)

var ctx map[string]string
var template := mustache.New()

ctx, err := godotenv.Read()

if err != nil {
  fmt.Fprintf(os.Stderr, "failed to open file: %s\n", err)
}

err := template.Parse(os.Stdin)

if err != nil {
  fmt.Fprintf(os.Stderr, "failed to parse template: %s\n", err)
}

template.Render(os.Stdout, ctx)

