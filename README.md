# OrderHub API

OrderHub API is a RESTful service built with Go (Golang) and Chi router. This API manages orders with functionalities such as creation, retrieval, update, and deletion. The application utilizes Redis as a database for storing order information and implements pagination for improved data handling.

## Features

- **Order Management:** CRUD operations for orders - creation, listing, retrieval by ID, updating, and deletion.
- **Pagination:** Efficiently handle large datasets with paginated results.
- **Redis Database:** Utilize Redis as the backend database for storing order information.
- **Dockerized Application:** Containerized using Docker for easy deployment and scalability.
- **Data Automation:** A Python script is included to automate data addition to the Redis database.
- **Deployment:** The API is deployed and accessible at [https://orderhub-api.onrender.com/orders](https://orderhub-api.onrender.com/orders). The Redis instance is also deployed in the cloud.

## Endpoints

- **Order Create:** `/orders/create` - Create a new order.
- **Order Get by ID:** `/orders/{id}` - Retrieve details of a specific order by ID.
- **Order Update by ID:** `/orders/{id}` - Update details of a specific order by ID.
- **Order Delete by ID:** `/orders/{id}` - Delete a specific order by ID.

## Prerequisites

- Go (Golang)
- Redis
- Docker
- Python (for data automation script)

## Installation and Usage

### Running Locally

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/OrderHub-API.git
    cd OrderHub-API
    ```

2. Set up Redis and ensure it's running.

3. Build and run the Go application:

    ```bash
    go build -o orderhub
    ./orderhub
    ```

4. Access the API endpoints locally at `http://localhost:PORT`.

### Docker Deployment

1. Build the Docker image:

    ```bash
    docker build -t orderhub-api .
    ```

2. Run the Docker container:

    ```bash
    docker run -p HOST_PORT:CONTAINER_PORT orderhub-api
    ```

3. Access the API endpoints via the specified port on the host machine.

## Data Automation

A Python script (`script.py`) is included to automate data addition to the Redis database. Run the script to add data to the database.

```bash
python3 data/script.py
