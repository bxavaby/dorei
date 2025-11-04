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
	notifier notify.Notifier
	matrix   *matrix.Matrix
	enabled  bool
	mu       sync.Mutex
}

// Create a new notifier from config
func New(userID, roomID, homeServer, accessToken string, enabled bool) (*Notifier, error) {
	n := &Notifier{enabled: enabled}
	if !enabled {
		return n, nil
	}

	matrixSvc, err := matrix.New("user-id", "room-id", "home-server", "access-token")
	if err != nil {
		return nil, err
	}

	notifier := notify.New()
	notifier.UseServices(matrixSvc)

	n.notifier = notifier
	n.matrix = matrixSvc

	return n, nil
}

// If enabled=true, send message to the room
func (n *Notifier) Send(ctx context.Context, noti string) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	if !n.enabled || n.notifier == nil {
		return nil
	}

	err := n.notifier.Send(ctx, "", noti)
	if err != nil {
		log.Printf("notifier.Send() failed: %v", err)
	}

	return err
}

// Dynamic toggle
func (n *Notifier) Enable(flag bool) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.enabled = flag
}

func (n *Notifier) IsEnabled() bool {
	n.mu.Lock()
	defer n.mu.Unlock()
	return n.enabled
}
