package cmd

import (
	"bufio"
	"fmt"
	"os"
	"project-app-inventaris-cli-fathoni/handler"
	"project-app-inventaris-cli-fathoni/model"
	"strings"
)

func AddCategory(handleInventaris handler.HandlerInventaris) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Add New Category")

	var name_category string
	for {
		fmt.Print("Category name: ")
		scanner.Scan()
		name_category = scanner.Text()
	
		if strings.TrimSpace(name_category) == "" {
			fmt.Println("category name must not be empty")
			continue
		}
		break
	}

	var description string
	for {
		fmt.Print("Description: ")
		scanner.Scan()
		description = scanner.Text()
	
		if strings.TrimSpace(description) == "" {
			fmt.Println("category description must not be empty")
			continue
		}
		break
	}
	
	category := model.Categories{
		Name: name_category,
		Description: description,
	}

	err := handleInventaris.AddCategory(&category)
	if err != nil {
		fmt.Println("Error:", err)
	}
}