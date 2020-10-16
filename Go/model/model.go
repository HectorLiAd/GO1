package model

import "fmt"

type Person struct {
    gorm.Model
    Id int
    Name string
    Age  int
}


/*func main() {
    
    
}*/

    /*s := Person{}
    s.Name = "Sean"
    s.Age = 42

    fmt.Println(s.Name)
    fmt.Println(s.Age)

    sp := &s
    fmt.Println(sp.Age)

    sp.Age = 51
    fmt.Println(s.Age)*/