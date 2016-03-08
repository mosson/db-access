package models

import (
	"db-access/db"
	"testing"

	"github.com/naoina/genmai"
)

func FixtureTable() {
	client := db.GetClient()

	if err := client.CreateTable(&Survey{}); err != nil {
		panic(err)
	}
}

func FixtureRecord(survey *Survey) {
	client := db.GetClient()

	_, err := client.Insert(survey)
	if err != nil {
		panic(err)
	}
}

func TestSurvey(t *testing.T) {
	db.Dial(&genmai.SQLite3Dialect{}, ":memory:")
	defer db.Close()

	FixtureTable()
	FixtureRecord(&Survey{
		Name: "test survey",
	})

	survey := FindSurvey(1)

	if survey == nil {
		t.Errorf("expected exists")
	}

	if survey.Name != "test survey" {
		t.Errorf("expected %v, actual %v", "test survey", survey.Name)
	}
}
