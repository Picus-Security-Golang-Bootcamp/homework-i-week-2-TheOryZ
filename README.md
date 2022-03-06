## Homework | Week 2

### Book list application with Go

This application has two functions.
1. Get all book list
2. Fetch detail information by book name


### list command
```
go run main.go list
```
This command is used to list all books.

### search command 
```
go run main.go search <bookName>
go run main.go search Lord of the Ring: The Return of 
```
This command is used to searching books

There is a .json file in the application where the book information is kept.

```json
{
    "Books": [
        {
            "Title": "Dune",
            "Author": "Frank Herbet",
            "Release_Date":"1965"
        },
        {
            "Title": "The Lord of The Rings",
            "Author": "J. R. R. Tolkien",
            "Release_Date": "1955"
        },
        {
            "Title": "1984",
            "Author": "George Orwell",
            "Release_Date": "1949"
        },
        {
            "Title": "Brave New World",
            "Author": "Aldous Huxley",
            "Release_Date": "1932"
        },
        {
            "Title": "Foundation",
            "Author": "Isaac Asimov",
            "Release_Date": "1942"
        }
    ]
}
```