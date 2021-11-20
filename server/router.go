package server

func Init() {
	e := Router()
	e.Logger.Fatal(e.Start(":8080"))
}
