// Code generated by Kitex v0.8.0. DO NOT EDIT.

package userservice

import (
	"context"
	user_service "entry/kitex_gen/user_service"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetUser(ctx context.Context, userID int64, callOptions ...callopt.Option) (r *user_service.User, err error)
	SaveUser(ctx context.Context, user *user_service.User, callOptions ...callopt.Option) (err error)
	DeleteUser(ctx context.Context, userID int64, callOptions ...callopt.Option) (err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) GetUser(ctx context.Context, userID int64, callOptions ...callopt.Option) (r *user_service.User, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUser(ctx, userID)
}

func (p *kUserServiceClient) SaveUser(ctx context.Context, user *user_service.User, callOptions ...callopt.Option) (err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SaveUser(ctx, user)
}

func (p *kUserServiceClient) DeleteUser(ctx context.Context, userID int64, callOptions ...callopt.Option) (err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteUser(ctx, userID)
}
