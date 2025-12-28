package interfaces

import "context"

type Node struct {
	ID        string
	PrivateIP string
	PublicIP  string
}

type CloudProvider interface {
	Provision(ctx context.Context, id string) (*Node, error)
	Terminate(ctx context.Context, id string) error
}
