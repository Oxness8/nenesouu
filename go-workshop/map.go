package main

import "fmt"

func main() {
    users := make(map[string]int)

    users["iness"] = 19
    users["idk"] = 15
    users["random"] = 21
    
    for name, age := range users {
        fmt.Printf("%s a %d ans.\n", name, age)
    }
}
