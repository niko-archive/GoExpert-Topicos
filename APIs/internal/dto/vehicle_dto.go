package dto

type VehicleOutput struct {
	ID             string  `json:"id"`
	Value          float64 `json:"value"`
	Brand          string  `json:"brand"`
	Model          string  `json:"model"`
	ModelYear      string  `json:"model_year"`
	Fuel           string  `json:"fuel"`
	FipeCode       string  `json:"fipe_code"`
	ReferenceMonth string  `json:"reference_month"`
	VehicleType    string  `json:"vehicle_type"`
}
