package main

import (
    // "errors"
    "fmt"
    // "golang-demo/test"
)

func main() {

    x := 42
    i := &x
    fmt.Printf("x: %d, type of x: %T\n", x, x)
    fmt.Printf("i: %d, type of i: %T\n", i, i)
    fmt.Printf("*i: %d, type of *i: %T\n", *i, *i)


    // a, b := 10, 2
    // result, err := test.Divide(a, b)
    // if err != nil {
    //     var divErr *test.DivisionError
    //     switch {
    //     case errors.As(err, &divErr):
    //         fmt.Printf("%d / %d is not mathematically valid: %s\n",
    //           divErr.IntA, divErr.IntB, divErr.Error())
    //     default:
    //         fmt.Printf("unexpected division error: %s\n", err)
    //     }
    //     return
    // }

    // fmt.Printf("%d / %d = %d\n", a, b, result)
}