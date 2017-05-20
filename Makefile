PACKAGES?=$$(go list ./... | grep -v vendor)
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

default: test vet

bin:
	./scripts/build

test: fmtcheck
	go test $(PACKAGES)

vet:
	@echo "go vet ."
	@go vet $(PACKAGES) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	./scripts/gofmtcheck

.PHONY: bin test vet fmt fmtcheck
