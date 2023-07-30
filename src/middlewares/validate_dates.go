package middlewares

import (
	"net/http"
	"time"
)

func ValidateDates(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startDateStr := r.URL.Query().Get("start_date")
		endDateStr := r.URL.Query().Get("end_date")

		// Verificar que las fechas no estén vacías
		if startDateStr == "" || endDateStr == "" {
			http.Error(w, "Las fechas no pueden estar vacías",
				http.StatusBadRequest)
			return
		}

		// Parsear las fechas a objetos time.Time
		startDate, err := time.Parse("2006-01-02", startDateStr)
		if err != nil {
			http.Error(w, "Fecha inicial inválida, debe tener el formato 'AAAA-MM-DD'",
				http.StatusBadRequest)
			return
		}

		endDate, err := time.Parse("2006-01-02", endDateStr)
		if err != nil {
			http.Error(w, "Fecha final inválida, debe tener el formato 'AAAA-MM-DD'",
				http.StatusBadRequest)
			return
		}

		// Verificar que la fecha inicial no sea mayor que la fecha final
		if startDate.After(endDate) {
			http.Error(w, "La fecha inicial no puede ser mayor que la fecha final",
				http.StatusBadRequest)
			return
		}

		// Si las validaciones son exitosas, llamamos al siguiente handler en la cadena
		next.ServeHTTP(w, r)
	})
}
