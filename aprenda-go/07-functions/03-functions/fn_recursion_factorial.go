package main


func main(){

	println(factorial(7))
}

func factorial(number int) int {

	if number == 1{
		return number
	}
	
	return number * factorial(number - 1)
}