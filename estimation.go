package costestimation

import (
	"context"
	"io"

	"github.com/cycloidio/cost-estimation/aws"
	"github.com/cycloidio/cost-estimation/cost"
	"github.com/cycloidio/cost-estimation/terraform"
)

// EstimateTerraformPlan is a helper function that reads a Terraform plan using the provided io.Reader,
// generates the prior and planned cost.State, and then creates a cost.Plan from them that is returned.
// It uses the Backend to retrieve the pricing data.
func EstimateTerraformPlan(ctx context.Context, backend Backend, plan io.Reader, providerInitializers ...terraform.ProviderInitializer) (*cost.Plan, error) {
	if len(providerInitializers) == 0 {
		providerInitializers = []terraform.ProviderInitializer{
			aws.TerraformProviderInitializer,
		}
	}

	tfplan := terraform.NewPlan(providerInitializers...)
	if err := tfplan.Read(plan); err != nil {
		return nil, err
	}

	priorQueries, err := tfplan.ExtractPriorQueries()
	if err != nil {
		return nil, err
	}
	prior, err := cost.NewState(ctx, backend, priorQueries)
	if err != nil {
		return nil, err
	}

	plannedQueries, err := tfplan.ExtractPlannedQueries()
	if err != nil {
		return nil, err
	}
	planned, err := cost.NewState(ctx, backend, plannedQueries)
	if err != nil {
		return nil, err
	}

	return cost.NewPlan(prior, planned), nil
}
