generate:
	@echo "Generate static files from assets"
	go run cmd/static/main.go assets pkg/generate

run: generate
	go run cmd/main/main.go -h 

clean:
	@echo "Remove old generated files, used only in debug mode"
	rm -rf test

debug: clean generate
	go run cmd/main/main.go -v=2 -logtostderr=true -f assets/api.json -o test