.PHONY: build-image build-image-no-cache clean

IMAGE_REPO=ghcr.io/plantonhq/project-planton/base-images/iac-runner
IMAGE_TAG?=local

# Build Docker image (builds all Pulumi modules to warm caches)
build-image:
	docker build -t $(IMAGE_REPO):$(IMAGE_TAG) .

# Build Docker image without cache (full rebuild)
build-image-no-cache:
	docker build --no-cache -t $(IMAGE_REPO):$(IMAGE_TAG) .

# Build with specific project-planton version
build-image-versioned:
	@if [ -z "$(VERSION)" ]; then \
		echo "Usage: make build-image-versioned VERSION=v1.2.3"; \
		exit 1; \
	fi
	docker build --build-arg PROJECT_PLANTON_REF=$(VERSION) \
		-t $(IMAGE_REPO):$(VERSION) .

# Clean up local images
clean:
	docker rmi $(IMAGE_REPO):$(IMAGE_TAG) 2>/dev/null || true
