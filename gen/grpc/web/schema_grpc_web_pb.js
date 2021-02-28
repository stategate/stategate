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
 *   !proto.stategate.Entity,
 *   !proto.stategate.Entity>}
 */
const methodDescriptor_EntityService_Edit = new grpc.web.MethodDescriptor(
  '/stategate.EntityService/Edit',
  grpc.web.MethodType.UNARY,
  proto.stategate.Entity,
  proto.stategate.Entity,
  /**
   * @param {!proto.stategate.Entity} request
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
 *   !proto.stategate.Entity,
 *   !proto.stategate.Entity>}
 */
const methodInfo_EntityService_Edit = new grpc.web.AbstractClientBase.MethodInfo(
  proto.stategate.Entity,
  /**
   * @param {!proto.stategate.Entity} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.Entity.deserializeBinary
);


/**
 * @param {!proto.stategate.Entity} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.stategate.Entity)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.stategate.Entity>|undefined}
 *     The XHR Node Readable Stream
 */
proto.stategate.EntityServiceClient.prototype.edit =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/stategate.EntityService/Edit',
      request,
      metadata || {},
      methodDescriptor_EntityService_Edit,
      callback);
};


/**
 * @param {!proto.stategate.Entity} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.stategate.Entity>}
 *     Promise that resolves to the response
 */
proto.stategate.EntityServicePromiseClient.prototype.edit =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/stategate.EntityService/Edit',
      request,
      metadata || {},
      methodDescriptor_EntityService_Edit);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.stategate.EventRef,
 *   !proto.stategate.Entity>}
 */
const methodDescriptor_EntityService_Revert = new grpc.web.MethodDescriptor(
  '/stategate.EntityService/Revert',
  grpc.web.MethodType.UNARY,
  proto.stategate.EventRef,
  proto.stategate.Entity,
  /**
   * @param {!proto.stategate.EventRef} request
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
 *   !proto.stategate.EventRef,
 *   !proto.stategate.Entity>}
 */
const methodInfo_EntityService_Revert = new grpc.web.AbstractClientBase.MethodInfo(
  proto.stategate.Entity,
  /**
   * @param {!proto.stategate.EventRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.Entity.deserializeBinary
);


/**
 * @param {!proto.stategate.EventRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.stategate.Entity)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.stategate.Entity>|undefined}
 *     The XHR Node Readable Stream
 */
proto.stategate.EntityServiceClient.prototype.revert =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/stategate.EntityService/Revert',
      request,
      metadata || {},
      methodDescriptor_EntityService_Revert,
      callback);
};


/**
 * @param {!proto.stategate.EventRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.stategate.Entity>}
 *     Promise that resolves to the response
 */
proto.stategate.EntityServicePromiseClient.prototype.revert =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/stategate.EntityService/Revert',
      request,
      metadata || {},
      methodDescriptor_EntityService_Revert);
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
 *   !proto.stategate.SearchEntityOpts,
 *   !proto.stategate.Entities>}
 */
const methodDescriptor_EntityService_Search = new grpc.web.MethodDescriptor(
  '/stategate.EntityService/Search',
  grpc.web.MethodType.UNARY,
  proto.stategate.SearchEntityOpts,
  proto.stategate.Entities,
  /**
   * @param {!proto.stategate.SearchEntityOpts} request
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
 *   !proto.stategate.SearchEntityOpts,
 *   !proto.stategate.Entities>}
 */
const methodInfo_EntityService_Search = new grpc.web.AbstractClientBase.MethodInfo(
  proto.stategate.Entities,
  /**
   * @param {!proto.stategate.SearchEntityOpts} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.Entities.deserializeBinary
);


/**
 * @param {!proto.stategate.SearchEntityOpts} request The
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
 * @param {!proto.stategate.SearchEntityOpts} request The
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
 *   !proto.stategate.StreamEventOpts,
 *   !proto.stategate.Event>}
 */
const methodDescriptor_EventService_Stream = new grpc.web.MethodDescriptor(
  '/stategate.EventService/Stream',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.stategate.StreamEventOpts,
  proto.stategate.Event,
  /**
   * @param {!proto.stategate.StreamEventOpts} request
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
 *   !proto.stategate.StreamEventOpts,
 *   !proto.stategate.Event>}
 */
const methodInfo_EventService_Stream = new grpc.web.AbstractClientBase.MethodInfo(
  proto.stategate.Event,
  /**
   * @param {!proto.stategate.StreamEventOpts} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.Event.deserializeBinary
);


/**
 * @param {!proto.stategate.StreamEventOpts} request The request proto
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
 * @param {!proto.stategate.StreamEventOpts} request The request proto
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


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.stategate.EventRef,
 *   !proto.stategate.Event>}
 */
const methodDescriptor_EventService_Get = new grpc.web.MethodDescriptor(
  '/stategate.EventService/Get',
  grpc.web.MethodType.UNARY,
  proto.stategate.EventRef,
  proto.stategate.Event,
  /**
   * @param {!proto.stategate.EventRef} request
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
 *   !proto.stategate.EventRef,
 *   !proto.stategate.Event>}
 */
const methodInfo_EventService_Get = new grpc.web.AbstractClientBase.MethodInfo(
  proto.stategate.Event,
  /**
   * @param {!proto.stategate.EventRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.Event.deserializeBinary
);


/**
 * @param {!proto.stategate.EventRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.stategate.Event)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.stategate.Event>|undefined}
 *     The XHR Node Readable Stream
 */
proto.stategate.EventServiceClient.prototype.get =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/stategate.EventService/Get',
      request,
      metadata || {},
      methodDescriptor_EventService_Get,
      callback);
};


/**
 * @param {!proto.stategate.EventRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.stategate.Event>}
 *     Promise that resolves to the response
 */
proto.stategate.EventServicePromiseClient.prototype.get =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/stategate.EventService/Get',
      request,
      metadata || {},
      methodDescriptor_EventService_Get);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.stategate.PeerServiceClient =
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
proto.stategate.PeerServicePromiseClient =
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
 *   !proto.stategate.Message,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_PeerService_Broadcast = new grpc.web.MethodDescriptor(
  '/stategate.PeerService/Broadcast',
  grpc.web.MethodType.UNARY,
  proto.stategate.Message,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.stategate.Message} request
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
 *   !proto.stategate.Message,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_PeerService_Broadcast = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.stategate.Message} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.stategate.Message} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.stategate.PeerServiceClient.prototype.broadcast =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/stategate.PeerService/Broadcast',
      request,
      metadata || {},
      methodDescriptor_PeerService_Broadcast,
      callback);
};


/**
 * @param {!proto.stategate.Message} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.stategate.PeerServicePromiseClient.prototype.broadcast =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/stategate.PeerService/Broadcast',
      request,
      metadata || {},
      methodDescriptor_PeerService_Broadcast);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.stategate.StreamMessageOpts,
 *   !proto.stategate.PeerMessage>}
 */
const methodDescriptor_PeerService_Stream = new grpc.web.MethodDescriptor(
  '/stategate.PeerService/Stream',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.stategate.StreamMessageOpts,
  proto.stategate.PeerMessage,
  /**
   * @param {!proto.stategate.StreamMessageOpts} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.PeerMessage.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.stategate.StreamMessageOpts,
 *   !proto.stategate.PeerMessage>}
 */
const methodInfo_PeerService_Stream = new grpc.web.AbstractClientBase.MethodInfo(
  proto.stategate.PeerMessage,
  /**
   * @param {!proto.stategate.StreamMessageOpts} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.PeerMessage.deserializeBinary
);


/**
 * @param {!proto.stategate.StreamMessageOpts} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.stategate.PeerMessage>}
 *     The XHR Node Readable Stream
 */
proto.stategate.PeerServiceClient.prototype.stream =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/stategate.PeerService/Stream',
      request,
      metadata || {},
      methodDescriptor_PeerService_Stream);
};


/**
 * @param {!proto.stategate.StreamMessageOpts} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.stategate.PeerMessage>}
 *     The XHR Node Readable Stream
 */
proto.stategate.PeerServicePromiseClient.prototype.stream =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/stategate.PeerService/Stream',
      request,
      metadata || {},
      methodDescriptor_PeerService_Stream);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.stategate.CacheServiceClient =
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
proto.stategate.CacheServicePromiseClient =
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
 *   !proto.stategate.Cache,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_CacheService_Set = new grpc.web.MethodDescriptor(
  '/stategate.CacheService/Set',
  grpc.web.MethodType.UNARY,
  proto.stategate.Cache,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.stategate.Cache} request
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
 *   !proto.stategate.Cache,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_CacheService_Set = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.stategate.Cache} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.stategate.Cache} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.stategate.CacheServiceClient.prototype.set =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/stategate.CacheService/Set',
      request,
      metadata || {},
      methodDescriptor_CacheService_Set,
      callback);
};


/**
 * @param {!proto.stategate.Cache} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.stategate.CacheServicePromiseClient.prototype.set =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/stategate.CacheService/Set',
      request,
      metadata || {},
      methodDescriptor_CacheService_Set);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.stategate.CacheRef,
 *   !proto.stategate.Cache>}
 */
const methodDescriptor_CacheService_Get = new grpc.web.MethodDescriptor(
  '/stategate.CacheService/Get',
  grpc.web.MethodType.UNARY,
  proto.stategate.CacheRef,
  proto.stategate.Cache,
  /**
   * @param {!proto.stategate.CacheRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.Cache.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.stategate.CacheRef,
 *   !proto.stategate.Cache>}
 */
const methodInfo_CacheService_Get = new grpc.web.AbstractClientBase.MethodInfo(
  proto.stategate.Cache,
  /**
   * @param {!proto.stategate.CacheRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stategate.Cache.deserializeBinary
);


/**
 * @param {!proto.stategate.CacheRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.stategate.Cache)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.stategate.Cache>|undefined}
 *     The XHR Node Readable Stream
 */
proto.stategate.CacheServiceClient.prototype.get =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/stategate.CacheService/Get',
      request,
      metadata || {},
      methodDescriptor_CacheService_Get,
      callback);
};


/**
 * @param {!proto.stategate.CacheRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.stategate.Cache>}
 *     Promise that resolves to the response
 */
proto.stategate.CacheServicePromiseClient.prototype.get =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/stategate.CacheService/Get',
      request,
      metadata || {},
      methodDescriptor_CacheService_Get);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.stategate.CacheRef,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_CacheService_Del = new grpc.web.MethodDescriptor(
  '/stategate.CacheService/Del',
  grpc.web.MethodType.UNARY,
  proto.stategate.CacheRef,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.stategate.CacheRef} request
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
 *   !proto.stategate.CacheRef,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_CacheService_Del = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.stategate.CacheRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.stategate.CacheRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.stategate.CacheServiceClient.prototype.del =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/stategate.CacheService/Del',
      request,
      metadata || {},
      methodDescriptor_CacheService_Del,
      callback);
};


/**
 * @param {!proto.stategate.CacheRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.stategate.CacheServicePromiseClient.prototype.del =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/stategate.CacheService/Del',
      request,
      metadata || {},
      methodDescriptor_CacheService_Del);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.stategate.MutexServiceClient =
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
proto.stategate.MutexServicePromiseClient =
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
 *   !proto.stategate.Mutex,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_MutexService_Lock = new grpc.web.MethodDescriptor(
  '/stategate.MutexService/Lock',
  grpc.web.MethodType.UNARY,
  proto.stategate.Mutex,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.stategate.Mutex} request
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
 *   !proto.stategate.Mutex,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_MutexService_Lock = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.stategate.Mutex} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.stategate.Mutex} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.stategate.MutexServiceClient.prototype.lock =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/stategate.MutexService/Lock',
      request,
      metadata || {},
      methodDescriptor_MutexService_Lock,
      callback);
};


/**
 * @param {!proto.stategate.Mutex} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.stategate.MutexServicePromiseClient.prototype.lock =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/stategate.MutexService/Lock',
      request,
      metadata || {},
      methodDescriptor_MutexService_Lock);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.stategate.MutexRef,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_MutexService_Unlock = new grpc.web.MethodDescriptor(
  '/stategate.MutexService/Unlock',
  grpc.web.MethodType.UNARY,
  proto.stategate.MutexRef,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.stategate.MutexRef} request
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
 *   !proto.stategate.MutexRef,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_MutexService_Unlock = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.stategate.MutexRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.stategate.MutexRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.stategate.MutexServiceClient.prototype.unlock =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/stategate.MutexService/Unlock',
      request,
      metadata || {},
      methodDescriptor_MutexService_Unlock,
      callback);
};


/**
 * @param {!proto.stategate.MutexRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.stategate.MutexServicePromiseClient.prototype.unlock =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/stategate.MutexService/Unlock',
      request,
      metadata || {},
      methodDescriptor_MutexService_Unlock);
};


module.exports = proto.stategate;

