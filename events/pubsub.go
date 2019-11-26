package events

import (
	"github.com/cskr/pubsub"
)

const (
	TOPIC_PAYLOAD_DELIVERED = "payload:delivered"
	TOPIC_PAYLOAD_CREATED   = "payload:created"
	TOPIC_PAYLOAD_DELETED   = "payload:deleted"

	TOPIC_ALIAS_CREATED = "alias:created"
	TOPIC_ALIAS_DELETED = "alias:deleted"

	TOPIC_USER_CREATED    = "user:created"
	TOPIC_USER_CONNECT    = "user:connect"
	TOPIC_USER_DISCONNECT = "user:disconnect"
	TOPIC_USER_DELETE     = "user:delete"

	TOPIC_LOOT_CREATED = "loot:created"
)

var (
	TOPIC_ALL []string
	Events    *pubsub.PubSub
)

func init() {
	//create the new pubsub events queue
	TOPIC_ALL = make([]string, 0)
	TOPIC_ALL = append(TOPIC_ALL, TOPIC_PAYLOAD_DELIVERED)
	TOPIC_ALL = append(TOPIC_ALL, TOPIC_PAYLOAD_CREATED)
	TOPIC_ALL = append(TOPIC_ALL, TOPIC_PAYLOAD_DELETED)
	TOPIC_ALL = append(TOPIC_ALL, TOPIC_ALIAS_CREATED)
	TOPIC_ALL = append(TOPIC_ALL, TOPIC_ALIAS_DELETED)

	Events = pubsub.New(0)
}
