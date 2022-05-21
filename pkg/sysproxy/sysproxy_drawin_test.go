//go:build darwin

package sysproxy

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestGetNetworkInterface(t *testing.T) {
	i := GetNetworkInterfaces("default")
	log.Println(i)

	i = GetNetworkInterfaces("all")
	log.Println(i)
}

func TestSetSysProxy(t *testing.T) {
	err := SetSysHttpProxy(":9191", "default")
	assert.Equal(t, nil, err)
	err = SetSysSocksProxy(":9192", "default")
	assert.Equal(t, nil, err)
}
func TestUnsetSysProxy(t *testing.T) {
	err := UnsetSysHttpProxy("default")
	assert.Equal(t, nil, err)
	err = UnsetSysSocksProxy("default")
	assert.Equal(t, nil, err)
}
