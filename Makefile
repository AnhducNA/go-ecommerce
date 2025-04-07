APP_NAME = server

hello: 
	echo "Hello"

run: 
	go run ./cmd/${APP_NAME}/main.go