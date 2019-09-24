package module

import (
	"github.com/emersion/go-message/textproto"
	"github.com/foxcpp/maddy/buffer"
)

// StatusCollector is an object that is passed by message source
// that is interested in intermediate status reports about partial
// delivery failures.
type StatusCollector interface {
	// SetStatus sets the error associated with the recipient.
	//
	// rcptTo should match exactly the value that was passed to the
	// AddRcpt, i.e. if any translations was made by the target,
	// they should not affect the rcptTo argument here.
	SetStatus(rcptTo string, err error)
}

// PartialDelivery is an optional interface that may be implemented
// by the object returned by DeliveryTarget.Start. See PartialDelivery.BodyNonAtomic
// documentation for details.
type PartialDelivery interface {
	// BodyNonAtomic is similar to Body method of the regular Delivery interface
	// with the except that it allows target to reject the body only for some
	// recipients by setting statuses using passed collector object.
	//
	// This interface is preferred by the LMTP endpoint and queue implementation
	// to ensure correct handling of partial failures.
	BodyNonAtomic(c StatusCollector, header textproto.Header, body buffer.Buffer)
}
