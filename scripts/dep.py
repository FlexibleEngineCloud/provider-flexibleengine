#!/usr/bin/env python3

# This script generate kubectl command to applys all necessary resources
# For determining the resources, read the file "<arg>.dep" and extract on each line the resource
# Resource is composed of 3 parts: <kind>.<apiVersion>
# Example: VPCSubnet.vpc.flexibleengine.upbound.io/v1beta1 -- Kind: VPCSubnet, APIVersion: vpc.flexibleengine.upbound.io/v1beta1
# The script require 1 argument: <groupResource>/<resource>
# Example: vpc/subnet.yaml

import sys
import os
import re

def main():

    fileToApply = []

    if len(sys.argv) != 2:
        print("Usage: python3 dep.py <groupResource>/<resource>")
        exit(1)

    # Get the groupResource and resource
    groupResource, resource = sys.argv[1].split("/")
    # # Get the directory of the resource
    # dir = os.path.dirname(os.path.realpath(resource))
    # # Get the filename of the resource
    # filename = os.path.basename(os.path.realpath(resource))
    # # Get the name of the resource
    # name = os.path.splitext(filename)[0]

    # Get the dependency file
    depFile = "examples/" + groupResource + "/" + resource + ".dep"
    # Open the dependency file
    f = open(depFile, "r")
    lines = f.readlines()
    # # Close the dependency file
    f.close()


    # For each line in the dependency file
    for i in range(len(lines)):
        line = lines[i].strip()
        # Get the kind and the APIVersion
        l = line.split(".")
        kind = l[0]
        # apiVersions end 
        apiVersion = ".".join(l[1:])
        # In resource split by dot and get the first part
        group = apiVersion.split(".")[0]

        # Finds in all files in the directory examples/ the file that contains the apiVersion and the kind
        for file in os.listdir("examples/" + group):
            if file.endswith(".yaml"):
                # Open the file
                x = open("examples/" + group + "/" + file, "r")
                # Read the content of the file
                content = x.read()
                # Close the file
                x.close()

                # If the file contains strict the apiVersion and the kind 
                # apiVersion: <apiVersion>
                # kind: <kind>
                # in File check first line is equal to apiVersion and second line is equal to kind
                if content.splitlines()[0] == "apiVersion: " + apiVersion and content.splitlines()[1] == "kind: " + kind:
                    fileToApply.append("examples/" + group + "/" + file)
                    
    # Print the kubectl command to apply all resources
    origin = sys.argv[1]
    fi = ""

    for fil in fileToApply:
        fi += f' -f {fil}'

    print(f'kubectl apply  -f examples/{origin}.yaml {fi}')


# Help func 
def help():
    print("This script generate kubectl command to applys all necessary resources\n")
    print("Usage: python3 dep.py <groupResource>/<resource>")
    print("Example: python3 dep.py vpc/subnet")
    print("--------------------------------------")
    print(".dep file example:")
    print("VPCSubnet.vpc.flexibleengine.upbound.io/v1beta1")
    print("VPC.vpc.flexibleengine.upbound.io/v1beta1")

if __name__ == "__main__":

    # If sys.argv is empty prompt to user to enter the argument
    if len(sys.argv) == 1:
        print("Enter name (nat/gateway): ")
        sys.argv.append(input())





    if len(sys.argv) != 2 or sys.argv[1] == "--help":
        help()
    else:
        main()

    # Exit 0 
    exit(0)