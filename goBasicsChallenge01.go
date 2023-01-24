/* Mike's first challenge
   using the fmt.Println() function
   print your name to the screen
   and a second with your City and State.
   Be sure your code executes when you run it with go run
   and it builds a binary when you execute go build.   */

package main
  
import "fmt"

var name, city, state string     


func main() {
        
        var name, city, state = "Mike", "Savannah", "Texas"

        fmt.Println("My name is: ", name)                    // "The value of the var rock is: 42"
        fmt.Println("I live in ", city + ", ", state)            // "The value of the var scissors is: 55"
}

