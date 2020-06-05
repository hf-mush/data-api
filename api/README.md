# API Specification

API Specification files by OpenAPI.

## Build

build API specification with HTML.

```
$ bash build_html.sh
```

## Convert Tools

use the following tools to build.

### redoc-cli

https://github.com/Redocly/redoc/tree/master/cli

```
$ redoc-cli bundle openapi.yaml --options.showExtensions --options.disableSearch --title "Title" -o openapi.html
```

### swagger2openapi

https://github.com/Mermade/oas-kit/tree/master/packages/swagger2openapi

```
$ swagger2openapi -y -o openapi.yml swagger.yaml
```