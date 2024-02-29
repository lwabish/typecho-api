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
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if tc.IsNewPost() {
		tc.PostDefault()
		if r := h.db.Create(tc); r.Error != nil || r.RowsAffected != 1 {
			c.JSON(http.StatusInternalServerError, r.Error.Error())
			return
		}
		c.JSON(http.StatusCreated, tc.Cid)
	} else {
		old := &models.TypechoContent{Cid: tc.Cid}
		if r := h.db.First(old); r.Error != nil || r.RowsAffected != 1 {
			c.JSON(http.StatusNotFound, r.Error.Error())
			return
		}
		old.UpdatePost(tc)
		if r := h.db.Save(old); r.Error != nil || r.RowsAffected != 1 {
			c.JSON(http.StatusInternalServerError, r.Error.Error())
		}
		c.JSON(http.StatusOK, tc.Cid)
	}
}

func (h *Handler) Setup(db *gorm.DB, l *log.Logger) {
	h.db = db
	h.Logger = l
}
