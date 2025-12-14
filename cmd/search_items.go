package cmd

import (
	"bufio"
	"fmt"
	"os"
	"project-app-inventaris-cli-fathoni/handler"
)

func SearchItems(handlerInventaris handler.HandlerInventaris) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Search Items by Keyword")

	fmt.Print("Input keyword: ")
	scanner.Scan()
	keyword := scanner.Text()

	err := handlerInventaris.SearchItems(keyword)
	if err != nil {
		fmt.Println("Error:", err)
	}
}