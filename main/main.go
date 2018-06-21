package main

import "fmt"

func mulitples() ([]int,int) {
	var numList []int
	var total int
	for i :=0; i < 1000; i++ {
		if(i % 3  == 0 || i % 5 == 0){
			numList = append(numList,i)
			total += i
		}
	}

	return numList, total
}


func main() {

	numbers, totalNum := mulitples()

	//Print all the numbers
	fmt.Println(numbers)

	//Print total
	fmt.Println("Total is: ", totalNum)

}
