## What kind of project is this?
This is a scalable API in which you can get data about books, get information about a specific book by id, record a new book, edit an existing book, delete existing books

## What technologies have I used?
- Golang
  - Gin (library)
  - Viper (library)
  - Gomn (library)
- PosgreSQL

## Why did I even start creating this project?
I created this project to study the Gin library, but also during the creation I learned new things for myself, such as: _PostgreSQL_, _.env_, _viper_, _gin_, _gomn_. This is a simple project that helped me learn a lot.

## How usage?
You can send requests:
- **GET** `/books/` - all info;
- **POST** `/books/` - if you send it with JSON information, then write a new book;
- **GET** `/books/id[0-9]` - info about a specific book;
- **PUT** `/books/id[0-9]` - update info about a specific book;
- **DELETE** `/books/id[0-9]` - delete info about a specific book.
