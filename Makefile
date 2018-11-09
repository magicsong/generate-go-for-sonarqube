generate:
	@echo "Generate static files from assets"
	go run cmd/static/main.go assets pkg/generate

run: generate
	go run cmd/main/main.go -h 

clean:
	@echo "Remove old generated files, used only in debug mode"
	rm -rf test

debug: clean generate
	go run cmd/main/main.go -logtostderr=true -v=1 -f assets/api.json -o test  -e "http://192.168.98.8:9000/api"

examples:
	go run cmd/response/main.go -v=2 -logtostderr=true -f assets/api.json -e "http://192.168.98.8:9000/api"