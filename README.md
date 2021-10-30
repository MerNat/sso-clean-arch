# SSO service (using clean arch)

### Running it
`go run .`

### Documentation
`[Swagger UI](http://0.0.0.0:8181/api/v1/docs/)`

### Getting the JWKS endpoint
`http://0.0.0.0:8181/api/v1/sso/jwks` [GET]

### Verifying the token

`http://0.0.0.0:8181/api/v1/sso/verify` [GET] [With header 'Authorization': 'Bearer **Token**']