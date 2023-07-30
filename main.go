package main

import (
	"fmt"
	"github.com/cantillo16/bia_energy/src/connection"
	"github.com/cantillo16/bia_energy/src/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	_ = godotenv.Load()

	connection.Migrate()
	router := mux.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
	}

	db := connection.Connect()
	fmt.Printf("Connected to database %s 🐬 \n", os.Getenv("DATABASE_NAME"))

	routes.SetupRoutes(router, db)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "template/index.html")
	})

	url := "http://" + srv.Addr

	fmt.Println("Server is running on", url+" ... 🚀")

	//err := openBrowser(url)
	//if err != nil {
	//	fmt.Println("No se pudo abrir el navegador:", err)
	//}

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func openBrowser(url string) error {
	var cmd *exec.Cmd

	switch os := runtime.GOOS; os {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default: // Linux y otros sistemas basados en Unix
		cmd = exec.Command("xdg-open", url)
	}

	return cmd.Start()
}
