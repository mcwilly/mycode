/*
Map several programming languages to their file extension,
the key Python should map to the value .py.
The key Golang should map to .go,
the key Java should map to .java,
the key Ansible should map to .yml,
the key C++ should map to .cpp. 

Once the map is constructed, display it on the screen. 
Then remove the language C++,
and add Julia and the key .jl. 

Display the new map after the change.
*/

package main


import "fmt"

func main () {
    // build our map
    var myMap = map[string]string{
        "Python": ".py",
        "Golang": ".go",
        "Java": ".java",
        "Ansible": ".yml",
        "C++": ".cpp",
    }

    // display myMap to screen
    fmt.Println(myMap)

    // remove c++ lanaguage
    delete(myMap, "C++")

    // add Julia
    myMap["Julia"] = ".jl"

    // display myMap to scrteen
    fmt.Println(myMap)

}
