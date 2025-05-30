# Golang API Product Management

This is a RESTful API built using Golang with the Gin framework, GORM as the ORM for MySQL.

### Features
* Create, Update, and Get Product
* Create, and Get a Product Transaction
* Create, Update, and Get Inventory
* Redis Memory Cache to Improve performance
* One to Many relations for Product ID
* GORM-based relational database modeling
* Upload a File with file size and format validation
* Download file with mutex for lock file, and buffer for memory management

# Technologies
![Golang](https://img.shields.io/badge/golang-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)  ![REST API](https://img.shields.io/badge/restapi-%23000000.svg?style=for-the-badge&logo=swagger&logoColor=white)  ![MySQL](https://img.shields.io/badge/mysql-%234479A1.svg?style=for-the-badge&logo=mysql&logoColor=white)  ![Redis](https://img.shields.io/badge/redis-%23DC382D.svg?style=for-the-badge&logo=redis&logoColor=white)  


Uses golang as main frameworks for build an API, with RESTful API for communication with mySQL database

* Golang (Gin-Gonic framework) 
* MySQL (Database)
* GORM (ORM for database operations)
* Mutex (Locking file on download)
* Buffer (Memory management for download file)
* Redis Memory Cache

# API Endpoints Documentation

| Method     | API Endpoint               | Description                                      | Table             |
|------------|----------------------------|--------------------------------------------------|-------------------|
| **POST**   | `/products`                | Create a new product                             | Products Table     |
| **GET**    | `/products/:id`            | Get details of a specific product by ID          | Products Table     |
| **PUT**    | `/products/:id`            | Update an existing product by ID                 | Products Table     |
| **DELETE** | `/products/:id`            | Delete a product by ID                           | Products Table     |
| **GET**    | `/inventories/:id`         | Get inventory details by ID                      | Inventory Table    |
| **PUT**    | `/inventories/:id`         | Update inventory details by ID                   | Inventory Table    |
| **POST**   | `/orders`                  | Create a new order                               | Orders Table       |
| **GET**    | `/orders/:id`              | Get details of a specific order by ID            | Orders Table       |
| **POST**   | `/upload`                  | Upload an Image of Product                              | -       |
| **GET**    | `/download`              | Download an Image of Product            | -      |


# Database Structure
```golang
type Produk struct {
	gorm.Model
	ID        uint   `json:"id"`
	Nama      string `json:"nama"`
	Deskripsi string `json:"deskripsi"`
	Harga     int    `json:"harga"`
	Kategori  string `json:"kategori"`
}

type Inventaris struct {
	gorm.Model
	IDProduk uint   `json:"id_produk"`
	Jumlah   int    `json:"jumlah"`
	Lokasi   string `json:"lokasi"`
}

type Pesanan struct {
	gorm.Model
	IDPesanan uint   `json:"id_pesanan"`
	IDProduk  uint   `json:"id_produk"`
	Jumlah    int    `json:"jumlah"`
	Tanggal   string `json:"tanggal"`
}
```

# Product Management Golang API Starter Guide

This guide will help you set up and run the Golang API using Gin-Gonic, GORM, and MySQL.

## Prerequisites

Ensure you have the following installed:

- [Golang](https://go.dev/dl/) (latest version)
- [MySQL](https://dev.mysql.com/downloads/)
- [Git](https://git-scm.com/)

## Initialization

Follow these steps to set up the project:

### 1. Initialize the Go Module
Run the following command in the project directory:

```sh
go mod init management-inventaris
```

### 2. Install Dependencies
Install the required packages:

```sh
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u github.com/joho/godotenv
go get github.com/redis/go-redis/v9  
```

### 3. Configure Database
Create a .env file in the root directory and add your database credentials:

```env
DB_USER=root
DB_PASSWORD=password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=management_inventaris
```

### 4. Start Redis on Port: 6379
Redis memory cache for resources management. Start using Docker / WSL

* Using Docker
```sh
 docker run --name redis-container -d -p 6379:6379 redis   
```
* Using WSL
```sh
sudo apt update
sudo apt install redis-server
sudo service redis-server start
```


### 5. Run API
```sh
go run main.go
```

### 6. Set Up Data Sample
Here is mySQL query for set up a data sample  [Here](https://github.com/shofwanshiddiq/product-management-api/blob/main/script_insert_database.sql)

### 7. Testing

Import this Postman Collection for Testing in Postman
https://github.com/shofwanshiddiq/product-management-api/blob/main/Assignment_Day_26%20.postman_collection.json


