package server

import (
	"littleserver/net/client"
)

type EventHandler interface {
	ConnectedOnClient(client *client.Client)
}
