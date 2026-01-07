.PHONY: clean

IMAGE_REPO=ghcr.io/plantonhq/project-planton/base-images/iac-runner
IMAGE_TAG?=local

# Note: The image is built by GitHub Actions workflow (.github/workflows/iac-runner-base-image.yml)
# which uses self-hosted runners with persistent Go cache storage.
# Local builds are not recommended as they won't have the pre-warmed caches.

# Clean up local images
clean:
	docker rmi $(IMAGE_REPO):$(IMAGE_TAG) 2>/dev/null || true
