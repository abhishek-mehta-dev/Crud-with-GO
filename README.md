Here’s the full **README.md** file content you can drop directly into your repo 👇

````markdown
# crud-with-go

A simple **CRUD API** built with **Go** and **MySQL**.  
This project demonstrates how to build a production-ready REST API using Golang with database integration.  

---

## 🚀 Features
- Create, Read, Update, Delete users  
- MySQL database integration  
- Environment variable support (`.env`)  
- Clean project structure  

---

## 📦 Requirements
- Go 1.20+  
- MySQL 8+  
- Git  

---

## ⚙️ Setup Instructions

### 1. Clone the repository
```bash
git clone https://github.com/your-username/crud-with-go.git
cd crud-with-go
````

### 2. Setup environment variables

Create a `.env` file in the project root:

```env
DB_USER=root
DB_PASSWORD=your_password
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=crud_with_go
```

### 3. Create the database

Login to MySQL and run:

```sql
CREATE DATABASE crud_with_go;

USE crud_with_go;

CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  email VARCHAR(100) UNIQUE NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 4. Run the project

```bash
go run main.go
```

The server will start on:

```
http://localhost:8080
```

---

## 📡 API Endpoints

| Method | Endpoint    | Description     |
| ------ | ----------- | --------------- |
| GET    | /users      | Get all users   |
| GET    | /users/{id} | Get user by ID  |
| POST   | /users      | Create new user |
| PUT    | /users/{id} | Update user     |
| DELETE | /users/{id} | Delete user     |

---

## 🛠️ Tech Stack

* **Go** – backend language
* **MySQL** – relational database
* **Gorilla Mux / Gin** – router (depending on what you choose)
* **godotenv** – load `.env` config

---

## 🤝 Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

---

```
