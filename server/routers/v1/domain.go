package v1

import (
	"context"
	"net/http"
	"strings"

	"github.com/crystal/groot/app"
	"github.com/crystal/groot/bean"
	"github.com/crystal/groot/db"
	"github.com/crystal/groot/domainscan"
	"github.com/crystal/groot/logging"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetDomainsByProject(c *gin.Context) {
	appG := app.Gin{C: c}
	project := c.Query("project")

	batch := []bean.Domain{}
	db.DomainCli.Find(context.Background(), bson.M{"project": project}).All(&batch)
	appG.Response(http.StatusOK, 0, batch)
}

func ScanDomain(c *gin.Context) {
	appG := app.Gin{C: c}
	var form bean.AddProject
	c.BindJSON(&form)
	line := strings.ReplaceAll(form.Domains, "\n", ",")
	param := bean.Param{
		Target:  line,
		Project: form.ProjectName,
	}
	s := domainscan.NewSubfinder(param)
	s.AsyncScan()
	a := domainscan.NewAssetfinder(param)
	a.AsyncScan()
	logging.RuntimeLog.Info(line)
	appG.Response(http.StatusOK, 0, form)
}

func GetProjects(c *gin.Context) {

}
