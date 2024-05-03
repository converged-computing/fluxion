import argparse
import os
import sys

from fluxion.protos import fluxion_pb2
from fluxion.client import FluxionClient
import fluxion.utils as utils
import time

# Config file from a few directories up
here = os.path.abspath(os.path.dirname(__file__))
root = here

# examples directory
root = os.path.dirname(root)


def get_parser():
    parser = argparse.ArgumentParser(description="🦩️ Fluxion Python client example")
    parser.add_argument(
        "--host", help="host of fluxion graph database", default="localhost:4242"
    )
    parser.add_argument(
        "--jobspec",
        help="Path to example jobspec to use",
        default=os.path.join(root, "go", "jobspec.yaml"),
    )
    parser.add_argument(
        "--cluster-nodes",
        help="Nodes to provide to initialize graph",
        default=os.path.join(root, "go", "cluster-nodes.json"),
    )
    return parser


def main():
    parser = get_parser()
    args, _ = parser.parse_known_args()
    cli = FluxionClient(host=args.host)

    # Read in the jobspec and nodes to raw text
    jobspec = utils.read_file(args.jobspec)
    nodes_jgf = utils.read_file(args.cluster_nodes)

    # Step 1: Init the fluxion graph
    response = cli.init(nodes_jgf)
    if response.status == fluxion_pb2.InitResponse.ResultType.INIT_SUCCESS:
        print("✅️ Init of Fluxion resource graph success!")
    else:
        sys.exit(f"Issue with init, return code {response.status}")

    # Step 2: Do a match
    response = cli.match(jobspec)
    if response.status == fluxion_pb2.MatchResponse.ResultType.MATCH_SUCCESS:
        print("✅️ Match of jobspec to Fluxion graph success!")
    else:
        sys.exit(f"Issue with match, return code {response.status}")

    print("😴️ Sleeping for 3 seconds before cancel...")
    time.sleep(3)

    # Step 3: cancel the job
    jobid = response.jobid
    response = cli.cancel(jobid=jobid)
    if response.status == fluxion_pb2.CancelResponse.ResultType.CANCEL_SUCCESS:
        print(f"✅️ Cancel of jobid {jobid} success!")
    else:
        sys.exit(f"Issue with cancel, return code {response.status}")

    print("👋️ That's all folks!")


if __name__ == "__main__":
    main()
