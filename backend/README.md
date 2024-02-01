## Backend
### Getting Started
1. [Install Go](https://go.dev/doc/install)
2. [Install Docker](https://docs.docker.com/get-docker/)
    1. Run mongoDB container
        ```bash
        docker run -d -p 27017:27017 --name mongodb mongodb/mongodb-community-server
        ```
3. Install dependencies
    ```bash
    go mod get
    ```
4. Run the server
    ```bash
    go run main.go
    ```