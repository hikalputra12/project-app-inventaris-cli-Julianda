package repository

import (
	"context"
	"inventory-system/model"

	"github.com/jackc/pgx/v5"
)

type RepositoryInventory struct {
	DB *pgx.Conn
}

type RepositoryInventoryInterface interface {
	ItemsMoreThan100Days() ([]model.TotalUsageDays, error)
	TotalInvesmentValue() ([]model.TotalInvestmentValue, error)
	InvesmentAndDepreciationValueByID(id int) (*model.Depreciation, error)
	FindInventoryByName(name string) (*model.Inventory, error)
}

func NewRepositoryInventory(db *pgx.Conn) RepositoryInventory {
	return RepositoryInventory{
		DB: db,
	}
}

// buat func reponya di sini
func (r *RepositoryInventory) ItemsMoreThan100Days() ([]model.TotalUsageDays, error) {
	query := `SELECT name,
	         price,
			CURRENT_DATE -purchase_date AS total_usage_days
			FROM inventory_items
			WHERE CURRENT_DATE- purchase_date >100;`
	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	totalUsageDays := make([]model.TotalUsageDays, 0)
	for rows.Next() {
		var t model.TotalUsageDays
		err := rows.Scan(&t.Name, &t.Price, &t.TotalUsageDays)
		if err != nil {
			return nil, err
		}
		totalUsageDays = append(totalUsageDays, t)

	}
	return totalUsageDays, nil
}

// display total invesment value
func (r *RepositoryInventory) TotalInvesmentValue() ([]model.TotalInvestmentValue, error) {
	query := `SELECT
    SUM(price * POWER(
				0.8,
		DATE_PART('year', AGE(CURRENT_DATE, purchase_date)))) AS total_investment_value
		FROM inventory_items`
	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var TotalInvesmentValue []model.TotalInvestmentValue
	for rows.Next() {
		var t model.TotalInvestmentValue
		err := rows.Scan(&t.TotalInvestmentValue)
		if err != nil {
			return nil, err
		}
		TotalInvesmentValue = append(TotalInvesmentValue, t)
	}
	return TotalInvesmentValue, nil
}

// display invesment dan depresiasi by id
func (r *RepositoryInventory) InvesmentAndDepreciationValueByID(id int) (*model.Depreciation, error) {
	query := `SELECT
    name,
    price AS initial_price,

    price * POWER(
        0.8,
        DATE_PART('year', AGE(CURRENT_DATE, purchase_date))
    ) AS investment_value,

    price - (
        price * POWER(
            0.8,
            DATE_PART('year', AGE(CURRENT_DATE, purchase_date))
        )
    ) AS depreciation

FROM inventory_items
WHERE inventory_items_id = $1;`

	rows, err := r.DB.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invesment *model.Depreciation
	for rows.Next() {
		var t model.Depreciation
		err := rows.Scan(&t.Name, &t.InitialPrice, &t.InvestmentValue, &t.Depreciation)
		if err != nil {
			return nil, err
		}
		invesment = &t
	}
	return invesment, nil
}

// display inventory by name
func (r *RepositoryInventory) FindInventoryByName(name string) (*model.Inventory, error) {
	query := `SELECT name,price,purchase_date FROM inventory_items
WHERE name =$1;`

	rows, err := r.DB.Query(context.Background(), query, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inventory *model.Inventory
	for rows.Next() {
		var t model.Inventory
		err := rows.Scan(&t.Name, &t.Price, &t.PurchaseDate)
		if err != nil {
			return nil, err
		}
		inventory = &t
	}
	return inventory, nil
}
