package serializer

type Pong struct {
	PongStr string `json:"pong";form:"pong"`
}

func BuildPong(pongStr string) Pong {
	return Pong{PongStr: pongStr}
}