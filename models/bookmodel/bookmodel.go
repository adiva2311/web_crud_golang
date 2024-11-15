package bookmodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Books{
	rows, err := config.DB.Query(`
	SELECT
		books.id,
		books.name,
		categories.name as category_name,
		books.stock,
		books.description,
		books.created_at,
		books.updated_at
	FROM books
	JOIN categories ON books.category_id = categories.id
	`)

	if err != nil{
		panic(err)
	}
	defer rows.Close()

	var books []entities.Books

	for rows.Next(){
		var book entities.Books
		err := rows.Scan(
			&book.Id,
			&book.Name,
			&book.Category_id.Name,
			&book.Stock,
			&book.Description,
			&book.Created_at,
			&book.Updated_at,
		)
		if err != nil{
			panic(err)
		}

		books = append(books, book)
	}

	return books

}

func Detail(Id int) entities.Books{
	row := config.DB.QueryRow(`
	SELECT
		books.id,
		books.name,
		categories.name as category_name,
		books.stock,
		books.description,
		books.created_at,
		books.updated_at
	FROM books
	JOIN categories ON books.category_id = categories.id
	WHERE books.id = ?`, Id)

	var book entities.Books
	err := row.Scan(
		&book.Id,
		&book.Name,
		&book.Category_id.Name,
		&book.Stock,
		&book.Description,
		&book.Created_at,
		&book.Updated_at,
	)

	if err != nil{
		panic(err)
	}
	
	return book
}

func Create(book entities.Books) bool{
	rows, err := config.DB.Exec(`
	INSERT INTO books(
		name, category_id, stock, description, created_at, updated_at 
	) VALUES (?, ?, ?, ?, ?, ?)`, book.Name, book.Category_id.Id, book.Stock, book.Description, book.Created_at, book.Updated_at)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := rows.LastInsertId()
	if err != nil{
		panic(err)
	}

	return lastInsertId > 0
}

func Update(id int, book entities.Books) bool{
	rows, err := config.DB.Exec(`
	UPDATE books SET 
		name = ?,
		category_id = ?,
		stock = ?,
		description = ?,
		updated_at = ?
	WHERE id = ?
	`,
		book.Name,
		book.Category_id.Id,
		book.Stock,
		book.Description,
		book.Updated_at,
		id,
	)

	if err != nil {
		panic(err)
	}

	result, err := rows.RowsAffected()
	if err != nil {
		panic(err)
	}

	// Kembalikan true jika ada baris yang diperbarui
	return result > 0
}

func Delete(id int) error{
	_, err := config.DB.Exec(`DELETE FROM books WHERE id = ?`, id)
	if err != nil{
		panic(err)
	}
	return err
}