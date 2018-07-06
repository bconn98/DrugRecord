package goDB

import (
	"fmt"
	"strings"
	"strconv"
)

func test() {
	orderfile := "DrugRecord/orders.txt"
	drugfile := "DrugRecord/Drugs.txt"
	orders, drugs := BuildMap(orderfile, drugfile)
	fmt.Println("Input as NDC,Date,Qty[,Initials,OrderDate,LogDate[,Script#]]")
	fmt.Println("Type \"Exit\" to quit")
	var text string
	fmt.Scan(&text)
	for ; text != "Exit" ; fmt.Scan(&text) {
		lst := strings.Split(text, ",")
		var order Order
		drug := drugs[lst[0]]
		//Get the date of the order
		date := MakeDate(lst[1], lst[2], lst[3])
		//Get the quantity of the order
		qty, _ := strconv.Atoi(lst[4])
		if len(lst) == 5 {
			emptyDrug := Drug{}
			if drug == emptyDrug {
				fmt.Print("Enter new drug name:")
				fmt.Scan(&text)
				drug = MakeDrug(text, lst[0], 0)
			}
			purchase := MakePurchase(drug, date, qty)
			drug = purchase.PurchasedDrug
			order = MakeOrder(purchase)
		} else {
			//Get the log date of the order
			lDate := MakeDate(lst[6], lst[7], lst[8])

			if len(lst) == 9 {
				audit := MakeAudit(drug, qty, lst[5], date, lDate)
				drug = audit.ADrug
				order = MakeOrder(audit)
			} else if len(lst) == 10 {
				prescription := MakePrescription(drug, date, qty, lst[5], lst[9], lDate)
				drug = prescription.OrderDrug
				order = MakeOrder(prescription)
			}
		}
		//Access order with order.ThisOrder
		drugs[lst[0]] = drug
		orders[lst[0]] = append(orders[lst[0]], order)
	}
	fmt.Println(orders)
}
