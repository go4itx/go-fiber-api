demo:
	go run cmd/demo/main.go --config=config/config.toml

demo-dev:
	fiber dev -a "--config=config/config.toml" -t "./cmd/demo"

admin:
	go run cmd/admin/main.go --config=config/config.toml

admin-dev:
	fiber dev -a "--config=config/config.toml" -t "./cmd/admin"
