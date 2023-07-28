package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetConsumptionHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	// crea una solicitud HTTP  utilizando el servidor de prueba
	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Ejecuta la solicitud
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// Comprueba el código de estado
	if res.StatusCode != http.StatusOK {
		t.Errorf("esperado %d; recibido %d", http.StatusOK, res.StatusCode)
	}

}
