package cmd

import (
	"fmt"
	"os"
	"project-app-inventaris-cli-fathoni/handler"
	"project-app-inventaris-cli-fathoni/utils"
)

func HomePage(handlerInventaris handler.HandlerInventaris) {
	for {
		fmt.Printf("\n%s=== INVENTORY APP MENU ===%s\n", utils.Yellow, utils.Reset)
		fmt.Println("1) Item Category Features")
		fmt.Println("2) Inventory Management Features")
		fmt.Println("3) List Items Must Be Replaced")
		fmt.Println("4) Investment and Depreciation Report")
		fmt.Println("5) Search Items by Keyword")
		fmt.Println("6) Exit")
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
			FiturKategori(handlerInventaris)
		case 2:
			ClearScreen()
			ManajemenBarang(handlerInventaris)
		case 3:
			ClearScreen()
			handlerInventaris.ItemReplaced()
		case 4:
			ClearScreen()
			InvestmentDepreciationReport(handlerInventaris)
		case 5:
			ClearScreen()
			SearchItems(handlerInventaris)
		case 6:
			fmt.Printf("%sExiting...%s\n", utils.Yellow, utils.Reset)
			os.Exit(0)
		default:
			fmt.Printf("%sInvalid option. Please choose between 1 - 6.%s\n", utils.Red, utils.Reset)
		}
	}

}
