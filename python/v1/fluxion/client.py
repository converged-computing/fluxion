import json
import os

import grpc

import fluxion.utils as utils
from fluxion.protos import fluxion_pb2, fluxion_pb2_grpc


class FluxionClient:
    """
    A FluxionClient is able to interact with a Fluxion cluster
    """

    def __init__(self, host="localhost:4242"):
        """
        Create a new rainbow client to interact with a rainbow cluster.
        """
        self.host = host

    def match(self, jobspec, request="allocate"):
        """
        Run a match for a jobspec
        """
        if os.path.exists(jobspec):
            jobspec = utils.read_file(jobspec)

        match_request = fluxion_pb2.MatchRequest(jobspec=jobspec, request=request, count=1)
        with grpc.insecure_channel(self.host) as channel:
            stub = fluxion_pb2_grpc.FluxionServiceStub(channel)
            response = stub.Match(match_request)
        return response

    def init(self, cluster_nodes):
        """
        Init Fluxion with a cluster nodes graph.
        """
        # Be forgiving if a filename is provided - try to read in
        if os.path.exists(cluster_nodes):
            cluster_nodes = utils.read_file(cluster_nodes)

        if not cluster_nodes:
            raise ValueError("Loaded cluster nodes are required")

        # These are the variables for our cluster - name for now
        request = fluxion_pb2.InitRequest(jgf=cluster_nodes)
        with grpc.insecure_channel(self.host) as channel:
            stub = fluxion_pb2_grpc.FluxionServiceStub(channel)
            response = stub.Init(request)
        return response

    def cancel(self, jobid):
        """
        Cancel a job
        """
        if not jobid:
            raise ValueError("A jobid is required")

        # These are the variables for our cluster - name for now
        request = fluxion_pb2.CancelRequest(jobID=jobid)
        with grpc.insecure_channel(self.host) as channel:
            stub = fluxion_pb2_grpc.FluxionServiceStub(channel)
            response = stub.Cancel(request)
        return response
