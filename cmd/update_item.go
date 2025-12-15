package cmd

import (
	"bufio"
	"fmt"
	"os"
	"project-app-inventaris-cli-fathoni/handler"
	"project-app-inventaris-cli-fathoni/model"
	"project-app-inventaris-cli-fathoni/utils"
	"strconv"
	"strings"
	"time"
)

func UpdateItem(handlerInventaris handler.HandlerInventaris) {
	scanner := bufio.NewScanner(os.Stdin)

	allItems, _ := handlerInventaris.ServiceInventaris.GetListItems()
	categories, err := handlerInventaris.ServiceInventaris.GetListCategories()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Update Item By ID")

	// Input ID Item
	var id int
	var existingItem model.Items
	var isItemFound bool

	for {
		fmt.Print("ID Item: ")
		scanner.Scan()
		itemId := scanner.Text()

		var err error
		id, err = strconv.Atoi(itemId)
		if err != nil {
			fmt.Println("Input ID invalid")
			continue
		}

		for _, item := range allItems {
			if item.Id == id {
				existingItem = item
				isItemFound = true
				break
			}
		}

		if !isItemFound {
			fmt.Println("Item ID not found")
			continue
		}
		break 
	}

	// Tampilkan data saat ini sebagai info
	fmt.Printf("\nCurrently editing an item: %s (ID: %d)\n", existingItem.Name, existingItem.Id)
	utils.PrintTabel(categories) // Tampilkan tabel kategori bantu user

	// Input category ID
	var newIdCategory int
	for {
		// Tampilkan value lama agar user tahu
		fmt.Printf("New Category ID (Current: %d, Enter to keep): ", existingItem.IdCategory)
		scanner.Scan()
		input := scanner.Text()

		if strings.TrimSpace(input) == "" {
			newIdCategory = existingItem.IdCategory
			// fmt.Println("➡ Menggunakan Category ID lama.")
			break
		}

		// Jika tidak kosong, validasi input baru
		val, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input must be integer")
			continue
		}

		// Cek apakah ID Kategori valid (ada di tabel kategori)
		isValidCategory := false
		for _, c := range categories {
			if c.Id == val {
				isValidCategory = true
				break
			}
		}

		if !isValidCategory {
			fmt.Println("Category ID not found")
			continue
		}

		newIdCategory = val
		break
	}

	// Input Name
	var newName string
	fmt.Printf("New Item Name (Current: %s, Enter to keep): ", existingItem.Name)
	scanner.Scan()
	input := scanner.Text()

	if strings.TrimSpace(input) == "" {
		// Jika kosong, pakai nama lama
		newName = existingItem.Name
	} else {
		// Jika ada isinya, pakai nama baru
		newName = input
	}

	// input price
	var newPrice float32
	for {
		fmt.Printf("New Price (Current: %.2f, Enter to keep): ", existingItem.Price)
		scanner.Scan()
		input := scanner.Text()

		if strings.TrimSpace(input) == "" {
			newPrice = existingItem.Price
			break
		}

		// Parse float
		p, err := strconv.ParseFloat(input, 32) // Pakai 32 bit sesuai model float32
		if err != nil {
			fmt.Println("price must be numbers")
			continue
		}

		if p < 0 {
			fmt.Println("price must be more than 0")
			continue
		}

		newPrice = float32(p)
		break
	}

	// input purchase date
	var newPurchaseDate time.Time
	for {
		// Format tanggal lama biar enak dibaca
		oldDateStr := existingItem.PurchaseDate.Format("2006-01-02")
		fmt.Printf("New Purchase Date (Current: %s, Enter to keep): ", oldDateStr)
		scanner.Scan()
		input := scanner.Text()

		if strings.TrimSpace(input) == "" {
			newPurchaseDate = existingItem.PurchaseDate
			break
		}

		// Parse date
		t, err := time.Parse("2006-01-02", input)
		if err != nil {
			fmt.Println("Format date invalid (YYYY-MM-DD)")
			continue
		}

		newPurchaseDate = t
		break
	}

	// input replaced
	var newReplaced bool
	for {
		fmt.Printf("Replaced Status (Current: %t, Enter to keep): ", existingItem.Replaced)
		scanner.Scan()
		input := scanner.Text()

		if strings.TrimSpace(input) == "" {
			newReplaced = existingItem.Replaced
			break
		}

		val, err := strconv.ParseBool(input)
		if err != nil {
			fmt.Println("input must be 'true' or 'false'")
			continue
		}

		newReplaced = val
		break
	}

	// update
	updateItem := &model.Items{
		Name:         newName,
		IdCategory:   newIdCategory,
		Price:        newPrice,
		PurchaseDate: newPurchaseDate,
		Replaced:     newReplaced,
	}

	fmt.Println("\nSave changes...")
	err = handlerInventaris.UpdateItem(id, updateItem)
	if err != nil {
		fmt.Println("Error update:", err)
	} else {
		fmt.Println("Success update item")
	}
}