package goDB

import (
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
)

/**
File: buildMap
Description: Creates a map of drugs and a map of orders from text files
@author Bryan Conn
@date 5/16/18
 */

 /**
 Function: makeDrugMap
 Description: Populates a map with drug structs
 @param fileName The file name containing drug information
  */
func makeDrugMap(fileName string) map[string]Drug {
	//Open file and make sure there are no errors
	drugFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	drugs := make(map[string]Drug)
	emptyDrug := Drug{}

	//Scanner to read the file line by line
	scanner := bufio.NewScanner(drugFile)
	scanner.Split( bufio.ScanLines )

	//Read each line and parse it into a drug
	check := scanner.Scan()
	for ;check != false; check = scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		qty, _ := strconv.Atoi(words[2])
		if drugs[words[1]] == emptyDrug {
			drug := MakeDrug(words[0], words[1], qty)
			drugs[words[1]] = drug
		} else {
			drugs[words[1]].UpdateQty(qty)
		}
	}
	drugFile.Close()
	return drugs
}

/**
Function: makeOrderMap
Description: Populates a map with order structs
@param fileName The name of the file containing past order information
@param drugs The map of drugs in the system
 */
func makeOrderMap(fileName string, drugs map[string]Drug) map[string][]Order {
	//Open file and check for erros
	orderFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	orders := make(map[string][]Order)

	//Scanner to read the file line by line
	scanner := bufio.NewScanner(orderFile)
	scanner.Split( bufio.ScanLines )

	/*
	Order all files
	NDC Date Qty ...
	 */
	check := scanner.Scan()
	for ;check != false; check = scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		//Get the drug
		drug := drugs[words[0]]
		//Get the date of the order
		date := MakeDate(words[1], words[2], words[3])
		//Get the quantity of the order
		qty, _ := strconv.Atoi(words[4])

		//Makes a purchase, audit, or prescription and adds it to
		//the list for the given drug
		var order Order
		if len(words) == 5 {
			purchase := MakePurchase(drug, date, qty)
			drug = purchase.PurchasedDrug
			order = MakeOrder(purchase)
		} else {
			//Get the log date of the order
			lDate := MakeDate(words[6], words[7], words[8])

			if len(words) == 9 {
				audit := MakeAudit(drug, qty, words[5], date, lDate)
				drug = audit.ADrug
				order = MakeOrder(audit)
			} else if len(words) == 10 {
				prescription := MakePrescription(drug, date, qty, words[5], words[9], lDate)
				drug = prescription.OrderDrug
				order = MakeOrder(prescription)
			}
		}
		//Access order with order.ThisOrder
		drugs[words[0]] = drug
		orders[words[0]] = append(orders[words[0]], order)
	}
	orderFile.Close()
	return orders
}

/**
Function: BuildMap
Description: Builds the drug and order map
@param orderFileName The file containing orders
@param drugsFileName The file containing drugs
 */
func BuildMap(orderFileName, drugsFileName string) (map[string][]Order, map[string]Drug){
	drugs := makeDrugMap(drugsFileName)
	return makeOrderMap(orderFileName, drugs), drugs
}