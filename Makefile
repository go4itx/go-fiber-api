demo:
	go run cmd/demo/main.go --config=config/config.toml

demo-docs:
	swag i -g ./internal/demo/controller/init.go --exclude internal/im,internal/admin  -o ./docs/demo

im:
	go run cmd/im/main.go --config=config/config.toml

im-docs:
	swag i -g ./internal/im/controller/init.go --exclude internal/demo,internal/admin  -o ./docs/im
	
cli:
	go run cmd/cli/main.go --name=project --module=all