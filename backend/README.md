## Backend
### Getting Started
1. [Install Go](https://go.dev/doc/install)
2. [Install Docker](https://docs.docker.com/get-docker/)
    1. Run mongoDB container
        ```bash
         docker run --name mongo -d -p 27017:27017 mongodb/mongodb-community-server:latest
        ```
3. Install dependencies
    ```bash
    go mod get
    ```
   
4. Copy the `.env.example` file to `.env` and fill in the fields to match your setup
    ```bash
    # For Mac/Linux
    cp .env.example .env
   
    # For Windows
    copy .env.example .env
    ```
5. Run the server
    ```bash
    go run main.go
    ```