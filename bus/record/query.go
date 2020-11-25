package recordQuery

import (
	"log"
	"medicarePartD/bus/db/database"
	"medicarePartD/gen/models"

	"errors"
)

func RecordQuery(id string) (*models.Record, error) {

	// end database object creation
	db, err := database.DatabaseCreation()
	if err != nil {
		return nil, err
	}
	// end database object creation

	// begin query
	var singleRecord models.Record
	rows, err := db.Query("SELECT Active, BeneCount, BeneCountGe65, BeneCountGe65Flag, BeneCountNum, City, DrugName, FirstName, Ge65SuppressFlag, GenericName, autoID, LastOrgName, Npi, Specialty, State, Total30DayFillCount, Total30DayFillCountGe65, TotalClaimCount, TotalClaimCountGe65, TotalClaimCountNum, TotalDaySupply, TotalDaySupplyGe65, TotalDrugCost, TotalDrugCostNum, TotalDrugcostGe65 FROM partd WHERE autoID=?", id)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("not found")
	}

	for rows.Next() {
		err = rows.Scan(&singleRecord.Active, &singleRecord.BeneCount, &singleRecord.BeneCountGe65, &singleRecord.BeneCountGe65Flag, &singleRecord.BeneCountNum, &singleRecord.City, &singleRecord.DrugName, &singleRecord.FirstName, &singleRecord.Ge65SuppressFlag, &singleRecord.GenericName, &singleRecord.ID, &singleRecord.LastName, &singleRecord.Npi, &singleRecord.Specialty, &singleRecord.State, &singleRecord.Total30DayFillCount, &singleRecord.Total30DayFillCountGe65, &singleRecord.TotalClaimCount, &singleRecord.TotalClaimCountGe65, &singleRecord.TotalClaimCountNum, &singleRecord.TotalDaySupply, &singleRecord.TotalDaySupplyGe65, &singleRecord.TotalDrugCost, &singleRecord.TotalDrugCostNum, &singleRecord.TotalDrugcostGe65)

		if err != nil {
			log.Println(err.Error())
			return nil, errors.New("not found")
		}
		// end query
	}
	return &singleRecord, nil
}
