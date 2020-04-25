package data

type DataService interface {
	InitDatabase(data [][] string)
	UpdateDatabase(data [][] string)
}