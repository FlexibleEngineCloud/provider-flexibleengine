# In python transform this text

# ---
# subcategory: "NAT Gateway (NAT)"
# ---

# # flexibleengine_nat_gateway_v2


# To

# ---
# subcategory: "NAT Gateway (NAT)"
# page_title: "flexibleengine_nat_gateway_v2"
# description: ""
# ---

# flexibleengine_nat_gateway_v2 -->

# For each .md file in the docs folder (.work/FlexibleEngineCloud/flexibleengine/docs/resources) do the following:
# 1. Read the file
# 2. Find the first line contain #
# 3. Add in the header (page_title, description)
# 4. Write the file

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
