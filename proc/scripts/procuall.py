#!/usr/bin/python
import os
from pathlib import Path

__doc__ ="""procuall.py: enumerate all processes run by user

usage: python procuall.py $USER
"""

def search_proc():
    """for each process under /proc
       parse the status file and try to match uid
       if there is a match, the process belongs to the user"""
    uid = os.getuid()
    all = Path("/proc")
    procs = all.glob("[0-9]*")
    processes = [p for p in procs]
    info = {}
    for proc in processes:
        status = proc.joinpath("status")
        with open(status) as f:
            lines = f.readlines()
            cmd = lines[0].split("\t")
            name = cmd[1].strip("\n")
            uidl = lines[8].split("\t")
            puid = uidl[1].strip("\n")
            if int(puid) == uid:
                info.update({name: uid})
    return info

def main():
    """Get info about running processes
    output the command being run and the pid"""
    info = search_proc()
    for k, v in info.items():
        print(f"cmd: {k}, pid: {v}")

if __name__ == '__main__':
    main()
