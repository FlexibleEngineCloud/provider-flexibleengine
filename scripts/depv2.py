#!/usr/bin/env python3

# This script generate kubectl command to applys all necessary resources
# For determining the resources, read the file "<arg>.dep" and extract on each line the resource
# Resource is composed of 3 parts: <kind>.<apiVersion>
# Example: VPCSubnet.vpc.flexibleengine.upbound.io/v1beta1 -- Kind: VPCSubnet, APIVersion: vpc.flexibleengine.upbound.io/v1beta1
# The script require 1 argument: <groupResource>/<resource>
# Example: vpc/subnet.yaml

import sys
import os
import yaml
from yaml.loader import SafeLoader
import re
from collections import defaultdict



def main():


    if len(sys.argv) != 2:
        print("Usage: python3 dep.py <groupResource>/<resource>")
        exit(1)

    debug = False

    fileToApply = []
    Labels = {}
    # Labels = defaultdict(dict)
    listLabels = {}
    listLabelsRequired = []

    groupResource, resource = sys.argv[1].split("/")

    # if file examples/<groupResource>/<resource>.yaml does not exist
    if not os.path.isfile("examples/" + groupResource + "/" + resource):
        print("Error: File examples/" + groupResource + "/" + resource + " does not exist")
        exit(1)

    # Read all yaml files in the directory examples/ and decode them
    # List all directory in the directory examples/
    for dir in os.listdir("examples/"):
        # dir is directory
        if os.path.isdir("examples/" + dir) and dir != "providerconfig" and dir != 'storeconfig' :
            # List all files in the directory examples/<dir>
            for file in os.listdir("examples/" + dir):
                # If the file is a yaml file
                if file.endswith(".yaml"):
                    with open("examples/" + dir + "/" + file, "r") as f:
                        y = list(yaml.load_all(f, Loader=SafeLoader))
                        # If the file is not empty
                        if y != None:

                            for i in y:
                                if i != None:
                                    if "metadata" in i and "labels" in i["metadata"] and "testing.upbound.io/example-name" in i["metadata"]["labels"]:
                                        # if i["metadata"]["labels"]["testing.upbound.io/example-name"] in listLabels and listLabels[i["metadata"]["labels"]["testing.upbound.io/example-name"]] != i["kind"] + "." + i["apiVersion"]:
                                        #     print("Error: Duplicate label testing.upbound.io/example-name: " + i["metadata"]["labels"]["testing.upbound.io/example-name"])
                                        #     print("Example: " + listLabels[i["metadata"]["labels"]["testing.upbound.io/example-name"]] + " and " + i["kind"] + "." + i["apiVersion"])
                                        #     exit(1)
                                        # else:

                                            # if Labels[i["kind"] + "." + i["apiVersion"]] == None:
                                            if i["kind"] + "." + i["apiVersion"] not in Labels:
                                                Labels[i["kind"] + "." + i["apiVersion"]] = {}

                                            Labels[i["kind"] + "." + i["apiVersion"]]["label"] = i["metadata"]["labels"]["testing.upbound.io/example-name"]
                                            Labels[i["kind"] + "." + i["apiVersion"]]["specs"] = []
                                            for k, v in i["spec"]["forProvider"].items():
                                                if v == None:
                                                    continue
                                                if isinstance(v, dict):
                                                    if "matchLabels" in v:
                                                        if "testing.upbound.io/example-name" in v["matchLabels"]:
                                                            Labels[i["kind"] + "." + i["apiVersion"]]["specs"].append(v["matchLabels"]["testing.upbound.io/example-name"])
                                                            if groupResource == dir and resource+".yaml" == file:
                                                                listLabelsRequired.append(v["matchLabels"]["testing.upbound.io/example-name"])

                            z = y[0]
                            if z["kind"] + "." + z["apiVersion"] not in Labels:
                                Labels[z["kind"] + "." + z["apiVersion"]] = {}

                            Labels[z["kind"] + "." + z["apiVersion"]]["file"] = "examples/" + dir + "/" + file

                        f.close()

    if debug:
        print("Labels:")
        print(Labels)
        # print("listLabels:")
        # print(listLabels)
        print("listLabelsRequired:")
        print(listLabelsRequired)

    for labelRequire in listLabelsRequired:
        for kind in Labels:
            if kind in Labels:
                if Labels[kind]["label"] in labelRequire:
                    for s in Labels[kind]["specs"]:
                        if s not in listLabelsRequired:
                            listLabelsRequired.append(s)

    for labelRequire in listLabelsRequired:
        for kind in Labels:
            if kind in Labels:
                if Labels[kind]["label"] in labelRequire:
                    fileToApply.append(Labels[kind]["file"])
                    break

    if debug:
        print("Labels:")
        print(Labels)
        print("listLabels:")
        print(listLabels)
        print("listLabelsRequired:")
        print(listLabelsRequired)
        print("fileToApply:")
        print(fileToApply)


    # Print the kubectl command to apply all resources
    origin = sys.argv[1]
    fi = ""

    for fil in fileToApply:
        fi += f' -f {fil}'

    print(f'kubectl apply -f examples/{origin}.yaml{fi}')

    return

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