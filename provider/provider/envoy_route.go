package provider

import (
	"context"

	"github.com/pulumi/pulumi-go-provider/infer"
)

// EnvoyRouteArgs defines the input properties for the EnvoyRoute resource.
type EnvoyRouteArgs struct {
	Name    string `pulumi:"name" validate:"required"`
	Cluster string `pulumi:"cluster" validate:"required"`
	Prefix  string `pulumi:"prefix" validate:"required"`
}

type EnvoyRouteOutputs struct {
	EnvoyRouteArgs
	RouteId string `pulumi:"routeId"`
}

type EnvoyRoute struct{}

func (r *EnvoyRoute) Create(ctx context.Context, req infer.CreateRequest[EnvoyRouteArgs]) (infer.CreateResponse[EnvoyRouteOutputs], error) {
	// TODO: Implement Envoy control plane interaction to create a route
	return infer.CreateResponse[EnvoyRouteOutputs]{}, nil
}

func (r *EnvoyRoute) Read(ctx context.Context, req infer.ReadRequest[EnvoyRouteArgs, EnvoyRouteOutputs]) (infer.ReadResponse[EnvoyRouteArgs, EnvoyRouteOutputs], error) {
	// TODO: Implement Envoy control plane interaction to read a route
	return infer.ReadResponse[EnvoyRouteArgs, EnvoyRouteOutputs]{}, nil
}

func (r *EnvoyRoute) Update(ctx context.Context, req infer.UpdateRequest[EnvoyRouteArgs, EnvoyRouteOutputs]) (infer.UpdateResponse[EnvoyRouteOutputs], error) {
	// TODO: Implement Envoy control plane interaction to update a route
	return infer.UpdateResponse[EnvoyRouteOutputs]{}, nil
}

func (r *EnvoyRoute) Diff(ctx context.Context, req infer.DiffRequest[EnvoyRouteArgs, EnvoyRouteOutputs]) (infer.DiffResponse, error) {
	// TODO: Implement Envoy control plane interaction to diff a route
	return infer.DiffResponse{}, nil
}

func (r *EnvoyRoute) Delete(ctx context.Context, req infer.DeleteRequest[EnvoyRouteOutputs]) error {
	// TODO: Implement Envoy control plane interaction to delete a route
	return nil
}
