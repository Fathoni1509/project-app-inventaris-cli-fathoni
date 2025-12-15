package cmd

import (
	"fmt"
	"project-app-inventaris-cli-fathoni/handler"
	"project-app-inventaris-cli-fathoni/utils"
)

func ManajemenBarang(handlerInventaris handler.HandlerInventaris) {
	for {
		fmt.Printf("\n%s=== Inventory Management Features ===%s\n", utils.Green, utils.Reset)
		fmt.Println("-------------------------------------")
		fmt.Println("1. View List Items")
		fmt.Println("2. Add New Item")
		fmt.Println("3. View Detail Item by ID")
		fmt.Println("4. Edit Item")
		fmt.Println("5. Delete Item")
		fmt.Println("6. Back to Home")
		fmt.Println("-------------------------------------")

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
			handlerInventaris.ListItems()
		case 2:
			ClearScreen()
			AddItem(handlerInventaris)
		case 3:
			ClearScreen()
			DetailItemId(handlerInventaris)
		case 4:
			ClearScreen()
			UpdateItem(handlerInventaris)
		case 5:
			ClearScreen()
			DeleteItem(handlerInventaris)
		case 6:
			ClearScreen()
			fmt.Printf("\n%sBack to home...%s\n", utils.Yellow, utils.Reset)
			HomePage(handlerInventaris)
		default:
			fmt.Printf("%sInvalid option. Please choose between 1 - 6.%s\n", utils.Red, utils.Reset)
		}
	}

}
