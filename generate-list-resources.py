#!/usr/bin/env python3

# In python get all .md files in this direcotry (.work/FlexibleEngineCloud/flexibleengine/docs/resources) and extract title for each file
# Then generate a list of resources in markdown format
# Check if title is already implemented in config/external_name.go
# Usage: python generate-list-resources.py > list-resources.md

import os


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
                                    print(
                                        '* [x] [' + line[2:-1] + '](https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/' + file[0:-3] + ')')
                                    break

                        if found == False:
                            print(
                                '* [ ] [' + line[2:-1] + '](https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/' + file[0:-3] + ')')
                        break
