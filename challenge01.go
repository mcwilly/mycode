/* Alta3 Research | RZFeeser
   Methods - defining         */

// Go program to illustrate the
// method with struct type receiver
package main
 
import "fmt"
 
// Author structure
type author struct {
    name      string
    branch    string
    books     int
    salary    int
}
 

// receiver function to increase value of books by 1
//create a method called bookup() that increases the value of author.books stored within a struct by 1.
func (a *author) bookup() {
    a.books ++

} 


// Method with a receiver
// of author type
func (a author) show() {
 
    fmt.Println("Author's Name: ", a.name)
    fmt.Println("Branch Name: ", a.branch)
    fmt.Println("Published articles: ", a.books)
    fmt.Println("Salary: ", a.salary)
}
 
// Main function
func main() {
 
    // Initializing the values
    // of the author structure
    res := author{
        name:      "RZFeeser",
        branch:    "Pennsylvania",
        books:     14,
        salary:    34000,
    }
 
    // Calling the method
    res.show()

    // run bookup
    //we should go from 14 to 15
    res.bokkup()

    // let's look at the new value
    res.show()


}


