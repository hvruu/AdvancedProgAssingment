package mysql

import (
	"AdvancedProgAssignment/clinic/pkg/models"
	"database/sql"
)

type ProductModel struct {
	DB *sql.DB
}

func (m *ProductModel) Highthree() ([]*models.Product, error) {
	stmt := `SELECT id, name, img_path, description, category, price FROM products
				 ORDER BY price ASC LIMIT 3`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []*models.Product{}
	for rows.Next() {
		p := &models.Product{}
		err = rows.Scan(&p.ID, &p.Name, &p.ImagePath, &p.Description, &p.Category, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
