package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net"
)

func main() {
	lic := LicenceGen()

	fmt.Println(lic)
}

// 注册
// func Register(lic string) bool{
//
// }

func LicenceCheck() bool {
	lic := LicenceGen()

	return true
}

func LicenceGen() (lic string) {

	interfaces, _ := net.Interfaces() //slice

	i := 0
	for _, inter := range interfaces {
		if inter.HardwareAddr.String() != "" {
			h := md5.New()
			h.Write([]byte(inter.HardwareAddr.String() + inter.Name))
			if i == 0 {
				lic = lic + hex.EncodeToString(h.Sum(nil))
			} else {
				lic = lic + "\n" + hex.EncodeToString(h.Sum(nil))
			}
			i++
		}
	}

	return

}
