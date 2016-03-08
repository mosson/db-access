package models

import "db-access/db"

// Survey サーベイ
type Survey struct {
	ID   uint   `db:"pk" json:"id"`
	Name string `json:"name"`
}

// FindSurvey idが一致するものを探す
func FindSurvey(id uint) *Survey {
	client := db.GetClient()
	var results []Survey
	if err := client.Select(&results, client.Where("id", "=", id)); err != nil {
		return nil
	}
	return &results[0]
}
