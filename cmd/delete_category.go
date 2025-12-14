package cmd

import (
	"bufio"
	"fmt"
	"os"
	"project-app-inventaris-cli-fathoni/handler"
	"strconv"
)

func DeleteCategory(handlerInventaris handler.HandlerInventaris) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Delete Category by ID")

	for {
		fmt.Print("Input ID category will be deleted: ")
		scanner.Scan()
		id_category := scanner.Text()
	
		id, err := strconv.Atoi(id_category)
		if err != nil {
			fmt.Println("input ID invalid")
			continue
		} else {
			err = handlerInventaris.DeleteCategory(id)
			if err != nil {
				fmt.Println("Error:", err)
			}
			break
		}
	}

}