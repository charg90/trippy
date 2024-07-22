package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)
func HandleValidationError(err error, ctx *gin.Context) {
    if validationErrors, ok := err.(validator.ValidationErrors); ok {
        errs := make(map[string]string)
        for _, ve := range validationErrors {
            fieldName := ve.Field()
            errMessage := fieldName + " is invalid"
            errs[fieldName] = errMessage
        }
        ctx.JSON(400, gin.H{"errors": errs})
        return
    }
 
    ctx.JSON(500, gin.H{"error": "Internal server error"})
}
