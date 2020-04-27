package services

type DataService interface {
	InitDatabase(data [][] string)
	UpdateDatabase(data [][] string)
}