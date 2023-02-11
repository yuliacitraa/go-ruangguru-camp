package main

import "fmt"

type Product struct {
	Name  string
	Price int
	Tax   int
}


// func (p *Product) getPrice(price ...int)  {

// }

//  func (p *Product) setTax(tax int) {
// 	p.Tax = tax
//  }


func MoneyChanges(amount int, products []Product) []int {
	subTotal := 0
	for _, v := range products {
		subTotal += v.Price + v.Tax
	}

	kembalian := amount - subTotal

	res := ExchangeCoin(kembalian)

	return res // TODO: replace this
}

func ExchangeCoin(amount int) []int {

	ribuan := 1000
	ratusan := [3]int{500,200,100}
	puluhan := [3]int{50,20,10}
	satuan := [2]int{5,1}

	res := []int{}

	for  {
		if amount == 0 {
			break
		}

		switch {
		case amount >= ribuan:
			res = append(res, ribuan)
			amount -= ribuan
		case amount < ribuan && amount >= ratusan[0]:
			res = append(res, ratusan[0])
			amount -= ratusan[0]
		case amount < ratusan[0] && amount >= ratusan[1]:
			res = append(res, ratusan[1])
			amount -= ratusan[1]
		case amount < ratusan[1] && amount >= ratusan[2]:
			res = append(res, ratusan[2])
			amount -= ratusan[2]
		case amount < ratusan[2] && amount >= puluhan[0]:
			res = append(res, puluhan[0])
			amount -= puluhan[0]
		case amount < puluhan[0] && amount >= puluhan[1]:
			res = append(res, puluhan[1])
			amount -= puluhan[1]
		case amount < ratusan[1] && amount >= puluhan[2]:
			res = append(res, puluhan[2])
			amount -= puluhan[2]
		case amount < puluhan[2] && amount >= satuan[0]:
			res = append(res, satuan[0])
			amount-= satuan[0]
		case amount < satuan[0] && amount >= satuan[1]:
			res = append(res, satuan[1])
			amount -= satuan[1]
		}
	}

	return res // TODO: replace this
}

func main() {
	instance := []Product{
		{
			Name:  "baju",
			Price: 5000,
			Tax:   500,
		},
		{
			Name:  "celana",
			Price: 3000,
			Tax:   300,
		},
	}

	res:=MoneyChanges(10000, instance)

	fmt.Println(res)
}
