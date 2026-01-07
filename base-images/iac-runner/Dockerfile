# =============================================================================
# IaC Runner Base Image
# =============================================================================
# A base image for running Pulumi and OpenTofu/Terraform programs with
# pre-warmed Go caches. Built by GitHub Actions using self-hosted runners
# with persistent cache storage.
#
# The Go caches are built in parallel on the runner and copied into this image,
# dramatically reducing cold start times for Go-based Pulumi programs.
#
# Usage:
#   FROM ghcr.io/plantonhq/project-planton/base-images/iac-runner:<version>
#
# Build: Triggered by .github/workflows/iac-runner-base-image.yml
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
# PRE-BUILT GO CACHES
# =============================================================================
# These caches were built in parallel on a self-hosted GitHub Actions runner
# with persistent PVC storage. This approach provides:
# - Incremental builds (only new/changed modules are compiled)
# - Persistent cache across workflow runs
# - Parallel compilation of all provider modules

# Create cache directories
RUN mkdir -p ${GOMODCACHE} ${GOCACHE}

# Copy pre-built caches from GitHub Actions workflow
COPY merged-cache/go-mod ${GOMODCACHE}
COPY merged-cache/go-build ${GOCACHE}

# Verify caches and display sizes
RUN echo "=== Go Cache Sizes ===" && \
    du -sh ${GOMODCACHE} && \
    du -sh ${GOCACHE} && \
    echo "" && \
    echo "IaC Runner base image ready with pre-warmed caches"

