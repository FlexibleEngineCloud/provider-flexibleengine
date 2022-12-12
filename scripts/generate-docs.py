#!/usr/bin/env python3

import os
import re

for file in os.listdir(".work/FlexibleEngineCloud/flexibleengine/docs/resources"):
    # Print file
    print(file)
    if file.endswith(".md"):
        with open(".work/FlexibleEngineCloud/flexibleengine/docs/resources/" + file, 'r') as f:
            lines = f.readlines()
            for i in range(len(lines)):
                if lines[i].startswith('#'):
                    # Extract the resource name
                    resource_name = lines[i].replace('#', '').strip()
                    # Add the in the header (page_title, description)
                    # Insert new line after the first line
                    lines.insert(2, 'page_title: "' +
                                 resource_name + '"\n')
                    lines.insert(2, 'description: ""\n')

                    with open(".work/FlexibleEngineCloud/flexibleengine/docs/resources/" + file, 'w') as f:
                        f.writelines(lines)
                    break
