run:
	go run cmd/main.go

clean:
	@echo "Remove old generated files, used only in debug mode"
	rm -rf test

debug: clean
	go run cmd/main.go -v=2 -logtostderr=true