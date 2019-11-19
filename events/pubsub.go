package events

import (
	"github.com/cskr/pubsub"
)

const TOPIC_PAYLOAD_DELIVERED = "payload:delivered"

var Events *pubsub.PubSub

func init() {
	//create the new pubsub events queue
	Events = pubsub.New(0)
}
