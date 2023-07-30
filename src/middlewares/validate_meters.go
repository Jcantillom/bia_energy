package middlewares

import (
	"github.com/cantillo16/bia_energy/src/utils"
	"net/http"
)

func ValidateMeters(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		meterIDs := utils.ParseMeterIDs(r.URL.Query().Get("meters_ids"))
		// Verificar si se especificaron medidores
		if len(meterIDs) == 0 {
			http.Error(w, "Debe especificar al menos un medidor", http.StatusBadRequest)
			return
		}
		// Si las validaciones son exitosas, llamamos al siguiente handler en la cadena
		next.ServeHTTP(w, r)
	})
}
