package cmd

import (
	"bufio"
	"fmt"
	"os"
	"project-app-inventaris-cli-fathoni/handler"
	"strconv"
)

func DetailCategoryId(handlerInventaris handler.HandlerInventaris) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Detail Category by ID")

	for {
		fmt.Print("Input ID category: ")
		scanner.Scan()
		id_category := scanner.Text()

		id, err := strconv.Atoi(id_category)
		if err != nil {
			fmt.Println("input ID invalid")
			continue
		} else {
			err = handlerInventaris.ListCategoryById(id)
			if err != nil {
				fmt.Println("Error:", err)
			}
			break
		}
	}
}