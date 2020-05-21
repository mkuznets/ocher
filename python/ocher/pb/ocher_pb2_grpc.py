# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from ocher.pb import ocher_pb2 as ocher_dot_pb_dot_ocher__pb2


class OcherStub(object):
    """Missing associated documentation comment in .proto file"""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.QueueTransactional = channel.stream_stream(
                '/ocher.Ocher/QueueTransactional',
                request_serializer=ocher_dot_pb_dot_ocher__pb2.Status.SerializeToString,
                response_deserializer=ocher_dot_pb_dot_ocher__pb2.Task.FromString,
                )


class OcherServicer(object):
    """Missing associated documentation comment in .proto file"""

    def QueueTransactional(self, request_iterator, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_OcherServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'QueueTransactional': grpc.stream_stream_rpc_method_handler(
                    servicer.QueueTransactional,
                    request_deserializer=ocher_dot_pb_dot_ocher__pb2.Status.FromString,
                    response_serializer=ocher_dot_pb_dot_ocher__pb2.Task.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'ocher.Ocher', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Ocher(object):
    """Missing associated documentation comment in .proto file"""

    @staticmethod
    def QueueTransactional(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_stream(request_iterator, target, '/ocher.Ocher/QueueTransactional',
            ocher_dot_pb_dot_ocher__pb2.Status.SerializeToString,
            ocher_dot_pb_dot_ocher__pb2.Task.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)
