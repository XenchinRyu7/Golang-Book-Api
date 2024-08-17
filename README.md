# Golang Book API

This is a simple CRUD API built with Go for managing books data. The API allows you to perform operations such as creating, reading, updating, and deleting book records. This project also includes database migrations and seed data.

## Features
- CRUD operations for books
- Database migrations
- Seed data for initial setup
- SQL database integration

## Table of Contents
1. [Requirements](#requirements)
2. [Installation](#installation)
3. [Database Setup](#database-setup)
4. [Running the Project](#running-the-project)
5. [API Endpoints](#api-endpoints)

## Requirements

Before you start, ensure you have the following installed on your machine:

- Go 1.20+
- SQL Database (MySQL, PostgreSQL, or SQLite)
- Git

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/golang-book-api.git
   cd golang-book-api
2. install the dependencies "$go mod tidy"
   
## Database Setup, 
1. Database name db_book
2. go run main.go migrate
3. go run main.go seed

## Running The Project
go run main.go

## API Endpoints
Here are the available API endpoints:

GET /books: List all books
GET /books/
: Get a book by ID
POST /books: Create a new book
PUT /books/
: Update a book by ID
DELETE /books/
: Delete a book by ID
