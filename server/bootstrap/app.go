package bootstrap

import (
	"github.com/panjf2000/ants/v2"
	"github.com/qiniu/qmgo"
	"github.com/sirupsen/logrus"
)

type Application struct {
	Env    *Env
	Logger *logrus.Logger
	Mongo  *qmgo.Client
	Pool   *ants.Pool
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Logger = NewLogger(app.Env)
	app.Mongo = NewMongoDataBase(app.Env, app.Logger)
	app.Pool = NewPool(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDbConnection(app.Mongo, app.Logger)
}
