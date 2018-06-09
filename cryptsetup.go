package cryptsetup

// #cgo LDFLAGS: -lcryptsetup -lluksmeta
// #include <errno.h>
// #include <libcryptsetup.h>
// #include <luksmeta.h>
import "C"

import (
	"fmt"
	"unsafe"
)

var (
	CRYPT_LUKS1 = MetaType(C.CRYPT_LUKS1)
)

type MetaType string

type CryptDevice struct {
	cpCryptDevice *C.struct_crypt_device
	path          string
}

type CryptVerifyInfo struct {
	cCryptVerifyInfo C.struct_crypt_params_verity
}

func NewDevice(devicePath string) (*CryptDevice, error) {
	dev := CryptDevice{path: devicePath}
	errCode := C.crypt_init(&dev.cpCryptDevice, C.CString(devicePath))
	if errCode != 0 {
		return nil, fmt.Errorf(`Got error while crypt_init("%v"), incorrect path or permission denied? Error code: %v`, devicePath, errCode)
	}
	return &dev, nil
}

func (dev *CryptDevice) LoadMeta(metaType MetaType) error {
	errCode := C.crypt_load(dev.cpCryptDevice, C.CString(string(metaType)), C.NULL)
	if errCode != 0 {
		return fmt.Errorf(`Got error while crypt_load(): %v`, errCode)
	}
	return nil
}

func (dev *CryptDevice) GetUUID() []byte {
	return []byte(C.GoString(C.crypt_get_uuid(dev.cpCryptDevice)))
}

/*func (dev *CryptDevice) SetLogCallback(f func(int, string, ...interface{}), args ...interface{}) {
	C.crypt_set_log_callback(dev.cpCryptDevice, func(level C.int, msg *C.char, usrptr unsafe.Pointer){
		f(level, msg, args...)
	}, C.NULL)
}*/

func (dev *CryptDevice) Dump() error {
	errCode := C.crypt_dump(dev.cpCryptDevice)
	if errCode != 0 {
		return fmt.Errorf(`Got error while crypt_dump(): %v`, errCode)
	}
	return nil
}

func (dev *CryptDevice) GetVerifyInfo() (*CryptVerifyInfo, error) {
	result := CryptVerifyInfo{}
	errCode := C.crypt_get_verity_info(dev.cpCryptDevice, &result.cCryptVerifyInfo)
	if errCode != 0 {
		return nil, fmt.Errorf(`Got error while crypt_get_verity_info(): %v`, errCode)
	}
	return &result, nil
}

func (dev *CryptDevice) LuksMetaInit() error {
	errCode := C.luksmeta_init(dev.cpCryptDevice)
	if errCode != 0 {
		return fmt.Errorf(`Got error while luksmeta_init(): %v: %v %v`, errCode, errCode == -C.EALREADY, errCode == -C.ENOSPC)
	}
	return nil
}

func (dev *CryptDevice) KeySlotMax(metaType MetaType) int {
	return int(C.crypt_keyslot_max(C.CString(string(metaType))))
}

func (dev *CryptDevice) LuksMetaLoad(slot int) error {
	var cUuid [16]C.uchar
	var buf [65536]byte
	errCode := C.luksmeta_load(dev.cpCryptDevice, C.int(slot), &cUuid[0],  (unsafe.Pointer)(&buf[0]), 65536)
	if errCode != 0 {
		return fmt.Errorf(`Got error while luksmeta_load() -> %v: %v: %v %v %v %v %v %v %v %v`, cUuid, errCode, errCode == -C.ENOENT, errCode == -C.EINVAL, errCode == -C.EBADSLT, errCode == -C.EKEYREJECTED, errCode == -C.EALREADY, errCode == -C.ENOSPC, errCode == -C.E2BIG, errCode == -C.ENODATA)
	}
	return nil
}

func (dev *CryptDevice) Free() {
	C.crypt_free(dev.cpCryptDevice)
}
