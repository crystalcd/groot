package bootstrap

import (
	"github.com/qiniu/qmgo"
)

type Application struct {
	Env   *Env
	Mongo *qmgo.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()

	InjectBeans(app.Env)

	app.Mongo = NewMongoDataBase(app.Env)

	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDbConnection(app.Mongo)
}
