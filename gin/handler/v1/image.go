package v1

import (
	"github.com/gin-gonic/gin"
)

type ImageUploadReq struct {
	ShopID int64 `form:"shop_id"`
}

func ImageUpload(c *gin.Context) {
	var req ImageUploadReq
	var err error
	err = c.ShouldBind(&req)

}


