package esockets

func init() {
	var esocket = Esocket{
		ID: "defaultHttp",
	}
	esocket.register()
}
