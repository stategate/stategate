/**
 * @fileoverview gRPC-Web generated client stub for eventgate
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
proto.eventgate = require('./schema_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.eventgate.EventGateServiceClient =
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
proto.eventgate.EventGateServicePromiseClient =
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
 *   !proto.eventgate.Object,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_EventGateService_SetObject = new grpc.web.MethodDescriptor(
  '/eventgate.EventGateService/SetObject',
  grpc.web.MethodType.UNARY,
  proto.eventgate.Object,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.eventgate.Object} request
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
 *   !proto.eventgate.Object,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_EventGateService_SetObject = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.eventgate.Object} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.eventgate.Object} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.eventgate.EventGateServiceClient.prototype.setObject =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/eventgate.EventGateService/SetObject',
      request,
      metadata || {},
      methodDescriptor_EventGateService_SetObject,
      callback);
};


/**
 * @param {!proto.eventgate.Object} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.eventgate.EventGateServicePromiseClient.prototype.setObject =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/eventgate.EventGateService/SetObject',
      request,
      metadata || {},
      methodDescriptor_EventGateService_SetObject);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.eventgate.ObjectRef,
 *   !proto.eventgate.Object>}
 */
const methodDescriptor_EventGateService_GetObject = new grpc.web.MethodDescriptor(
  '/eventgate.EventGateService/GetObject',
  grpc.web.MethodType.UNARY,
  proto.eventgate.ObjectRef,
  proto.eventgate.Object,
  /**
   * @param {!proto.eventgate.ObjectRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.eventgate.Object.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.eventgate.ObjectRef,
 *   !proto.eventgate.Object>}
 */
const methodInfo_EventGateService_GetObject = new grpc.web.AbstractClientBase.MethodInfo(
  proto.eventgate.Object,
  /**
   * @param {!proto.eventgate.ObjectRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.eventgate.Object.deserializeBinary
);


/**
 * @param {!proto.eventgate.ObjectRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.eventgate.Object)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.eventgate.Object>|undefined}
 *     The XHR Node Readable Stream
 */
proto.eventgate.EventGateServiceClient.prototype.getObject =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/eventgate.EventGateService/GetObject',
      request,
      metadata || {},
      methodDescriptor_EventGateService_GetObject,
      callback);
};


/**
 * @param {!proto.eventgate.ObjectRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.eventgate.Object>}
 *     Promise that resolves to the response
 */
proto.eventgate.EventGateServicePromiseClient.prototype.getObject =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/eventgate.EventGateService/GetObject',
      request,
      metadata || {},
      methodDescriptor_EventGateService_GetObject);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.eventgate.StreamOpts,
 *   !proto.eventgate.Event>}
 */
const methodDescriptor_EventGateService_StreamEvents = new grpc.web.MethodDescriptor(
  '/eventgate.EventGateService/StreamEvents',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.eventgate.StreamOpts,
  proto.eventgate.Event,
  /**
   * @param {!proto.eventgate.StreamOpts} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.eventgate.Event.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.eventgate.StreamOpts,
 *   !proto.eventgate.Event>}
 */
const methodInfo_EventGateService_StreamEvents = new grpc.web.AbstractClientBase.MethodInfo(
  proto.eventgate.Event,
  /**
   * @param {!proto.eventgate.StreamOpts} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.eventgate.Event.deserializeBinary
);


/**
 * @param {!proto.eventgate.StreamOpts} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.eventgate.Event>}
 *     The XHR Node Readable Stream
 */
proto.eventgate.EventGateServiceClient.prototype.streamEvents =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/eventgate.EventGateService/StreamEvents',
      request,
      metadata || {},
      methodDescriptor_EventGateService_StreamEvents);
};


/**
 * @param {!proto.eventgate.StreamOpts} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.eventgate.Event>}
 *     The XHR Node Readable Stream
 */
proto.eventgate.EventGateServicePromiseClient.prototype.streamEvents =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/eventgate.EventGateService/StreamEvents',
      request,
      metadata || {},
      methodDescriptor_EventGateService_StreamEvents);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.eventgate.SearchOpts,
 *   !proto.eventgate.Events>}
 */
const methodDescriptor_EventGateService_SearchEvents = new grpc.web.MethodDescriptor(
  '/eventgate.EventGateService/SearchEvents',
  grpc.web.MethodType.UNARY,
  proto.eventgate.SearchOpts,
  proto.eventgate.Events,
  /**
   * @param {!proto.eventgate.SearchOpts} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.eventgate.Events.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.eventgate.SearchOpts,
 *   !proto.eventgate.Events>}
 */
const methodInfo_EventGateService_SearchEvents = new grpc.web.AbstractClientBase.MethodInfo(
  proto.eventgate.Events,
  /**
   * @param {!proto.eventgate.SearchOpts} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.eventgate.Events.deserializeBinary
);


/**
 * @param {!proto.eventgate.SearchOpts} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.eventgate.Events)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.eventgate.Events>|undefined}
 *     The XHR Node Readable Stream
 */
proto.eventgate.EventGateServiceClient.prototype.searchEvents =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/eventgate.EventGateService/SearchEvents',
      request,
      metadata || {},
      methodDescriptor_EventGateService_SearchEvents,
      callback);
};


/**
 * @param {!proto.eventgate.SearchOpts} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.eventgate.Events>}
 *     Promise that resolves to the response
 */
proto.eventgate.EventGateServicePromiseClient.prototype.searchEvents =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/eventgate.EventGateService/SearchEvents',
      request,
      metadata || {},
      methodDescriptor_EventGateService_SearchEvents);
};


module.exports = proto.eventgate;

