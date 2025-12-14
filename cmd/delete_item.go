package cmd

import (
	"bufio"
	"fmt"
	"os"
	"project-app-inventaris-cli-fathoni/handler"
	"strconv"
)

func DeleteItem(handlerInventaris handler.HandlerInventaris) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Delete Item by ID")

	for {
		fmt.Print("Input ID item will be deleted: ")
		scanner.Scan()
		id_item := scanner.Text()
	
		id, err := strconv.Atoi(id_item)
		if err != nil {
			fmt.Println("input ID invalid")
			continue
		} else {
			err = handlerInventaris.DeleteItem(id)
			if err != nil {
				fmt.Println("Error:", err)
			}
			break
		}
	}

}