package main

import (
	lConf "go-blogs-api/configs"
	lUtls "go-blogs-api/utils"
	"log"
	"os"
)

func init() {
	lUtls.SetGlobals()
	lConf.InitDB()
}

func main() {
	defer lUtls.HandlePanic()

	APP := lConf.InitFiber()

	lUtls.InstallRouter(APP)

	hostUrl := os.Getenv("APP_URL")
	hostPort := os.Getenv("APP_PORT")
	listenUrl := hostUrl + ":" + hostPort

	// Start server on http://localhost:3000
	log.Fatal(APP.Listen(listenUrl))
	// fmt.Println(DBO)
}
