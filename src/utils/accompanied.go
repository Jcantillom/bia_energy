package utils

func SumarArreglo(arr []float64) float64 {
	sum := 0.0
	for _, value := range arr {
		sum += value
	}
	return sum
}

// GetSumarizedData calcula el acumulado de los valores por mes de acuerdo al campo especificado y los devuelve en un arreglo.
func GetSumarizedData(dataGraphGroup []*ConsumptionGroup, field string) []float64 {
	var sumarizedData []float64

	for _, group := range dataGraphGroup {
		switch field {
		case "ActiveEnergy":
			sumarizedData = append(sumarizedData, SumarArreglo(group.ActiveEnergy))
		case "ReactiveEnergy":
			sumarizedData = append(sumarizedData, SumarArreglo(group.ReactiveEnergy))
		case "Solar":
			sumarizedData = append(sumarizedData, SumarArreglo(group.Solar))
		}
	}

	return sumarizedData
}
