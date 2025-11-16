default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

# Run unit tests
.PHONY: test
test:
	go test ./... -v $(TESTARGS) -timeout 30s

# Generate documentation
.PHONY: generate
generate:
	go generate ./...

# Format code
.PHONY: fmt
fmt:
	gofmt -s -w -e .
	terraform fmt -recursive ./examples/

# Build provider locally
.PHONY: build
build:
	go build -o terraform-provider-${PROVIDER_NAME}

# Install provider for local testing
.PHONY: install
install: build
	mkdir -p ~/.terraform.d/plugins/localhost/providers/${PROVIDER_NAME}/0.0.1/darwin_arm64
	mv terraform-provider-${PROVIDER_NAME} ~/.terraform.d/plugins/localhost/providers/${PROVIDER_NAME}/0.0.1/darwin_arm64/

# Lint code
.PHONY: lint
lint:
	golangci-lint run

# Validate documentation
.PHONY: docs
docs:
	tfplugindocs generate --provider-name ${PROVIDER_NAME}
	tfplugindocs validate
