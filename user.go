package boomer

type User interface {
	OnStart()
	OnStop()
	GetAllTasks() []*Task
}

type UserInstFunc func() (User, error)
