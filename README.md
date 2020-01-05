# magro

Go Tracing library parody of C MACRO `__LINE__`, `__FUNCTION__`, `__FILE__`.

## Example

``` golang
import (
    ...
    m "github.com/Napat/magro"
)

func example() {
    fmt.Printf("File fullpath: %s\n", m.FileFullpath())
    fmt.Printf("File: %s\n", m.File())
    fmt.Printf("Function: %s\n", m.Function())
    fmt.Printf("Line: %d\n", m.Line())
    fmt.Printf("Where: %s\n", m.Where())
    fmt.Printf("Info: %s\n", m.Info())
}
```
