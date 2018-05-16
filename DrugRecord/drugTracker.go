package main

import (
	"fmt"
)

func main() {
	orderfile := "DrugRecord/test.txt"
	drugfile := "DrugRecord/Drugs.txt"
	orders, _ := BuildMap(orderfile, drugfile)
	fmt.Println(orders)

}
