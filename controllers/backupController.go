package controllers

import (
	"net/http"
	"message_backup/models"
	"message_backup/validation"
	"message_backup/businessLogic"
	"gopkg.in/gin-gonic/gin.v1"
)


func MsgBackup(c *gin.Context) {

	deviceKey, userId, err := validation.ValidateHeaders(c)
	if err.Status != http.StatusOK {
		res:= models.MessagesResponse{0 ,[]string{"device key or user-id header missing"},0,0,[]models.Success{},[]models.Invalid{}}
		c.JSON(400, res)
		return
	}


	msg,err := validation.JsonStructureValidation(c)
	if err.Status != http.StatusOK {
		res:= models.MessagesResponse{0 ,[]string{"Invalid Json Body"},0,0,[]models.Success{},[]models.Invalid{}}
		c.JSON(400, res)
		return
	}

	msg1,err:= validation.JsonSignatureValidation(msg)
	if err.Status != http.StatusOK {
		res:= models.MessagesResponse{0 ,[]string{err.Error},0,0,[]models.Success{},[]models.Invalid{}}
		c.JSON(400, res)
		return
	}

	valid,invalid,partial,err := validation.RequestValidation(msg1)
	if err.Status != http.StatusOK {
		res:= models.MessagesResponse{0 ,[]string{err.Error},0,0,[]models.Success{},[]models.Invalid{}}
		c.JSON(400, res)
		return
	}

	response, err := businessLogic.PutInCass(userId,deviceKey,valid)
	if err.Status != http.StatusOK {
		res:= models.MessagesResponse{0 ,[]string{err.Error},0,0,[]models.Success{},[]models.Invalid{}}
		c.JSON(err.Status, res)
		return
	}


	if partial{
		response.Status = 2
		response.Invalid = invalid
		c.JSON(202, response)
		return

	} else {
		response.Status = 1
		response.Invalid = invalid
		c.JSON(200, response)
		return
	}


}
