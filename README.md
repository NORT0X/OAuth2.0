# OAuth 2.0 Three-Legged-Authorization 

This project demonstrates a simple OAuth 2.0 flow, including an **Authorization Server**, **Resource Server**, and **Client Server**. It allows you to authenticate using OAuth 2.0, check access tokens, and access protected resources through the Resource Server.

## Project Structure

- **auth-server**: The Authorization Server responsible for issuing tokens.
- **resource-server**: The Resource Server responsible for serving protected resources.
- **client-server**: The Client Server that interacts with the Authorization Server to get the access token and access protected resources.

## How to Run with Docker

1. Clone the repository:

    ```bash
    git clone https://github.com/NORT0X/OAuth2.0.git
    cd OAuth2.0
    ```

2. Build and run the Docker containers:

    ```bash
    docker-compose up --build
    ```

This will build the images for the Authorization Server, Resource Server, and Client Server and start them in separate containers.

To bring down the containers:

```bash
docker-compose down
```

## How to Test

1. Send a GET request to localhost:8081/login?username=test to initiate the authentication and obtain the access token:

    ```bash
    curl -X GET "http://localhost:8081/login?username=test"
    ```
2. Finally, use the obtained access token to access a protected resource from the Resource Server. Replace <access_token> with the actual token you received: 

    ```bash
    curl -X GET "http://localhost:8081/resource/get" \
    -H "Authorization: Bearer <access_token>"
    ```

