package azure

import (
	"context"
	"fmt"

)

type Provider struct {
	Config *AzureCloudSpec
}

func (p *Provider) Provision(ctx context.Context, id string) (error) {
	fmt.Printf("[Azure] Provisioning instance %s with type %s (Spot: %v)\n", id, p.Config.InstanceType, p.Config.Spot)
	return nil
}

func (p *Provider) Terminate(ctx context.Context, id string) error {
	fmt.Printf("[Azure] Terminating instance %s\n", id)
	return nil
}

