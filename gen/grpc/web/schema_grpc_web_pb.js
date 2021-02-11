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
proto.stategate.StateGateServiceClient =
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
proto.stategate.StateGateServicePromiseClient =
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
const methodDescriptor_StateGateService_SetObject = new grpc.web.MethodDescriptor(
  '/stategate.StateGateService/SetObject',
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
const methodInfo_StateGateService_SetObject = new grpc.web.AbstractClientBase.MethodInfo(
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
proto.stategate.StateGateServiceClient.prototype.setObject =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/stategate.StateGateService/SetObject',
      request,
      metadata || {},
      methodDescriptor_StateGateService_SetObject,
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
proto.stategate.StateGateServicePromiseClient.prototype.setObject =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/stategate.StateGateService/SetObject',
      request,
      metadata || {},
      methodDescriptor_StateGateService_SetObject);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.stategate.ObjectRef,
 *   !proto.stategate.Object>}
 */
const methodDescriptor_StateGateService_GetObject = new grpc.web.MethodDescriptor(
  '/stategate.StateGateService/GetObject',
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
const methodInfo_StateGateService_GetObject = new grpc.web.AbstractClientBase.MethodInfo(
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
proto.stategate.StateGateServiceClient.prototype.getObject =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/stategate.StateGateService/GetObject',
      request,
      metadata || {},
      methodDescriptor_StateGateService_GetObject,
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
proto.stategate.StateGateServicePromiseClient.prototype.getObject =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/stategate.StateGateService/GetObject',
      request,
      metadata || {},
      methodDescriptor_StateGateService_GetObject);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.stategate.StreamOpts,
 *   !proto.stategate.Event>}
 */
const methodDescriptor_StateGateService_StreamEvents = new grpc.web.MethodDescriptor(
  '/stategate.StateGateService/StreamEvents',
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
const methodInfo_StateGateService_StreamEvents = new grpc.web.AbstractClientBase.MethodInfo(
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
proto.stategate.StateGateServiceClient.prototype.streamEvents =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/stategate.StateGateService/StreamEvents',
      request,
      metadata || {},
      methodDescriptor_StateGateService_StreamEvents);
};


/**
 * @param {!proto.stategate.StreamOpts} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.stategate.Event>}
 *     The XHR Node Readable Stream
 */
proto.stategate.StateGateServicePromiseClient.prototype.streamEvents =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/stategate.StateGateService/StreamEvents',
      request,
      metadata || {},
      methodDescriptor_StateGateService_StreamEvents);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.stategate.SearchOpts,
 *   !proto.stategate.Events>}
 */
const methodDescriptor_StateGateService_SearchEvents = new grpc.web.MethodDescriptor(
  '/stategate.StateGateService/SearchEvents',
  grpc.web.MethodType.UNARY,
  proto.stategate.SearchOpts,
  proto.stategate.Events,
  /**
   * @param {!proto.stategate.SearchOpts} request
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
 *   !proto.stategate.SearchOpts,
 *   !proto.stategate.Events>}
 */
const methodInfo_StateGateService_SearchEvents = new grpc.web.AbstractClientBase.MethodInfo(
  proto.stategate.Events,
  /**
   * @param {!proto.stategate.SearchOpts} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.Events.deserializeBinary
);


/**
 * @param {!proto.stategate.SearchOpts} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.stategate.Events)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.stategate.Events>|undefined}
 *     The XHR Node Readable Stream
 */
proto.stategate.StateGateServiceClient.prototype.searchEvents =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/stategate.StateGateService/SearchEvents',
      request,
      metadata || {},
      methodDescriptor_StateGateService_SearchEvents,
      callback);
};


/**
 * @param {!proto.stategate.SearchOpts} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.stategate.Events>}
 *     Promise that resolves to the response
 */
proto.stategate.StateGateServicePromiseClient.prototype.searchEvents =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/stategate.StateGateService/SearchEvents',
      request,
      metadata || {},
      methodDescriptor_StateGateService_SearchEvents);
};


module.exports = proto.stategate;

