package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func findKeyPath(node *yaml.Node, targetKey string, currentPath string) string {
	if node.Kind == yaml.DocumentNode || node.Kind == yaml.SequenceNode {
		for idx, child := range node.Content {
			childPath := fmt.Sprintf("%s[%d]", currentPath, idx)
			result := findKeyPath(child, targetKey, childPath)
			if result != "" {
				return result
			}
		}
	} else if node.Kind == yaml.MappingNode {
		for i := 0; i < len(node.Content)-1; i += 2 {
			keyNode := node.Content[i]
			valueNode := node.Content[i+1]

			key := keyNode.Value
			childPath := fmt.Sprintf("%s.%s", currentPath, key)
			if key == targetKey {
				return childPath
			}

			result := findKeyPath(valueNode, targetKey, childPath)
			if result != "" {
				return result
			}
		}
	}

	return ""
}

func main() {
	// Vérifier si les arguments sont présents
	if len(os.Args) != 3 {
		fmt.Println("How to use : ", os.Args[0], "<key to search> <file path>")
		os.Exit(1)
	}

	// Récupérer les arguments
	targetKey := os.Args[1]
	yamlFilePath := os.Args[2]

	// Lire le contenu du fichier YAML
	yamlContent, err := ioutil.ReadFile(yamlFilePath)
	if err != nil {
		fmt.Println("Error while reading YAML file:", err)
		os.Exit(1)
	}

	// Parsage du contenu YAML
	var root yaml.Node
	err = yaml.Unmarshal(yamlContent, &root)
	if err != nil {
		fmt.Println("Error while parsing YAML:", err)
		os.Exit(1)
	}

	// Rechercher le chemin jusqu'à la clé
	keyPath := findKeyPath(&root, targetKey, "")

	if keyPath != "" {
		// Supprimer le préfixe du chemin s'il commence par "."
		keyPath = strings.TrimPrefix(keyPath, ".")
		fmt.Printf("Path to the Key '%s' is : %s\n", targetKey, keyPath)
	} else {
		fmt.Printf("Key '%s' was not found in the file.\n", targetKey)
	}
}
