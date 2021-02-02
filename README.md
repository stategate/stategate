# eventgate

A secure API gateway for CloudEvents applications

[![GoDoc](https://godoc.org/github.com/autom8ter/eventgate?status.svg)](https://godoc.org/github.com/autom8ter/eventgate/cep-client-go)

- [API Documentation](https://autom8ter.github.io/eventgate/)
                                        
## Features
- [x] [Headless](https://en.wikipedia.org/wiki/Headless_software)
- [x] [Stateless](https://nordicapis.com/defining-stateful-vs-stateless-web-services/)
- [x] [gRPC](https://grpc.io/) support
    - [protobuf schema](schema.proto)
- [x] [graphQL](https://graphql.org/) support `/graphql`
    - [graphQL schema](schema.graphql)
- [x] REST support `/`
    - [open api schema](schema.swagger.json)
- [x] Serve gRPC, graphQL, & REST on same port
- [x] Authentication - JWT/OAuth with remote JWKS verification
- [x] Authorization - [Rego](https://www.openpolicyagent.org/docs/latest/policy-language/) based Authorization engine
- [x] Autogenerated Client gRPC SDK's
    - [x] Go
        - [![GoDoc](https://godoc.org/github.com/autom8ter/eventgate?status.svg)](https://godoc.org/github.com/autom8ter/eventgate/cep-client-go)
    - [x] [Node](./gen/grpc/node)
    - [x] [PHP](./gen/grpc/php)
    - [x] [C#](./gen/grpc/csharp)
    - [x] [Java](./gen/grpc/java)
    - [x] [gRPC Web](./gen/grpc/web)
- [x] Structured JSON Logs
- [x] Metrics Server(prometheus/pprof)
- [x] [Sample Kubernetes Manifest](k8s.yaml)
- [x] Pluggable Backends
    - [x] Nats
    - [ ] Kafka
    

## Command Line

```
eventgate -h
Usage of eventgate:
      --config string   path to config file (env: EVENTGATE_CONFIG) (default "config.yaml")
```

## Sample Config


```yaml
# port to serve on. metrics server is started on this port+1
port: 8820
# enable debug logs
debug: true
nats_url: "0.0.0.0:4444"
# json web keys uri for authentication
jwks_uri: ""
# rego policy for request authorization - this one allows any request
request_policy:
  rego_policy: |-
    package eventgate.authz

    default allow = true
  # query the allow variable
  rego_query: "data.eventgate.authz.allow"
# rego policy for response authorization - this one allows any request
response_policy:
  rego_policy: |-
    package eventgate.authz

    default allow = true
  # query the allow variable
  rego_query: "data.eventgate.authz.allow"

```

## Notes

      
