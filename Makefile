run:
	go run cmd/main.go
watch:
	air
init_ent:
	go run entgo.io/ent/cmd/ent init $(ent)
gen_ent:
	go generate ./ent