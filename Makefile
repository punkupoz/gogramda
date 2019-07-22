build:
	rm bin/* ; \
	for CMD in `ls cmd`; do \
		go build -o bin/`echo $$CMD | cut -d'.' -f1` ./cmd/$$CMD; \
	done
test:
	go test -v ./...
deploy:
	sls deploy function -f graphql
log:
	sls logs -t -f graphql
thuxem:
	bin/graphql