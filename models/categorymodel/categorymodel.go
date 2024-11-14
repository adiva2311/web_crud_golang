package categorymodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Categories {
	rows, err := config.DB.Query(`SELECT * FROM categories`)
	if err != nil{
		panic(err)
	}
	defer rows.Close()

	var categories[]entities.Categories
	for rows.Next(){
		var category entities.Categories
		err := rows.Scan(&category.Id, &category.Name, &category.Created_at, &category.Updated_at)
		if err != nil{
			panic(err)
		}
		categories = append(categories, category)
	}

	return categories
}

func Create(category entities.Categories) bool {
	result, err := config.DB.Exec(`
		INSERT INTO categories (name, created_at, updated_at)
		VALUES (?, ?, ?)`,
		category.Name, category.Created_at, category.Updated_at,
	)
	if err != nil{
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil{
		panic(err)
	}

	return lastInsertId > 0
}

func Detail(Id int) entities.Categories{
	row := config.DB.QueryRow(`SELECT id, name FROM categories WHERE id = ?`, Id)

	var category entities.Categories
	err := row.Scan(&category.Id, &category.Name)
	if err != nil{
		panic(err.Error())
	}
	return category
}

func Update(Id int, category entities.Categories) bool{
	query, err := config.DB.Exec(`UPDATE categories SET name = ?, updated_at = ? WHERE id = ?`, category.Name, category.Updated_at, Id)
	if err != nil{
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil{
		panic(err)
	}

	return result > 0
}

func Delete(Id int) error {
	_,err := config.DB.Exec(`DELETE FROM categories WHERE id = ?`, Id)
	return err
}