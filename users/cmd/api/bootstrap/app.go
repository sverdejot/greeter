package bootstrap

var users map[int]string = map[int]string{
	1: "Samuel",
}

func Run() {
	ListenGrpc(":8081")
}
