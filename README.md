# Golang Microservices for Product & Login Management  
This is a Microservices RESTful API built using Golang with Gin frameworks, GORM as ORM for MySQL, deployed using Docker compose alongside with message brokers using RabbitMQ

## Technologies  
![Golang](https://img.shields.io/badge/golang-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)  ![REST API](https://img.shields.io/badge/restapi-%23000000.svg?style=for-the-badge&logo=swagger&logoColor=white)  ![MySQL](https://img.shields.io/badge/mysql-%234479A1.svg?style=for-the-badge&logo=mysql&logoColor=white)  ![Redis](https://img.shields.io/badge/redis-%23DC382D.svg?style=for-the-badge&logo=redis&logoColor=white)  ![Docker](https://img.shields.io/badge/docker-%230087C6.svg?style=for-the-badge&logo=docker&logoColor=white)  ![RabbitMQ](https://img.shields.io/badge/rabbitmq-%23FF6600.svg?style=for-the-badge&logo=rabbitmq&logoColor=white)  


## Services
This microservices contains 3 seperate services, user login management, product management and message consumer.


### 1. Product Management

Features
* Create, Update, and Get Product with Message Publishing
* Create, and Get a Product Transaction with Message Publishing
* Create, Update, and Get Inventory with Message Publishing
* Redis Memory Cache to Improve performance
* One to Many relations for Product ID
* GORM-based relational database modeling
* Upload a File with file size and format validation
* Download file with mutex for lock file, and buffer for memory management

### 2. Golang API User Login Authorization

Features
* User registration with SHA-256 password hashing
* User login with JWT-based authentication
* Middleware for protected routes
* CRUD operations for users, posts, and tags
* Many-to-Many relationship between posts and tags
* GORM-based relational database modeling

### 3. Message Consumer
* Consume transaction, and login messages from RabbitMQ services

# EndPoint Documentation
# API Endpoints Documentation

| Method     | API Endpoint               | Description                                      | Services             |
|------------|----------------------------|--------------------------------------------------|-------------------|
| **POST**   | `localhost:8082/products`                | Create a new product                             | Product Management     |
| **GET**    | `localhost:8082/products/:id`            | Get details of a specific product by ID          | Product Management     |
| **PUT**    | `localhost:8082/products/:id`            | Update an existing product by ID                 | Product Management     |
| **DELETE** | `localhost:8082/products/:id`            | Delete a product by ID                           | Product Management     |
| **GET**    | `localhost:8082/inventories/:id`         | Get inventory details by ID                      | Product Management    |
| **PUT**    | `localhost:8082/inventories/:id`         | Update inventory details by ID                   | Product Management    |
| **POST**   | `localhost:8082/orders`                  | Create a new order                               | Product Management       |
| **GET**    | `localhost:8082/orders/:id`              | Get details of a specific order by ID            | Product Management       |
| **POST**   | `localhost:8082/upload`                  | Upload an Image of Product                              | Product Management       |
| **GET**    | `localhost:8082/download`              | Download an Image of Product            | Product Management      |
