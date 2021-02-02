/**
 * @fileoverview gRPC-Web generated client stub for cloudEventsProxy
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
proto.cloudEventsProxy = require('./schema_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.cloudEventsProxy.CloudEventsServiceClient =
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
proto.cloudEventsProxy.CloudEventsServicePromiseClient =
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
 *   !proto.cloudEventsProxy.CloudEventInput,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_CloudEventsService_Send = new grpc.web.MethodDescriptor(
  '/cloudEventsProxy.CloudEventsService/Send',
  grpc.web.MethodType.UNARY,
  proto.cloudEventsProxy.CloudEventInput,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.cloudEventsProxy.CloudEventInput} request
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
 *   !proto.cloudEventsProxy.CloudEventInput,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_CloudEventsService_Send = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.cloudEventsProxy.CloudEventInput} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.cloudEventsProxy.CloudEventInput} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.cloudEventsProxy.CloudEventsServiceClient.prototype.send =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/cloudEventsProxy.CloudEventsService/Send',
      request,
      metadata || {},
      methodDescriptor_CloudEventsService_Send,
      callback);
};


/**
 * @param {!proto.cloudEventsProxy.CloudEventInput} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.cloudEventsProxy.CloudEventsServicePromiseClient.prototype.send =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/cloudEventsProxy.CloudEventsService/Send',
      request,
      metadata || {},
      methodDescriptor_CloudEventsService_Send);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.cloudEventsProxy.CloudEventInput,
 *   !proto.cloudEventsProxy.CloudEvent>}
 */
const methodDescriptor_CloudEventsService_Request = new grpc.web.MethodDescriptor(
  '/cloudEventsProxy.CloudEventsService/Request',
  grpc.web.MethodType.UNARY,
  proto.cloudEventsProxy.CloudEventInput,
  proto.cloudEventsProxy.CloudEvent,
  /**
   * @param {!proto.cloudEventsProxy.CloudEventInput} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cloudEventsProxy.CloudEvent.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.cloudEventsProxy.CloudEventInput,
 *   !proto.cloudEventsProxy.CloudEvent>}
 */
const methodInfo_CloudEventsService_Request = new grpc.web.AbstractClientBase.MethodInfo(
  proto.cloudEventsProxy.CloudEvent,
  /**
   * @param {!proto.cloudEventsProxy.CloudEventInput} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cloudEventsProxy.CloudEvent.deserializeBinary
);


/**
 * @param {!proto.cloudEventsProxy.CloudEventInput} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.cloudEventsProxy.CloudEvent)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.cloudEventsProxy.CloudEvent>|undefined}
 *     The XHR Node Readable Stream
 */
proto.cloudEventsProxy.CloudEventsServiceClient.prototype.request =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/cloudEventsProxy.CloudEventsService/Request',
      request,
      metadata || {},
      methodDescriptor_CloudEventsService_Request,
      callback);
};


/**
 * @param {!proto.cloudEventsProxy.CloudEventInput} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.cloudEventsProxy.CloudEvent>}
 *     Promise that resolves to the response
 */
proto.cloudEventsProxy.CloudEventsServicePromiseClient.prototype.request =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/cloudEventsProxy.CloudEventsService/Request',
      request,
      metadata || {},
      methodDescriptor_CloudEventsService_Request);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.cloudEventsProxy.ReceiveRequest,
 *   !proto.cloudEventsProxy.CloudEvent>}
 */
const methodDescriptor_CloudEventsService_Receive = new grpc.web.MethodDescriptor(
  '/cloudEventsProxy.CloudEventsService/Receive',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.cloudEventsProxy.ReceiveRequest,
  proto.cloudEventsProxy.CloudEvent,
  /**
   * @param {!proto.cloudEventsProxy.ReceiveRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cloudEventsProxy.CloudEvent.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.cloudEventsProxy.ReceiveRequest,
 *   !proto.cloudEventsProxy.CloudEvent>}
 */
const methodInfo_CloudEventsService_Receive = new grpc.web.AbstractClientBase.MethodInfo(
  proto.cloudEventsProxy.CloudEvent,
  /**
   * @param {!proto.cloudEventsProxy.ReceiveRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cloudEventsProxy.CloudEvent.deserializeBinary
);


/**
 * @param {!proto.cloudEventsProxy.ReceiveRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.cloudEventsProxy.CloudEvent>}
 *     The XHR Node Readable Stream
 */
proto.cloudEventsProxy.CloudEventsServiceClient.prototype.receive =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/cloudEventsProxy.CloudEventsService/Receive',
      request,
      metadata || {},
      methodDescriptor_CloudEventsService_Receive);
};


/**
 * @param {!proto.cloudEventsProxy.ReceiveRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.cloudEventsProxy.CloudEvent>}
 *     The XHR Node Readable Stream
 */
proto.cloudEventsProxy.CloudEventsServicePromiseClient.prototype.receive =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/cloudEventsProxy.CloudEventsService/Receive',
      request,
      metadata || {},
      methodDescriptor_CloudEventsService_Receive);
};


module.exports = proto.cloudEventsProxy;

