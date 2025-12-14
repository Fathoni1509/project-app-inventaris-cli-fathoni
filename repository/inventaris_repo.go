package repository

import (
	"context"
	"project-app-inventaris-cli-fathoni/database"
	"project-app-inventaris-cli-fathoni/model"
)

type RepositoryInventarisInterface interface {
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

type RepositoryInventaris struct {
	DB database.PgxIface
}

func NewRepoInventaris(db database.PgxIface) RepositoryInventaris {
	return RepositoryInventaris{
		DB: db,
	}
}

// get list categories
func (repo *RepositoryInventaris) GetListCategories() ([]model.Categories, error) {
	query := `SELECT category_id, name_category, description FROM categories ORDER BY category_id`

	rows, err := repo.DB.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	var listCategories []model.Categories
	var list model.Categories
	for rows.Next() {
		err := rows.Scan(&list.Id, &list.Name, &list.Description)
		if err != nil {
			return nil, err
		}
		listCategories = append(listCategories, list)
	}

	return listCategories, nil
}

// add category
func (repo *RepositoryInventaris) AddCategory(category *model.Categories) error {
	query := `INSERT INTO categories (name_category, description) VALUES
	($1, $2) RETURNING category_id`

	err := repo.DB.QueryRow(context.Background(), query,
		category.Name,
		category.Description,
	).Scan(&category.Id)

	if err != nil {
		return err
	}

	return nil
}

// get list category by ID
func (repo *RepositoryInventaris) GetListCategoryById(category_id int) (model.Categories, error) {
	query := `SELECT category_id, name_category, description 
		FROM categories
		WHERE category_id=$1`
	
	var category model.Categories

	err := repo.DB.QueryRow(context.Background(), query, category_id).Scan(&category.Id, &category.Name, &category.Description)

	if err != nil {
		return model.Categories{}, err
	}

	return category, nil
}

// update category by ID
func (repo *RepositoryInventaris) UpdateCategory(category_id int, category *model.Categories) error {
	query := `UPDATE categories
		SET name_category=$1, description=$2
		WHERE category_id=$3`
	
	_, err := repo.DB.Exec(context.Background(), query,
		category.Name,
		category.Description,
		category_id,
	)

	if err != nil {
		return err
	}

	return nil
}

// delete category by ID
func (repo *RepositoryInventaris) DeleteCategory(category_id int) error {
	query := `DELETE FROM categories
		WHERE category_id=$1`
	
	_, err := repo.DB.Exec(context.Background(), query, category_id)

	if err != nil {
		return err
	}

	return nil
}

// get list all inventory items
func (repo *RepositoryInventaris) GetListItems() ([]model.Items, error) {
	query := `SELECT 
			i.item_id AS item_id, 
			i.category_id,
			i.name_item AS name,
			c.name_category AS category, 
			price, 
			purchase_date,
			replaced,
			CURRENT_DATE - purchase_date AS total_usage
		FROM categories c
		JOIN items i ON c.category_id = i.category_id
		ORDER BY i.item_id ASC`
	
	rows, err := repo.DB.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	var listItems []model.Items
	var list model.Items
	for rows.Next() {
		err := rows.Scan(&list.Id, &list.IdCategory, &list.Name, &list.Category, &list.Price, &list.PurchaseDate, &list.Replaced, &list.TotalUsage)
		if err != nil {
			return nil, err
		}
		listItems = append(listItems, list)
	}

	return listItems, nil
}

// add item
func (repo *RepositoryInventaris) AddItem(item *model.Items) error {
	query := `INSERT INTO items(category_id, name_item, price, purchase_date, replaced)
		VALUES ($1, $2, $3, $4, $5)`

	err := repo.DB.QueryRow(context.Background(), query,
		item.IdCategory,
		item.Name,
		item.Price,
		item.PurchaseDate,
		item.Replaced,
	).Scan(&item.Id)

	if err != nil {
		return err
	}

	return nil
}

// get list inventory item by ID
func (repo *RepositoryInventaris) GetListItemById(item_id int) (model.Items, error) {
	query := `SELECT 
		i.item_id AS item_id, 
		i.name_item AS name,
		c.name_category AS category, 
		price,
		purchase_date,
		replaced, 
		CURRENT_DATE - purchase_date AS total_usage
	FROM categories c
	JOIN items i ON c.category_id = i.category_id
	WHERE i.item_id = $1`

	var item model.Items
	
	err := repo.DB.QueryRow(context.Background(), query, item_id).Scan(&item.Id, &item.Name, &item.Category, &item.Price, &item.PurchaseDate, &item.Replaced, &item.TotalUsage)

	if err != nil {
		return model.Items{}, err
	}

	return item, nil
}

// update item by ID
func (repo *RepositoryInventaris) UpdateItem(item_id int, item *model.Items) error {
	query := `UPDATE items
		SET
			name_item = $1 ,
			category_id = $2, 
			price = $3, 
			purchase_date = $4,
			replaced = $5
		WHERE item_id = $6`
	
	_, err := repo.DB.Exec(context.Background(), query,
		item.Name,
		item.IdCategory,
		item.Price,
		item.PurchaseDate,
		item.Replaced,
		item_id,
	)

	return err
}

// delete item by ID
func (repo *RepositoryInventaris) DeleteItem(item_id int) error {
	query := `DELETE FROM items
		WHERE item_id=$1`
	
	_, err := repo.DB.Exec(context.Background(), query, item_id)

	if err != nil {
		return err
	}

	return nil
}

// get items must be replaced after 100 days and replaced is true
func (repo *RepositoryInventaris) ItemReplaced() ([]model.Items, error) {
	query := `SELECT 
			i.item_id AS item_id, 
			i.name_item AS name,
			c.name_category AS category, 
			price, 
			purchase_date,
			replaced,
			CURRENT_DATE - purchase_date AS total_usage
		FROM categories c
		JOIN items i ON c.category_id = i.category_id
		WHERE CURRENT_DATE - purchase_date > 100 AND replaced = TRUE
		ORDER BY i.item_id ASC`

	rows, err := repo.DB.Query(context.TODO(), query)

	if err != nil {
		return nil, err
	}

	var listReplaced []model.Items
	var list model.Items
	for rows.Next() {
		err := rows.Scan(&list.Id, &list.Name, &list.Category, &list.Price, &list.PurchaseDate, &list.Replaced ,&list.TotalUsage)

		if err != nil {
			return nil, err
		}

		listReplaced = append(listReplaced, list)
	}

	return listReplaced, nil
}

// get total invest
func (repo *RepositoryInventaris) TotalInvest() (model.TotalInvest, error) {
	query := `SELECT 
			COUNT(name_item) AS total_item,
			SUM (price) AS total_purchase_price,

			SUM(
				ROUND (
					price * power(0.8, 
					EXTRACT(YEAR FROM AGE(CURRENT_DATE, purchase_date)))
				) 
			) AS total_current_investment,
			
			SUM(
				ROUND (
					price - (price * power(0.8, 
					EXTRACT(YEAR FROM AGE(CURRENT_DATE, purchase_date)))
				))
			) AS total_depreciation_value
			
		FROM items`

	var totalInvest model.TotalInvest
	err := repo.DB.QueryRow(context.Background(), query).Scan(&totalInvest.TotalItem, &totalInvest.TotalPurchasePrice, &totalInvest.TotalInvest, &totalInvest.TotalDepreciation)

	if err != nil {
		return model.TotalInvest{}, err
	}

	return totalInvest, nil
}

// get invest and depreciation value of all items
func (repo *RepositoryInventaris) GetInvestValueItems() ([]model.Items, error) {
	query := `SELECT 
			item_id, 
			name_item, 
			price AS purchase_price, 
			purchase_date,
			EXTRACT(YEAR FROM AGE(CURRENT_DATE, purchase_date)) AS age_item,
			
			ROUND (
				price * power(0.8, 
				EXTRACT(YEAR FROM AGE(CURRENT_DATE, purchase_date)))
			) AS current_investment,
			
			ROUND (
				price - (price * power(0.8, 
				EXTRACT(YEAR FROM AGE(CURRENT_DATE, purchase_date)))
			)) AS depreciation_value
			
		FROM items
		ORDER BY item_id`

	rows, err := repo.DB.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	var investValues []model.Items
	var list model.Items
	for rows.Next() {
		err := rows.Scan(&list.Id, &list.Name, &list.Price, &list.PurchaseDate, &list.AgeItem, &list.CurrentInvestment, &list.DepreciationValue)

		if err != nil {
			return nil, err
		}

		investValues = append(investValues, list)
	}

	return investValues, nil
}

// get invest and depreciation value of item by ID
func (repo *RepositoryInventaris) GetInvestValueItemById(item_id int) (model.Items, error) {
	query := `SELECT 
			item_id, 
			name_item, 
			price AS purchase_price, 
			purchase_date,
			EXTRACT(YEAR FROM AGE(CURRENT_DATE, purchase_date)) AS age_item,
			
			ROUND (
				price * power(0.8, 
				EXTRACT(YEAR FROM AGE(CURRENT_DATE, purchase_date)))
			) AS current_investment,
			
			ROUND (
				price - (price * power(0.8, 
				EXTRACT(YEAR FROM AGE(CURRENT_DATE, purchase_date)))
			)) AS depreciation_value
			
		FROM items
		WHERE item_id = $1`

	var investValue model.Items
	err := repo.DB.QueryRow(context.Background(), query, item_id).Scan(&investValue.Id, &investValue.Name, &investValue.Price, &investValue.PurchaseDate, &investValue.AgeItem, &investValue.CurrentInvestment, &investValue.DepreciationValue)

	if err != nil {
		return model.Items{}, err
	}

	return investValue, nil
}

// search items by keyword
func (repo *RepositoryInventaris) SearchItems(keyword string) ([]model.Items, error) {
	query := `SELECT 
			i.item_id AS item_id, 
			i.name_item AS name,
			c.name_category AS category, 
			price, 
			purchase_date,
			CURRENT_DATE - purchase_date AS total_usage
		FROM categories c
		JOIN items i ON c.category_id = i.category_id
		WHERE i.name_item ILIKE $1
		ORDER BY i.item_id ASC`
	
	search := "%" + keyword + "%"
	rows, err := repo.DB.Query(context.Background(), query, search)
	
	if err != nil {
		return nil, err
	}

	var listItems []model.Items
	var list model.Items
	for rows.Next() {
		err := rows.Scan(&list.Id, &list.Name, &list.Category, &list.Price, &list.PurchaseDate, &list.TotalUsage)
		if err != nil {
			return nil, err
		}
		listItems = append(listItems, list)
	}

	return listItems, nil
}