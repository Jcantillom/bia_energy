package utils

// SumarAcomulados suma los valores de los acumulados en cada arreglo de ConsumptionGroup.
func SumarAcomulados(dataGraphGroup []ConsumptionGroup) {
	for i := range dataGraphGroup {
		dataGraphGroup[i].ActiveEnergy = []float64{sumarArreglo(dataGraphGroup[i].ActiveEnergy)}
		dataGraphGroup[i].ReactiveEnergy = []float64{sumarArreglo(dataGraphGroup[i].ReactiveEnergy)}
		dataGraphGroup[i].Solar = []float64{sumarArreglo(dataGraphGroup[i].Solar)}
	}
}

// sumarArreglo suma los elementos de un arreglo de float64.
func sumarArreglo(arr []float64) float64 {
	sum := 0.0
	for _, value := range arr {
		sum += value
	}
	return sum
}
