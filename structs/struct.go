package main

import (
	"fmt"
)

// ex1
type modeOfComms struct {
	internet bool
	sms bool
}

type messageToSent struct {
	sender string
	reciever string
	number int
	cost float64 
	mode modeOfComms
	
}
func main1() {
	messageInformation := messageToSent{}
	messageInformation.number = 7822650126
	messageInformation.sender = "Mike"
	messageInformation.reciever = "Alice"
	messageInformation.cost = 0.33
	messageInformation.mode.internet = false
	messageInformation.mode.sms = true

	fmt.Println(messageInformation)
}
// e2 anonymous struct 

func main2() {
	myCar := struct {
		model int
		maker string
	} {
		model: 2021,
		maker: "Honda",
	}
	fmt.Println(myCar)
}

// ex3 embedded struct vs nested struct
type wheel struct {
	frontTyrePrice int
	rareTyrePrice int
}
type Car struct {
	model int
	company string
	tyreInfo wheel
}
type specs struct {
	os string
	manufactured int
}
type device struct {
	name string
	price int
	specs
}

func main3() {
	carOne := Car{
		model: 265,
		company: "Tesla Motors",
		tyreInfo: wheel{frontTyrePrice: 23, rareTyrePrice:55 },
	}
	deviceOne := device{
		name: "iPhone 15",
		price: 80000,
		specs: specs{
			os: "ios",
			manufactured: 2023,
		},
	}
	// here we see the ease of accessing the keys of embedded struct over nested struct
	fmt.Println(carOne.tyreInfo.frontTyrePrice ,deviceOne.os)
}
// ex 4 methods

type authenticationInformation struct {
	uuid int
	username string
	password string
}
func (authInfo authenticationInformation) getBasicAuth() string {
	return fmt.Sprintf(
		"Authentication: %d %s %s",
		authInfo.uuid,
		authInfo.username,
		authInfo.password,
	)
}
func main() {
	user := authenticationInformation{
		uuid: 265,
		username: "bodanabeel",
		password: "N766N85c088@bodanabeel",
	}
	fmt.Println(user.getBasicAuth())
}