A simple library for taking a string such as "-10..5,7,10" and getting a
set where you can test for inclusion and also add/delete more ranges

Example usage:

```
package main

import (
        "fmt"
        "github.com/eest/numrange"
        "log"
)

func main() {

        is, err := numrange.ParseIntSet("-10..5,7,10")
        if err != nil {
                log.Fatal(err)
        }

        err = is.Add("20,30..35")
        if err != nil {
                log.Fatal(err)
        }

        err = is.Del("34")
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println("inRange(25)", is.InRange(25))
        fmt.Println("inRange(40)", is.InRange(40))

        fmt.Println(is)
}
```

Output:
```
inRange(25) true
-10..5,7,10,20,30..33,35
```
