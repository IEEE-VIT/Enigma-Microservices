# Enigma Microservices
"Enigma" (by IEEE-VIT) implemented as a microservices architecture using Golang

> Note : This project is still in a development phase and was implemented as an experiment. This may not be used in production.

## Architecture
![Architecture](docs/architecture.jpg)

## Production Setup
To serve the application, follow the steps:

* Clone the repository

  ```
  $ git clone https://github.com/mayankshah1607/Enigma-Microservices.git
  ```
* Export the following environment variables, or create a .env file in the root directory of the project :
  ```
  AUTH_SERVICE_NAME=""
  AUTH_SERVICE_PORT=""
  ADMIN_SERVICE_NAME=""
  ADMIN_SERVICE_PORT=""
  SUBMISSION_SERVICE_NAME=""
  SUBMISSION_SERVICE_PORT=""
  DB_URI=""
  DB_NAME=""
  JWT_KEY=""
  ```

* ```
  $ cd Enigma-Microservices
  ```
  
 * Make sure you have [Docker](https://docs.docker.com/v17.09/engine/installation/) and [docker-compose](https://docs.docker.com/compose/install/) set up before executing the next command :
 
    ```
    $ docker-compose up --build
    ```

## Development Setup
To develop this project, simply clone it and `cd` into the root directory. Open your favourite text editor and start coding!

## API Documentation
### 1. Admin
This route creates a new Question in the database
```
POST /admin/submit
Request Body => JSON({
    text: String, 
    image_url: String,
    answer: String
   })
Response => JSON({
    Status: Boolean,
    Message: String
  })
```

This route deletes a question from the database
```
POST /admin/delete
Request Body => JSON({
    id: String
   })
Response => JSON({
    Status: Boolean,
    Message: String
  })
```

### 2. Auth
This route creates a new user
```
POST /auth/sign-up
Request body => JSON({
    name: String,
    email: String,
    university: String,
    password: String,
  })
Response => JSON({
  Status: Boolean,
  Message: String
})
```

This route authenticates an existing user and returns a cookie
```
POST /auth/sign-up
Request body => JSON({
    email: String
    password: String,
  })
Response => JSON({
  Status: Boolean,
  Message: String
})
```

### 3. Submission
This route accepts an answer for a given question, and checks if the answer is correct
```
POST /submission/submit
Request body => JSON({
    q_id: String,
    answer: String
  })
Response => JSON({
    Status: Boolean,
    Message: String
  })
```

