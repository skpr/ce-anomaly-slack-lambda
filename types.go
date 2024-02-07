package main

// Event is a struct that represents a cost explorer anomaly event.
type Event struct {
	AccountID          string `json:"accountId"`
	AnomalyDetailsLink string `json:"anomalyDetailsLink"`
	AnomalyID          string `json:"anomalyId"`
	DimensionalValue   string `json:"dimensionalValue"`
	Impact             Impact `json:"impact"`
}

// Impact is a struct that represents the impact of a cost explorer anomaly event.
type Impact struct {
	TotalImpactPercentage float64 `json:"totalImpactPercentage"`
}
