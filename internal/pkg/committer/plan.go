package committer

import (
	"context"
	"cloud.google.com/go/spanner"

	"github.com/MediStatTech/commitplan"
	//"github.com/MediStatTech/commitplan/drivers/spanner"
)

// PlanCommitter wraps commitplan.Plan and provides a typed Apply method.
type PlanCommitter struct {
	client *spanner.Client
}

// New creates a new PlanCommitter with the given Spanner client.
func New(client *spanner.Client) *PlanCommitter {
	return &PlanCommitter{client: client}
}


func (c *PlanCommitter) Apply(ctx context.Context, plan *commitplan.Plan) error {
	if plan == nil {
		return nil
	}
	_, err := c.client.Apply(ctx, plan.Mutations()) // assuming you have Mutations() getter
	return err
}

