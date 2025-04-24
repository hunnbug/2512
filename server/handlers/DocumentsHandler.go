package handlers

import (
	"main/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PersonalCard(ctx *gin.Context) {
	_, err := tools.CheckParamID(ctx)
	if err != nil {
		return
	}

	// tools.CreatePersonalCard()

	ctx.JSON(http.StatusOK, gin.H{"output:": "document created."})
}
