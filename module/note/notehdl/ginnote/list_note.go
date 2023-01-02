package ginnote

import (
	"fooddlv/common"
	"fooddlv/module/note/notemodel"
	"fooddlv/module/note/noterepo"
	"fooddlv/module/note/notestorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListNote(appCtx common.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var p notemodel.ListParam

		if err := c.ShouldBind(&p); err != nil && err.Error() != "EOF" {
			panic(err)
		}

		p.Fulfill()

		db := appCtx.GetDBConnection()

		store := notestorage.NewMysqlStore(db)
		repo := noterepo.NewListNoteRepo(store)

		result, err := repo.ListNote(c.Request.Context(), &p.Paging, p.ListFilter)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, p.Paging, p.ListFilter))
	}
}
