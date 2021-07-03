# appserver
Simple Single Page Application server written in Go Gin that can forward API calls to some backend

- forwards API calls from `/api/*` to `API_HOST` env var.
- serves `public` folder
- serves `public/index.html` if no route is found.
