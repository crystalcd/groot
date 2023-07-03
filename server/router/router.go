package router

import (
	"time"

	"github.com/crystal/groot/global"
	"github.com/crystal/groot/pool"
	"github.com/crystal/groot/tools"

	"github.com/gin-gonic/gin"
)

const Domain = "hackerone.com"

func Rcon(c *gin.Context) {
	global.G_LOG.Info("start Recon")
	domains := []string{"slack.com"}
	jobParam := pool.JobParam{
		Domains: domains,
		Project: "test-groot",
	}
	job := pool.Job{
		ID:      time.Now().UnixNano(),
		JobFunc: tools.Dosubfinder,
		Param:   jobParam,
	}
	pool.ThreadPool.AddJob(job)
	job2 := pool.Job{
		ID:      time.Now().UnixNano(),
		JobFunc: tools.Doassertfinder,
		Param:   jobParam,
	}
	pool.ThreadPool.AddJob(job2)
	job3 := pool.Job{
		ID:      time.Now().UnixNano(),
		JobFunc: tools.Dokatana,
		Param:   jobParam,
	}
	pool.ThreadPool.AddJob(job3)
	job4 := pool.Job{
		ID:      time.Now().UnixNano(),
		JobFunc: tools.Donaabu,
		Param:   jobParam,
	}
	pool.ThreadPool.AddJob(job4)
	job5 := pool.Job{
		ID:      time.Now().UnixNano(),
		JobFunc: tools.Dorobots,
		Param:   jobParam,
	}
	pool.ThreadPool.AddJob(job5)
	job6 := pool.Job{
		ID:      time.Now().UnixNano(),
		JobFunc: tools.Dowaybackurls,
		Param:   jobParam,
	}
	pool.ThreadPool.AddJob(job6)
	job7 := pool.Job{
		ID:      time.Now().UnixNano(),
		JobFunc: tools.Dowebanayze,
	}
	pool.ThreadPool.AddJob(job7)

	global.G_LOG.Info("submit all jobs")
	c.JSON(200, gin.H{"Message": "Hello World"})
}
