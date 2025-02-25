package bookModel

import "strings"

type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// TODO DBに移動する
var books = []Book{
	{ID: 1, Name: "React入門", Status: "在庫あり"},
	{ID: 2, Name: "TypeScript入門", Status: "貸出中"},
	{ID: 3, Name: "Next.js入門", Status: "返却済"},
}

// GetAll ブックを取得するという振る舞いはモデルに対しての直接的な振る舞いではない為レシーバーにしない
func GetAll(k string) []Book {
	if k == "" {
		return books
	}

	var filteredBooks = make([]Book, 0)
	for _, book := range books {
		if strings.Contains(book.Name, k) {
			filteredBooks = append(filteredBooks, book)
		}
	}

	return filteredBooks
}

func (b *Book) Add() Book {
	b.ID = len(books) + 1
	b.Status = "在庫あり"

	books = append(books, *b)

	return *b
}

func (b *Book) Update() Book {
	for _, book := range books {
		if book.ID == b.ID {
			b.Name = book.Name
		}
	}

	for _, book := range books {
		if book.ID == b.ID {
			book = *b
		}
	}

	return *b
}
