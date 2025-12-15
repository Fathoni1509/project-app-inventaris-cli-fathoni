package cmd

import (
	"fmt"
	"project-app-inventaris-cli-fathoni/handler"
	"project-app-inventaris-cli-fathoni/utils"
)

func FiturKategori(handlerInventaris handler.HandlerInventaris) {
	for {
		fmt.Printf("\n%s===  Item Category Features  ===%s\n", utils.Blue, utils.Reset)
		fmt.Println("------------------------------- ")
		fmt.Println("1. View List Categories")
		fmt.Println("2. Add New Category")
		fmt.Println("3. View Detail Category by ID")
		fmt.Println("4. Edit Category")
		fmt.Println("5. Delete Category")
		fmt.Println("6. Back to Home")
		fmt.Println("------------------------------- ")

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
			handlerInventaris.ListCategories()
		case 2:
			ClearScreen()
			AddCategory(handlerInventaris)
		case 3:
			ClearScreen()
			DetailCategoryId(handlerInventaris)
		case 4:
			ClearScreen()
			UpdateCategory(handlerInventaris)
		case 5:
			ClearScreen()
			DeleteCategory(handlerInventaris)
		case 6:
			ClearScreen()
			fmt.Printf("\n%sBack to home...%s\n", utils.Yellow, utils.Reset)
			HomePage(handlerInventaris)
		default:
			fmt.Printf("%sInvalid option. Please choose between 1 - 6.%s\n", utils.Red, utils.Reset)
		}
	}

}
