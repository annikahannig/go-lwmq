package lwmq

/**
 * Light Weight Message Queueing
 * implements a simple pubsub model and can persisted
 * by using a store.
 *
 * (c) 2015 Matthias Hannig
 */

/**
 * A consumer is a subscriber to a queue
 * it can be added and removed at runtime.
 *
 * If a queue currently has no active consumers,
 * messages will be queued for delivery.
 */
type Consumer interface {
	// Handle messages
	HandleMessage(*Message) error
}
