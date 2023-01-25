/* Alta3 Research | RZFeeser
try writing a program that iterates (loops) over the arguments passed in via the command line. Display them on the screen with some indication of what position they correspond to.
*/

package main
 
import (
    "fmt"
    "os"
)
 
func main() {

    // use the len() function to return length
    argLength := len(os.Args[1:])
    
    // use fmt.Printf() to format string
    fmt.Printf("There are %d arguments.", argLength) 

    for i, a := range os.Args[1:] {
        fmt.Printf("Argument # %d is %s\n", i+1, a)

    }
}

