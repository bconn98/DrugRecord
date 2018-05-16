package main

import (
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
)

func makeDrugMap(fileName string) map[string]Drug {
	drugFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	drugs := make(map[string]Drug)

	scanner := bufio.NewScanner(drugFile)

	scanner.Split( bufio.ScanLines )
	check := scanner.Scan()
	emptyDrug := Drug{}
	for ;check != false; check = scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		qty, _ := strconv.Atoi(words[2])
		if drugs[words[0]] == emptyDrug {
			drug := makeDrug(words[0], words[1], qty)
			drugs[words[0]] = drug
		} else {
			drugs[words[0]] = makeDrug(words[0], words[1], drugs[words[0]].Quantity + qty)
		}
	}
	drugFile.Close()
	return drugs
}

func makeOrderMap(fileName string, drugs map[string]Drug) map[string][]Order {
	orderFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	orders := make(map[string][]Order)

	scanner := bufio.NewScanner(orderFile)
	scanner.Split( bufio.ScanLines )

	check := scanner.Scan()

	/*
	Order all files
	Name Date Qty ...
	 */
	for ;check != false; check = scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		drug := drugs[words[0]]
		month, _ := strconv.Atoi(words[1])
		day, _ := strconv.Atoi(words[2])
		year, _ := strconv.Atoi(words[3])
		qty, _ := strconv.Atoi(words[4])
		date := Date{month, day, year}
		var order Order
		if len(words) == 5 {
			purchase := makePurchase(drug, date, qty)
			order = makeOrder(Prescription{}, purchase)
		} else if len(words) == 12 {
			month, _ = strconv.Atoi(words[9])
			day, _ = strconv.Atoi(words[10])
			year, _ = strconv.Atoi(words[11])
			aQty, _ := strconv.Atoi(words[7])
			audit := words[8]
			lDate := Date{month, day, year}

			prescription := makePrescritption(drug, date, qty, words[5], words[6], aQty, lDate)
			order = makeOrder(prescription, Purchase{})
		}
		orders[words[0]] = append(orders[words[0]], order)
	}
	orderFile.Close()
	return orders
}

func BuildMap(orderFileName, drugsFileName string) map[string][]Order{
	drugs := makeDrugMap(drugsFileName)

	return makeOrderMap(orderFileName, drugs)
}
