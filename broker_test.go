package bokchoy_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/hurngchunlee/bokchoy"
)

func TestBroker_Redis(t *testing.T) {
	is := assert.New(t)
	ctx := context.Background()

	_, err := bokchoy.New(ctx, bokchoy.Config{
		Broker: bokchoy.BrokerConfig{
			Type: "redis",
			Redis: bokchoy.RedisConfig{
				Type: "sentinel",
			},
		},
	}, bokchoy.WithInitialize(false))
	is.NoError(err)

	_, err = bokchoy.New(ctx, bokchoy.Config{
		Broker: bokchoy.BrokerConfig{
			Type: "redis",
			Redis: bokchoy.RedisConfig{
				Type: "cluster",
			},
		},
	}, bokchoy.WithInitialize(false))

	is.NoError(err)
}
