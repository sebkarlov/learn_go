// Package counters provides alert counter support
package counters

// alertCounter is an unexporting type that contains an integer counter for alerts
type alertCounter int

// New creates and returns values of the unexported type alertCouner
func New(value int) alertCounter {
	return alertCounter(value)
}
