# stategate

A pluggable "Application State Gateway" that enforces the [Event Sourcing Pattern](https://microservices.io/patterns/data/event-sourcing.html) for securely persisting & broadcasting application state changes

[![GoDoc](https://godoc.org/github.com/stategate/stategate?status.svg)](https://godoc.org/github.com/stategate/stategate/stategate-client-go)

- [API Documentation](https://stategate.github.io/stategate/)

## API Services/Methods

```proto
// EntityService serves API methods to clients that modify/query the current state of an entity
// An Entity is a single object with a type, domain, key, and k/v values
service EntityService {
  // Set sets the current state value of an entity, adds it to the event log, then broadcast the event to all interested consumers(EventService.Stream)
  rpc Set(Entity) returns(google.protobuf.Empty) {
    option (google.api.http) = {
      post: "/api/entity/ref/{domain}/{type}/{key}"
    };
  }
  // Edit overwrites the k/v pairs present in the entity request without replacing the entire entity.
  // It then adds the state change to the event log, then broadcast the event to all interested consumers(EventService.Stream)
  // Edit returns the current state of the Entity after the mutation.
  rpc Edit(Entity) returns(Entity){
    option (google.api.http) = {
      patch: "/api/entity/ref/{domain}/{type}/{key}"
    };
  }
  // Revert reverts an Entity to a previous version of itself
  // Reverting an entity dispatches another event since it is a state change
  rpc Revert(EventRef) returns(Entity) {
    option (google.api.http) = {
      put: "/api/entity/ref/{domain}/{type}/{key}/revert"
    };
  }
  // Get gets an entity's current state
  rpc Get(EntityRef) returns(Entity) {
    option (google.api.http) = {
      get: "/api/entity/ref/{domain}/{type}/{key}"
    };
  }
  // Del hard deletes an entity from current state store, adds it's state prior to deletion to the event log, then broadcast the event to all interested consumers(EventService.Stream)
  // an Entity may be recovered via querying the Event store for historical records of the deleted Entity.
  rpc Del(EntityRef) returns(google.protobuf.Empty) {
    option (google.api.http) = {
      delete: "/api/entity/ref/{domain}/{type}/{key}"
    };
  }
  // Search queries the current state of entities
  rpc Search(SearchEntityOpts) returns(Entities) {
    option (google.api.http) = {
      get: "/api/entity/search"
    };
  }
}

// EventService serves API methods related to stategate Event Consumers
// Events are automatically emitted from mutations made from State mutations within the EntityService
service EventService {
  // Stream creates an event stream/subscription to state changes to entities in real time. Glob matching is supported.
  rpc Stream(StreamEventOpts) returns(stream Event) {
    option (google.api.http) = {
      get: "/api/events/stream"
    };
  }
  // Search queries historical events - every historical state change to an entity may be queried.
  rpc Search(SearchEventOpts) returns(Events) {
    option (google.api.http) = {
      get: "/api/events/search"
    };
  }
  // Get gets a single event
  rpc Get(EventRef) returns(Event) {
    option (google.api.http) = {
      get: "/api/events/ref/{domain}/{type}/{key}/{id}"
    };
  }
}

// PeerService provides a means for clients to communicate directly with one another WITHOUT making any state changes.
// Please note that all messages transported via the PeerService are not persisted in any way.
service PeerService {
  // Broadcast broadcasts a message to N subscribers(clients calling Stream)
  rpc Broadcast(Message) returns(google.protobuf.Empty){
    option (google.api.http) = {
      post: "/api/peers/broadcast"
      body: "*"
    };
  }
  // Stream consumes/streams messages from message producers(clients calling broadcast)
  rpc Stream(StreamMessageOpts) returns(stream PeerMessage){
    option (google.api.http) = {
      get: "/api/peers/stream"
    };
  }
}

// CacheService is for persisting short lived values in memory for performance-critical operations
service CacheService {
  // Set sets a value in the cache
  rpc Set(Cache) returns(google.protobuf.Empty){
    option (google.api.http) = {
      put: "/api/cache/ref/{domain}/{key}"
    };
  }
  // Get gets a value from the cache
  rpc Get(CacheRef) returns(Cache) {
    option (google.api.http) = {
      get: "/api/cache/ref/{domain}/{key}"
    };
  }
  // Del deletes a value from the cache
  rpc Del(CacheRef) returns(google.protobuf.Empty) {
    option (google.api.http) = {
      delete: "/api/cache/ref/{domain}/{key}"
    };
  }
}

// MutexService offers distributed locking capabilities for client's that need to coordinate with peer services.
service MutexService {
  // Lock locks a value for a period of time if it is not locked already.
  // If it is already locked, an error will be returned
  // It is best practice for client's to call Unlock when the distributed lock operation is completed instead of relying on the TTL
  rpc Lock(Mutex) returns(google.protobuf.Empty) {
    option (google.api.http) = {
      post: "/api/mutex/ref/{domain}/{key}/lock"
    };
  }
  // Unlock unlocks the key(if it's currently locked) so that it may be locked again.
  // It is best practice for client's to call Unlock when the distributed lock operation is completed instead of relying on the TTL
  rpc Unlock(MutexRef) returns(google.protobuf.Empty) {
    option (google.api.http) = {
      post: "/api/mutex/ref/{domain}/{key}/unlock"
    };
  }
}

```
                                        
## Features

- [x] Capture all changes to an application's state(entities) as a sequence of events - Event Sourcing(EntityService/EventService)
- [x] High Performance Pubsub Interface Service(PeerService) 
- [x] High Performance Caching Interface(CacheService)
- [x] High Performance Distributed Locking Interface(MutexService)
- [x] Stateless & horizontally scaleable
- [x] Native [gRPC](https://grpc.io/) support
    - [protobuf schema](schema.proto)
- [x] Embedded REST support `/` (transcoding)
    - [open api schema](schema.swagger.json)
- [x] Embedded GraphQL support `/api/graphql` (transcoding)
    - [graphQL schema](schema.graphql)
- [x] Embedded [grpcweb](https://grpc.io/docs/platforms/web/basics/) support (transcoding)
- [x] Metrics Server(prometheus/pprof)
- [x] Authentication - JWT/OAuth with remote [JWKS](https://auth0.com/docs/tokens/json-web-tokens/json-web-key-sets) verification
- [x] Authorization - [Rego](https://www.openpolicyagent.org/docs/latest/policy-language/) based Authorization engine
- [x] Autogenerated Client gRPC SDK's
    - [x] Go
        - [![GoDoc](https://godoc.org/github.com/stategate/stategate?status.svg)](https://godoc.org/github.com/stategate/stategate/stategate-client-go)
    - [x] [Node](./gen/grpc/node)
    - [x] [PHP](./gen/grpc/php)
    - [x] [C#](./gen/grpc/csharp)
    - [x] [Java](./gen/grpc/java)
    - [x] [gRPC Web](./gen/grpc/web)
    - [x] [Python](./gen/grpc/python)
    - [ ] Ruby
- [x] Structured JSON Logs
- [x] [Sample Kubernetes Manifest](k8s.yaml)
- [x] [Sample Docker Compose](docker-compose.yml)
- [x] Pluggable ["Storage" Providers](internal/storage/providers.go)
    - [x] MongoDb
        - [x] fully-tested
    - [ ] PostgreSQL
    - [ ] MySQL
    - [ ] Cassandra
- [x] Pluggable ["Cache" Providers](internal/cache/providers.go)
    - [x] Redis
        - [x] fully-tested
    


## Goals

- [x] Create a simple API interface for storing state(entities) and subscribing to state changes(events) using pluggable cache & storage providers
- [x] Capture all changes to an application's state/entities as a sequence of events.
- [x] Safe to swap backend providers without changing client-side code
- [x] Type-safe client's generated in many languages
- [x] Safe to expose to the public internet due to fine-grained authentication/authorization model.
- [x] Capture a persistant, immutable historical record of all state changes to entities using a pluggable storage provider
- [x] Revert/Rollback an entity to any previous version of itself at any point in time
- [x] Store identity(jwt.claims) & timestamp in event logs to capture who is changing what & when
- [x] Easy deployment model - fully configureable via environmental variables
- [x] Create complex client applications with stategate as their only dependency
- [ ] Create serverless deployment model for stategate client applications

## Design

Stategate was designed with EventSourcing in mind

What is Event Sourcing?

> Event sourcing persists the state of a business entity such an Order or a Customer as a sequence of state-changing events. Whenever the state of a business entity changes, a new event is appended to the list of events. Since saving an event is a single operation, it is inherently atomic. The application reconstructs an entity’s current state by replaying the events. 

> Applications persist events in an event store, which is a database of events. The store has an API for adding and retrieving an entity’s events. The event store also behaves like a message broker. It provides an API that enables services to subscribe to events. When a service saves an event in the event store, it is delivered to all interested subscribers. 

![Event-Sourcing](./stategate.png)


### Primitives

#### Entity

An entity represents a single record(k/v pairs) with a unique key with a given [type](https://en.wikipedia.org/wiki/Type_system), belonging to a particular [domain](https://en.wikipedia.org/wiki/Domain-driven_design)

        
        // Entity represents a single record(k/v pairs) with a unique key with a given [type](https://en.wikipedia.org/wiki/Type_system), belonging to a particular [domain](https://en.wikipedia.org/wiki/Domain-driven_design)
        // EventService clients should use the EntityService to persist & interact with the current state of an entity.
        message Entity {
          // the entity's business domain(ex: accounting)
          // must not be empty or contain spaces
          string domain =1[(validator.field) = {regex : "^\\S+$"}];
          // the entity's type (ex: user)
          // must not be empty or contain spaces
          string type =2[(validator.field) = {regex : "^\\S+$"}];
          // the entity's key (unique within type). 
          // must not be empty or contain spaces
          string key =3[(validator.field) = {regex : "^\\S+$"}];
          // the entity's values (k/v pairs)
          google.protobuf.Struct values = 4[(validator.field) = {msg_exists : true}];
        }

#### Events
 
Event is primitive that represents a single state change to an entity


        // Event is primitive that represents a single state change to an entity
        // Events are persisted to history & broadcasted to interested consumers(Stream) any time an entity is created/modified/deleted
        // Events are immutable after creation and may be searched.
        // EventService client's may search events to query previous state of an entity(s)
        message Event {
          // identifies the event(uuid v4).
          string id = 1[(validator.field) = {uuid_ver : 4}];
          // state of an Entity after it has been mutated
          Entity entity = 2[(validator.field) = {msg_exists : true}];
          // the invoked method that triggered the event
          string method =5[(validator.field) = {string_not_empty : true}];
          // the authentication claims of the event producer.
          google.protobuf.Struct claims =3[(validator.field) = {msg_exists : true}];
          // timestamp(ns) of when the event was received.
          int64 time =4[(validator.field) = {int_gt : 0}];
        }

#### Messages  

Message is a non-persisted message passed between Peers as a means of communication

    // Message is an arbitrary message to be delivered to a Peer
    // Messages are NOT persisted and should only be used to communicate with other Peers
    message Message {
      // the message's business domain(ex: accounting)
      // must not be empty or contain spaces
      string domain =1[(validator.field) = {regex : "^\\S+$"}];
      // the message's channel(ex: general)
      // must not be empty or contain spaces
      string channel =2[(validator.field) = {regex : "^\\S+$"}];
      // message's type (ex: comment)
      // must not be empty or contain spaces
      string type =3[(validator.field) = {regex : "^\\S+$"}];
      // the body of the message(k/v values).
      google.protobuf.Struct body =4[(validator.field) = {msg_exists : true}];
    }

## Environmental Variables

.env files are loaded if found in the same directory as stategate

```yaml
# port to serve on (optional). defaults to 8080
STATEGATE_PORT=8080
# enable debug logs (optional)
STATEGATE_DEBUG=true
# disable all authentication & authorization(jwks, request policies, response policies) (optional)
STATEGATE_AUTH_DISABLED=false
# tls cert file (optional)
STATEGATE_TLS_CERT_FILE=/tmp/certs/stategate.cert
# tls key file (optional)
STATEGATE_TLS_KEY_FILE=/tmp/certs/stategate.key

# JSON Web Key Set remote URI used for fetching jwt signing keys for verification/validation (optional)
STATEGATE_JWKS_URI=https://www.googleapis.com/oauth2/v3/certs

# base64 encoded OPA rego policy executed on inbound requests from clients (optional)
STATEGATE_REQUEST_POLICY=cGFja2FnZSBzdGF0ZWdhdGUuYXV0aHoKCmRlZmF1bHQgYWxsb3cgPSB0cnVl
# base64 encoded OPA rego policy executed on responses sent to clients (optional)
STATEGATE_RESPONSE_POLICY=cGFja2FnZSBzdGF0ZWdhdGUuYXV0aHoKCmRlZmF1bHQgYWxsb3cgPSB0cnVl

# storage provider configuration(JSON) options: [mongo] REQUIRED
STATEGATE_STORAGE_PROVIDER={ "name": "mongo", "database": "testing", "addr": "mongodb://localhost:27017/testing" }

# cache provider configuration(JSON) options: [redis] REQUIRED
STATEGATE_CACHE_PROVIDER={ "name": "redis", "addr": "localhost:6379", "user": "xxx", "password": "xxxxxxxxxx" }

# CORS options for accessing stategate from the browser
STATEGATE_CORS_ALLOW_ORIGINS=*
STATEGATE_CORS_ALLOW_METHODS=POST,GET,PUT,DELETE
STATEGATE_CORS_ALLOW_HEADERS=*
```

## Implementation Details

- [Storage Provider Implementations](./internal/storage)
- [Cache Provider Implementations](./internal/cache)
- [Auth](./internal/auth)
- [ListenAndServe](./internal/server)
- [Errors](./internal/errorz)
- [Go Client SDK](./stategate-client-go)
- [Generated Code](./gen)
- [Testing Framework](./internal/testing/framework)
- [Makefile](./Makefile)

### Storage Providers

supported providers: [mongo]

- A stategate storage provider is a pluggable, 3rd party database storage service. 
- Storage providers provide persistance for all current entities/events and should be scaled independently of stategate instances.

[interface](internal/storage/providers.go)
```go

// EntityProvider provides logic for querying/persisting entities
type EntityProvider interface {
	SetEntity(ctx context.Context, entity *stategate.Entity) *errorz.Error
	EditEntity(ctx context.Context, entity *stategate.Entity) (*stategate.Entity, *errorz.Error)
	SearchEntities(ctx context.Context, ref *stategate.SearchEntityOpts) (*stategate.Entities, *errorz.Error)
	DelEntity(ctx context.Context, ref *stategate.EntityRef) *errorz.Error
	GetEntity(ctx context.Context, ref *stategate.EntityRef) (*stategate.Entity, *errorz.Error)
}

// EventProvider provides logic for querying/persisting events
type EventProvider interface {
	SaveEvent(ctx context.Context, event *stategate.Event) *errorz.Error
	SearchEvents(ctx context.Context, ref *stategate.SearchEventOpts) (*stategate.Events, *errorz.Error)
	GetEvent(ctx context.Context, ref *stategate.EventRef) (*stategate.Event, *errorz.Error)
}

// Provider is an event & entity provider
type Provider interface {
	EventProvider
	EntityProvider
	Close() error
}

```
   
### Cache Providers

supported providers: [redis]

- A stategate cache provider is a pluggable, 3rd party caching & message-queue service. 
- Cache providers provide a way for stategate to store ephemeral data & broadcast events to itself while scaling horizontally. 

[interface](internal/cache/providers.go)
  
```go

// ChannelProvider acts as dependency injection for broadcasting messages to stategate instances as they fan out
type ChannelProvider interface {
	PublishEvent(ctx context.Context, event *stategate.Event) *errorz.Error
	GetEventChannel(ctx context.Context) (chan *stategate.Event, *errorz.Error)
	PublishMessage(ctx context.Context, message *stategate.PeerMessage) *errorz.Error
	GetMessageChannel(ctx context.Context) (chan *stategate.PeerMessage, *errorz.Error)
}

// CacheProvider acts as dependency injection for caching ephemeral data 
type CacheProvider interface {
	Get(ctx context.Context, ref *stategate.CacheRef) (*stategate.Cache, *errorz.Error)
	Set(ctx context.Context, value *stategate.Cache) *errorz.Error
	Del(ctx context.Context, value *stategate.CacheRef) *errorz.Error
	
}

// MutexProvider acts as dependency injection for distributed mutex operations
type MutexProvider interface {
	Lock(ctx context.Context, ref *stategate.Mutex) *errorz.Error
	Unlock(ctx context.Context, value *stategate.MutexRef) *errorz.Error
}

// Provider is a channel, cache, & mutex provider
type Provider interface {
	CacheProvider
	ChannelProvider
	MutexProvider
	Close() error
}
```
    
- Cache providers should be scaled independently of stategate instances.

## Authorization

### Request Authorization Policies
TODO

## Response Authorization Policies
TODO

## Authentication
### Remote JWKS URI

https://auth0.com/docs/tokens/json-web-tokens/json-web-key-sets

TODO

## Sample GraphQL queries

```graphql
mutation {
  setEntity(input: {
    domain: "accounting",
    type: "user"
    key: "colemanword@gmail.com",
    values: {
      first_name: "Coleman"
      last_name: "Word"
    }
  })
}
```

## FAQ


