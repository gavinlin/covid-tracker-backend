package data

import (
	"fmt"
)

type fakeDataService struct {}

func NewFakeDataService() DataService {
	return &fakeDataService {}
}

func (s *fakeDataService) InitDatabase(data [][]string) {
	fmt.Printf("init fake database")
}

func (s *fakeDataService) UpdateDatabase(data [][]string) {
	fmt.Printf("update fake database")
}