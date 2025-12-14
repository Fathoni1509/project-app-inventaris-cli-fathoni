package cmd

import (
	"fmt"
	"os"
	"project-app-inventaris-cli-fathoni/handler"
)

func HomePage(handlerInventaris handler.HandlerInventaris) {
	for {
		fmt.Println("\n=== INVENTORY APP MENU ===")
		fmt.Println("***  Fitur Kategori Barang  ***")
		fmt.Println("-------------------------------")
		fmt.Println("1. View List Categories")
		fmt.Println("2. Add New Category")
		fmt.Println("3. View Detail Category by ID")
		fmt.Println("4. Edit Category")
		fmt.Println("5. Delete Category")

		fmt.Printf("\n\n*** Manajemen Barang Inventaris ***\n")
		fmt.Println("-------------------------------")
		fmt.Println("6. View List Items")
		fmt.Println("7. Add New Item")
		fmt.Println("8. View Detail Item by ID")
		fmt.Println("9. Edit Item")
		fmt.Println("10. Delete Item")

		fmt.Printf("\n\n*** Barang yang Perlu Diganti ***\n")
		fmt.Println("-------------------------------")
		fmt.Println("11. View List Items Must Be Replaced")

		fmt.Printf("\n\n*** Laporan Investasi dan Depresiasi ***\n")
		fmt.Println("-------------------------------")
		fmt.Println("12. View Total Invest")
		fmt.Println("13. View Invest Value All items")
		fmt.Println("14. View Invest Value Item by ID")

		fmt.Printf("\n\n*** Fitur Pencarian ***\n")
		fmt.Println("-------------------------------")
		fmt.Println("15. Search Items by Keyword")


		fmt.Println("16. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
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
			handlerInventaris.ListItems()
		case 7:
			ClearScreen()
			AddItem(handlerInventaris)
		case 8:
			ClearScreen()
			DetailItemId(handlerInventaris)
		case 9:
			ClearScreen()
			UpdateItem(handlerInventaris)
		case 10:
			ClearScreen()
			DeleteItem(handlerInventaris)
		case 11:
			ClearScreen()
			handlerInventaris.ItemReplaced()
		case 12:
			ClearScreen()
			handlerInventaris.TotalInvest()
		case 13:
			ClearScreen()
			handlerInventaris.InvestValueItems()
		case 14:
			ClearScreen()
			DetailInvestId(handlerInventaris)
		case 15:
			ClearScreen()
			SearchItems(handlerInventaris)
		case 16:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please choose between 1 - 16.")
		}
	}

}
