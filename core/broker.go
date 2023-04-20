package core

import log "github.com/sirupsen/logrus"

type Broker[T any] struct {
	stopCh    chan struct{}
	publishCh chan T
	subCh     chan chan T
	unsubCh   chan chan T
	subs      map[chan T]struct{}
}

func NewBroker[T any]() *Broker[T] {
	return &Broker[T]{
		stopCh:    make(chan struct{}),
		publishCh: make(chan T, 1),
		subCh:     make(chan chan T, 1),
		unsubCh:   make(chan chan T, 1),
		subs:      map[chan T]struct{}{},
	}
}

func (b *Broker[T]) Start() {
	for {
		select {
		case <-b.stopCh:
			return
		case msgCh := <-b.subCh:
			b.subs[msgCh] = struct{}{}
		case msgCh := <-b.unsubCh:
			delete(b.subs, msgCh)
		case msg := <-b.publishCh:
			for msgCh := range b.subs {
				// msgCh is buffered, use non-blocking send to protect the broker:
				log.Debugf("%v <- %v", msgCh, msg)
				select {
				case msgCh <- msg:
				default:
				}
			}
		}
	}
}

func (b *Broker[T]) Close() {
	close(b.stopCh)
	close(b.publishCh)
	for subch := range b.subs {
		close(subch)
	}
}

func (b *Broker[T]) Wait() {
	<-b.stopCh
}

func (b *Broker[T]) Subscribe() chan T {
	msgCh := make(chan T, 1)
	b.subCh <- msgCh
	return msgCh
}

func (b *Broker[T]) Unsubscribe(msgCh chan T) {
	b.unsubCh <- msgCh
}

func (b *Broker[T]) Publish(msg T) {
	b.publishCh <- msg
}
