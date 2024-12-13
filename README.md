## What kind of project is this?
This is a scalable API in which you can get data about books, get information about a specific book by id, record a new book, edit an existing book, delete existing books. A minimum AAA server (Authentication, Authorization, and Accounting Server) and JWT have also been added.

## What technologies have I used?
- Golang
  - Gin
  - Viper
  - Gomn
  - jwt-go
  - bcrypt
- PostgreSQL

## Why did I even start creating this project?
I created this project to study the Gin library, but also during the creation I learned new things for myself, such as: _PostgreSQL_, _.env_, _viper_, _gin_, _gomn_. This is a simple project that helped me learn a lot.

## How usage?
You can send requests:
- **GET** `/books/` - all info;
- **POST** `/books/` - if you send it with JSON information, then write a new book;
- **GET** `/books/:id` - info about a specific book;
- **PUT** `/books/:id` - update info about a specific book;
- **DELETE** `/books/:id` - delete info about a specific book.

Also you can request:
- **GET** `/user/me` - information about you;
- **POST** `/user/auth/register` - account registration;
- **POST** `/user/auth/login` - log in to your account;
- **POST** `/user/auth/refresh` - refresh the access token;
- **GET** `/user/admin/get_all_info` - getting information about all users;
- **POST** `/user/admin/change_role/:id` - change the role by id.