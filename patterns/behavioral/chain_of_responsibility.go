package behavioral

import (
	"fmt"
)

// HandlerInterface defined a handler for handling a given handleID.
type Handler interface {
	Name() string
	Next() Handler
	HandleID() int
 	Handle(handleID int) string
}

// Handler is ...
type handler struct {
	name string
	next Handler
	handleID int
}

// NewHandler returns a new Handler.
func NewHandler(name string, next Handler, handleID int) Handler {
	return &handler{name, next, handleID}
}

// Handle handles a given handleID.
func (h *handler) Handle(handleID int) string {
	if h.handleID == handleID {
		return fmt.Sprintf("%s handled %d", h.name, handleID)
	}
	if h.next == nil {
		return "no handler can handle it."
	}
	return h.next.Handle(handleID)
}

func (h *handler) HandleID() int {
	return h.handleID
}

func (h *handler) Name() string {
	return h.name
}

func (h *handler) Next() Handler {
	return h.next
}