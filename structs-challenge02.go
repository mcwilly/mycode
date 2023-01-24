/*
Create a program that creates a struct called Astro. 

It should track the name (string),
  age (int),
  mission (string),
  and if they are needed on the next mission, isneeded (bool).

Create 3 structs, and populate them with fictitious data. 

When you are done, print your records to the screen.


Take the previous challenge, and place your three (3) Astro structures inside of a slice named p.
Display the slice when you're done.
Furthermore, use the slice to print only the string, SpaceX Crew-3
*/


package main

import "fmt"

type Astro struct {
  name     string
  age      int
  mission  string
  isneeded bool
}

func main() {
    myAstro1 := Astro{name:"Bob", age: 48, mission: "Moonbase Alpha", isneeded: true}
    myAstro2 := Astro{name:"Nancy", age: 34, mission: "Mars 123", isneeded: false}
    myAstro3 := Astro{name: "Lou", age: 51, mission: "Homeward", isneeded: true}


    fmt.Println(myAstro1)
    fmt.Println(myAstro2)
    fmt.Println(myAstro3)

    // create slice p
    p := []Astro{myAstro1, myAstro2, myAstro3}

    // now we will display the slice
    fmt.Println(p)

    // now we will display myAstro3's mission name
    fmt.Println(p[2].mission)



}
