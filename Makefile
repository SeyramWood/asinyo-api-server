run:
	go run cmd/main.go
prod-d:
	docker compose -f docker-compose-prod.yml up --build -d
prod:
	docker compose -f docker-compose-prod.yml up --build
prod-down:
	docker compose -f docker-compose-prod.yml down -v
watch-d:
	docker compose -f docker-compose-dev.yml up -d
watch:
	docker compose -f docker-compose-dev.yml up
watch-down:
	docker compose -f docker-compose-dev.yml down -v
init_ent:
	go run entgo.io/ent/cmd/ent init $(ent)
gen_ent:
	go generate ./ent