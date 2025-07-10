# envreader

A minimal Go package for reading simple `.env` files and setting key-value pairs as environment variables at runtime.

- `.env` file MUST follow the yaml format

**ex.**  
```
# this is a comment
API_KEY=abc123
DEBUG=true
```

## Usage
import package name and call Load() function

**ex.**
```
import "github.com/harrison-blake/envreader"

func main() {
    err := envreader.Load("./config/.env")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(os.Getenv("MY_KEY"))
}
```

`stout: some_key123987324`
