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
proto.stategate.ObjectServiceClient =
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
proto.stategate.ObjectServicePromiseClient =
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
 *   !proto.stategate.Object,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_ObjectService_Set = new grpc.web.MethodDescriptor(
  '/stategate.ObjectService/Set',
  grpc.web.MethodType.UNARY,
  proto.stategate.Object,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.stategate.Object} request
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
 *   !proto.stategate.Object,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_ObjectService_Set = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.stategate.Object} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.stategate.Object} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.stategate.ObjectServiceClient.prototype.set =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/stategate.ObjectService/Set',
      request,
      metadata || {},
      methodDescriptor_ObjectService_Set,
      callback);
};


/**
 * @param {!proto.stategate.Object} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.stategate.ObjectServicePromiseClient.prototype.set =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/stategate.ObjectService/Set',
      request,
      metadata || {},
      methodDescriptor_ObjectService_Set);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.stategate.ObjectRef,
 *   !proto.stategate.Object>}
 */
const methodDescriptor_ObjectService_Get = new grpc.web.MethodDescriptor(
  '/stategate.ObjectService/Get',
  grpc.web.MethodType.UNARY,
  proto.stategate.ObjectRef,
  proto.stategate.Object,
  /**
   * @param {!proto.stategate.ObjectRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.Object.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.stategate.ObjectRef,
 *   !proto.stategate.Object>}
 */
const methodInfo_ObjectService_Get = new grpc.web.AbstractClientBase.MethodInfo(
  proto.stategate.Object,
  /**
   * @param {!proto.stategate.ObjectRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.Object.deserializeBinary
);


/**
 * @param {!proto.stategate.ObjectRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.stategate.Object)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.stategate.Object>|undefined}
 *     The XHR Node Readable Stream
 */
proto.stategate.ObjectServiceClient.prototype.get =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/stategate.ObjectService/Get',
      request,
      metadata || {},
      methodDescriptor_ObjectService_Get,
      callback);
};


/**
 * @param {!proto.stategate.ObjectRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.stategate.Object>}
 *     Promise that resolves to the response
 */
proto.stategate.ObjectServicePromiseClient.prototype.get =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/stategate.ObjectService/Get',
      request,
      metadata || {},
      methodDescriptor_ObjectService_Get);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.stategate.ObjectRef,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_ObjectService_Del = new grpc.web.MethodDescriptor(
  '/stategate.ObjectService/Del',
  grpc.web.MethodType.UNARY,
  proto.stategate.ObjectRef,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.stategate.ObjectRef} request
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
 *   !proto.stategate.ObjectRef,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_ObjectService_Del = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.stategate.ObjectRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.stategate.ObjectRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.stategate.ObjectServiceClient.prototype.del =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/stategate.ObjectService/Del',
      request,
      metadata || {},
      methodDescriptor_ObjectService_Del,
      callback);
};


/**
 * @param {!proto.stategate.ObjectRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.stategate.ObjectServicePromiseClient.prototype.del =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/stategate.ObjectService/Del',
      request,
      metadata || {},
      methodDescriptor_ObjectService_Del);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.stategate.SearchObjectOpts,
 *   !proto.stategate.Objects>}
 */
const methodDescriptor_ObjectService_Search = new grpc.web.MethodDescriptor(
  '/stategate.ObjectService/Search',
  grpc.web.MethodType.UNARY,
  proto.stategate.SearchObjectOpts,
  proto.stategate.Objects,
  /**
   * @param {!proto.stategate.SearchObjectOpts} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.Objects.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.stategate.SearchObjectOpts,
 *   !proto.stategate.Objects>}
 */
const methodInfo_ObjectService_Search = new grpc.web.AbstractClientBase.MethodInfo(
  proto.stategate.Objects,
  /**
   * @param {!proto.stategate.SearchObjectOpts} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.Objects.deserializeBinary
);


/**
 * @param {!proto.stategate.SearchObjectOpts} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.stategate.Objects)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.stategate.Objects>|undefined}
 *     The XHR Node Readable Stream
 */
proto.stategate.ObjectServiceClient.prototype.search =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/stategate.ObjectService/Search',
      request,
      metadata || {},
      methodDescriptor_ObjectService_Search,
      callback);
};


/**
 * @param {!proto.stategate.SearchObjectOpts} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.stategate.Objects>}
 *     Promise that resolves to the response
 */
proto.stategate.ObjectServicePromiseClient.prototype.search =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/stategate.ObjectService/Search',
      request,
      metadata || {},
      methodDescriptor_ObjectService_Search);
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

