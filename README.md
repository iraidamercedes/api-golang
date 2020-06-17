# :star2: About
This is an API service made with Go using cockroachDB.

## :ballot_box_with_check: Requirements.

Before deployment, be sure to have Docker and docker-compose installed in your computer.

I you already have it, install the cockroach image with the command: ```docker pull cockroachdb/cockroach```

# :crystal_ball:Use

To deploy this app follow these steps:

1. Run this image using docker-compose command

    ```bash
    $ docker-compose build && docker-compose up
    ```
2. Run the main file:

    - Inside docker container:
    ```bash
    $ docker exec -it golang_app bash
    root@65664456464/:go$ go run main.go
    ```
    - In you computer
    ```bash
    docker exec -it golang_app bash -c "go run main.go"
    ```

Navigate to localhost:8000 to see the API working.