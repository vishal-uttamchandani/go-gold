This repo will contain:

- rest api
- openapi spec
- versioning of module
- linter config
- makefile
- github actions to:
  - checkout code
  - run linter
  - run tests
  - build
  - create docker image
  - push image to registry
- configuration based on
  - environment variables
  - vault
- TLS
- Timeouts on the rest server
- Testing of rest server handlers
- Authentication and authorization using JWT