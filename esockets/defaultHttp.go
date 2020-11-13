package esockets

func init() {
	// Create the esocket as a local variable
	var esocket = Esocket{
		ID: "defaultHttp",
		onInit: func(es *Esocket) {
			println(es.ID)
		},
	}
	// Register the esocket so that it can be listed and used
	esocket.register()
}
