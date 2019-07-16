build:
	for CMD in `ls cmd`; do \
		go build -o bin/`echo $$CMD | cut -d'.' -f1` ./cmd/$$CMD; \
	done
test:
	go test -v ./...
deploy:
	sls deploy