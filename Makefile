.PHONY: build clean deploy gomodgen


GO=env GOOS=linux go build 
GOFLGAS=-s -w
GOSRC=$(wildcard */handler.go)

$(info GOSRC = $(GOSRC))


build:
	export GO111MODULE=on
	@ for src in $(GOSRC) ; do \
		target=$$(dirname $$src); \
		echo $(GO) -ldflags=$(GOFLAGS) -o bin/$$target $$src;\
		$(GO) -ldflags=$(GOFLAGS) -o bin/$$target $$src;\
	done


clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh