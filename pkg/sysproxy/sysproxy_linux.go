//go:build linux

package sysproxy

func SetSysHttpProxy(listenAddr, mode string) error {
	return nil
}

func UnsetSysHttpProxy(mode string) error {
	return nil
}

func SetSysSocksProxy(listenAddr, mode string) error {
	return nil
}

func UnsetSysSocksProxy(mode string) error {
	return nil
}
