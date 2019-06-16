package models

type Session struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	StartTimeMillis    int    `json:"startTimeMillis"`
	EndTimeMillis      int    `json:"endTimeMillis"`
	ModifiedTimeMillis int    `json:"modifiedTimeMillis"`
	ActivityType       int    `json:"activtyType"`
}
