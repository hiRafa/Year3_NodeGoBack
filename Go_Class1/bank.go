package main

import (
	"fmt"

	"example.com/banking/fileops"
)

const accountBalanceFile = "account-balance.txt"

func main() {
	accountBalance, error := fileops.GetStringFloatFromFile(accountBalanceFile)

	if error != nil {
		fmt.Println("ERROR")
		fmt.Println(error)
		panic("Stop in the name of love")
	}

	for {
		options()

		var choice int
		fmt.Scan(&choice)

		wantsCheckBalance := choice == 1

		if wantsCheckBalance {
			fmt.Println("your balance: ", accountBalance)
		} else if choice == 2 {
			fmt.Println("deposit amount before: ", accountBalance)

			var depositAmount float64
			fmt.Print("deposit amount: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 1 {
				return
			}

			accountBalance += depositAmount
			fmt.Println("deposit amount after deposit: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		} else if choice == 3 {
			fmt.Println("withdrawl amount before: ", accountBalance)

			var withdrawlAmount float64
			fmt.Print("withdrawl amount: ")
			fmt.Scan(&withdrawlAmount)

			if withdrawlAmount <= 1 {
				return
			}

			if withdrawlAmount >= accountBalance {
				fmt.Print("not enough credits")
				return
			}

			accountBalance -= withdrawlAmount
			fmt.Println("withdrawl amount after withdrawl: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		} else {
			fmt.Println("Go bitch")
			break
		}
	}

	fmt.Println("Out of the loop, bitch. Go Go")
}
