package main

import "fmt"

func main(){
	var num int
	fmt.Print("Masukan bilangan: ")
	fmt.Scanln(&num)

	if isMod7(num){
		fmt.Printf("%d adalah kelipatan 7. \n" , num)
	}else{
		fmt.Printf("%d Bukan kelipatan 7. \n", num)
	}
}

func isMod7(num int) bool{
	if num % 7 == 0 {
		return true
	}
	return false
}