package service

import (
	"errors"
	"project-app-inventaris-cli-fathoni/model"
	"project-app-inventaris-cli-fathoni/repository"
)

type ServiceInventarisInterface interface {
	GetListCategories() ([]model.Categories, error)
	AddCategory(category *model.Categories) error
	GetListCategoryById(category_id int) (model.Categories, error)
	UpdateCategory(category_id int, category *model.Categories) error
	DeleteCategory(category_id int) error
	GetListItems() ([]model.Items, error)
	AddItem(item *model.Items) error
	GetListItemById(item_id int) (model.Items, error)
	UpdateItem(item_id int, item *model.Items) error
	DeleteItem(item_id int) error
	ItemReplaced() ([]model.Items, error)
	TotalInvest() (model.TotalInvest, error)
	GetInvestValueItems() ([]model.Items, error)
	GetInvestValueItemById(item_id int) (model.Items, error)
	SearchItems(keyword string) ([]model.Items, error)
}

type ServiceInventaris struct {
	RepoInventaris repository.RepositoryInventarisInterface
}

func NewServiceInventaris(repoInventaris repository.RepositoryInventarisInterface) ServiceInventaris {
	return ServiceInventaris{
		RepoInventaris: repoInventaris,
	}
}

// service get list categories
func (serviceInventaris *ServiceInventaris) GetListCategories() ([]model.Categories, error) {
	return serviceInventaris.RepoInventaris.GetListCategories()
}

// service add category
func (serviceInventaris *ServiceInventaris) AddCategory(category *model.Categories) error {
	return serviceInventaris.RepoInventaris.AddCategory(category)
}

// service get list category by ID
func (serviceInventaris *ServiceInventaris) GetListCategoryById(category_id int) (model.Categories, error) {
	return serviceInventaris.RepoInventaris.GetListCategoryById(category_id)
}

// service update category by ID
func (serviceInventaris *ServiceInventaris) UpdateCategory(category_id int, category *model.Categories) error {
	categories, err := serviceInventaris.GetListCategories()
	if err != nil {
		return err
	}
	for _, c := range categories {
		if category_id == c.Id {
			return serviceInventaris.RepoInventaris.UpdateCategory(category_id, category)
		}
	}

	return errors.New("id category not found")
}

// service delete category by ID
func (serviceInventaris *ServiceInventaris) DeleteCategory(category_id int) error {
	categories, err := serviceInventaris.GetListCategories()
	if err != nil {
		return err
	}
	for _, c := range categories {
		if category_id == c.Id {
			return serviceInventaris.RepoInventaris.DeleteCategory(category_id)
		}
	}

	return errors.New("id category not found")
}

// service get list all inventory items
func (serviceInventaris *ServiceInventaris) GetListItems() ([]model.Items, error) {
	return serviceInventaris.RepoInventaris.GetListItems()
}

// service add item
func (serviceInventaris *ServiceInventaris) AddItem(item *model.Items) error {
	return serviceInventaris.RepoInventaris.AddItem(item)
}

// service get list inventory item by ID
func (serviceInventaris *ServiceInventaris) GetListItemById(item_id int) (model.Items, error) {
	return serviceInventaris.RepoInventaris.GetListItemById(item_id)
}

// service update item by ID
func (serviceInventaris *ServiceInventaris) UpdateItem(item_id int, item *model.Items) error {
	items, err := serviceInventaris.GetListItems()
	if err != nil {
		return err
	}
	for _, i := range items {
		if item_id == i.Id {
			return serviceInventaris.RepoInventaris.UpdateItem(item_id, item)
		}
	}

	return errors.New("id item not found")
}

// service delete item by ID
func (serviceInventaris *ServiceInventaris) DeleteItem(item_id int) error {
	items, err := serviceInventaris.GetListItems()
	if err != nil {
		return err
	}
	for _, i := range items {
		if item_id == i.Id {
			return serviceInventaris.RepoInventaris.DeleteItem(item_id)
		}
	}

	return errors.New("id item not found")
}

// service get items must be replaced after 100 days and replaced is true
func (serviceInventaris *ServiceInventaris) ItemReplaced() ([]model.Items, error) {
	return serviceInventaris.RepoInventaris.ItemReplaced()
}

// service get total invest
func (serviceInventaris *ServiceInventaris) TotalInvest() (model.TotalInvest, error) {
	return serviceInventaris.RepoInventaris.TotalInvest()
}

// service get invest and depreciation value of all items
func (serviceInventaris *ServiceInventaris) GetInvestValueItems() ([]model.Items, error) {
	return serviceInventaris.RepoInventaris.GetInvestValueItems()
}

// service get invest and depreciation value of item by ID
func (serviceInventaris *ServiceInventaris) GetInvestValueItemById(item_id int) (model.Items, error) {
	return serviceInventaris.RepoInventaris.GetInvestValueItemById(item_id)
}

// service search items by keyword
func (serviceInventaris *ServiceInventaris) SearchItems(keyword string) ([]model.Items, error) {
	return serviceInventaris.RepoInventaris.SearchItems(keyword)
}