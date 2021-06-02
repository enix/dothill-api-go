package mock

func StartServer() {
	r := NewRouter()
	r.Run()
}
