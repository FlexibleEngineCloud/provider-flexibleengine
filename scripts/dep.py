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
from prompt_toolkit import prompt
from prompt_toolkit.completion import Completer, Completion, PathCompleter
from prompt_toolkit.document import Document
# from prompt_toolkit.contrib.completers import PathCompleter


def main():

    if len(sys.argv) != 2:
        print("Usage: python3 dep.py <groupResource>/<resource>")
        exit(1)

    debug = False

    fileToApply = []
    Labels = {}
    # Labels = defaultdict(dict)
    listLabelsRequired = []
    listRefsRequired = []

    groupResource, resource = sys.argv[1].split("/")

    # if file examples/<groupResource>/<resource>.yaml does not exist
    if not os.path.isfile("examples/" + groupResource + "/" + resource + ".yaml"):
        print("Error: File examples/" + groupResource +
              "/" + resource + ".yaml does not exist")
        exit(1)

    # Read all yaml files in the directory examples/ and decode them
    # List all directory in the directory examples/
    for dir in os.listdir("examples/"):
        # dir is directory
        if os.path.isdir("examples/" + dir) and dir != "providerconfig" and dir != 'storeconfig':
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
                                        if i["kind"] + "." + i["apiVersion"] not in Labels:
                                            Labels[i["kind"] + "." +
                                                   i["apiVersion"]] = {}

                                        Labels[i["kind"] + "." + i["apiVersion"]
                                               ]["label"] = i["metadata"]["labels"]["testing.upbound.io/example-name"]
                                        Labels[i["kind"] + "." + i["apiVersion"]
                                               ]["name"] = i["metadata"]["name"]
                                        Labels[i["kind"] + "." +
                                               i["apiVersion"]]["specs"] = []
                                        Labels[i["kind"] + "." +
                                               i["apiVersion"]]["refs"] = []
                                        for k, v in i["spec"]["forProvider"].items():
                                            if v == None:
                                                continue
                                            if isinstance(v, list):
                                                for item in v:
                                                    if isinstance(item, dict):
                                                        for key in item:
                                                            if isinstance(item[key], dict):
                                                                if "matchLabels" in item[key]:
                                                                    if "testing.upbound.io/example-name" in item[key]["matchLabels"]:
                                                                        Labels[i["kind"] + "." + i["apiVersion"]]["specs"].append(
                                                                            item[key]["matchLabels"]["testing.upbound.io/example-name"])
                                                                        if groupResource == dir and resource+".yaml" == file:
                                                                            listLabelsRequired.append(
                                                                                item[key]["matchLabels"]["testing.upbound.io/example-name"])
                                                            if isinstance(k, str):
                                                                if re.match(r"^.*Ref(s)?", k):
                                                                    Labels[i["kind"] + "." +
                                                                           i["apiVersion"]]["refs"].append(item[key])
                                                                    if groupResource == dir and resource+".yaml" == file:
                                                                        listRefsRequired.append(
                                                                            item[key])

                                            if isinstance(v, dict):
                                                if "matchLabels" in v:
                                                    if "testing.upbound.io/example-name" in v["matchLabels"]:
                                                        Labels[i["kind"] + "." + i["apiVersion"]]["specs"].append(
                                                            v["matchLabels"]["testing.upbound.io/example-name"])
                                                        if groupResource == dir and resource+".yaml" == file:
                                                            listLabelsRequired.append(
                                                                v["matchLabels"]["testing.upbound.io/example-name"])

                            z = y[0]
                            if z["kind"] + "." + z["apiVersion"] not in Labels:
                                Labels[z["kind"] + "." + z["apiVersion"]] = {}

                            Labels[z["kind"] + "." + z["apiVersion"]
                                   ]["file"] = "examples/" + dir + "/" + file

                        f.close()

    if debug:
        print("Labels:")
        print(Labels)
        print("listLabelsRequired:")
        print(listLabelsRequired)

    listLabelsRequiredFound = []

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
                    listLabelsRequiredFound.append(Labels[kind]["name"])
                    fileToApply.append(Labels[kind]["file"])
                    break

    # If the resource is not found in the list of required resources
    foundErrors = False
    for labelRequire in listLabelsRequired:
        if labelRequire not in listLabelsRequiredFound:
            foundErrors = True
            print("ResourceSelect not found: " + labelRequire)

    listRefsRequiredFound = []

    for refRequire in listRefsRequired:
        for kind in Labels:
            if kind in Labels:
                if Labels[kind]["name"] in refRequire:
                    for s in Labels[kind]["refs"]:
                        if s not in listRefsRequired:
                            listRefsRequired.append(s)

    for refRequire in listRefsRequired:
        for kind in Labels:
            if kind in Labels:
                if Labels[kind]["name"] in refRequire:
                    listRefsRequiredFound.append(Labels[kind]["name"])
                    fileToApply.append(Labels[kind]["file"])
                    break

    for refRequire in listRefsRequired:
        if refRequire not in listRefsRequiredFound:
            foundErrors = True
            print("ResourceRef not found: " + refRequire)

    if foundErrors:
        print("\nError: Some resources are not found")
        print("Please check the list of required resources in the examples directory")
        sys.exit(1)

    if debug:
        print("Labels:")
        print(Labels)
        print("listLabelsRequired:")
        print(listLabelsRequired)
        print("listRefsRequired:")
        print(listRefsRequired)
        print("fileToApply:")
        print(fileToApply)

    # Print the kubectl command to apply all resources
    origin = sys.argv[1]
    fi = ""

    for fil in fileToApply:
        fi += f' -f {fil}'

    kubeFiles = f'-f examples/{origin}.yaml{fi}'
    kubeApplyCmd = f'kubectl apply {kubeFiles}'
    kubeDeleteCmd = f'kubectl delete {kubeFiles}'

    print(kubeApplyCmd)

    apply = input("Execute kubectl apply command (Y/n) : ") or "Y"
    if apply == "Y" or apply == "y" or apply == "yes":
        os.system(kubeApplyCmd)
    else:
        delete = input("Execute kubectl delete command (Y/n) : ") or "n"
        if delete == "Y" or delete == "y" or delete == "yes":
            os.system(kubeDeleteCmd)

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


class CommandCompleter(Completer):
    def __init__(self):
        self.path_completer = PathCompleter()
        self.commands = []
        for dir in os.listdir("examples/"):
            # dir is directory
            if os.path.isdir("examples/" + dir) and dir != "providerconfig" and dir != 'storeconfig':
                # List all files in the directory examples/<dir>
                for file in os.listdir("examples/" + dir):
                    # If the file is a yaml file
                    if file.endswith(".yaml"):
                        # File name without extension
                        self.commands.append(dir + "/" + file[:-5])

        # self.commands = ["vpc/vpc", "vpc/vpcsubnet", "command2"]

    def get_completions(self, document, complete_event):
        if document.text.startswith(".x "):
            sub_doc = Document(document.text[3:])
            yield from (Completion(completion.text, completion.start_position, display=completion.display)
                        for completion
                        in self.path_completer.get_completions(sub_doc, complete_event))
        else:
            yield from (Completion(cmd, -document.cursor_position)
                        for cmd in self.commands
                        if cmd.startswith(document.text))


if __name__ == "__main__":

    # If sys.argv is empty prompt to user to enter the argument
    if len(sys.argv) == 1:
        other_args = {}
        inputObj = prompt(
            "Find example to apply (nat/gateway) >>> ", completer=CommandCompleter(), **other_args)

        sys.argv.append(inputObj)
        print("")

    if len(sys.argv) != 2 or sys.argv[1] == "--help":
        help()
    else:
        main()

    # Exit 0
    exit(0)
