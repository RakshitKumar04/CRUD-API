# CRUD API

This project is a simple RESTful API built using Go, which allows users to manage a collection of movies. The API supports basic CRUD (Create, Read, Update, Delete) operations and is built using the `gorilla/mux` router.

## Getting Started

### Prerequisites

- Go 1.16 or later
- A terminal or command prompt
- Postman

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/RakshitKumar04/CRUD-API.git
   cd CRUD-API
   ```

2. **Install dependencies**:

   The project uses the `gorilla/mux` package. You can install it by running:

   ```bash
   go install github.com/gorilla/mux@latest
   ```

3. **Run the application**:

   ```bash
   go run main.go
   ```

   The server will start on `http://localhost:8000`.

## API Endpoints

![IMG_0464](https://github.com/user-attachments/assets/60cdfca4-91bc-480c-957c-2558220d369a)

### 1. Get All Movies

- **Endpoint**: `/movies`
- **Method**: `GET`
- **Description**: Retrieves a list of all movies.
- **Response**: JSON array of movie objects.

### 2. Get a Single Movie

- **Endpoint**: `/movies/{id}`
- **Method**: `GET`
- **Description**: Retrieves a movie by its ID.
- **Response**: JSON object of the requested movie.

### 3. Create a New Movie

- **Endpoint**: `/movies`
- **Method**: `POST`
- **Description**: Adds a new movie to the collection.
- **Request Body**: JSON object containing `isbn`, `title`, and `director` (which includes `firstname` and `lastname`).
- **Response**: JSON object of the newly created movie.

### 4. Update a Movie

- **Endpoint**: `/movies/{id}`
- **Method**: `PUT`
- **Description**: Updates an existing movie by its ID.
- **Request Body**: JSON object containing the updated `isbn`, `title`, and `director`.
- **Response**: JSON object of the updated movie.

### 5. Delete a Movie

- **Endpoint**: `/movies/{id}`
- **Method**: `DELETE`
- **Description**: Deletes a movie by its ID.
- **Response**: JSON array of the remaining movies.

## Testing the API with Postman

Postman is an excellent tool for testing and interacting with APIs. You can use Postman to test all the endpoints provided by this API.

1. **Import the API**: You can manually create requests in Postman or import a collection if available.

2. **Test Endpoints**: 
   - Use the `GET` request to retrieve movies.
   - Use the `POST` request to add new movies.
   - Use the `PUT` request to update existing movies.
   - Use the `DELETE` request to remove movies.

3. **Request Headers**: Ensure the `Content-Type` header is set to `application/json` for requests that send or receive JSON data.

4. **Request Body**: For `POST` and `PUT` requests, use the "raw" option in the body tab of Postman, and select `JSON` as the format.

## Project Structure

```bash
crud-api/
│
├── main.go       # Main application file
├── README.md     # This file
└── go.mod        # Go module file for dependency management
└── go.sum        # Checksum file for dependencies
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- The `gorilla/mux` package for providing the routing functionality.
