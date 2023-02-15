package main

import "fmt"

func main() {
    arr := [3]int{1, 2, 3}
    slice := arr[:]

    fmt.Println("Tab :", arr)
    fmt.Println("Slice :", slice)
}
