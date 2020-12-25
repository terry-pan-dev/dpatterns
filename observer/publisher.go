package observer

import "fmt"

// Publisher interface
type Publisher interface {
	GetFlag() int
	SetFlag(int)
	Register(Observer)
	Unregister(Observer)
	NotifyAll()
}

// Subject class
type Subject struct {
	Flag      int
	Observers map[*Observer]*Observer
}

// NewSubject constructor
func NewSubject() *Subject {
	return &Subject{Observers: make(map[*Observer]*Observer)}
}

// GetFlag method
func (s *Subject) GetFlag() int {
	return s.Flag
}

// SetFlag method
func (s *Subject) SetFlag(flag int) {
	fmt.Println("Setting Flag")
	s.Flag = flag
}

// Register method
func (s *Subject) Register(observer *Observer) {
	fmt.Println("Register Observer")
	s.Observers[observer] = observer
}

// Unregister method
func (s *Subject) Unregister(observer *Observer) {
	fmt.Println("Unregister Observer")
	delete(s.Observers, observer)
}

// NotifyAll method
func (s *Subject) NotifyAll() {
	for _, observer := range s.Observers {
		(*observer).Update(s.GetFlag())
	}
}
