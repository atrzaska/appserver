# appserver
Simple Single Page Application server written in Go Gin that can forward API calls to some backend

- forwards API calls from `/api/*` to `API_HOST` env var.
- serves `public` folder
- serves `public/index.html` if no route is found.
- optionally specify public folder as first argument to the program

## Installation

go get -u github.com/atrzaska/appserver

## Usage

This will serve files from `public` folder

    API_HOST=localhost:4000 appserver

You can also specify public folder

    API_HOST=localhost:4000 appserver dist
