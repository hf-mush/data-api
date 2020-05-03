# Data API

A repository for learning domain-driven design and golang development.

## Specification

* spec : `api/openapi.yaml`
* build openapi.html by redoc-cli : `bash api/build_html.sh`

## Installation

```
# Get packages.
go get github.com/hf-mush/data-api/cmd/dataapi

# Install command.
go install github.com/hf-mush/data-api/cmd/dataapi

# Install vendor packages.
cd src/github.com/hf-mush/data-api && dep ensure
```

## Usage

* launch api : `bin/dataapi`
