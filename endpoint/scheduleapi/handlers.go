package scheduleapi

import (
	"net/http"
	"schapi/common"
	"schapi/usecase/schedule"

	"schapi/endpoint/authapi"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func FindAllSchedule(c *gin.Context) {
	u := authapi.GetLoginUser(c)

	res := schedule.FetchAllschedulesRequest(u)
	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})

}

func AddSchedule(c *gin.Context) {
	u := authapi.GetLoginUser(c)
	var r schedule.ScheduleRequest
	if err := c.ShouldBindBodyWith(&r, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	schedule.AddScheduleRequest(r, u)

	c.JSON(http.StatusOK, gin.H{
		"message": common.Stored,
	})
}

func ChangeSchedule(c *gin.Context) {
	var r schedule.ScheduleRequest
	if err := c.ShouldBindBodyWith(&r, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	schedule.ChangeScheduleRequest(r)
	c.JSON(200, gin.H{
		"message": common.Modified,
	})
}

func Deleteschedule(c *gin.Context) {

	var r schedule.ScheduleRequest
	if err := c.ShouldBindBodyWith(&r, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	schedule.DeleteScheduleRequest(r.ID)

	c.JSON(200, gin.H{
		"message": common.Deleted,
	})
}
