package event

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/legattes/cartola-legates/pkg/uow"
)

type ProcessEventStrategy interface {
	Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error
}
