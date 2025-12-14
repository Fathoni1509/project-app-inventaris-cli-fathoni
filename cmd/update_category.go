package cmd

import (
	"bufio"
	"fmt"
	"os"
	"project-app-inventaris-cli-fathoni/handler"
	"project-app-inventaris-cli-fathoni/model"
	"strconv"
	"strings"
)

func UpdateCategory(handlerInventaris handler.HandlerInventaris) {
	scanner := bufio.NewScanner(os.Stdin)

	category, _ := handlerInventaris.ServiceInventaris.GetListCategories()
	
	fmt.Println("Update Category by ID")

	var id int
	var err error

	for {
		fmt.Print("ID Category: ")
		scanner.Scan()
		categoryId := scanner.Text()
	
		id, err = strconv.Atoi(categoryId)
	
		if err != nil {
			fmt.Println("input ID invalid")
			continue
		}
		break
	}

	var newName string
	fmt.Print("New Category Name: ")
	scanner.Scan()
	newName = scanner.Text()

	if strings.TrimSpace(newName) == "" {
		for _, c := range category {
			if id == c.Id {
				newName = c.Name
				break
			}
		}
	}

	var newDescription string
	fmt.Print("New Category Description: ")
	scanner.Scan()
	newDescription = scanner.Text()

	if strings.TrimSpace(newDescription) == "" {
		for _, c := range category {
			if id == c.Id {
				newDescription = c.Description
				break
			}
		}
	}

	updateCategory := &model.Categories {
		Name: newName,
		Description: newDescription,
	}

	err = handlerInventaris.UpdateCategory(id, updateCategory)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}