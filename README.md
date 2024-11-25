# Go Tutorials

This repository contains a collection of tutorials and examples for learning the Go programming language. Each tutorial is designed to help you understand different aspects of Go, from basic syntax to advanced features.

## Contents

- **Introduction to Go**: Learn the basics of Go, including installation, setup, and your first Go program.
- **Go Syntax**: Understand the syntax and structure of Go programs.
- **Data Types**: Explore the various data types available in Go.
- **Control Structures**: Learn about loops, conditionals, and other control structures in Go.
- **Functions**: Understand how to define and use functions in Go.
- **Concurrency**: Dive into Go's powerful concurrency model with goroutines and channels.
- **Error Handling**: Learn how to handle errors effectively in Go.
- **Standard Library**: Explore some of the most commonly used packages in Go's standard library.
- **RSS Aggregator**: A project to aggregate RSS feeds.

## Getting Started

To get started with these tutorials, clone the repository and follow the instructions in each tutorial's README file.

```sh
git clone https://github.com/johnsiver94/go-tutorials.git
cd go-tutorials
```

## RSS Aggregator

The `rss-aggregator` project is designed to aggregate RSS feeds and provide an API to interact with the aggregated data.

### Features

- **User Authentication**: Middleware for authenticating users using API keys.
- **Feed Management**: Create, read, and delete RSS feeds.
- **Feed Follows**: Follow and unfollow RSS feeds.
- **Scraping**: Periodically scrape RSS feeds for new posts.
- **Error Handling**: Comprehensive error handling for API endpoints.

### Project Structure

- `main.go`: Entry point of the application.
- `handler_feed.go`: Handlers for feed-related endpoints.
- `handler_feed_follows.go`: Handlers for feed follow-related endpoints.
- `handler_readiness.go`: Handler for readiness check.
- `handler_user.go`: Handlers for user-related endpoints.
- `middleware_auth.go`: Middleware for user authentication.
- `internal/database`: Contains database queries and models.
- `scraper.go`: Logic for scraping RSS feeds.
- `rss.go`: Logic for parsing RSS feeds.
- `sqlc.yaml`: Configuration for SQLC (SQL code generator).

### Running the Project

1. Set up the environment variables in `.env` file.
2. Run the project:

```sh
cd rss-aggregator
go run main.go && ./rssagg
```

Happy coding!
