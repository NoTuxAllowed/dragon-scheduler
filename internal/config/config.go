package config

import "encoding/json"

type ConfigManifest struct {
	Version  string
	Kind     string
	Metadata ObjectMetadata
	Spec     json.RawMessage
}

type ObjectMetadata struct {
	Name string
}

type ClusterSpecV1 struct {
	CheckTimer    string
	CloudProvider string
	Spec          json.RawMessage
}