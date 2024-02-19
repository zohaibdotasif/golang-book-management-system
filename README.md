# golang-book-management-system
A simple book store management system which allows to create, read, update, and delete books

# routes
1. `GET /books` to get the list of all the books
2. `GET /books/{bookId}` to get a book by id
3. `POST /books` to create a new book
4. `PUT /books/{bookId}` to update fields of an existing book
5. `DELETE /books/{bookId}` to delete an existing book

# dependencies
1. [MySql](https://hub.docker.com/_/mysql) `v8.3.0` - RDBMS
