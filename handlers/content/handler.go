package content

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/lwabish/typecho-api/models"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var (
	Hdl = Handler{}
)

type Handler struct {
	db *gorm.DB
	*log.Logger
}

func (h *Handler) Publish(c *gin.Context) {
	tc := &models.TypechoContent{}
	if err := c.ShouldBindWith(tc, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
}

func (h *Handler) Setup(db *gorm.DB, l *log.Logger) {
	h.db = db
	h.Logger = l
}
