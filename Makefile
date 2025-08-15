server:
	nodemon --watch "./**/*.go" --ext go,json --signal SIGTERM --exec "cross-env APP_ENV=dev go run main.go"