package main

import (
	"fmt"

	"github.com/xaionaro-go/cryptsetup"
)

func checkError(err error) {
	if err == nil {
		return
	}
	panic(err)
}

func main() {
	dev, err := cryptsetup.NewDevice("/dev/loop0")
	checkError(err)
	defer dev.Free()
	err = dev.LoadMeta(cryptsetup.CRYPT_LUKS1)
	checkError(err)
	fmt.Println(dev.Dump())
	fmt.Println(dev.GetUUID())
	fmt.Println(dev.KeySlotMax(cryptsetup.CRYPT_LUKS1))
	fmt.Println(dev.LuksMetaLoad(0))
}
