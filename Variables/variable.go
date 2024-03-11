package main

import "fmt"

func main() {
	// declaration
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var name string

	var username string = "gedion"
	var password string = "2321"

	fmt.Println("Username and password: ", username, password)

	fmt.Printf(
		"%v %f %v %q\n",
		smsSendingLimit,
		costPerSMS,
		hasPermission,
		name,
	)

	// short variable declaration
	msg := "This is a message."
	fmt.Println(msg)

	penniesPerText := 2.5
	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)

	// same line declaration
	averageOpenRate, displayMessage := .35, "is the average open rate of your messages"

	fmt.Println(averageOpenRate, displayMessage)

	// constant
	// NB: constants don't use short declaration syntax
	const premiumPlanName = "Premium Plan"

	// computed constants:
	// constants must be know at compile time
	const secondsInMinute = 60
	const minutesInHour = 60

	const secondsInHour = minutesInHour * secondsInMinute

	fmt.Println("number of seconds in an hour:", secondsInHour)
	// That said you cant declare a constant that can only be computed at run-time
}
