package services

import (
	"fmt"
)

type fakeDataService struct {}

func NewFakeDataService() DataService {
	return &fakeDataService {}
}

func (s *fakeDataService) InitDatabase(confirmedData [][] string, recoveredData [][]string, deathData [][]string) {
	fmt.Println("init fake database")
}

func (s *fakeDataService) UpdateDatabase(confirmedData [][] string, recoveredData [][]string, deathData [][]string) {
	fmt.Println("update fake database")
}