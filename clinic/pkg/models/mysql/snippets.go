package mysql

import (
	"database/sql"
	"errors"
	"fmt"

	"AdvancedProgAssignment/clinic/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, category, imagePath string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, category, image_path, created_at)
	VALUES(?, ?, ?, ?, UTC_TIMESTAMP())`
	result, err := m.DB.Exec(stmt, title, content, category, imagePath)
	if err != nil {
		fmt.Println("Error executing SQL statement:", err)
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *SnippetModel) InsertMenu(meal_name, weekday, quantity string) (int, error) {
	stmt := `INSERT INTO canteen_menu(meal_name, weekday, quantity)
	VALUES(?, ?, ?)`
	result, err := m.DB.Exec(stmt, meal_name, weekday, quantity)
	if err != nil {
		fmt.Println("Error executing SQL statement:", err)
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *SnippetModel) GetMenu() ([]*models.Menu, error) {
	stmt := `SELECT * FROM canteen_menu`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []*models.Menu{}
	for rows.Next() {
		s := &models.Menu{}
		err = rows.Scan(&s.ID, &s.Meal_name, &s.Weekday, &s.Quantity)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT id, title,content, category, image_path, created_at FROM snippets
	WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)
	s := &models.Snippet{}
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Category, &s.ImagePath, &s.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := `SELECT id, title, content, category, image_path, created_at FROM snippets
				 ORDER BY created_at DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []*models.Snippet{}
	for rows.Next() {
		s := &models.Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Category, &s.ImagePath, &s.Created)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
