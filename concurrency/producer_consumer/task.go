package producer_consumer

// Add whatever you want
type Task struct {
	Name string
}

func (t *Task) String() string {
	return t.Name
}
