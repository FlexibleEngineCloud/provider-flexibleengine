#!/usr/bin/env python3

# In python get all .md files in this direcotry (.work/FlexibleEngineCloud/flexibleengine/docs/resources) and extract title for each file
# Then generate a list of resources in markdown format
# Check if title is already implemented in config/external_name.go
# Usage: python generate-list-resources.py > list-resources.md

import os


print('# Resources')
# Count number of resources implemented and not implemented
countImplemented = 0
countNotImplemented = 0

linesImplementedPrint = []
linesNotImplementedPrint = []

for root, dirs, files in os.walk('.work/FlexibleEngineCloud/flexibleengine/docs/resources'):
    for file in files:
        if file.endswith('.md'):
            with open('.work/FlexibleEngineCloud/flexibleengine/docs/resources/'+file) as f:
                for line in f:
                    if line.startswith('# flexibleengine'):
                        # Define var found to check if resource is already implemented
                        found = False
                        with open('config/external_name.go') as x:
                            for l in x:

                                # if line containe the resource name
                                if l.find(line[2:-1]) != -1:
                                    found = True
                                    countImplemented += 1
                                    linesImplementedPrint.append(
                                        '* [x] [' + line[2:-1] + '](https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/' + file[0:-3] + ')')
                                    break

                        if found == False:
                            countNotImplemented += 1
                            linesNotImplementedPrint.append(
                                '* [ ] [' + line[2:-1] + '](https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/' + file[0:-3] + ')')
                        break

print('## Implemented (' + str(countImplemented) + ')')
for line in linesImplementedPrint:
    print(line)

print('## Not implemented (' + str(countNotImplemented) + ')')
for line in linesNotImplementedPrint:
    print(line)
