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

# array to store lines to print by group
# array array
linesImplementedPrint = {}
linesNotImplementedPrint = {}

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

# Print Last update
print('\nLast update: ' + os.popen('date').read())

# Calculate percentage of resources implemented
percentage = (countImplemented /
              (countImplemented + countNotImplemented)) * 100
print('\n## Resources implemented: ' + str(countImplemented) + '/' + str((countImplemented + countNotImplemented)) +
      ' (' + str(round(percentage, 2)) + '%)\n')

# Print resources implemented by group
for group in linesImplementedPrint:
    # calculate percentage of resources implemented by group
    percentage = (len(linesImplementedPrint[group]) / (
        len(linesImplementedPrint[group]) + len(linesNotImplementedPrint[group]))) * 100
    print('### ' + group + ' (' + str(round(percentage, 2)) + '%)\n')
    print('#### Implemented (' +
          str(len(linesImplementedPrint[group])) + ') -- Not implemented (' +
          str(len(linesNotImplementedPrint[group])) + ')\n')
    for line in linesImplementedPrint[group]:
        print(line)
    for line in linesNotImplementedPrint[group]:
        print(line)
    print('\n')

# print('\n## Resources not implemented (' + str(countNotImplemented) + ')\n')
# # Print resources not implemented by group
# for group in linesNotImplementedPrint:
#     print('### ' + group + '\n')
#     for line in linesNotImplementedPrint[group]:
#         print(line)
#     print('\n')
