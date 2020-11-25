package record

import (
	recordQuery "medicarePartD/bus/record"
	"medicarePartD/gen/models"
	"strings"

	"github.com/go-openapi/runtime/middleware"
)

func HandleSingleRecord(params GetSingleRecordParams) middleware.Responder {

	recordQuery, err := recordQuery.RecordQuery(params.ID)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return NewGetSingleRecordNotFound().WithPayload(&models.ResponseCode{
				Code:    "404",
				Message: "not found"})
		}
		if strings.Contains(err.Error(), "forbidden") {
			return NewGetSingleRecordForbidden().WithPayload(&models.ResponseCode{
				Code:    "403",
				Message: "you do not have permission to access this"})
		}
		return NewGetSingleRecordInternalServerError().WithPayload(&models.ResponseCode{
			Code:    "500",
			Message: "server error"})

	}
	return NewGetSingleRecordOK().WithPayload(recordQuery)
}
