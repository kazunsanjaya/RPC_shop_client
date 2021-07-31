package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Entry struct {
	Name string
	Price float64
	Amount  float32
}

func main() {
	var strReply string
	var nameList []string
	var price float64
	var amount float32

	var client, err = rpc.DialHTTP("tcp", "localhost:1234")

	if err != nil {
		log.Fatal("Failed to connect: ", err)
	}

	a := Entry{"Carrot", 200, 12}
	b := Entry{"Spinach", 50, 4}
	c := Entry{"Pumpkin", 90, 24}

	//get list of available vegetables
	client.Call("VSERVER.ListVegetableNames","",&nameList)
	fmt.Println("Vegetable List: ", nameList)

	//get the price of a vegetable by name
	client.Call("VSERVER.GetPriceByName", "Carrot", &price)
	fmt.Println("Carrot price: ", price)

	//get the amount of a vegetable by name
	client.Call("VSERVER.GetAmountByName", "Spinach", &amount)
	fmt.Println("Spinach amount: ", amount)

	//add a new vegetable name to the server to be added to the server file
	client.Call("VSERVER.AddEntry", c, &strReply)
	fmt.Println("Entry added:  ", strReply)

	//add new price or available amount for a given vegetable to be updated in the server file
	client.Call("VSERVER.UpdateEntry", a, &strReply)
	fmt.Println("Entry update:  ", strReply)
	client.Call("VSERVER.UpdateEntry", b, &strReply)
	fmt.Println("Entry update:  ", strReply)

	//price of "a" after update
	client.Call("VSERVER.GetPriceByName", "Carrot", &price)
	fmt.Println("Carrot price after update: ", price)

	//amount of "b" after update
	client.Call("VSERVER.GetAmountByName", "Spinach", &amount)
	fmt.Println("Spinach amount after update: ", amount)


	//get list of available vegetables after update
	client.Call("VSERVER.ListVegetableNames","",&nameList)
	fmt.Println("Vegetable List after update: ", nameList)

}