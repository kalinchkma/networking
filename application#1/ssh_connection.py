import paramiko
import os.path
import time
import sys
import re

# checking username/password file
user_file = input("\n# Enter user file path and name (e.g. /demo/file.txt): ")

# verifying the validity of the USERNAME/PASSWORD file
if os.path.isfile(user_file) == True:
    print("\n* Username/Password file is valid: \n")
else:
    print(f"\n* File {user_file} does not exist : (Please check and try again.\n)")
    sys.exit()

# checking commands files
cmd_file = input("\n# Enter commands file path and name (e.g. /demo/file.txt): ")

if os.path.isfile(cmd_file) == True:
    print("\n* Commands file is valid: \n")
else:
    print(f"\n* File {cmd_file} does not exist : (Please check and try again.\n)")
    sys.exit()


# Open SSHv2 connection to the device
def ssh_connection(ip):

    global user_file
    global cmd_file

