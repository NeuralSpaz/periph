package apds9960

func (d *Dev) read(buf []byte) error {
	return d.c.Tx(nil, buf)
}

// readReg is similar to Read but it reads from a register.
func (d *Dev) readReg(reg byte, buf []byte) error {
	return d.c.Tx([]byte{reg}, buf)
}

// write writes the buffer to the Dev. If it is required to write to a
// specific register, the register should be passed as the first byte in the
// given buffer.
func (d *Dev) write(buf []byte) (err error) {
	return d.c.Tx(buf, nil)
}

// writeReg is similar to Write but writes to a register.
func (d *Dev) writeReg(reg byte, buf []byte) (err error) {
	// TODO(jbd): Do not allocate, not optimal.
	return d.c.Tx(append([]byte{reg}, buf...), nil)
}

const (
	enableReg                byte = 0x80
	aTimeReg                      = 0x81
	wTimeReg                      = 0x83
	lightIntLowThresholdReg       = 0x84
	ailtHighReg                   = 0x85
	lightIntHighThresholdReg      = 0x86
	apds9960Reg_AIHTH             = 0x87
	proxIntLowThresholdReg        = 0x89
	proxIntHighThresholdReg       = 0x8B
	apds9960Reg_PERS              = 0x8C
	apds9960Reg_CONFIG1           = 0x8D
	apds9960Reg_PPULSE            = 0x8E
	apds9960Reg_CONTROL           = 0x8F
	apds9960Reg_CONFIG2           = 0x90
	apds9960Reg_ID                = 0x92
	apds9960Reg_STATUS            = 0x93
	apds9960Reg_CDATAL            = 0x94
	apds9960Reg_CDATAH            = 0x95
	apds9960Reg_RDATAL            = 0x96
	apds9960Reg_RDATAH            = 0x97
	apds9960Reg_GDATAL            = 0x98
	apds9960Reg_GDATAH            = 0x99
	apds9960Reg_BDATAL            = 0x9A
	apds9960Reg_BDATAH            = 0x9B
	apds9960Reg_PDATA             = 0x9C
	apds9960Reg_POFFSET_UR        = 0x9D
	apds9960Reg_POFFSET_DL        = 0x9E
	apds9960Reg_CONFIG3           = 0x9F
	gestureEnterThresholdReg      = 0xA0
	gestureExitThresholdReg       = 0xA1
	apds9960Reg_GCONF1            = 0xA2
	apds9960Reg_GCONF2            = 0xA3
	apds9960Reg_GOFFSET_U         = 0xA4
	apds9960Reg_GOFFSET_D         = 0xA5
	apds9960Reg_GOFFSET_L         = 0xA7
	apds9960Reg_GOFFSET_R         = 0xA9
	apds9960Reg_GPULSE            = 0xA6
	apds9960Reg_GCONF3            = 0xAA
	apds9960Reg_GCONF4            = 0xAB
	apds9960Reg_GFLVL             = 0xAE
	apds9960Reg_GSTATUS           = 0xAF
	apds9960Reg_IFORCE            = 0xE4
	apds9960Reg_PICLEAR           = 0xE5
	apds9960Reg_CICLEAR           = 0xE6
	apds9960Reg_AICLEAR           = 0xE7
	apds9960Reg_GFIFO_U           = 0xFC
	apds9960Reg_GFIFO_D           = 0xFD
	apds9960Reg_GFIFO_L           = 0xFE
	apds9960Reg_GFIFO_R           = 0xFF
)
