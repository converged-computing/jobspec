
COMMONENVVAR=GOOS=$(shell uname -s | tr A-Z a-z)
RELEASE_VERSION?=v$(shell date +%Y%m%d)-$(shell git describe --tags --match "v*")

.PHONY: all
all: example1 example2 example3 example4 example5 example6 createnew exp1 exp2

.PHONY: build
build: 
	go mod tidy
	mkdir -p ./examples/v1/bin
	mkdir -p ./examples/experimental/bin
	mkdir -p ./examples/nextgen/v1/bin

# Build examples
.PHONY: createnew
createnew: build
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o ./examples/v1/bin/createnew examples/v1/createnew/example.go

.PHONY: example1
example1: build
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o ./examples/v1/bin/example1 examples/v1/example1/example.go

.PHONY: example2
example2: build
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o ./examples/v1/bin/example2 examples/v1/example2/example.go

.PHONY: example3
example3: build
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o ./examples/v1/bin/example3 examples/v1/example3/example.go

.PHONY: example4
example4: build
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o ./examples/v1/bin/example4 examples/v1/example4/example.go

.PHONY: example5
example5: build
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o ./examples/v1/bin/example5 examples/v1/example5/example.go

.PHONY: example6
example6: build
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o ./examples/v1/bin/example6 examples/v1/example6/example.go

.PHONY: exp1
exp1: build
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o ./examples/experimental/bin/example1 examples/experimental/example1/example.go

.PHONY: exp2
exp2: build
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o ./examples/experimental/bin/example2 examples/experimental/example2/example.go

.PHONY: ng1
ng1: build
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o ./examples/nextgen/v1/bin/example1 examples/nextgen/v1/example1/example.go


.PHONY: test
test: build all
	./examples/v1/bin/example1
	./examples/v1/bin/example2
	./examples/v1/bin/example3
	./examples/v1/bin/example4
	./examples/v1/bin/example5
	./examples/v1/bin/example6
	./examples/v1/bin/createnew
	./examples/experimental/bin/example1
	./examples/experimental/bin/example2
	./examples/nextgen/v1/bin/example1

.PHONY: clean
clean:
	rm -rf ./examples/bin/*