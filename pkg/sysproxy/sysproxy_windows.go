//go:build windows

package sysproxy

import (
	"github.com/Trisia/gosysproxy"
)

func SetSysHttpProxy(listenAddr, mode string) error {
	return gosysproxy.SetGlobalProxy(listenAddr)
}

func UnsetSysHttpProxy(mode string) error {
	return gosysproxy.Off()
}

func SetSysSocksProxy(listenAddr, mode string) error {
	return nil
}

func UnsetSysSocksProxy(mode string) error {
	return nil
}
