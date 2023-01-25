/* RZFeeser | Alta3 Research
   switch - case and default  */
/*
Rewrite example, ~/switch01.go so that any minor version would match on our cases (for example, go1.17.3 should match what currently only exactly matches case "go.1.17"). Hint: Check the strings package.
*/



package main

import (
    "fmt"
    "runtime"
    "strings"
)

func main() {

    // init gover                
    var gover string = runtime.Version()
    switch {
    case strings.Contains(gover, "go1.19"):  // if contains "go1.19", do the code below and then STOP
        fmt.Println("Go Ver 1.19.x - Released in 2022")
    case strings.Contains(gover, "go1.16"), strings.Contains(gover, "go1.17"), strings.Contains(gover, "go1.18"):                 // if contains "go1.16, 17, 18", do the code below then STOP
        fmt.Println("You are using an older version")
        default:                       // in all other cases run the code below
        fmt.Println("I cannot make a recommendation.")
    }
}

