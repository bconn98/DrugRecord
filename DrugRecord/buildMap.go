package main

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
		month, _ := strconv.Atoi(words[1])
		day, _ := strconv.Atoi(words[2])
		year, _ := strconv.Atoi(words[3])
		date := MakeDate(month, day, year)
		//Get the quantity of the order
		qty, _ := strconv.Atoi(words[4])

		//Makes a purchase, audit, or prescription and adds it to
		//the list for the given drug
		var order Order
		if len(words) == 5 {
			purchase := MakePurchase(drug, date, qty)
			order = MakeOrder(purchase)
		} else if len(words) == 8 {
			//Get the log date of the order
			month, _ = strconv.Atoi(words[7])
			day, _ = strconv.Atoi(words[8])
			year, _ = strconv.Atoi(words[9])
			lDate := MakeDate(month, day, year)

			audit := MakeAudit(drug, qty, words[5], date, lDate)
			order = MakeOrder(audit)
		} else if len(words) == 10 {
			//Get the log date of the order
			month, _ = strconv.Atoi(words[7])
			day, _ = strconv.Atoi(words[8])
			year, _ = strconv.Atoi(words[9])
			lDate := MakeDate(month, day, year)

			prescription := MakePrescritption(drug, date, qty, words[5], words[6], lDate)
			order = MakeOrder(prescription)
		}
		//Access order with order.ThisOrder
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