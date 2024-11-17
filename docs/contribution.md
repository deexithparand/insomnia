# Contributing to Insomnia

Thank you for considering contributing to Insomnia! This guide will help you understand our project structure and guide you through the contribution process. 

```
insomnia/
├── cmd/                       # Command definitions for CLI
│   ├── root.go                # Base/root command
│   ├── workspace/             # Commands for workspaces
│   │   ├── workspace.go       # Parent command
│   │   ├── list.go            # Subcommand: list workspaces
│   │   ├── create.go          # Subcommand: create workspace
│   │   ├── connect.go         # Subcommand: connect workspace
│   │   ├── delete.go          # Subcommand: delete workspace
│   ├── endpoint/              # Commands for endpoints
│   │   ├── endpoint.go        # Parent command
│   │   ├── list.go            # Subcommand: list endpoints
│   │   ├── create.go          # Subcommand: create endpoint
│   │   ├── delete.go          # Subcommand: delete endpoint
├── config/                    # Configuration utilities
│   ├── config.go              # Manage CLI state (e.g., current workspace)
├── database/                  # Database-related code
│   ├── models.go              # GORM models (Workspace, Endpoint, etc.)
│   ├── migrations/            # Database migration files
│   ├── seed.go                # Database seeding logic
├── utils/                     # Utility functions (optional)
│   ├── printer.go             # Common output functions for CLI
├── main.go                    # Entry point
├── go.mod                     # Dependencies
├── go.sum                     # Dependency checksums
```

---

## Project Structure

The project is organized as follows:

### Root Directory
- **`main.go`**: The entry point for the CLI application.
- **`go.mod`**: Dependency management file for Go modules.
- **`go.sum`**: Checksums for the project's dependencies.

---

### Key Directories

#### 1. **`cmd/`**
This folder contains the command definitions for the CLI. Each subfolder corresponds to a specific feature or resource.

- **`root.go`**: Defines the base/root command for the CLI.
- **`workspace/`**: Commands related to workspaces.
  - **`workspace.go`**: The parent command for managing workspaces.
  - **`list.go`**: Subcommand to list all workspaces.
  - **`create.go`**: Subcommand to create a new workspace.
  - **`connect.go`**: Subcommand to connect to a workspace.
  - **`delete.go`**: Subcommand to delete a workspace.
- **`endpoint/`**: Commands related to endpoints.
  - **`endpoint.go`**: The parent command for managing endpoints.
  - **`list.go`**: Subcommand to list all endpoints.
  - **`create.go`**: Subcommand to create a new endpoint.
  - **`delete.go`**: Subcommand to delete an endpoint.

---

#### 2. **`config/`**
Manages the configuration utilities for the CLI.

- **`config.go`**: Handles CLI state, such as the current active workspace.

---

#### 3. **`database/`**
Contains all database-related logic.

- **`models.go`**: Defines the GORM models, such as `Workspace` and `Endpoint`.
- **`migrations/`**: Stores database migration files.
- **`seed.go`**: Contains logic for seeding initial data into the database.

---

#### 4. **`utils/`**
Optional utilities to keep the codebase clean and reusable.

- **`printer.go`**: Common functions for output formatting and printing to the terminal.

---

## How to Contribute

We welcome contributions in the form of bug reports, feature requests, and pull requests.

### 1. **Fork the Repository**
Start by forking the repository to your GitHub account.

### 2. **Clone Your Fork**
```bash
git clone https://github.com/<your-username>/insomnia.git
cd insomnia
```

### 3. **Create a New Branch**
Follow the naming conventions for your branch:
- Bug fix: `fix/<issue-description>`
- Feature: `feature/<feature-name>`
- Hotfix: `hotfix/<hotfix-name>`

Example:
```bash
git checkout -b feature/add-endpoint-edit
```

### 4. **Make Changes**
- Follow the folder structure conventions.
- Keep your code clean and document functions where needed.

### 5. **Test Your Changes**
Before submitting, ensure that your changes work as expected:
- Test CLI commands manually.
- Verify database migrations and seeds.

### 6. **Commit and Push**
```bash
git add .
git commit -m "Add feature: Edit endpoint"
git push origin feature/add-endpoint-edit
```

### 7. **Submit a Pull Request**
- Go to the original repository on GitHub.
- Click on "New Pull Request."
- Provide a clear description of the changes.

---

## Coding Guidelines

1. **Write Clean Code**  
   Follow Go's best practices and conventions.
   
2. **Keep Commands Modular**  
   Each command should focus on a single responsibility.

3. **Database Models**  
   Ensure all models are defined in `database/models.go` and linked to migrations.

4. **Document Your Code**  
   Add comments for functions and complex logic.

5. **Testing**  
   Test commands and database functionality before submitting.

---

## Issues and Support

If you encounter any issues or have questions, feel free to open an issue on GitHub. We are happy to help!

---

Thank you for contributing to Insomnia! 🚀