FROM mcr.microsoft.com/devcontainers/go:1.25

RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    pkg-config \
    librdkafka-dev \
 && rm -rf /var/lib/apt/lists/*
