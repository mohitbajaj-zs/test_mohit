package Calculator

import (
	"errors"
	"fmt"
)

func calculator(a, b float64, ch string) (float64,error) {
	return 0.0,nil
}
func add(a, b float64) float64{
	return a+b
}
func subtract(a, b float64)float64{

		return a-b

}

func multiply(a, b float64)float64{

	return a*b
}

func division(a, b float64)(float64,error){

	if b!=0{
	return a/b,nil
}
	return 0,errors.New("divide by zero error")

}

func modulo(a, b int)(int,error){

	if b!=0{
		return a%b,nil
	}
	return 0,errors.New("divide by zero error")

}

func main(){
	a:=5
	b:=6.0
	ch:="+"
	fmt.Println(calculator(a,b,ch))
}
