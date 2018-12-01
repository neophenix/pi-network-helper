package internal

import (
	"github.com/kidoman/embd"
	"github.com/kidoman/embd/controller/hd44780"
	_ "github.com/kidoman/embd/host/rpi" // load raspberry pi driver
	"math"
)

var lcd *hd44780.HD44780

func init() {
	embd.InitI2C()

	bus := embd.NewI2CBus(1)

	lcd, _ = hd44780.NewI2C(
		bus,
		0x27,
		hd44780.PCF8574PinMap,
		hd44780.RowAddress16Col,
		hd44780.TwoLine,
	)
}

// CloseLCD wraps hd44780.Close() and embd.CloseI2C()
func CloseLCD() {
	lcd.DisplayOff()
	lcd.BacklightOff()
	lcd.Close()
	embd.CloseI2C()
}

// SetupLCD clears the screen, then turns the display and backlight on
func SetupLCD() {
	lcd.Clear()
	lcd.SetMode(hd44780.TwoLine)
	lcd.DisplayOn()
	lcd.BacklightOn()
}

// writeString writes a string to the display on the given line and starting at the given col
// if col == -1 we try to center the string
func WriteString(msg string, row int, col int) {
	if lcd == nil {
		return
	}

	if col == -1 {
		col = int(math.Floor(float64((16 - len(msg)) / 2)))
	}

	if len(msg) > 16 {
		msg = msg[0:16]
	}

	lcd.SetCursor(col, row)
	for _, b := range []byte(msg) {
		lcd.WriteChar(b)
	}
}
