package provider

import (
	"context"

	"github.com/pulumi/pulumi-go-provider/infer"
)

// EnvoyListenerArgs defines the input properties for the EnvoyListener resource.
type EnvoyListenerArgs struct {
	Name    string `pulumi:"name" validate:"required"`
	Address string `pulumi:"address" validate:"required"`
	Port    int    `pulumi:"port" validate:"required"`
}

type EnvoyListenerOutputs struct {
	EnvoyListenerArgs
	ListenerId string `pulumi:"listenerId"`
}

type EnvoyListener struct{}

func (l *EnvoyListener) Create(ctx context.Context, req infer.CreateRequest[EnvoyListenerArgs]) (infer.CreateResponse[EnvoyListenerOutputs], error) {
	// TODO: Implement Envoy control plane interaction to create a listener
	return infer.CreateResponse[EnvoyListenerOutputs]{}, nil
}

func (l *EnvoyListener) Read(ctx context.Context, req infer.ReadRequest[EnvoyListenerArgs, EnvoyListenerOutputs]) (infer.ReadResponse[EnvoyListenerArgs, EnvoyListenerOutputs], error) {
	// TODO: Implement Envoy control plane interaction to read a listener
	return infer.ReadResponse[EnvoyListenerArgs, EnvoyListenerOutputs]{}, nil
}

func (l *EnvoyListener) Update(ctx context.Context, req infer.UpdateRequest[EnvoyListenerArgs, EnvoyListenerOutputs]) (infer.UpdateResponse[EnvoyListenerOutputs], error) {
	// TODO: Implement Envoy control plane interaction to update a listener
	return infer.UpdateResponse[EnvoyListenerOutputs]{}, nil
}

func (l *EnvoyListener) Diff(ctx context.Context, req infer.DiffRequest[EnvoyListenerArgs, EnvoyListenerOutputs]) (infer.DiffResponse, error) {
	// TODO: Implement Envoy control plane interaction to diff a listener
	return infer.DiffResponse{}, nil
}

func (l *EnvoyListener) Delete(ctx context.Context, req infer.DeleteRequest[EnvoyListenerOutputs]) error {
	// TODO: Implement Envoy control plane interaction to delete a listener
	return nil
}
