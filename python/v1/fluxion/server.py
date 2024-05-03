import logging
from concurrent import futures

import grpc

from fluxion.protos import fluxion_pb2_grpc


class FluxionServicer(fluxion_pb2_grpc.FluxionServicer):
    """
    Unimplemented server - let us know if you need this, but largely
    you should use Go for the serve and python for a client only.
    """

    pass


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    api_pb2_grpc.add_RainbowSchedulerServicer_to_server(FluxionServicer(), server)
    server.add_insecure_port("[::]:4242")
    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    logging.basicConfig()
    serve()
