package observer

// MyObserverAppDemo driver
func MyObserverAppDemo() {
	var email Observer = NewObserver1("email")
	var newsFeed Observer = NewObserver2("news feed")
	subject := NewSubject()
	subject.Register(&email)
	subject.Register(&newsFeed)
	subject.NotifyAll()
	subject.Unregister(&email)
	subject.SetFlag(100)
	subject.NotifyAll()
}
