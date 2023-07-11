package v1

import (
	"context"

	"github.com/crystal/groot/bean"
	"github.com/crystal/groot/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetDomainsByProject(c *gin.Context) {
	project := c.Query("project")

	batch := []bean.Domain{}
	db.DomainCli.Find(context.Background(), bson.M{"project": project}).All(&batch)
	c.JSON(200, batch)
}

func ScanDomain(c * gin.Context) {

}

func GetProjects(c *gin.Context) {
	
}
