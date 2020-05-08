# Data API

A repository for learning domain-driven design and golang development.

## Specification

* spec : `api/openapi.yaml`
* build openapi.html by redoc-cli : `bash api/build_html.sh`

## Installation

```
# Get packages.
go get github.com/shuufujita/data-api/cmd/dataapi

# Install command.
go install github.com/shuufujita/data-api/cmd/dataapi

# Install vendor packages.
cd src/github.com/shuufujita/data-api && dep ensure
```

## Generate RSA Keys

* create private key : `openssl genrsa 1024 > private-key.pem`
* create public key : `openssl rsa -in private-key.pem -pubout -out public-key.pem`

## Usage

* launch api : `bin/dataapi`
