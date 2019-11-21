# numrange: range parsing library
[![Build Status](https://travis-ci.com/eest/numrange.svg?branch=master)](https://travis-ci.com/eest/numrange)
[![codecov](https://codecov.io/gh/eest/numrange/branch/master/graph/badge.svg)](https://codecov.io/gh/eest/numrange)
[![Go Report Card](https://goreportcard.com/badge/github.com/eest/numrange)](https://goreportcard.com/report/github.com/eest/numrange)

A simple library for taking a string such as "-10..5,7,10" and getting a
set where you can test for inclusion and also add/delete more ranges

Inspired by the [Number::Range](https://metacpan.org/pod/Number::Range) perl module.

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
inRange(40) false
-10..5,7,10,20,30..33,35
```
