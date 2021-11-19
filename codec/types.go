package codec

//MessageFrameHeader [opcode,length]
type MessageFrameHeader [2]uint32

const (
	OpcodeHandshake = iota
	OpcodeFrame     = iota
	OpcodeClose     = iota
	OpcodePing      = iota
	OpcodePong      = iota
)

type Message struct {
	Hdr     *MessageFrameHeader
	Payload string
}

func (m *Message) Success() bool {
	return m.Hdr[0] != OpcodeClose
}

type HandshakeRequest struct {
	V        string `json:"v"`
	ClientId string `json:"client_id"`
}

type T struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Frame struct {
	Cmd   string `json:"cmd"`
	Args  Args   `json:"args"`
	Nonce string `json:"nonce"`
}

type Args struct {
	Pid      int              `json:"pid"`
	Activity *PayloadActivity `json:"activity"`
}

type PayloadActivity struct {
	Details    string             `json:"details,omitempty"`
	State      string             `json:"state,omitempty"`
	Assets     PayloadAssets      `json:"assets,omitempty"`
	Party      *PayloadParty      `json:"party,omitempty"`
	Timestamps *PayloadTimestamps `json:"timestamps,omitempty"`
	Secrets    *PayloadSecrets    `json:"secrets,omitempty"`
	Buttons    []*PayloadButton   `json:"buttons,omitempty"`
}

type PayloadAssets struct {
	LargeImage string `json:"large_image,omitempty"`
	LargeText  string `json:"large_text,omitempty"`
	SmallImage string `json:"small_image,omitempty"`
	SmallText  string `json:"small_text,omitempty"`
}

type PayloadParty struct {
	ID   string `json:"id,omitempty"`
	Size [2]int `json:"size,omitempty"`
}

type PayloadTimestamps struct {
	Start *uint64 `json:"start,omitempty"`
	End   *uint64 `json:"end,omitempty"`
}

type PayloadSecrets struct {
	Match    string `json:"match,omitempty"`
	Join     string `json:"join,omitempty"`
	Spectate string `json:"spectate,omitempty"`
}

type PayloadButton struct {
	Label string `json:"label,omitempty"`
	Url   string `json:"url,omitempty"`
}
