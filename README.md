# Temperature Conversion API - Go

A simple and efficient RESTful API built with Go for converting temperatures between Celsius, Fahrenheit, and Kelvin.  
This project demonstrates structuring, parsing JSON, error handling, and building scalable web services with Go.

## Features

- Convert temperatures between **Celsius**, **Fahrenheit**, and **Kelvin**
- Receives input as JSON and returns well-structured JSON responses
- Handles invalid units and incorrect requests gracefully
- Easy to run, test, and integrate with any client (Python, curl, Postman, etc.)

## How it works

The API exposes a single endpoint (`/convert`) that accepts POST requests containing a temperature value and its unit.  
It validates the input and responds with all three temperature representations.

### **Request Example**

`POST /convert`  
Header: `Content-Type: application/json`  
Body:
{
"value": 100,
"unit": "celsius"
}



### **Response Example**

{
"celsius": 100,
"fahrenheit": 212,
"kelvin": 373.15
}



**Supported units:** `celsius`, `fahrenheit`, `kelvin`

When an error occurs (e.g., invalid unit):

{
"error": "Invalid unit. Use celsius, fahrenheit or kelvin."
}


## How to Run

1. **Run the server**
go run main.go


The server will start at `http://localhost:8080/convert`.

2. **Test using curl, Postman, or a simple Python script.**

## Example Usage with Curl

curl -X POST -H "Content-Type: application/json"
-d '{"value": 37, "unit": "celsius"}'
http://localhost:8080/convert



*Developed for learning purposes and as a portfolio project in Go. Feel free to adapt and improve!*
