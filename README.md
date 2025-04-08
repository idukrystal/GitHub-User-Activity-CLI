# GitHub-User-Activity-CLI

**GitHub User Activity CLI** is a command-line interface (CLI) tool written in Go, designed to fetch and display activity data from GitHub users. It allows you to view key information such as repositories, commits, pull requests, and other statistics directly from your terminal.

## Features

- Fetch GitHub user activity data.
- Display repositories, commits, pull requests, and issues a user has contributed to.
- Simple and intuitive CLI for monitoring your GitHub activity.
- Activity data is displayed in a structured format in the terminal.

## Installation

### Prerequisites

Ensure you have the following installed:

- [Go 1.x](https://golang.org/doc/install) or higher
- `git` (for cloning the repository)

### Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/idukrystal/GitHub-User-Activity-CLI.git
   ```

2. Navigate to the project directory:

   ```bash
   cd GitHub-User-Activity-CLI
   ```

3. Build the Go program:

   ```bash
   go build github_activity_cli
   ```
   
4. You can now run the program directly from the terminal.
   ```bash
   ./github_activity_cli <username>
   ```

