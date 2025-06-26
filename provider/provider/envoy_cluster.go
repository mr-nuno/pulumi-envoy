package provider

import (
	"context"

	"github.com/pulumi/pulumi-go-provider/infer"
)

// EnvoyClusterArgs defines the input properties for the EnvoyCluster resource.
type EnvoyClusterArgs struct {
	Name string `pulumi:"name" validate:"required"`
	Type string `pulumi:"type" validate:"required"`
}

type EnvoyClusterOutputs struct {
	EnvoyClusterArgs
	ClusterId string `pulumi:"clusterId"`
}

type EnvoyCluster struct{}

func (c *EnvoyCluster) Create(ctx context.Context, req infer.CreateRequest[EnvoyClusterArgs]) (infer.CreateResponse[EnvoyClusterOutputs], error) {
	// TODO: Implement Envoy control plane interaction to create a cluster
	return infer.CreateResponse[EnvoyClusterOutputs]{}, nil
}

func (c *EnvoyCluster) Read(ctx context.Context, req infer.ReadRequest[EnvoyClusterArgs, EnvoyClusterOutputs]) (infer.ReadResponse[EnvoyClusterArgs, EnvoyClusterOutputs], error) {
	// TODO: Implement Envoy control plane interaction to read a cluster
	return infer.ReadResponse[EnvoyClusterArgs, EnvoyClusterOutputs]{}, nil
}

func (c *EnvoyCluster) Update(ctx context.Context, req infer.UpdateRequest[EnvoyClusterArgs, EnvoyClusterOutputs]) (infer.UpdateResponse[EnvoyClusterOutputs], error) {
	// TODO: Implement Envoy control plane interaction to update a cluster
	return infer.UpdateResponse[EnvoyClusterOutputs]{}, nil
}

func (c *EnvoyCluster) Diff(ctx context.Context, req infer.DiffRequest[EnvoyClusterArgs, EnvoyClusterOutputs]) (infer.DiffResponse, error) {
	// TODO: Implement Envoy control plane interaction to diff a cluster
	return infer.DiffResponse{}, nil
}

func (c *EnvoyCluster) Delete(ctx context.Context, req infer.DeleteRequest[EnvoyClusterOutputs]) error {
	// TODO: Implement Envoy control plane interaction to delete a cluster
	return nil
}
