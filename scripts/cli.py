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
import signal
from prompt_toolkit import prompt
from prompt_toolkit.completion import Completer, Completion, PathCompleter
from prompt_toolkit.document import Document
from prompt_toolkit.key_binding import KeyBindings
# from prompt_toolkit.contrib.completers import PathCompleter


def generateListTested():

    def load_gvks(path, loader):
        types = set()
        for root, _, files in os.walk(path):
            for f in files:
                if f.endswith(".yaml"):
                    with open(os.path.join(root, f)) as s:
                        for t in yaml.safe_load_all(s):
                            for gvk in loader(t):
                                types.add(gvk)
        return types

    def load_crd_type(t):
        kind = t["spec"]["names"]["kind"]
        group = t["spec"]["group"]
        for v in t["spec"]["versions"]:
            yield f'{kind}.{group}/{v["name"]}'

    exceptions = {
        "provider-flexibleengine": {
            'ProviderConfigUsage.flexibleengine.upbound.io/v1beta1',
        },
    }

    try:
        exception_set = exceptions["provider-flexibleengine"]
    except KeyError:
        exception_set = set()
    known_crd_types = load_gvks("package/crds/", load_crd_type)

    example_types = load_gvks("examples/", lambda t: [] if t is None or not {"kind", "apiVersion"}.issubset(t.keys())
                              else [f'{t["kind"]}.{t["apiVersion"]}'])

    tested_crd_types = example_types.union(exception_set)

    # for known_crd_types group by group
    # ProtectionGroup.sdrs.flexibleengine.upbound.io/v1beta1
    # ProtectionGroup is the kind
    # sdrs is the group
    # v1beta1 is the version

    mapPrint = {}
    for known_crd_type in known_crd_types:
        # Extract the group from the known_crd_type
        group = known_crd_type.split(".")[1]
        # Extract the kind from the known_crd_type
        kind = known_crd_type.split(".")[0]
        # If the group is not in the mapPrint
        if group not in mapPrint:
            # Add the group as a key and the kind as a value
            mapPrint[group] = []
        tested = ":x:"
        if known_crd_type in example_types:
            tested = ":white_check_mark:"
        # Add the kind to the group
        mapPrint[group].append(f"|{kind}|{known_crd_type}|{tested}|")

    # Write list of resources in file
    with open('list-tested.md', 'w') as f:
        f.write('# Resources tested\n')
        f.write('\nLast update: ' + os.popen('date').read())
        for group in mapPrint:
            f.write(f'\n## {group}\n')
            f.write(f'|Kind|CRD|Tested|\n')
            f.write(f'|---|---|---|\n')
            for kind in mapPrint[group]:
                f.write(kind+'\n')
    print("> Successfully generated list of resources tested")


def generateListResources():
    # Count number of resources implemented and not implemented
    countImplemented = 0
    countNotImplemented = 0

    # array to store lines to print by group
    # array array
    linesImplementedPrint = {}
    linesNotImplementedPrint = {}
    linesDeprecatedPrint = []

    for root, dirs, files in os.walk('.work/FlexibleEngineCloud/flexibleengine/docs/resources'):
        for file in files:
            if file.endswith('.md'):
                with open('.work/FlexibleEngineCloud/flexibleengine/docs/resources/'+file) as f:
                    for line in f:
                        if line.startswith('# flexibleengine'):

                            # remove \ in line
                            line = line.replace('\\', '')

                            # identify group of resource (ex: Elastic Cloud Server (ECS))
                            # find line contain 'subcategory: "Elastic Cloud Server (ECS)"' and extract group name
                            group = ''
                            with open('.work/FlexibleEngineCloud/flexibleengine/docs/resources/'+file) as x:
                                for l in x:
                                    if l.find('subcategory: "') != -1:
                                        group = l[14:-2]
                                        break

                            if group not in linesImplementedPrint:
                                linesImplementedPrint[group] = []

                            if group not in linesNotImplementedPrint:
                                linesNotImplementedPrint[group] = []

                            # Define var found to check if resource is already implemented
                            found = False
                            with open('config/external_name.go') as x:
                                for l in x:

                                    # if line containe the resource name
                                    if l.find(line[2:-1]) != -1:
                                        found = True
                                        countImplemented += 1
                                        linesImplementedPrint[group].append(
                                            '* [x] [' + line[2:-1] + '](https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/' + file[0:-3] + ')')
                                        break

                            if found == False:
                                countNotImplemented += 1
                                linesNotImplementedPrint[group].append(
                                    '* [ ] [' + line[2:-1] + '](https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/' + file[0:-3] + ')')
                            break

    # If exist remove group Deprecated
    if 'Deprecated' in linesImplementedPrint:
        linesDeprecatedPrint.append(linesImplementedPrint['Deprecated'])
        countImplemented -= len(linesImplementedPrint['Deprecated'])
        del linesImplementedPrint['Deprecated']

    if 'Deprecated' in linesNotImplementedPrint:
        linesDeprecatedPrint.append(linesNotImplementedPrint['Deprecated'])
        countNotImplemented -= len(linesNotImplementedPrint['Deprecated'])
        del linesNotImplementedPrint['Deprecated']

    percentageG = (countImplemented /
                   (countImplemented + countNotImplemented)) * 100

    # Write list of resources in file
    with open('list-resources.md', 'w') as f:
        f.write('# Resources\n')
        f.write('\nLast update: ' + os.popen('date').read())
        f.write('\n## Resources implemented: ' + str(countImplemented) + '/' + str((countImplemented + countNotImplemented)) +
                ' (' + str(round(percentageG, 2)) + '%)\n')
        for group in linesImplementedPrint:
            percentage = (len(linesImplementedPrint[group]) / (
                len(linesImplementedPrint[group]) + len(linesNotImplementedPrint[group]))) * 100
            f.write('### ' + group + ' (' + str(round(percentage, 2)) + '%)\n')
            f.write('#### Implemented (' +
                    str(len(linesImplementedPrint[group])) + '/' +
                    str(len(linesNotImplementedPrint[group]) +
                        len(linesImplementedPrint[group])) + ')\n')
            for line in linesImplementedPrint[group]:
                f.write(line + '\n')
            for line in linesNotImplementedPrint[group]:
                f.write(line + '\n')
            f.write('\n')
        if len(linesDeprecatedPrint) > 0:
            f.write('---------------------------------------\n')
            f.write('## Deprecated\n')
            for group in linesDeprecatedPrint:
                for line in group:
                    f.write(line + '\n')
                f.write('\n')

    print("> Successfully generated list of resources")
    print("Implemented " + str(countImplemented) + "/" + str((countImplemented +
                                                              countNotImplemented)) + " (" + str(round(percentageG, 2)) + "%) resources")


def kubectl():

    if len(sys.argv) != 3:
        print("Usage: python3 cli.py kubectl <groupResource>/<resource>")
        exit(1)

    debug = False

    fileToApply = []
    Labels = {}
    # Labels = defaultdict(dict)
    listLabelsRequired = []
    listRefsRequired = []

    groupResource, resource = sys.argv[2].split("/")

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
                    listLabelsRequiredFound.append(Labels[kind]["label"])
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

    if debug:
        print("listLabelsRequiredFound:")
        print(listLabelsRequiredFound)

    if foundErrors:
        print("\nError: Some resources are not found")
        print("Please check the list of required resources in the examples directory")
        sys.exit(1)

    if debug:
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


class CommandCompleterCmd(Completer):
    def __init__(self):
        self.path_completer = PathCompleter()
        self.commands = ["help", "listResources",
                         "listTested", "kubectl", "quit"]

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


class CommandCompleterKubectl(Completer):
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
        other_args_cmd = {}
        inputObj = prompt(
            "\nWelcome to CLI toolbox\n(press tab for available commands)\nCmd >>> ", completer=CommandCompleterCmd(), **other_args_cmd)

        sys.argv.append(inputObj)
        print("")

    # Switch  sys.argv[1]
    match sys.argv[1]:
        case "kubectl":
            if len(sys.argv) == 2:
                other_args_kubectl = {}
                inputObj = prompt("Find example to apply (nat/gateway) >>> ",
                                  completer=CommandCompleterKubectl(), **other_args_kubectl)
                sys.argv.append(inputObj)
            kubectl()
        case "help":
            help()
        case "listResources":
            generateListResources()
        case "listTested":
            generateListTested()
        case "exit" | "quit":
            exit(0)
        # default
        case _:
            help()

    # Exit 0
    exit(0)
