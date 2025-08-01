Go Shopping Cart API
Hey there!  Welcome to the Go Shopping Cart API project.

This is a complete backend server for a simple e-commerce application, built from the ground up using Go and the awesome Gin framework. It handles everything you'd expect: creating user accounts, logging in securely with JWTs, managing a product inventory, and, of course, letting users add items to a shopping cart and place an order.

It's designed to be a solid foundation for any e-commerce website or a great learning project for understanding how backend systems work.

Features 
Secure User Authentication: Full user signup and login system using passwords hashed with bcrypt and stateless authentication with JSON Web Tokens (JWT).

Complete Shopping Flow: Users can add items to a personal cart and "check out" by converting that cart into a permanent order.

Protected API Routes: Critical endpoints (like adding to a cart or placing an order) are protected. You can't access them without a valid token, thanks to our auth middleware.

Product Management: Simple endpoints for adding new products to the store and viewing all available products.

Solid Database Foundation: Uses GORM as an ORM to communicate with a SQLite database, making data management clean and simple. It even handles creating the database schema for you automatically!

Tech Stack 
Go: The core programming language.

Gin: For routing and handling HTTP requests.

GORM: As the ORM for all database interactions.

SQLite: As the file-based database.

JWT-Go: For generating and validating JSON Web Tokens.

Bcrypt: For securely hashing user passwords.

Getting Started
Ready to run the project yourself? Just follow these steps.

Prerequisites
Make sure you have these installed on your machine first.

Go (version 1.22.0 or newer).

A C Compiler (like GCC): This is a requirement for the SQLite library we use.

On Windows, the easiest way is to install MinGW-w64.

On macOS, you can install Xcode Command Line Tools by running xcode-select --install.

On Linux (like Ubuntu), you can run sudo apt-get install build-essential.

Installation & Setup
Clone the project to your local machine.

git clone <your-repository-url>
cd shopping-cart-backend

Install the dependencies. This command fetches all the required packages for the project.

go mod tidy

Run the server! This is the magic command. The CGO_ENABLED=1 part is essential for SQLite to work.

On macOS or Linux:

CGO_ENABLED=1 go run main.go

On Windows (using PowerShell):

$env:CGO_ENABLED=1; go run main.go

You should see a  Server starting on port 8080 message. Your API is now live and ready to accept requests!

API Endpoint Guide 
Here's a full guide to all the available endpoints. You can use a tool like Postman or Thunder Client to interact with them.

Authentication
POST /api/signup
Creates a new user account.

Request Body:

{
    "username": "myuser",
    "password": "password123"
}

Success Response (200 OK):

{
    "message": "User created successfully"
}

POST /api/login
Logs in a user and returns a JWT to use for authenticated requests.

Request Body:

{
    "username": "myuser",
    "password": "password123"
}

Success Response (200 OK):

{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.very.long.string"
}

Products
POST /api/items
Adds a new product to the store.

Authentication: Bearer Token required.

Request Body:

{
    "name": "Cool T-Shirt",
    "price": 24.99
}

Success Response (200 OK): The full product object that was just created, including its new ID.

GET /api/items
Gets a list of all products available in the store.

Authentication: None needed.

Success Response (200 OK): An array of all product objects.

Shopping Cart
POST /api/carts
Adds a product to your personal shopping cart. If you don't have an active cart, one will be created for you.

Authentication: Bearer Token required.

Request Body:

{
    "item_id": 1
}

Success Response (200 OK): Your entire cart object, showing all the items inside it.

Orders
POST /api/orders
Converts your active shopping cart into a final order.

Authentication: Bearer Token required.

Request Body:

{
    "cart_id": 1
}

Success Response (200 OK): A success message along with the final order object.

{
    "message": "Order placed successfully!",
    "order": { ... }
}

GET /api/orders
Gets a list of all completed orders in the system.

Authentication: Bearer Token required.

Success Response (200 OK): An array of all order objects.