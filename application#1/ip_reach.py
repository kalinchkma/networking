import sys
import subprocess

# checking IP reachability
def ip_reach(iplist):

    for ip in iplist:
        ip = ip.rstrip("\n")

        ping_reply = subprocess.call('ping %s -c 2' % (ip), stdout = subprocess.DEVNULL, stderr = subprocess.DEVNULL, shell=True)

        if ping_reply == 0:
            print(f"\n* {ip} is reachable :)\n")
            continue
        else:
            print(f"\n* {ip} not reachable :(check connectivity and try again.)")
            sys.exit()


