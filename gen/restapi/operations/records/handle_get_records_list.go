package records

import (
	recordsQuery "medicarePartD/bus/records"
	"medicarePartD/gen/models"
	"strings"

	"github.com/go-openapi/runtime/middleware"
)

func HandleRecordsList(params GetRecordsListParams) middleware.Responder {
	recordsQuery, err := recordsQuery.RecordsQuery(params.City, params.State, params.DrugName, params.Specialty)
	if err != nil {
		if strings.Contains(err.Error(), "forbidden") {
			return NewGetRecordsListForbidden().WithPayload(&models.ResponseCode{
				Code:    "403",
				Message: "you do not have permission to access this"})
		}
		if strings.Contains(err.Error(), "not found") {
			return NewGetRecordsListNotFound().WithPayload(&models.ResponseCode{
				Code:    "404",
				Message: "not found"})
		}
		return NewGetRecordsListInternalServerError().WithPayload(&models.ResponseCode{
			Code:    "500",
			Message: "server error"})
	}
	return NewGetRecordsListOK().WithPayload(recordsQuery)
}
