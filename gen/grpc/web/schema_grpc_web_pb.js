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
 *   !proto.eventgate.Event,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_EventGateService_Send = new grpc.web.MethodDescriptor(
  '/eventgate.EventGateService/Send',
  grpc.web.MethodType.UNARY,
  proto.eventgate.Event,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.eventgate.Event} request
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
 *   !proto.eventgate.Event,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_EventGateService_Send = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.eventgate.Event} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.eventgate.Event} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.eventgate.EventGateServiceClient.prototype.send =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/eventgate.EventGateService/Send',
      request,
      metadata || {},
      methodDescriptor_EventGateService_Send,
      callback);
};


/**
 * @param {!proto.eventgate.Event} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.eventgate.EventGateServicePromiseClient.prototype.send =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/eventgate.EventGateService/Send',
      request,
      metadata || {},
      methodDescriptor_EventGateService_Send);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.eventgate.Event,
 *   !proto.eventgate.Event>}
 */
const methodDescriptor_EventGateService_Request = new grpc.web.MethodDescriptor(
  '/eventgate.EventGateService/Request',
  grpc.web.MethodType.UNARY,
  proto.eventgate.Event,
  proto.eventgate.Event,
  /**
   * @param {!proto.eventgate.Event} request
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
 *   !proto.eventgate.Event,
 *   !proto.eventgate.Event>}
 */
const methodInfo_EventGateService_Request = new grpc.web.AbstractClientBase.MethodInfo(
  proto.eventgate.Event,
  /**
   * @param {!proto.eventgate.Event} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.eventgate.Event.deserializeBinary
);


/**
 * @param {!proto.eventgate.Event} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.eventgate.Event)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.eventgate.Event>|undefined}
 *     The XHR Node Readable Stream
 */
proto.eventgate.EventGateServiceClient.prototype.request =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/eventgate.EventGateService/Request',
      request,
      metadata || {},
      methodDescriptor_EventGateService_Request,
      callback);
};


/**
 * @param {!proto.eventgate.Event} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.eventgate.Event>}
 *     Promise that resolves to the response
 */
proto.eventgate.EventGateServicePromiseClient.prototype.request =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/eventgate.EventGateService/Request',
      request,
      metadata || {},
      methodDescriptor_EventGateService_Request);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.eventgate.Filter,
 *   !proto.eventgate.Event>}
 */
const methodDescriptor_EventGateService_Receive = new grpc.web.MethodDescriptor(
  '/eventgate.EventGateService/Receive',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.eventgate.Filter,
  proto.eventgate.Event,
  /**
   * @param {!proto.eventgate.Filter} request
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
 *   !proto.eventgate.Filter,
 *   !proto.eventgate.Event>}
 */
const methodInfo_EventGateService_Receive = new grpc.web.AbstractClientBase.MethodInfo(
  proto.eventgate.Event,
  /**
   * @param {!proto.eventgate.Filter} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.eventgate.Event.deserializeBinary
);


/**
 * @param {!proto.eventgate.Filter} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.eventgate.Event>}
 *     The XHR Node Readable Stream
 */
proto.eventgate.EventGateServiceClient.prototype.receive =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/eventgate.EventGateService/Receive',
      request,
      metadata || {},
      methodDescriptor_EventGateService_Receive);
};


/**
 * @param {!proto.eventgate.Filter} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.eventgate.Event>}
 *     The XHR Node Readable Stream
 */
proto.eventgate.EventGateServicePromiseClient.prototype.receive =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/eventgate.EventGateService/Receive',
      request,
      metadata || {},
      methodDescriptor_EventGateService_Receive);
};


module.exports = proto.eventgate;

