#!/usr/bin/env python3

# In python get all .md files in this direcotry (.work/FlexibleEngineCloud/flexibleengine/docs/resources) and extract title for each file
# Then generate a list of resources in markdown format
# Check if title is already implemented in config/external_name.go
# Usage: python generate-list-resources.py > list-resources.md

import os

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
