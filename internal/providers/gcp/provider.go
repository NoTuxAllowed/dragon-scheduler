package gcp

import (
	"context"
	"fmt"
)

type Provider struct {
	Config *GCPCloudSpec
}

func (p *Provider) Provision(ctx context.Context, id string) error {
	fmt.Printf("[GCP] Provisioning instance %s with type %s (Spot: %v)\n", id, p.Config.InstanceType, p.Config.Spot)
	return nil
}

func (p *Provider) Terminate(ctx context.Context, id string) error {
	fmt.Printf("[GCP] Terminating instance %s\n", id)
	return nil
}
