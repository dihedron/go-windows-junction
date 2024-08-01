// +build windows

package junction

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows"

	"github.com/Microsoft/go-winio"
)

func create(target, mountPt string) error {
	_target, err := filepath.Abs(target)
	if err != nil {
		return fmt.Errorf("%s: %s", target, err)
	}
	_mountPt, err := windows.UTF16PtrFromString(mountPt)
	if err != nil {
		return fmt.Errorf("%s: %s", mountPt, err)
	}

	err = os.Mkdir(mountPt, 0777)
	if err != nil {
		return fmt.Errorf("%s: %s", mountPt, err)
	}
	ok := false
	defer func() {
		if !ok {
			os.Remove(mountPt)
		}
	}()

	handle, err := windows.CreateFile(_mountPt,
		windows.GENERIC_WRITE,
		0,
		nil,
		windows.OPEN_EXISTING,
		windows.FILE_FLAG_BACKUP_SEMANTICS,
		0)
	if err != nil {
		return fmt.Errorf("%s: %s", mountPt, err)
	}
	defer windows.CloseHandle(handle)

	rp := winio.ReparsePoint{
		Target:       _target,
		IsMountPoint: true,
	}

	data := winio.EncodeReparsePoint(&rp)

	var size uint32

	err = windows.DeviceIoControl(
		handle,
		FSCTL_SET_REPARSE_POINT,
		&data[0],
		uint32(len(data)),
		nil,
		0,
		&size,
		nil)

	if err != nil {
		return fmt.Errorf("windows.DeviceIoControl: %s", err)
	}
	ok = true
	return nil
}
