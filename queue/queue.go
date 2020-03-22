// Package queue creates a ItemQueue data structure for the Item type
package queue

import (
    "sync"
    "github.com/cheekybits/genny/generic"
)

// Item the type of the queue
type Item generic.Type

// ItemQueue the queue of Items
type ItemQueue struct {
    Items []Item
    Lock  sync.RWMutex
}

// New creates a new ItemQueue
func (s *ItemQueue) New() *ItemQueue {
    s.Items = []Item{}
    return s
}

// Enqueue adds an Item to the end of the queue
func (s *ItemQueue) Enqueue(t Item) {
    s.Lock.Lock()
    s.Items = append(s.Items, t)
    s.Lock.Unlock()
}

// Dequeue removes an Item from the start of the queue
func (s *ItemQueue) Dequeue() *Item {
    s.Lock.Lock()
    item := s.Items[0]
    s.Items = s.Items[1:len(s.Items)]
    s.Lock.Unlock()
    return &item
}

// Front returns the item next in the queue, without removing it
func (s *ItemQueue) Front() *Item {
    s.Lock.RLock()
    item := s.Items[0]
    s.Lock.RUnlock()
    return &item
}

// IsEmpty returns true if the queue is empty
func (s *ItemQueue) IsEmpty() bool {
    return len(s.Items) == 0
}

// Size returns the number of Items in the queue
func (s *ItemQueue) Size() int {
    return len(s.Items)
}