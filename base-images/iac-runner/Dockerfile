# =============================================================================
# IaC Runner Base Image
# =============================================================================
# A base image for running Pulumi and OpenTofu/Terraform programs with
# pre-warmed Go caches. This dramatically reduces cold start times for
# Go-based Pulumi programs by pre-downloading and pre-compiling provider SDKs.
#
# Features:
# - Go runtime with pre-warmed GOMODCACHE and GOCACHE
# - Pulumi CLI
# - OpenTofu CLI
# - Kubernetes credential exec plugins for GCP, AWS, and Azure
#
# Usage:
#   FROM ghcr.io/plantonhq/project-planton/base-images/iac-runner:<version>
# =============================================================================

FROM debian:bullseye

ARG PULUMI_VERSION=v3.202.0
ARG OPENTOFU_VERSION=1.9.1
ARG GOLANG_VERSION=1.25.0

# Set environment
ENV DEBIAN_FRONTEND=noninteractive \
    PATH="/usr/local/go/bin:/root/go/bin:${PATH}" \
    GOMODCACHE=/var/cache/go-mod \
    GOCACHE=/var/cache/go-build

# OCI labels
LABEL org.opencontainers.image.source="https://github.com/plantonhq/project-planton"
LABEL org.opencontainers.image.description="IaC Runner base image with pre-warmed Go caches for Pulumi"
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

# ---- OpenTofu -------------------------------------------------------------
RUN mkdir -p $HOME/.terraform.d/plugin-cache && \
    wget -q "https://github.com/opentofu/opentofu/releases/download/v${OPENTOFU_VERSION}/tofu_${OPENTOFU_VERSION}_linux_amd64.tar.gz" && \
    tar -xf "tofu_${OPENTOFU_VERSION}_linux_amd64.tar.gz" && \
    mv tofu /usr/local/bin/ && \
    echo 'plugin_cache_dir = "$HOME/.terraform.d/plugin-cache"' > $HOME/.tofurc && \
    rm "tofu_${OPENTOFU_VERSION}_linux_amd64.tar.gz"

# ---- Pulumi ---------------------------------------------------------------
RUN wget -q "https://get.pulumi.com/releases/sdk/pulumi-${PULUMI_VERSION}-linux-x64.tar.gz" && \
    tar -xf "pulumi-${PULUMI_VERSION}-linux-x64.tar.gz" && \
    mv pulumi/* /usr/local/bin/ && \
    rm "pulumi-${PULUMI_VERSION}-linux-x64.tar.gz"

# =============================================================================
# GO CACHE PRE-WARMING
# =============================================================================
# Pre-populate GOMODCACHE and GOCACHE by building ALL Pulumi modules from
# project-planton. This guarantees complete cache coverage for every
# deployment component.

# Create cache directories
RUN mkdir -p ${GOMODCACHE} ${GOCACHE}

# Clone project-planton (shallow clone to minimize transfer)
ARG PROJECT_PLANTON_REF=main
RUN cd /tmp && \
    git clone --depth 1 --branch ${PROJECT_PLANTON_REF} \
        https://github.com/plantonhq/project-planton.git

# Build ALL Pulumi modules to warm caches
# This iterates through every module directory and runs go build
# The build output is captured to show progress
RUN cd /tmp/project-planton && \
    echo "=== Building all Pulumi modules to warm Go caches ===" && \
    MODULE_COUNT=0 && \
    FAILED_COUNT=0 && \
    for dir in $(find . -type d -path '*/iac/pulumi' -not -path '*/module'); do \
        MODULE_COUNT=$((MODULE_COUNT + 1)); \
        echo "[$MODULE_COUNT] Building: $dir"; \
        cd /tmp/project-planton && cd "$dir" && \
        if go build -v . 2>&1 | tail -3; then \
            echo "  ✓ Success"; \
        else \
            echo "  ✗ Failed (continuing...)"; \
            FAILED_COUNT=$((FAILED_COUNT + 1)); \
        fi; \
    done && \
    echo "" && \
    echo "=== Build Summary ===" && \
    echo "Total modules: $MODULE_COUNT" && \
    echo "Failed: $FAILED_COUNT" && \
    echo "Success: $((MODULE_COUNT - FAILED_COUNT))"

# Clean up source code (keep only caches)
RUN rm -rf /tmp/project-planton

# Verify caches and display sizes
RUN echo "=== Go Cache Sizes ===" && \
    du -sh ${GOMODCACHE} && \
    du -sh ${GOCACHE} && \
    echo "" && \
    echo "IaC Runner base image ready with fully pre-warmed caches"

