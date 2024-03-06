package mysql

import (
	"database/sql"
	"errors"
	"fmt"

	"AdvancedProgAssignment/clinic/pkg/models"
)

type NewsModel struct {
	DB *sql.DB
}

func (m *NewsModel) Insert(title, content, category, imagePath string) (int, error) {
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

func (m *NewsModel) Get(id int) (*models.News, error) {
	stmt := `SELECT id, title,content, category, image_path, created_at FROM snippets
	WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)
	s := &models.News{}
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

func (m *NewsModel) Latest() ([]*models.News, error) {
	stmt := `SELECT id, title, content, category, image_path, created_at FROM snippets
				 ORDER BY created_at DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []*models.News{}
	for rows.Next() {
		s := &models.News{}
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
