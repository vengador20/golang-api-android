package adapter

import "context"

type Adapter interface {
	New()
	Disconnect(ctx context.Context)
	FindId(ctx context.Context, id string) (map[string]interface{}, error)
	FindEmail(ctx context.Context, email string) (map[string]interface{}, error)
	UpdateOne(ctx context.Context, filter, update map[string]interface{}) error
}
