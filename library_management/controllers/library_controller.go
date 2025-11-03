package controllers

import (
	"fmt"
	"library_management/models"
	"library_management/services"
)

func StartLibraryConsole() {
	lib := services.NewLibrary()
	lib.Members[1] = models.Member{ID: 1, Name: "Alice"}

	for {
		fmt.Println("\n--- Library Menu ---")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("0. Exit")
		fmt.Print("Enter choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id int
			var title, author string
			fmt.Print("Enter book ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter title: ")
			fmt.Scan(&title)
			fmt.Print("Enter author: ")
			fmt.Scan(&author)
			lib.AddBook(models.Book{ID: id, Title: title, Author: author})
			fmt.Println("Book added successfully!")

		case 2:
			var id int
			fmt.Print("Enter book ID to remove: ")
			fmt.Scan(&id)
			lib.RemoveBook(id)
			fmt.Println("Book removed successfully!")

		case 3:
			var bookID int
			fmt.Print("Enter book ID to borrow: ")
			fmt.Scan(&bookID)
			err := lib.BorrowBook(bookID, 1)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book borrowed successfully!")
			}

		case 4:
			var bookID int
			fmt.Print("Enter book ID to return: ")
			fmt.Scan(&bookID)
			err := lib.ReturnBook(bookID, 1)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book returned successfully!")
			}

		case 5:
			fmt.Println("Available books:")
			for _, b := range lib.ListAvailableBooks() {
				fmt.Printf("%d - %s by %s\n", b.ID, b.Title, b.Author)
			}

		case 6:
			fmt.Println("Borrowed books:")
			for _, b := range lib.ListBorrowedBooks(1) {
				fmt.Printf("%d - %s by %s\n", b.ID, b.Title, b.Author)
			}

		case 0:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}
