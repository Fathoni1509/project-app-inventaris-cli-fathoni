package cmd

import (
	"fmt"
	"project-app-inventaris-cli-fathoni/handler"
	"project-app-inventaris-cli-fathoni/utils"
)

func InvestmentDepreciationReport(handlerInventaris handler.HandlerInventaris) {
	for {
		fmt.Printf("\n%s=== Investment and Depreciation Report ===%s\n", utils.Magenta, utils.Reset)
		fmt.Println("------------------------------------------")
		fmt.Println("1. View Total Invest")
		fmt.Println("2. View Invest Value All items")
		fmt.Println("3. View Invest Value Item by ID")
		fmt.Println("4. Back to Home")
		fmt.Println("------------------------------------------")

		fmt.Print("Choose an option: ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Printf("%sInvalid input. Please enter a number.%s\n", utils.Red, utils.Reset)
			continue
		}

		switch choice {
		case 1:
			ClearScreen()
			handlerInventaris.TotalInvest()
		case 2:
			ClearScreen()
			handlerInventaris.InvestValueItems()
		case 3:
			ClearScreen()
			DetailInvestId(handlerInventaris)
		case 4:
			ClearScreen()
			fmt.Printf("\n%sBack to home...%s\n", utils.Yellow, utils.Reset)
			HomePage(handlerInventaris)
		default:
			fmt.Printf("%sInvalid option. Please choose between 1 - 4.%s\n", utils.Red, utils.Reset)
		}
	}

}
