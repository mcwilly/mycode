/*
Create a new struct, of type nasaMission.
To create this type, include people mapped to []Astro
   number mapped to int
   and message mapped to string

Once you've created this, create a nasaMission struct named s.
  Initialize people as p
  init as 3
  and message as success

Display s (with and without its key fields) to end your program.
*/

package main

import "fmt"

type Astro struct {
    name     string
    age      int
    mission  string
    isneeded bool
}

// this is our new struct
type nasaMission struct {
        people []Astro
        number int
        message string
}

func main() {
    
    myAstro1 := Astro{name:"Bob", age: 48, mission: "Moonbase Alpha", isneeded: true}
    myAstro2 := Astro{name:"Nancy", age: 34, mission: "Mars 123", isneeded: false}
    myAstro3 := Astro{name: "Lou", age: 51, mission: "Homeward", isneeded: true}


    p := []Astro{myAstro1, myAstro2, myAstro3}

    fmt.Println(p)

    // select data from p
    fmt.Println(p[1].mission)

    // initialize nassMission as s
    s := nasaMission{p, 3, "success"}

    // display s with fields
    fmt.Println(s)

    // display s withOUT fields
    fmt.Printf("%+v", s)

}
