package filmsrepo

type dependencies struct {
	Mongo pgrepoadapter
}

var dp dependencies
var defaultDependency = dp.Mongo

//NewAccountRepo instantiates user service using a specified dependency
func NewFilmsRepo(dependency string) (service interface{}) {

	switch dependency {
	case "pgdb":
		return dp.Mongo
	}
	return defaultDependency
}
