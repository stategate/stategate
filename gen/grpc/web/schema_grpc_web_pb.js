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
proto.stategate.StateServiceClient =
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
proto.stategate.StateServicePromiseClient =
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
 *   !proto.stategate.State,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_StateService_Set = new grpc.web.MethodDescriptor(
  '/stategate.StateService/Set',
  grpc.web.MethodType.UNARY,
  proto.stategate.State,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.stategate.State} request
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
 *   !proto.stategate.State,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_StateService_Set = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.stategate.State} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.stategate.State} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.stategate.StateServiceClient.prototype.set =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/stategate.StateService/Set',
      request,
      metadata || {},
      methodDescriptor_StateService_Set,
      callback);
};


/**
 * @param {!proto.stategate.State} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.stategate.StateServicePromiseClient.prototype.set =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/stategate.StateService/Set',
      request,
      metadata || {},
      methodDescriptor_StateService_Set);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.stategate.StateRef,
 *   !proto.stategate.State>}
 */
const methodDescriptor_StateService_Get = new grpc.web.MethodDescriptor(
  '/stategate.StateService/Get',
  grpc.web.MethodType.UNARY,
  proto.stategate.StateRef,
  proto.stategate.State,
  /**
   * @param {!proto.stategate.StateRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.State.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.stategate.StateRef,
 *   !proto.stategate.State>}
 */
const methodInfo_StateService_Get = new grpc.web.AbstractClientBase.MethodInfo(
  proto.stategate.State,
  /**
   * @param {!proto.stategate.StateRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.State.deserializeBinary
);


/**
 * @param {!proto.stategate.StateRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.stategate.State)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.stategate.State>|undefined}
 *     The XHR Node Readable Stream
 */
proto.stategate.StateServiceClient.prototype.get =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/stategate.StateService/Get',
      request,
      metadata || {},
      methodDescriptor_StateService_Get,
      callback);
};


/**
 * @param {!proto.stategate.StateRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.stategate.State>}
 *     Promise that resolves to the response
 */
proto.stategate.StateServicePromiseClient.prototype.get =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/stategate.StateService/Get',
      request,
      metadata || {},
      methodDescriptor_StateService_Get);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.stategate.StateRef,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_StateService_Del = new grpc.web.MethodDescriptor(
  '/stategate.StateService/Del',
  grpc.web.MethodType.UNARY,
  proto.stategate.StateRef,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.stategate.StateRef} request
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
 *   !proto.stategate.StateRef,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_StateService_Del = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.stategate.StateRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.stategate.StateRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.stategate.StateServiceClient.prototype.del =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/stategate.StateService/Del',
      request,
      metadata || {},
      methodDescriptor_StateService_Del,
      callback);
};


/**
 * @param {!proto.stategate.StateRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.stategate.StateServicePromiseClient.prototype.del =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/stategate.StateService/Del',
      request,
      metadata || {},
      methodDescriptor_StateService_Del);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.stategate.SearchStateOpts,
 *   !proto.stategate.StateValues>}
 */
const methodDescriptor_StateService_Search = new grpc.web.MethodDescriptor(
  '/stategate.StateService/Search',
  grpc.web.MethodType.UNARY,
  proto.stategate.SearchStateOpts,
  proto.stategate.StateValues,
  /**
   * @param {!proto.stategate.SearchStateOpts} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.StateValues.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.stategate.SearchStateOpts,
 *   !proto.stategate.StateValues>}
 */
const methodInfo_StateService_Search = new grpc.web.AbstractClientBase.MethodInfo(
  proto.stategate.StateValues,
  /**
   * @param {!proto.stategate.SearchStateOpts} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.StateValues.deserializeBinary
);


/**
 * @param {!proto.stategate.SearchStateOpts} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.stategate.StateValues)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.stategate.StateValues>|undefined}
 *     The XHR Node Readable Stream
 */
proto.stategate.StateServiceClient.prototype.search =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/stategate.StateService/Search',
      request,
      metadata || {},
      methodDescriptor_StateService_Search,
      callback);
};


/**
 * @param {!proto.stategate.SearchStateOpts} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.stategate.StateValues>}
 *     Promise that resolves to the response
 */
proto.stategate.StateServicePromiseClient.prototype.search =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/stategate.StateService/Search',
      request,
      metadata || {},
      methodDescriptor_StateService_Search);
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

