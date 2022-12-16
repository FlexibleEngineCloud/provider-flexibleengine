#!/usr/bin/env python3

import yaml
import os
import sys


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

# Example usage: check-examples.py <CRD dir> <example manifests dir>
if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("Example usage: check-examples.py <CRD dir> <example manifests "
              "dir>")
        sys.exit(1)
    try:
        exception_set = exceptions["provider-flexibleengine"]
    except KeyError:
        exception_set = set()
    known_crd_types = load_gvks(sys.argv[1], load_crd_type)

    example_types = load_gvks(sys.argv[2], lambda t: [] if t is None or not {"kind", "apiVersion"}.issubset(t.keys())
        else [f'{t["kind"]}.{t["apiVersion"]}'])

    tested_crd_types = example_types.union(exception_set)


    # for known_crd_types group by group
    # ProtectionGroup.sdrs.flexibleengine.upbound.io/v1beta1
    # ProtectionGroup is the kind
    # sdrs is the group
    # v1beta1 is the version

    mapPrint = {}
    # Define map 


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

    # Print the map by group
    for group in mapPrint:
        print(f'\n## {group}')
        print(f'|Kind|CRD|Tested|')
        print(f'|---|---|---|')
        for kind in mapPrint[group]:
            print(kind)







    # diff = known_crd_types.difference(example_types.union(exception_set))
    # if len(diff) == 0:
    #     print("All CRDs have at least one example...")
    #     print(f'Exceptions allowed for: {exception_set}')
    #     sys.exit(0)
    # print(f'Please add example manifests for the following types: {diff}')
    # sys.exit(2)