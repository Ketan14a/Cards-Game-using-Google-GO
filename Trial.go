package main

import "fmt"

func main() {

	Myslice := []int{1,2,3,4,5}

	Part := Myslice[0:3]
	Part2 := Myslice[2:5]
	Front := Myslice[3:]
	Rear := Myslice[:3]

	fmt.Println(Part)
	fmt.Println(Part2)
	fmt.Println(Front)
	fmt.Println(Rear)

}