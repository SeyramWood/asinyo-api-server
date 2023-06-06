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
	@echo Stopping down Docker images if running...
	docker compose -f docker-compose-dev.yml down -v
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
prune-image:
	@echo Removing all images without at least one container associated to them...
	docker image prune -a
	@echo Docker images pruned!
prune-system:
	@echo Removing all images without at least one container associated to them...
	docker system prune -a
	@echo Docker images pruned!
new-ent:
	@echo Creating ent Model ...
	go run entgo.io/ent/cmd/ent new $(ent)
	@echo Model created successfully!
gen-ent:
	@echo Gnerating ent Model ...
	go generate ./ent
	@echo Model generated successfully!
jwt-cert:
	openssl genrsa -out mnt/cert/jwt_rsa 4096
	openssl rsa -in mnt/cert/jwt_rsa -pubout -out mnt/cert/jwt_rsa.pub