package handler

import (
	"fmt"
	"project-app-inventaris-cli-fathoni/model"
	"project-app-inventaris-cli-fathoni/service"
	"project-app-inventaris-cli-fathoni/utils"
)

type HandlerInventaris struct {
	ServiceInventaris service.ServiceInventarisInterface
}

func NewHandlerInventaris(serviceInventaris service.ServiceInventarisInterface) HandlerInventaris {
	return HandlerInventaris{
		ServiceInventaris: serviceInventaris,
	}
}

// handler get list categories
func (handlerInventaris *HandlerInventaris) ListCategories() error {

	// call service get list categories
	categories, err := handlerInventaris.ServiceInventaris.GetListCategories()
	if err != nil {
		return err
	}

	fmt.Printf("=== List Categories ===\n\n")
	
	// cetak table
	utils.PrintTabel(categories)
	return nil
}

// handler add category
func (handlerInventaris *HandlerInventaris) AddCategory(category *model.Categories) error {
	// call service add category
	err := handlerInventaris.ServiceInventaris.AddCategory(category)
	if err != nil {
		return err
	}

	fmt.Println("Success add category")

	return nil
}

// handler list category by ID
func (handlerInventaris *HandlerInventaris) ListCategoryById(category_id int) error {
	// call service list category by ID
	category, err := handlerInventaris.ServiceInventaris.GetListCategoryById(category_id)

	if err != nil {
		return err
	}

	fmt.Printf("\n=== List Category ===\n\n")
	
	// cetak table
	utils.PrintTabel([]model.Categories{category})

	return nil

}

// handler update category by ID
func (handlerInventaris *HandlerInventaris) UpdateCategory(category_id int, category *model.Categories) error {
	// call service update category by ID
	err := handlerInventaris.ServiceInventaris.UpdateCategory(category_id, category)

	if err != nil {
		return err
	}

	fmt.Println("Success update category, with ID:", category_id)

	return nil
}

// handler delete category by ID
func (handleInventaris *HandlerInventaris) DeleteCategory(category_id int) error {
	// call service delete category by id
	err := handleInventaris.ServiceInventaris.DeleteCategory(category_id)

	if err != nil {
		return err
	}

	fmt.Println("Success delete category, with ID:", category_id)
	return nil
}

// handler get list items
func (handlerInventaris *HandlerInventaris) ListItems() error {

	// call service get list items
	items, err := handlerInventaris.ServiceInventaris.GetListItems()
	if err != nil {
		return err
	}

	fmt.Printf("=== List Items ===\n\n")
	
	// cetak table
	utils.PrintTabelItems(items)
	return nil
}

// handler add item
func (handlerInventaris *HandlerInventaris) AddItem(item *model.Items) error {
	// call service add item
	err := handlerInventaris.ServiceInventaris.AddItem(item)
	if err != nil {
		return err
	}

	fmt.Println("Success add item")

	return nil
}

// handler list item by ID
func (handlerInventaris *HandlerInventaris) ListItemsById(item_id int) error {
	// call service list item by ID
	item, err := handlerInventaris.ServiceInventaris.GetListItemById(item_id)

	if err != nil {
		return err
	}

	fmt.Printf("\n=== List Item ===\n\n")
	
	// cetak table
	utils.PrintTabelItems([]model.Items{item})

	return nil

}

// handler update item by ID
func (handlerInventaris *HandlerInventaris) UpdateItem(item_id int, item *model.Items) error {
	// call service update item by ID
	err := handlerInventaris.ServiceInventaris.UpdateItem(item_id, item)

	if err != nil {
		return err
	}

	return nil
}

// handler delete item by ID
func (handleInventaris *HandlerInventaris) DeleteItem(item_id int) error {
	// call service delete item by id
	err := handleInventaris.ServiceInventaris.DeleteItem(item_id)

	if err != nil {
		return err
	}

	fmt.Println("Success delete item, with ID:", item_id)
	return nil
}

// handler item replace
func (handlerInventaris *HandlerInventaris) ItemReplaced() error {
	// call service item replaced
	items, err := handlerInventaris.ServiceInventaris.ItemReplaced()
	if err != nil {
		return err
	}

	fmt.Printf("=== List Items Must Be Replaced ===\n\n")
	
	// cetak table
	utils.PrintTabelItems(items)
	return nil
}

// handler total invest
func (handlerInventaris *HandlerInventaris) TotalInvest() error {
	// call service total invest
	totalInvest, err := handlerInventaris.ServiceInventaris.TotalInvest()

	if err != nil {
		return err
	}

	fmt.Printf("=== Total Invest Value ===\n\n")
	
	// cetak table
	utils.PrintTabelInvest([]model.TotalInvest{totalInvest})
	return nil
}

// handler invest value items
func (handlerInventaris *HandlerInventaris) InvestValueItems() error {
	// call service total invest
	totalInvest, err := handlerInventaris.ServiceInventaris.GetInvestValueItems()

	if err != nil {
		return err
	}

	fmt.Printf("=== Invest Value Each Items ===\n\n")
	
	// cetak table
	utils.PrintTabelInvestItems(totalInvest)
	return nil
}

// handler invest value of item by ID
func (handlerInventaris *HandlerInventaris) InvestValueItemById(item_id int) error {
	// call service total invest
	totalInvest, err := handlerInventaris.ServiceInventaris.GetInvestValueItemById(item_id)

	if err != nil {
		return err
	}

	fmt.Printf("=== Invest Value Item ===\n\n")
	
	// cetak table
	utils.PrintTabelInvestItems([]model.Items{totalInvest})
	return nil
}

// handler search items by keyword
func (handlerInventaris *HandlerInventaris) SearchItems(keyword string) error {
	// call service search items
	items, err := handlerInventaris.ServiceInventaris.SearchItems(keyword)

	if err != nil {
		return err
	}

	fmt.Printf("=== Search Results ===\n\n")

	// cetak table
	utils.PrintTabelItems(items)

	return nil
}