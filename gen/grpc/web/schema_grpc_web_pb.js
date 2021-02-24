/**
 * @fileoverview gRPC-Web generated client stub for stategate
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_api_annotations_pb = require('./google/api/annotations_pb.js')

var google_protobuf_struct_pb = require('google-protobuf/google/protobuf/struct_pb.js')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')

var google_protobuf_any_pb = require('google-protobuf/google/protobuf/any_pb.js')

var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var github_com_mwitkow_go$proto$validators_validator_pb = require('./github.com/mwitkow/go-proto-validators/validator_pb.js')
const proto = {};
proto.stategate = require('./schema_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.stategate.EntityServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.stategate.EntityServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.stategate.Entity,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_EntityService_Set = new grpc.web.MethodDescriptor(
  '/stategate.EntityService/Set',
  grpc.web.MethodType.UNARY,
  proto.stategate.Entity,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.stategate.Entity} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.stategate.Entity,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_EntityService_Set = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.stategate.Entity} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.stategate.Entity} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.stategate.EntityServiceClient.prototype.set =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/stategate.EntityService/Set',
      request,
      metadata || {},
      methodDescriptor_EntityService_Set,
      callback);
};


/**
 * @param {!proto.stategate.Entity} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.stategate.EntityServicePromiseClient.prototype.set =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/stategate.EntityService/Set',
      request,
      metadata || {},
      methodDescriptor_EntityService_Set);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.stategate.EntityRef,
 *   !proto.stategate.Entity>}
 */
const methodDescriptor_EntityService_Get = new grpc.web.MethodDescriptor(
  '/stategate.EntityService/Get',
  grpc.web.MethodType.UNARY,
  proto.stategate.EntityRef,
  proto.stategate.Entity,
  /**
   * @param {!proto.stategate.EntityRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.Entity.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.stategate.EntityRef,
 *   !proto.stategate.Entity>}
 */
const methodInfo_EntityService_Get = new grpc.web.AbstractClientBase.MethodInfo(
  proto.stategate.Entity,
  /**
   * @param {!proto.stategate.EntityRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.Entity.deserializeBinary
);


/**
 * @param {!proto.stategate.EntityRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.stategate.Entity)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.stategate.Entity>|undefined}
 *     The XHR Node Readable Stream
 */
proto.stategate.EntityServiceClient.prototype.get =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/stategate.EntityService/Get',
      request,
      metadata || {},
      methodDescriptor_EntityService_Get,
      callback);
};


/**
 * @param {!proto.stategate.EntityRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.stategate.Entity>}
 *     Promise that resolves to the response
 */
proto.stategate.EntityServicePromiseClient.prototype.get =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/stategate.EntityService/Get',
      request,
      metadata || {},
      methodDescriptor_EntityService_Get);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.stategate.EntityRef,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_EntityService_Del = new grpc.web.MethodDescriptor(
  '/stategate.EntityService/Del',
  grpc.web.MethodType.UNARY,
  proto.stategate.EntityRef,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.stategate.EntityRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.stategate.EntityRef,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_EntityService_Del = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.stategate.EntityRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.stategate.EntityRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.stategate.EntityServiceClient.prototype.del =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/stategate.EntityService/Del',
      request,
      metadata || {},
      methodDescriptor_EntityService_Del,
      callback);
};


/**
 * @param {!proto.stategate.EntityRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.stategate.EntityServicePromiseClient.prototype.del =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/stategate.EntityService/Del',
      request,
      metadata || {},
      methodDescriptor_EntityService_Del);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.stategate.SearchEntitiesOpts,
 *   !proto.stategate.Entities>}
 */
const methodDescriptor_EntityService_Search = new grpc.web.MethodDescriptor(
  '/stategate.EntityService/Search',
  grpc.web.MethodType.UNARY,
  proto.stategate.SearchEntitiesOpts,
  proto.stategate.Entities,
  /**
   * @param {!proto.stategate.SearchEntitiesOpts} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.Entities.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.stategate.SearchEntitiesOpts,
 *   !proto.stategate.Entities>}
 */
const methodInfo_EntityService_Search = new grpc.web.AbstractClientBase.MethodInfo(
  proto.stategate.Entities,
  /**
   * @param {!proto.stategate.SearchEntitiesOpts} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.Entities.deserializeBinary
);


/**
 * @param {!proto.stategate.SearchEntitiesOpts} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.stategate.Entities)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.stategate.Entities>|undefined}
 *     The XHR Node Readable Stream
 */
proto.stategate.EntityServiceClient.prototype.search =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/stategate.EntityService/Search',
      request,
      metadata || {},
      methodDescriptor_EntityService_Search,
      callback);
};


/**
 * @param {!proto.stategate.SearchEntitiesOpts} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.stategate.Entities>}
 *     Promise that resolves to the response
 */
proto.stategate.EntityServicePromiseClient.prototype.search =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/stategate.EntityService/Search',
      request,
      metadata || {},
      methodDescriptor_EntityService_Search);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.stategate.EventServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.stategate.EventServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.stategate.StreamOpts,
 *   !proto.stategate.Event>}
 */
const methodDescriptor_EventService_Stream = new grpc.web.MethodDescriptor(
  '/stategate.EventService/Stream',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.stategate.StreamOpts,
  proto.stategate.Event,
  /**
   * @param {!proto.stategate.StreamOpts} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.Event.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.stategate.StreamOpts,
 *   !proto.stategate.Event>}
 */
const methodInfo_EventService_Stream = new grpc.web.AbstractClientBase.MethodInfo(
  proto.stategate.Event,
  /**
   * @param {!proto.stategate.StreamOpts} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.Event.deserializeBinary
);


/**
 * @param {!proto.stategate.StreamOpts} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.stategate.Event>}
 *     The XHR Node Readable Stream
 */
proto.stategate.EventServiceClient.prototype.stream =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/stategate.EventService/Stream',
      request,
      metadata || {},
      methodDescriptor_EventService_Stream);
};


/**
 * @param {!proto.stategate.StreamOpts} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.stategate.Event>}
 *     The XHR Node Readable Stream
 */
proto.stategate.EventServicePromiseClient.prototype.stream =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/stategate.EventService/Stream',
      request,
      metadata || {},
      methodDescriptor_EventService_Stream);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.stategate.SearchEventOpts,
 *   !proto.stategate.Events>}
 */
const methodDescriptor_EventService_Search = new grpc.web.MethodDescriptor(
  '/stategate.EventService/Search',
  grpc.web.MethodType.UNARY,
  proto.stategate.SearchEventOpts,
  proto.stategate.Events,
  /**
   * @param {!proto.stategate.SearchEventOpts} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.Events.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.stategate.SearchEventOpts,
 *   !proto.stategate.Events>}
 */
const methodInfo_EventService_Search = new grpc.web.AbstractClientBase.MethodInfo(
  proto.stategate.Events,
  /**
   * @param {!proto.stategate.SearchEventOpts} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.Events.deserializeBinary
);


/**
 * @param {!proto.stategate.SearchEventOpts} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.stategate.Events)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.stategate.Events>|undefined}
 *     The XHR Node Readable Stream
 */
proto.stategate.EventServiceClient.prototype.search =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/stategate.EventService/Search',
      request,
      metadata || {},
      methodDescriptor_EventService_Search,
      callback);
};


/**
 * @param {!proto.stategate.SearchEventOpts} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.stategate.Events>}
 *     Promise that resolves to the response
 */
proto.stategate.EventServicePromiseClient.prototype.search =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/stategate.EventService/Search',
      request,
      metadata || {},
      methodDescriptor_EventService_Search);
};


module.exports = proto.stategate;

