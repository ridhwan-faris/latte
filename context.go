package latte

import (
	"log"

	"github.com/go-playground/validator/v10"
)

var (
	// Normal colors
	NBlack   = []byte{'\033', '[', '3', '0', 'm'}
	NRed     = []byte{'\033', '[', '3', '1', 'm'}
	NGreen   = []byte{'\033', '[', '3', '2', 'm'}
	NYellow  = []byte{'\033', '[', '3', '3', 'm'}
	NBlue    = []byte{'\033', '[', '3', '4', 'm'}
	NMagenta = []byte{'\033', '[', '3', '5', 'm'}
	NCyan    = []byte{'\033', '[', '3', '6', 'm'}
	NWhite   = []byte{'\033', '[', '3', '7', 'm'}
	// Bright colors
	BBlack   = []byte{'\033', '[', '3', '0', ';', '1', 'm'}
	BRed     = []byte{'\033', '[', '3', '1', ';', '1', 'm'}
	BGreen   = []byte{'\033', '[', '3', '2', ';', '1', 'm'}
	BYellow  = []byte{'\033', '[', '3', '3', ';', '1', 'm'}
	BBlue    = []byte{'\033', '[', '3', '4', ';', '1', 'm'}
	BMagenta = []byte{'\033', '[', '3', '5', ';', '1', 'm'}
	BCyan    = []byte{'\033', '[', '3', '6', ';', '1', 'm'}
	BWhite   = []byte{'\033', '[', '3', '7', ';', '1', 'm'}

	Reset = []byte{'\033', '[', '0', 'm'}
)

func (c *Context) Info(message interface{}) {
	log.Println(string(BGreen), message, string(Reset))
}

type Context struct {
	Debug         bool
	Configuration Configuration
	Validator     *validator.Validate
	MySqlDsn      string
	PostgreDsn    string
	RedisDsn      string
	MongoDsn      string
}
