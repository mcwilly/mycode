package main

import (
    "fmt"
)

func main () {
    
    // let's create our variables and set them
    const uri = "https://example.org:6001/v2/snacks?"

    var r = "req=snickers"
    var q = "quantity=1"
    var s = "size=king"

    // now let's create the string
    response := fmt.Sprintf("%s%s%s%s", uri, r, q, s)
    fmt.Println(response)
}
