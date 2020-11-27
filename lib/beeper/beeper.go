package beeper

// BeeperInterface represents the methods all beepers should implement
type BeeperInterface interface {
	Init() error
	Beep()
	Destroy()
}
