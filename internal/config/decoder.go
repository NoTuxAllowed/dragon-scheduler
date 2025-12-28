package config

import (
	"fmt"
	"os"

	"sigs.k8s.io/yaml"
)

func LoadManifest(filePath string) (string, any, error) {
	// 1. Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", nil, err
	}

	// 2. Unmarshal the envelope only
	var m ConfigManifest
	if err := yaml.Unmarshal(data, &m); err != nil {
		return "", nil, err
	}

	// 3. Decode the spec based on the Kind
	switch m.Kind {
	case "Cluster":
		var s ClusterSpecV1
		err := yaml.Unmarshal(m.Spec, &s)
		return m.Kind, s, err

	default:
		return m.Kind, nil, fmt.Errorf("unknown kind: %s", m.Kind)
	}
}
