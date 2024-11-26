package sensor

import (
	"time"

	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
)

const (
	shtc3Address = 0x70 // I2C address for SHTC3
)

type Measurement struct {
	TemperatureF float64
	TemperatureC float64
	Humidity     float64
	Timestamp    int64
	Location     string
}

func ReadShtc3(location string) (*Measurement, error) {
	if _, err := host.Init(); err != nil {
		return nil, err
	}

	bus, err := i2creg.Open("")
	if err != nil {
		return nil, err
	}
	defer bus.Close()

	dev := i2c.Dev{Bus: bus, Addr: shtc3Address}

	wakeupCmd := []byte{0x35, 0x17}
	if err := dev.Tx(wakeupCmd, nil); err != nil {
		return nil, err
	}

	measureCmd := []byte{0x7C, 0xA2}
	if err := dev.Tx(measureCmd, nil); err != nil {
		return nil, err
	}

	// Delay to ensure data comes back from sensor before reading.
	time.Sleep(50 * time.Millisecond)

	// Read data
	data := make([]byte, 6)
	if err := dev.Tx(nil, data); err != nil {
		return nil, err
	}

	temperatureC := float64(uint16(data[0])<<8|uint16(data[1]))*175.0/65536.0 - 45.0
	humidity := float64(uint16(data[3])<<8|uint16(data[4])) * 100.0 / 65536.0

	measurement := &Measurement{
		Location:     location,
		Timestamp:    time.Now().Unix(),
		TemperatureC: temperatureC,
		TemperatureF: (temperatureC * 9 / 5) + 32,
		Humidity:     humidity,
	}

	return measurement, nil
}
