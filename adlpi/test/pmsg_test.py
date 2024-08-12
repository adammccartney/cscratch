#!/usr/bin/python3
import argparse
import logging
import shlex
import subprocess as sp

__doc__ = """Simple use cases for POSIX message queues.

Assumes that the various pmsg binaries are installed on the system.
"""

logger = logging.getLogger(__name__)

def global_env():
    env = {}
    env.update({
        "pmsg_send_cmd": lambda mq, msg: f"pmsg_send {mq} {msg}",
        "pmsg_receive_cmd": lambda mq: f"pmsg_receive {mq}",
        })
    return env


environment = global_env()


def pmsg_send(mq: str, msg: str):
    """Send MSG to POSIX message queue MQ"""
    try:
        global environment
        pmsg_send_cmd = environment.get("pmsg_send_cmd")
        cmd = pmsg_send_cmd(mq, msg)
        cmd_tokens = shlex.split(cmd)
        proc = sp.run(cmd_tokens, capture_output=True)
        logger.info(f"queued message: {msg} to {mq}")
    except Exception as e:
        logger.error(f"failed to queue message, exit with: {e}")
        if proc:
            logger.error(f"Subprocess exited with {proc.stderr}")


def pmsg_receive(mq: str) -> str:
    """Receive a message from the POSIX message queue MQ

    return the message as a string
    """
    try:
        global environment
        pmsg_receive_cmd = environment.get("pmsg_receive_cmd")
        cmd = pmsg_receive_cmd(mq)
        cmd_tokens = shlex.split(cmd)
        proc = sp.run(cmd_tokens, capture_output=True)
        result = proc.stdout
        logger.info(f"received message: {result}")
        return result
    except Exception as e:
        logger.error(f"failed to receive message, exit with {e}")



def parse_arguments():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("-q", "--message-queue", help="Name of posix message queue (must exist on system already)")
    parser.add_argument("-m", "--message", help="Message to send to queue")
    return parser.parse_args()

        
def main():
    args = parse_arguments()
    pmsg_send(args.message_queue, args.message)
    result = pmsg_receive(args.message_queue)
    print(result)

if __name__ == '__main__':
    main()
