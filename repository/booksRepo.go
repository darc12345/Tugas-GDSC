package repository

import (
	"fmt"
	"log"
	"main/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *DBrepo) GetBooksRepo(c *gin.Context) ([]model.Book, error) {
	db := h.db
	rows, err := db.Query("SELECT id,title,author,published_at,created_at,updated_at FROM books")
	if err != nil {
		log.Printf("Failed to retrieve products from database. Original Error: %s", err)
		return nil, fmt.Errorf("error fetching product list")
	}
	defer rows.Close()

	var books []model.Book

	// Iterate through the result set
	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.PublishedAt, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			log.Printf("Failed to scan product data. Original Error: %s", err)
			return nil, fmt.Errorf("error processing product data")
		}
		books = append(books, book)
	}

	// Check for errors during rows iteration
	if err = rows.Err(); err != nil {
		log.Printf("Error during rows iteration for products. Original Error: %s", err)
		return nil, fmt.Errorf("error processing product data")
	}

	return books, nil
}

func (h *DBrepo) GetBookByIDRepo(c *gin.Context) (model.Book, error) {
	db := h.db
	id := c.Param("id")
	var book model.Book

	// Query product by ID
	err := db.QueryRow("SELECT id,title,author,published_at,created_at,updated_at FROM books WHERE id = ?", id).
		Scan(&book.Id, &book.Title, &book.Author, &book.PublishedAt, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		log.Printf("Failed to retrieve product by ID from database. Product ID: %s, Original Error: %s", id, err)
		return model.Book{}, fmt.Errorf("error fetching product details")
	}

	return book, nil
}
func (h *DBrepo) getBookByID(id string) (model.Book, error) {
	db := h.db
	var book model.Book

	// Query product by ID
	err := db.QueryRow("SELECT id,title,author,published_at,created_at,updated_at FROM books WHERE id = ?", id).
		Scan(&book.Id, &book.Title, &book.Author, &book.PublishedAt, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		log.Printf("Failed to retrieve product by ID from database. Product ID: %s, Original Error: %s", id, err)
		return model.Book{}, fmt.Errorf("error fetching product details")
	}

	return book, nil
}
func (h *DBrepo) PostBooksRepo(c *gin.Context) (model.Book, error) {
	db := h.db
	book := model.Book{}

	// Extract product data from the request
	err := c.BindJSON(&book)
	if err != nil {
		log.Printf("failed to bind JSON to product struct. Original Error: %s", err)
		return model.Book{}, fmt.Errorf("error extracting the product data")
	}

	// Insert the product into the database
	_, err = db.Exec("INSERT INTO books(id,title,author,published_at,created_at,updated_at) VALUES(?,?,?,?,NOW(), NOW());",
		book.Id, book.Title, book.Author, book.PublishedAt)
	if err != nil {
		log.Printf("failed to insert product into the database. Product ID: %s. Original Error: %s", string(book.Id), err)
		return model.Book{}, fmt.Errorf("failed to insert the product into the database")
	}
	bookByID, err := h.getBookByID(string(book.Id))
	if err != nil {
		return model.Book{}, nil
	}
	return bookByID, nil

}
func (h *DBrepo) DeleteBooksRepo(c *gin.Context) error {
	db := h.db
	id := c.Param("id")

	// Delete the product from the database
	_, err := db.Exec("DELETE FROM books WHERE id=?;", id)
	if err != nil {
		log.Printf("Failed to delete product from the database. Product ID: %s. Original Error: %s", id, err)
		return fmt.Errorf("failed to delete the product from the database")
	}

	return nil

}
func (h *DBrepo) PutBooksRepo(c *gin.Context) (model.Book, error) {
	book := model.Book{}

	// Extract product data from the request
	err := c.BindJSON(&book)
	if err != nil {
		log.Printf("Failed to bind JSON to product struct. Original Error: %s", err)
		return model.Book{}, fmt.Errorf("error extracting the product data")
	}

	// Use product ID from the URL parameter
	idStr := (c.Param("id"))
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("error during id conversion during put: %s", err)
		return model.Book{}, fmt.Errorf("error during conversion of id into an integer")
	}
	book.Id = uint(idInt)

	// Update the product in the database
	db := h.db
	_, err = db.Exec(`
	UPDATE books 
	SET 
	title=COALESCE(?, title),
	author=COALESCE(?,author),
	published_at=COALESCE(?,published_at),
	updated_at=NOW() 
	WHERE id=?;`,
		book.Title, book.Author, book.PublishedAt, book.Id)
	if err != nil {
		log.Printf("Failed to update product in the database. Product ID: %s. Original Error: %s", string(book.Id), err)
		return model.Book{}, fmt.Errorf("failed to update the product in the database")
	}
	bookByID, err := h.getBookByID(string(book.Id))
	if err != nil {
		return model.Book{}, nil
	}
	return bookByID, nil
}
