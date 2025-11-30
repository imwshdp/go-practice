# Go Practice Projects

Welcome to the **go-practice** repository. This repo contains a collection of personal projects showcasing various basics, practices and http apps using Golang.

## 📂 Repository Structure

The repository follows a **one-folder-per-project** structure, making it easy to navigate and explore each project individually. Each project folder typically includes:

- **source code**: Core project files and implementation (.go files).
- **go.mod/go.sum**: Go module files for dependency management.
- **documentation (maybe)**: README or notes with a description of the project and usage instructions.
- **examples**: Sample usage or test files (if applicable).

## ⚙️ Running Projects

Each project may have a different setup. Please refer to the specific README.md file in each project folder for detailed instructions.

### For Simple Projects

If the project is a single Go file, you can run it directly using:

```shell
go run [filename].go
```

or run main.go with:

```shell
go run .
```

## 📚 List of Projects

Below is a brief overview of the projects included in this repository:

- **learning** folder:
- - **_basics_**: core concepts demonstrating Go fundamentals (io, structs, pointers, concurrency, etc.)
- - **_stepik-course_**: practices and homeworks from course of go web development from VK Team

<br>

- **backend-dev** folder:
- - **_http-based_**: simple demo setup of go api with postgresql docker container (mostly for training)
- - **_rest_**: first simple go rest api writing practice
- - **_e-commerce-api_**: e-com project with modern pre-production tech stack such as goose for migrations and pgx as postgres driver with sqlc repositories layer generation

<br>

- **extra** folder:
- - **_vs-typescript_**: some cases of comparing Golang with TypeScript language (sorry, I am front-end dev who migrated to full-stack community)
