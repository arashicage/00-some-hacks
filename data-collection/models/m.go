package models

import (
	"time"
)

type FPXX struct {
	HASH        string
	FPDM        string
	FPHM        string
	KPRQ        time.Time
	KPJE        float64
	XFNSRSBH    string
	GFNSRSBH    string
	XFMC        string
	GFMC        string
	XFDZDH      string
	GFdZDH      string
	XFSJSWJG_DM string
	XFDSSWJG_DM string
	XFQXSWJG_DM string
	GFSJSWJG_DM string
	GFDSSWJG_DM string
	GFQXSWJG_DM string
	TSRQ        string
	JSRQ        time.Time
	// ------------------------ //
	HWXX
}

type HWXX struct {
	HWMC string
	JE   float64
	SE   float64
	SL   float64
	HWBM string
	QDBZ string
}
