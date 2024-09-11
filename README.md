# Getting Started

This project demonstrates a basic setup of a RESTful API with user management using Go, Gin, GORM, PostgreSQL, and Google Wire. It supports basic CRUD operations.


## 1. Installation

Use below command to install and move to root directory

```bash
git clone https://github.com/Abhishek-Shukla1989/Go--Hotel-Management.git
cd go--hotel-management
```

## 2. Install dependency
```bash
go mod tidy
```
## 3. Make sure Postgresql installed and running

```bash
createdb restdb
```
## 4. Setup env variable

```bash
db_dsn = your_database_url
POST = add your port i.e 8000/9000
```

## 2. Run the project
```bash
air // this will run the project 
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
