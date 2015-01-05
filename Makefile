
.PHONY: all
all: list

#-----------------------------------------
# list tasks

.PHONY: list
list:
	@grep -e "^[^.	 ].*:" Makefile

#-----------------------------------------
# prepare project

.PHONY: prepare
prepare:
	goop install
	mkdir -p .vendor/src/github.com/s-shin
	ln -s ../../../.. .vendor/src/github.com/s-shin/gobbs

#-----------------------------------------
# build

gobbs:
	goop exec go-bindata \
	-pkg="bindata" \
	-o="bindata/production.go" \
	-tags="production" \
	-ignore=".+\\.go" \
	./config/... ./tmpl/...
	goop go build -tags "production" gobbs.go

.PHONY: build
build: gobbs

#-----------------------------------------
# archive

gobbs.zip: build
	goop go run -tags=production script/init.go
	zip -r gobbs.zip gobbs production.sqlite3 static/

.PHONY: archive
archive: gobbs.zip
	goop go run -tags=production script/init.go
	zip -r gobbs.zip gobbs production.sqlite3 static/

#-----------------------------------------
# clean

.PHONY: clean
clean:
	rm -f gobbs gobbs.zip

#-----------------------------------------
# debug

.PHONY: run
run:
	goop go run gobbs.go

.PHONY: initdb
initdb:
	goop go run script/init.go

#-----------------------------------------
# test

define dotest
	@echo "Test $(1)*_test.go"
	@find $1 -name "*_test.go" | \
	xargs -n 1 goop go test -tags "test"
endef

.PHONY: test test_config test_validation test_service test_util
test: test_config test_validation test_service test_util

test_config:
	$(call dotest,"config/")

test_validation:
	$(call dotest,"validation/")

test_service:
	$(call dotest,"service/")

test_util:
	$(call dotest,"util/")

#-----------------------------------------
# fmt

.PHONY: fmt
fmt:
	@goop go fmt ./...
