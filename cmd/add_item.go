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

func AddItem(handlerInventaris handler.HandlerInventaris) {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Add New Item")

	// tampilkan list kategori
	categories, err := handlerInventaris.ServiceInventaris.GetListCategories()

	if err != nil {
		fmt.Println("Error:", err)
	}

	utils.PrintTabel(categories)

	// input id
	var id_category int
	for {
		fmt.Print("Input ID category: ")
		scanner.Scan()
		id := scanner.Text()

		id_category, err = strconv.Atoi(id)

		if err != nil {
			fmt.Println("input ID invalid")
			continue
		}

		isFound := false
		for _, c := range categories {
			if id_category == c.Id {
				isFound = true
				break
			}
		}

		if isFound {
			break
		} else {
			fmt.Println("id category not found")
		}
	}

	// input item name
	var name_item string
	for {
		fmt.Print("Item name: ")
		scanner.Scan()
		name_item = scanner.Text()
	
		if strings.TrimSpace(name_item) == "" {
			fmt.Println("item name must not be empty")
			continue
		}
		break
	}

	// input price
	var price float32
	for {
		fmt.Print("Price (ex: 100000): ")
		scanner.Scan()
		input := scanner.Text()
	
		// Parse ke float
        p, err := strconv.ParseFloat(input, 64)
        if err != nil {
            fmt.Println("Error:", err)
            continue
        }
        
        if p < 0 {
             fmt.Println("Error: price must be more than 0")
             continue
        }

        price = float32(p)
        break
	}

	// input purchase date
	var purchaseDate time.Time
    for {
        fmt.Print("Purchase date (YYYY-MM-DD): ")
        scanner.Scan()
        input := scanner.Text()

        t, err := time.Parse("2006-01-02", input)
        if err != nil {
            fmt.Println("Error: date format invalid")
            continue
        }

        purchaseDate = t
        break
    }

	// input status replaced
	var replaced bool
	for {
        fmt.Print("Replaced (true/false): ")
        scanner.Scan()
        input := scanner.Text()

        val, err := strconv.ParseBool(input)
        if err != nil {
            fmt.Println("Error: input must be 'true' or 'false'")
            continue
        }

        replaced = val
        break
    }

	item := model.Items{
		IdCategory: id_category,
		Name: name_item,
		Price: price,
		PurchaseDate: purchaseDate,
		Replaced: replaced,
	}

	err = handlerInventaris.AddItem(&item)
	if err != nil {
		fmt.Println("Error:", err)
	}
}