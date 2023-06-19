package selenium

type DevNull struct {
}

func (d *DevNull) Write(p []byte) (n int, err error) {
	n = len(p)
	err = nil
	return
}

func (d *DevNull) Read(p []byte) (n int, err error) {
	n = 0
	err = nil
	return
}

func NewDevNull() *DevNull {
	return &DevNull{}
}
