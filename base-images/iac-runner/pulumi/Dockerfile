# =============================================================================
# IaC Runner - Pulumi Variant (Per-Provider)
# =============================================================================
# A base image for running Go-based Pulumi programs with pre-warmed Go caches
# for a specific cloud provider. Built by GitHub Actions using self-hosted
# runners with persistent cache storage.
#
# The Go caches are built on the runner for a single provider and copied into
# this image, dramatically reducing cold start times for that provider's SDK.
#
# Usage:
#   FROM ghcr.io/plantonhq/project-planton/base-images/iac-runner:pulumi-aws-v0.0.1
#   FROM ghcr.io/plantonhq/project-planton/base-images/iac-runner:pulumi-gcp-v0.0.1
#   etc.
#
# Build: Triggered by .github/workflows/iac-runner-pulumi.yml
# =============================================================================

FROM debian:bullseye

ARG PULUMI_VERSION=v3.202.0
ARG GOLANG_VERSION=1.25.0

# Set environment
ENV DEBIAN_FRONTEND=noninteractive \
    PATH="/usr/local/go/bin:/root/go/bin:${PATH}" \
    GOMODCACHE=/var/cache/go-mod \
    GOCACHE=/var/cache/go-build

# OCI labels
LABEL org.opencontainers.image.source="https://github.com/plantonhq/project-planton"
LABEL org.opencontainers.image.description="IaC Runner Pulumi variant with pre-warmed Go caches"
LABEL org.opencontainers.image.licenses="Apache-2.0"

EXPOSE 8080

# ---- Base packages --------------------------------------------------------
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        ca-certificates \
        debian-archive-keyring \
        git \
        wget && \
    rm -rf /var/lib/apt/lists/*

# ---- Go -------------------------------------------------------------------
RUN wget -q "https://go.dev/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz" -O go.tar.gz && \
    tar -xf go.tar.gz -C /usr/local && \
    rm go.tar.gz

# ---- Pulumi ---------------------------------------------------------------
RUN wget -q "https://get.pulumi.com/releases/sdk/pulumi-${PULUMI_VERSION}-linux-x64.tar.gz" && \
    tar -xf "pulumi-${PULUMI_VERSION}-linux-x64.tar.gz" && \
    mv pulumi/* /usr/local/bin/ && \
    rm "pulumi-${PULUMI_VERSION}-linux-x64.tar.gz"

# =============================================================================
# PRE-BUILT GO CACHES (Provider-Specific)
# =============================================================================
# These caches were built on a self-hosted GitHub Actions runner with
# persistent PVC storage for a single cloud provider. This approach provides:
# - Smaller image size (only one provider's dependencies)
# - Faster pulls (2-3GB vs 15GB)
# - Incremental builds across workflow runs

# Create cache directories
RUN mkdir -p ${GOMODCACHE} ${GOCACHE}

# Copy pre-built caches from GitHub Actions workflow
# These are populated by the workflow before docker build
COPY merged-cache/go-mod ${GOMODCACHE}
COPY merged-cache/go-build ${GOCACHE}

# Verify caches and display sizes
RUN echo "=== Go Cache Sizes ===" && \
    du -sh ${GOMODCACHE} && \
    du -sh ${GOCACHE} && \
    echo "" && \
    echo "IaC Runner Pulumi variant ready with pre-warmed caches"

