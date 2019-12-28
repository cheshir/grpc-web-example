/**
 * @fileoverview gRPC-Web generated client stub for currenttime
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')
const proto = {};
proto.currenttime = require('./currenttimer_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.currenttime.CurrentTimerClient =
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
proto.currenttime.CurrentTimerPromiseClient =
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
 *   !proto.google.protobuf.Empty,
 *   !proto.currenttime.CurrentTimeResponse>}
 */
const methodDescriptor_CurrentTimer_GetCurrentTime = new grpc.web.MethodDescriptor(
  '/currenttime.CurrentTimer/GetCurrentTime',
  grpc.web.MethodType.UNARY,
  google_protobuf_empty_pb.Empty,
  proto.currenttime.CurrentTimeResponse,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.currenttime.CurrentTimeResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.google.protobuf.Empty,
 *   !proto.currenttime.CurrentTimeResponse>}
 */
const methodInfo_CurrentTimer_GetCurrentTime = new grpc.web.AbstractClientBase.MethodInfo(
  proto.currenttime.CurrentTimeResponse,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.currenttime.CurrentTimeResponse.deserializeBinary
);


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.currenttime.CurrentTimeResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.currenttime.CurrentTimeResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.currenttime.CurrentTimerClient.prototype.getCurrentTime =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/currenttime.CurrentTimer/GetCurrentTime',
      request,
      metadata || {},
      methodDescriptor_CurrentTimer_GetCurrentTime,
      callback);
};


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.currenttime.CurrentTimeResponse>}
 *     A native promise that resolves to the response
 */
proto.currenttime.CurrentTimerPromiseClient.prototype.getCurrentTime =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/currenttime.CurrentTimer/GetCurrentTime',
      request,
      metadata || {},
      methodDescriptor_CurrentTimer_GetCurrentTime);
};


module.exports = proto.currenttime;

