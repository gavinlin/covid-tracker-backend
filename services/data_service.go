package services

type DataService interface {
	InitDatabase(confirmedData [][] string, recoveredData [][]string, deathData [][]string)
	UpdateDatabase(confirmedData [][] string, recoveredData [][]string, deathData [][]string)
}