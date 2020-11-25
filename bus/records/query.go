package recordsQuery

import (
	"errors"
	"log"
	"medicarePartD/bus/db/database"
	"medicarePartD/gen/models"
)

func RecordsQuery(City, State, DrugName, Specialty *string) ([]*models.Records, error) {

	db, err := database.DatabaseCreation()
	if err != nil {
		return nil, err
	}

	var qCity string
	if City != nil {
		qCity = *City
	} else {
		qCity = "%"
	}
	var qState string
	if State != nil {
		qState = *State
	} else {
		qState = "%"
	}
	var qDrugName string
	if DrugName != nil {
		qDrugName = *DrugName
	} else {
		qDrugName = "%"
	}
	var qSpecialty string
	if Specialty != nil {
		qSpecialty = *Specialty
	} else {
		qSpecialty = "%"
	}

	if City == nil && State == nil && DrugName == nil && Specialty == nil {
		return nil, errors.New("not found")
	}

	rows, err := db.Query("SELECT City, State, drugName, GenericName, lastOrgName, Specialty, autoid, firstName FROM partd WHERE city LIKE ? AND state LIKE ? AND drugName LIKE ? AND specialty LIKE ? Limit 10", qCity, qState, qDrugName, qSpecialty)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("not found")
	}

	defer rows.Close()

	var RecordsArray []*models.Records
	for rows.Next() {
		var recordsList models.Records
		err = rows.Scan(&recordsList.City, &recordsList.State, &recordsList.DrugName, &recordsList.GenericName, &recordsList.LastName, &recordsList.Specialty, &recordsList.ID, &recordsList.FirstName)
		if err != nil {
			log.Println(err.Error())
			return nil, errors.New("not found")
		}
		RecordsArray = append(RecordsArray, &recordsList)
	}

	return RecordsArray, nil
}
