package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Books struct {
	Books []Book `json:"Books"`
}
type Book struct {
	Title        string `json:"Title"`
	Author       string `json:"Author"`
	Release_Date string `json:"Release_Date"`
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Please enter a command.")
	} else {
		command := strings.ToLower(args[0])
		commandValid := ValidationCommand(command)
		if !commandValid {
			fmt.Println("Please enter a valid command.")
		} else {
			if command == "list" {
				var books Books
				GetAllBoks(&books)
				for i := 0; i < len(books.Books); i++ {
					fmt.Println(strings.Repeat("*", 25))
					fmt.Println("Book Title: " + books.Books[i].Title)
					fmt.Println("Book Author: " + books.Books[i].Author)
					fmt.Println("Book Release Date: " + books.Books[i].Release_Date)
				}
				fmt.Println(strings.Repeat("*", 25))
				fmt.Println("Books count: ", len(books.Books))
			} else if command == "search" && len(args) < 2 {
				fmt.Println("Please enter a valid book title.")
			} else {
				bookTitle := args[1:]
				_bookTitle := strings.Join(bookTitle, " ")
				book, err := GetByTitle(_bookTitle)
				if err != nil {
					fmt.Println("An error occurred while fetching the book detail. Error :" + err.Error())
				} else if len(book.Title) <= 0 {
					fmt.Println(strings.Repeat("*", 25))
					fmt.Println("No information was found for the book you were looking for.")
					fmt.Println(strings.Repeat("*", 25))
				} else {
					fmt.Println(strings.Repeat("*", 25))
					fmt.Println("Book Title: " + book.Title)
					fmt.Println("Book Author: " + book.Author)
					fmt.Println("Book Release Date: " + book.Release_Date)
					fmt.Println(strings.Repeat("*", 25))
				}
			}
		}
	}
}

//ValidationCommand: Function that checks the validity of the command entered by the user. Returns (true/false).
func ValidationCommand(command string) bool {
	switch command {
	case "list":
		return true
	case "search":
		return true
	default:
		return false
	}
}

//GetAllBooks: Function that returns all books
func GetAllBoks(books *Books) {
	jsonFile, err := os.Open("db/books.json")
	if err != nil {
		fmt.Println(err)
	} else {
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &books)
	}
	defer jsonFile.Close()
}

//GetByTitle: Returns the details of the entered book.
func GetByTitle(s string) (Book, error) {
	jsonFile, err := os.Open("db/books.json")
	var book Book
	if err != nil {
		return book, err
	} else {
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var books Books
		json.Unmarshal(byteValue, &books)
		for i := 0; i < len(books.Books); i++ {
			if strings.ToLower(books.Books[i].Title) == strings.ToLower(s) {
				book.Title = books.Books[i].Title
				book.Author = books.Books[i].Author
				book.Release_Date = books.Books[i].Release_Date
				break
			}
		}
	}
	defer jsonFile.Close()
	return book, nil
}
