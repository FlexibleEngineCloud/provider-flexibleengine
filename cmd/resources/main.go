package main

import (
	"fmt"

	"github.com/fatih/color"

	tfProvider "github.com/FlexibleEngineCloud/terraform-provider-flexibleengine/flexibleengine"
	cm "github.com/azrod/common-go"

	crConfig "github.com/FrangipaneTeam/provider-flexibleengine/config"
)

//nolint:all
func main() {
	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)

	green := color.New(color.FgGreen)
	boldGreen := green.Add(color.Bold)

	blue := color.New(color.FgBlue)
	boldBlue := blue.Add(color.Bold)

	blue.Printf("Checking resources\n")
	fmt.Printf("=====================================\n\n")

	tfSchema := *tfProvider.Provider()
	blue.Printf("* Found %v resources in terraform provider\n", len(tfSchema.ResourcesMap))
	crImplemented, crIgnored := crConfig.ExternalNameConfiguredAndIgnored()
	blue.Printf("* Found %v resources implemented in crossplace provider\n", len(crImplemented))
	blue.Printf("* Found %v resources ignored in crossplace provider\n\n", len(crIgnored))
	blue.Printf("Total %v/%v resources implemented\n", len(crImplemented)+len(crIgnored), len(tfSchema.ResourcesMap))

	listOfResourcesNotImplemented := []string{}
	listOfResourcesToDeprecated := []string{}
	listOfResourcesDeprecated := []string{}
	listOfResourcesDeletedInTerraform := []string{}
	listOfResourcesImplementedAndIgnored := []string{}

	for k, v := range tfSchema.ResourcesMap {
		if _, ok := cm.Find(crImplemented, k); !ok {
			// Resource is not implemented
			if _, ok := cm.Find(crIgnored, k); !ok {
				// Resource is not ignored
				if v.DeprecationMessage != "" {
					// Resource is deprecated but not ignored
					listOfResourcesDeprecated = append(listOfResourcesDeprecated, k)
					continue
				}
				listOfResourcesNotImplemented = append(listOfResourcesNotImplemented, k)
			}
		} else {
			// Resource is implemented
			if v.DeprecationMessage == "" {
				// Resource is not deprecated
				continue
			}
			// Resource is deprecated
			listOfResourcesToDeprecated = append(listOfResourcesToDeprecated, k)
		}
	}

	crRess := []string{}
	crRess = append(crRess, crIgnored...)
	crRess = append(crRess, crImplemented...)

	for _, r := range crRess {
		if _, ok := tfSchema.ResourcesMap[r]; !ok {
			// Resource is not in terraform provider
			listOfResourcesDeletedInTerraform = append(listOfResourcesDeletedInTerraform, r)
		}
	}

	for _, r := range crIgnored {
		if _, ok := cm.Find(crImplemented, r); ok {
			// Resource is in crIgnored and in crImplemented
			listOfResourcesImplementedAndIgnored = append(listOfResourcesImplementedAndIgnored, r)
		}
	}

	if len(listOfResourcesNotImplemented) == 0 && len(listOfResourcesToDeprecated) == 0 && len(listOfResourcesDeprecated) == 0 && len(listOfResourcesDeletedInTerraform) == 0 && len(listOfResourcesImplementedAndIgnored) == 0 {
		boldGreen.Println("\nAll resources are implemented")
	} else {
		if len(listOfResourcesNotImplemented) > 0 {
			red.Printf("\n> %v resources are not implemented (Add resources to ExternalName):\n", len(listOfResourcesNotImplemented))
			for _, r := range listOfResourcesNotImplemented {
				boldGreen.Print("    + ")
				fmt.Println(r)
			}
		}
		if len(listOfResourcesToDeprecated) > 0 {
			red.Printf("\n> %v resources are to be deprecated (Move resources to skipList):\n", len(listOfResourcesToDeprecated))
			for _, r := range listOfResourcesToDeprecated {
				boldRed.Print("   -")
				boldGreen.Print("+")
				fmt.Println(r)
			}
		}
		if len(listOfResourcesDeprecated) > 0 {
			red.Printf("\n> %v resources are deprecated (Add resources in skipList):\n", len(listOfResourcesDeprecated))
			for _, r := range listOfResourcesDeprecated {
				boldGreen.Print("    + ")
				fmt.Println(r)
			}
		}
		if len(listOfResourcesDeletedInTerraform) > 0 {
			boldRed.Printf("\n> %v resources are deleted in terraform (Remove resources in crossplane):\n", len(listOfResourcesDeletedInTerraform))
			for _, r := range listOfResourcesDeletedInTerraform {
				boldRed.Print("    - ")
				fmt.Println(r)
			}
		}
		if len(listOfResourcesImplementedAndIgnored) > 0 {
			boldRed.Printf("\n> %v resources are implemented and ignored:\n", len(listOfResourcesImplementedAndIgnored))
			for _, r := range listOfResourcesImplementedAndIgnored {
				if v, ok := tfSchema.ResourcesMap[r]; ok {
					if v.DeprecationMessage != "" {
						boldRed.Print("    - ")
						fmt.Printf("%s (Marked deprecated)\n", r)
					} else {
						boldBlue.Print("    ? ")
						fmt.Println(r)
					}
				}
			}
		}
	}
}
