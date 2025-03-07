# OrderPack

OrderPack is a Golang-based API that calculates the optimal distribution of complete packs to fulfill an order. The application follows these rules:

1. Only whole packs can be sent. Packs cannot be broken open.
2. Within the constraints of rule 1, ship out the least amount of items to fulfill the order.
3. Within the constraints of rules 1 and 2, use as few packs as possible to fulfill each order.

For example, with the default pack sizes (250, 500, 1000, 2000, 5000):

**Test Case 1:**  
Items Ordered: 1  
Expected Output: 1 x 250

**Test Case 2:**  
Items Ordered: 250  
Expected Output: 1 x 250

**Test Case 3:**  
Items Ordered: 251  
Expected Output: 1 x 500

**Test Case 4:**  
Items Ordered: 501  
Expected Output: 1 x 500, 1 x 250

**Test Case 5:**  
Items Ordered: 12001  
Expected Output: 2 x 5000, 1 x 2000, 1 x 250

The API can also handle custom pack sizes. For instance, with pack sizes of 23, 31, and 53 for an order of 500,000 items, the desired output is:  
2 x 23, 7 x 31, 9429 x 53

---

## Table of Contents

- [Features](#features)
- [Architecture](#architecture)
- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
- [Deployment](#deployment)
- [License](#license)

---

## Features

- **Optimal Pack Distribution:** Calculates the minimum total items shipped while using the fewest packs.
- **Flexible Pack Sizes:** Works with both default and custom pack sizes.
- **RESTful API:** Built using Golang with a clean, modular architecture.

---


---

## Installation

### Prerequisites

- Go 1.24.1
- Git
- Docker for containerized deployment

### Steps

1. Clone the repository:
   ```
    git clone https://github.com/bebefabian/orderpack.git
      ```
2. Install dependencies:
      ```
   go mod tidy
      ```
3. Run the application:
      ```
    go run cmd/main.go


## Deployment

### Docker Deployment

1. Build the Docker image:
   ```
   make docker-build

2. Run the Docker container:
   ```
   make docker-run

### Health Check Endpoint
```
curl -X GET http://localhost:8080/packs
```
### Get Available Pack Sizes
```
curl -X GET http://localhost:8080/packs

```
### Update Pack Sizes
```
curl -X POST http://localhost:8080/packs -H "Content-Type: application/json" -d '[250,500,1000,2000,5000]'
```
Response
```json
{"message": "Pack sizes updated", "packs": [5000,2000,1000,500,250]}
```
### Update Pack Sizes
```
curl -X GET "http://localhost:8080/calculate?quantity=12001"
```
Response
```json
{
  "orderQuantity": 12001,
  "packs": [
    { "packSize": 5000, "quantity": 2 },
    { "packSize": 2000, "quantity": 1 },
    { "packSize": 250, "quantity": 1 }
  ]
}
```

# OrderPack UI

OrderPack UI is a simple React-based frontend that interacts with the OrderPack backend API to calculate the optimal distribution of complete packs to fulfill an order. This application allows users to view, update, and calculate pack distributions in a user-friendly interface.
This UI app does not have any logic or validation on it. It's purpose is just to interact with the backend.
## Installation

### Prerequisites
- Node.js (version 14 or higher)
- npm (Node Package Manager)
### Steps

1 **Move to Ui dir:**
```
cd ui/orderpack-ui
```

2 **Install dependencies:**
```
npm install
```
## Usage

### Local Development
1. Start the development server
```
npm start
```

### Environment Variables
- The frontend uses an environment variable `REACT_APP_API_BASE_URL` to determine the backend API URL.
- Create a `.env` file in the project root with the following content (for local development):
  REACT_APP_API_BASE_URL=http://localhost:8080