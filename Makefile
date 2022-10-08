run:
	go run cmd/main.go
prod-d:
	@echo Starting Docker images for production in detatch mode...
	docker compose -f docker-compose-prod.yml up --build -d
	@echo Docker images started!
prod:
	@echo Starting Docker images for production...
	docker compose -f docker-compose-prod.yml up --build
	@echo Docker images started!
prod-down:
	@echo Stopping Docker images ...
	docker compose -f docker-compose-prod.yml down -v
	@echo Docker images stopped!
restart:
	@echo Restarting Docker images in watch mode...
	docker compose -f docker-compose-dev.yml up --build
	@echo Docker images started!
watch:
	@echo Starting Docker images in watch mode...
	docker compose -f docker-compose-dev.yml up
	@echo Docker images started!
watch-d:
	@echo Starting Docker images in detatch mode...
	docker compose -f docker-compose-dev.yml up -d
	@echo Docker images started!
watch-down:
	@echo Stopping Docker images ...
	docker compose -f docker-compose-dev.yml down -v
	@echo Docker images stopped!
init-ent:
	@echo Creating ent Model ...
	go run entgo.io/ent/cmd/ent init $(ent)
	@echo Model created successfully!
gen-ent:
	@echo Gnerating ent Model ...
	go generate ./ent
	@echo Model generated successfully!