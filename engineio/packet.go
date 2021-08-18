package engineio

type PacketType []byte

var (
	Open    PacketType = []byte("0") // 0
	Close   PacketType = []byte("1") // 1
	Ping    PacketType = []byte("2") // 2
	Pong    PacketType = []byte("3") // 3
	Message PacketType = []byte("4") // 4
	Upgrade PacketType = []byte("5") // 5
	Noop    PacketType = []byte("6") // 6
)
