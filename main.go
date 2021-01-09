package main

import "fmt"

// factory method

type PayMethod interface {
	Pay()
}

type Paypal struct{}

func (p Paypal) Pay() {
	fmt.Println("Pagado por paypal")
}

type Cash struct {}

func (c Cash) Pay() {
	fmt.Println("Pagado por efectivo")
}

type CreditCard struct {}

func (c CreditCard) Pay() {
	fmt.Println("Pagado por tarjeta de credito")
}

func Factory(method uint) PayMethod {
	switch method {
	case 1:
		return Paypal{}
	case 2:
		return Cash{}
	case 3:
		return CreditCard{}
	default:
		return nil
	}
}

func main() {
	var method uint
	fmt.Println("Digite el numero de metodo")
	fmt.Println(("\t 1:Paypal 2:Efectivo 3:Tarjeta de credito"))
	_, err := fmt.Scanln(&method) // Scanln manda dos parametros el numero y el error
	if err != nil {
		panic("debe digitar un metodo valido")
	}

	if method > 3 {
		panic("debe digitar un metodo valido")
	}

	payMethod := Factory(method)
	payMethod.Pay()
}
