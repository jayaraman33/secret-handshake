package secret

const (
	wink          uint = 0b1
	doubleBlink   uint = 0b10
	closeYourEyes uint = 0b100
	jump          uint = 0b1000
	reverse       uint = 0b10000
)

func Handshake(code uint) []string {
	var actions []string
	if code&wink == wink {
		actions = append(actions, "wink")
	}
	if code&doubleBlink == doubleBlink {
		actions = append(actions, "double blink")
	}
	if code&closeYourEyes == closeYourEyes {
		actions = append(actions, "close your eyes")
	}
	if code&jump == jump {
		actions = append(actions, "jump")
	}
	if code&reverse == reverse {
		for i, j := 0, len(actions)-1; i < j; i, j = i+1, j-1 {
			actions[i], actions[j] = actions[j], actions[i]
		}
	}

	return actions
}
