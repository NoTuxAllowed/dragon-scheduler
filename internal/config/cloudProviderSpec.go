package config

type awsCloudSpec struct {
	VMDefinitions map[string]VMSpec
}

type azureCloudSpec struct {
}

type gcpCloudSpec struct {
}

type VMSpec struct {
	Name                string
	Ephemeral           bool
	Scale               string
	Spot                bool
	InstanceType        string
	BackupInstanceTypes []string
}
