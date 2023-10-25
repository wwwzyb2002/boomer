package boomer

type User interface {
	OnStart()
	OnStop()
	GetAllTasks() []*Task
}

type UserInstFunc func() (User, error)

var globalUserInstFuncMap = make(map[string]UserInstFunc)

func RegisterUserInstance(userType string, userInstFunc UserInstFunc) {
	globalUserInstFuncMap[userType] = userInstFunc
}

func GetUserInstFunc(userType string) UserInstFunc {
	return globalUserInstFuncMap[userType]
}
