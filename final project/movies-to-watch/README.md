## Prerequisites

Please download and install the following in order to run the app
* Nodejs - [https://nodejs.org/en/download/](https://nodejs.org/en/download/)
* npm
    ```sh
    npm install npm@latest -g
    ```
* Go - [https://go.dev/doc/install](https://go.dev/doc/install)


## Installation

#### Frontend
1. From the main folder, go to frontend folder
    ```sh
    cd frontend
    ```
2. Install npm packages
    ```sh
    npm install
    ```


## Run the App (Without Using Docker)

#### Backend
1. From the main folder, go to /backend folder
    ```sh
    cd backend
    ```
2. Run the backend
    ```sh
    go run main.go
    ```
#### Frontend
1. From the main folder, go to /frontend folder
    ```sh
    cd frontend
    ```
2. Run the frontend
    ```sh
    npm start
    ```
Go to [http://localhost:3000](http://localhost:3000) and start using the app!


## Run the App (Using Docker)

1. From the main folder, build the services
    ```sh
    docker-compose build
    ```
2. Create and start the containers
    ```sh
    docker-compose up
    ```
3. Go to [http://localhost:3000](http://localhost:3000) and start using the app!