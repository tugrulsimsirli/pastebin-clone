
# Pastebin Clone

This is a backend application that allows users to share text or code snippets securely, similar to the functionality of Pastebin.

## Features

- **JWT-based Authentication**: Secure user authentication using JWT, with the secret key in `.env`.
- **Snippet Management**: Full CRUD (Create, Read, Update, Delete) operations for snippets.
- **Public and Private Snippets**: Users can mark their snippets as private or public.
- **View Count Tracking**: Track how many times a snippet has been viewed.
- **Environment-based Configuration**: Configuration options through `.env` and `config.yml`.

## Technologies

- **Golang (Echo)**: Web framework for API development.
- **GORM**: ORM library for Golang.
- **PostgreSQL**: Database system for user and snippet data.
- **JWT Authentication**: For secure user access control.
- **Docker**: Containerization for environment consistency.

## Project Structure

```
.
├── cmd
│   └── main.go                 # Main entry point for the application
├── configs
│   └── config.yml              # YAML-based configuration settings (contains DB config)
├── .env                        # Contains JWT secret only
├── internal
│   ├── bootstrap
│   │   └── handler_bootstrapper.go # Initializes application handlers
│   ├── db
│   │   ├── db.go                 # Database connection setup
│   │   └── data-models
│   │       ├── user.go           # Database model for users
│   │       └── snippet.go        # Database model for snippets
│   ├── http
│   │   ├── docs                  # Swagger API documentation
│   │   ├── handlers
│   │   │   ├── auth_handler.go   # Handles authentication routes
│   │   │   ├── snippet_handler.go# Handles snippet routes
│   │   │   └── user_handler.go   # Handles user routes
│   │   ├── middlewares
│   │   │   └── jwt_middleware.go # Secures routes with JWT
│   │   ├── models                # Request and response models
│   │       ├── user.go           # User-related models
│   │       ├── auth.go           # Auth-related models
│   │       ├── snippet.go        # Snippet-related models
│   │       └── generic.go        # Generic response models
│   ├── mapper
│   │   └── mapper.go             # Logic for mapping data between layers
│   ├── repositories
│   │   ├── user_repository.go    # User data repository
│   │   ├── auth_repository.go    # Authentication repository
│   │   ├── snippet_repository.go # Snippet data repository
│   │   └── dto                   # Data transfer objects (DTOs)
│   │       ├── snippetDto.go     # Snippet DTO
│   │       └── userDto.go        # User DTO
│   ├── services
│   │   ├── auth_service.go       # Business logic for authentication
│   │   ├── snippet_service.go    # Business logic for snippets
├── Dockerfile                   # Dockerfile for containerizing the app
├── docker-compose.yml           # Docker Compose file for managing the API and PostgreSQL containers
├── Makefile                     # Contains swagger generation and build commands
└── README.md                    # Project documentation
```

## Installation

1. **Clone the repository**:

```bash
git clone https://github.com/tugrulsimsirli/pastebin-clone.git
cd pastebin-clone
```

2. **Configure environment variables**:

Ensure that both `.env` and `config.yml` are properly configured:

- `.env`:
```
JWT_SECRET_KEY=your_secret_key
```

- `config.yml` (example):
```yaml
db:
  host: "db"
  user: "pastebin"
  password: "pastebin"
  dbname: "pastebin"
  port: "5432"
  sslmode: "disable"
```

3. **Run the application using Makefile**:

```bash
make build
```

4. **Access the application**:

   - API documentation is available at `http://localhost:8080/swagger/index.html`

## Endpoints

### Auth Endpoints:
- `POST /api/v1/auth/register`: Register a new user.
- `POST /api/v1/auth/login`: Log in a user and receive a JWT token.
- `POST /api/v1/auth/refresh-token`: Refresh the JWT token.

### Snippet Endpoints:
- `GET /api/v1/snippet`: Retrieve all snippets for the authenticated user.
- `GET /api/v1/snippet/user/{userId}`: Retrieve all snippets created by a specific user.
- `GET /api/v1/snippet/{id}`: Retrieve a specific snippet by ID.
- `POST /api/v1/snippet`: Create a new snippet.
- `PATCH /api/v1/snippet/{id}`: Update a specific snippet by ID.
- `DELETE /api/v1/snippet/{id}`: Delete a specific snippet by ID.

### User Endpoints:
- `GET /api/v1/user`: Retrieve details for the authenticated user.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
