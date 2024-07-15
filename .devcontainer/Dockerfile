FROM mcr.microsoft.com/devcontainers/go:1.22
# install ca certificates in the container
COPY .devcontainer/all-ca-certs.crt /usr/local/share/ca-certificates/all-ca-certs.crt
RUN update-ca-certificates
# update container packages
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install git