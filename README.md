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
- Docker & docker-compose

## Why did I even start creating this project?
I created this project to study the Gin library, but also during the creation I learned new things for myself, such as: _PostgreSQL_, _.env_, _viper_, _gin_, _gomn_. This is a simple project that helped me learn a lot.

## Fast start
Copying the project from the repository
```bash
git clone https://github.com/kreipikc/golang-book-api.git
```

Moving on to the project
```bash
cd <your_path>/golang-book-api
```

Copying data from `.example.env` to `.env`and then change the necessary parameters in `.env` (we can do this in `.example.env` before copying - it doesn't matter)
```bash
cat .example.env > .env
```

Next, we build the project using `docker-compose`
```bash
docker-compose up --build -d
```

Well done!

## Description endpoints
You can send requests for module `books`:
- **GET** `/books/` - all info;
- **POST** `/books/` - if you send it with JSON information, then write a new book;
- **GET** `/books/:id` - info about a specific book;
- **PUT** `/books/:id` - update info about a specific book;
- **DELETE** `/books/:id` - delete info about a specific book.

Also module `user`:
- **GET** `/user/me` - information about you;
- **POST** `/user/auth/register` - account registration;
- **POST** `/user/auth/login` - log in to your account;
- **POST** `/user/auth/refresh` - refresh the access token;
- **GET** `/user/admin/get_all_info` - getting information about all users;
- **POST** `/user/admin/change_role/:id` - change the role by id.

## Structure ptoject
```
golang-book-api         
│   .env                  # Main .env
│   .example.env          # Example .env for github
│   .gitignore            
│   docker-compose.yml 
│   Dockerfile
│   go.mod
│   go.sum
│   README.md
│
├───cmd                   # Main applications for this project
│       main.go           # Main file
│
└───pkg                       # Library code that's ok to use by external applications
    ├───books                 # Code for module 'books'
    │       add_book.go
    │       controller.go     # Controller with all endpoints
    │       delete_book.go
    │       get_book.go
    │       get_books.go
    │       update_book.go
    │
    ├───common
    │   ├───database          # Module for db 
    │   │       db.go         # Database init code
    │   │
    │   └───models            # Models for tables from db
    │           Book.go
    │           User.go
    │
    └───users                 # Code for module 'users'
            all_user_info.go
            auth.go
            change_role.go
            controller.go     # Controller with all endpoints
            login.go
            me.go
            refresh.go
            register.go
```