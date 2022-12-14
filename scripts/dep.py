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

# In specific directory find all.dep file and return content
# Return list of content
def loadDepFiles(dir, resource):

    content = []

    for file in os.listdir("examples/" + dir):
        if file.find(f'{resource}.dep') != -1 or file.find('all.dep') != -1:
            x = loadDepFile(dir, file)
            for i in range(len(x)):
                content.append(x[i])

    return content

def loadDepFile(dir, filename):
    content = []
    f = open("examples/" + dir + "/" + filename, "r")
    x = f.readlines()
    for i in range(len(x)):
        content.append(x[i])
    f.close()

    return content

def main():

    fileToApply = []

    if len(sys.argv) != 2:
        print("Usage: python3 dep.py <groupResource>/<resource>")
        exit(1)

    # Get the groupResource and resource
    groupResource, resource = sys.argv[1].split("/")

    lines = loadDepFiles(groupResource, resource)
    lines = list(dict.fromkeys(lines))

    # For each line in the dependency file
    for i in range(len(lines)):
        l = lines[i].strip().split(".")
        kind = l[0]
        apiVersion = ".".join(l[1:])
        group = apiVersion.split(".")[0]

        # Finds in all files in the directory examples/ the file that contains the apiVersion and the kind
        for file in os.listdir("examples/" + group):
            if file.endswith(".yaml"):
                x = open("examples/" + group + "/" + file, "r")
                content = x.readlines()
                x.close()

                # If the file contains strict the apiVersion and the kind 
                # apiVersion: <apiVersion>
                # kind: <kind>
                # in File check first line is equal to apiVersion and second line is equal to kind
                foundApiVersion = False
                foundKind = False

                if content[0].strip() == f"apiVersion: {apiVersion}":
                    foundApiVersion = True
                if content[1].strip() == f"kind: {kind}":
                    foundKind = True

                if foundApiVersion and foundKind:
                    fileToApply.append("examples/" + group + "/" + file)
                    
    # Print the kubectl command to apply all resources
    origin = sys.argv[1]
    fi = ""

    for fil in fileToApply:
        fi += f' -f {fil}'

    print(f'kubectl apply -f examples/{origin}.yaml{fi}')


# Help func 
def help():
    print("This script generate kubectl command to applys all necessary resources\n")
    print("Usage: python3 dep.py <groupResource>/<resource>")
    print("Example: python3 dep.py vpc/subnet")
    print("--------------------------------------")
    print(".dep file example:")
    print("VPCSubnet.vpc.flexibleengine.upbound.io/v1beta1")
    print("VPC.vpc.flexibleengine.upbound.io/v1beta1")
    print("--------------------------------------")
    print("all.dep file is a special file that contains all resources that are necessary to apply in the specific directory")
    print("all.dep file example:")
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