gen:
	- go run github.com/99designs/gqlgen generate

run:
	- go run .

.PHONY: gen run