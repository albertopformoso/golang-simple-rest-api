package app

type Msg struct {
	Mess string `json:"message"`
}

type allMsg []Msg

func Message(m string) (myMessage allMsg) {
	myMessage = allMsg {
		{
			Mess: m,
		},
	}
	return
}