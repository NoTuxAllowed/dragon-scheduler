package aws

import (
	"context"
	"fmt"

	"github.com/NoTuxAllowed/dragon-scheduler/internal/interfaces"
)

type Provider struct {
	Config *AwsCloudSpec
}

func (p *Provider) Provision(ctx context.Context, id string) (*interfaces.Node, error) {
	fmt.Printf("[AWS] Provisioning instance %s with type %s (Spot: %v)\n", id, p.Config.InstanceType, p.Config.Spot)
	return &interfaces.Node{
		ID:        id,
		PrivateIP: "10.0.0.1",
	},
	nil
}

func (p *Provider) Terminate(ctx context.Context, id string) error {
	fmt.Printf("[AWS] Terminating instance %s\n", id)
	return nil
}
