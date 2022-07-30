package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"markettracker.com/pkg/command"
	"markettracker.com/tracker/internal/replicate"
)

type BvcAssetRequest struct {
	Date     string  `json:"date"`
	Exchange string  `json:"exchange"`
	Price    float32 `json:"price"`
}

func BvcAsset(cmdBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req BvcAssetRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		// TODO
		cmd := replicate.NewReplicateCommand(time.Now(), req.Exchange, req.Price)
		if err := cmdBus.Dispatch(ctx, cmd); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Bvc Asset sent",
		})
	}
}
