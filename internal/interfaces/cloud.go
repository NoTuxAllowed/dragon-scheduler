package interfaces

import "context"


type CloudProvider interface {
	Provision(ctx context.Context) error
	Terminate(ctx context.Context) error
}
