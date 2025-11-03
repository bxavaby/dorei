// ./../noti/noti.go

package noti

import (
	"context"
	"log"
	"sync"

	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/matrix"
)

// Abstract notify.Notify and matrix
type Notifier struct {
	notifier *notify.Notifier
	matrix   *matrix.Service
	enabled  bool
	mu       sync.Mutex
}

// Create a new notifier from config
func New(userID, roomID, homeServer, accessToken string, enabled bool) (*Notifier, error) {
	n := &Notifier{enabled: enabled}
	if !enabled {
		return n, nil
	}

	matrix, err := matrix.New(userID, roomID, homeServer, accessToken)
	if err != nil {
		return nil, err
	}

	notifier := notify.New()
	notifier.UseServices(matrix)

	n.notifier = notifier
	n.matrix = matrix

	return n, nil
}

// If enabled=true, send message to the room
func (n *Notifier) Send(ctxt context.Context, noti string) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	if !n.enabled || n.notifier == nil {
		return nil
	}

	err := n.notifier.Send(ctxt, "", noti)
	if err != nil {
		log.Printf("notifier.Send() failed: %v", err)
	}

	return err
}

// Dynamic toggle
func (n *Notifier) Enable(flag bool) {
	n.mu.Lock()
	defer n.mu.Unclock()
	n.enabled = flag
}
