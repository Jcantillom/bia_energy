package middlewares

import (
	"net/http"
)

func ValidateKindPeriod(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		kindPeriod := r.URL.Query().Get("kind_period")

		// Verificar que el campo "kind_period" no esté vacío
		if kindPeriod == "" {
			http.Error(w, "El campo 'kind_period' no puede estar vacío", http.StatusBadRequest)
			return
		}

		// Verificar que el valor del campo "kind_period" sea válido
		validKindPeriods := map[string]bool{
			"monthly": true,
			"daily":   true,
			"weekly":  true,
		}

		if !validKindPeriods[kindPeriod] {
			http.Error(w, "El valor del campo 'kind_period' es inválido", http.StatusBadRequest)
			return
		}

		// Si las validaciones son exitosas, llamamos al siguiente handler en la cadena
		next.ServeHTTP(w, r)
	})
}
