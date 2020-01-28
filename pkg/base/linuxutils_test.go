package base

import (
	"testing"

	"eos2git.cec.lab.emc.com/ECS/baremetal-csi-plugin.git/pkg/mocks"
	"github.com/stretchr/testify/assert"
)

func TestLinuxUtils_LsblkSuccess(t *testing.T) {
	e := mocks.NewMockExecutor(map[string]mocks.CmdOut{LsblkCmd: mocks.LsblkTwoDevices})
	l := NewLinuxUtils(e)

	out, err := l.Lsblk(DriveTypeDisk)
	assert.Nil(t, err)
	assert.NotNil(t, out)
	assert.Equal(t, 2, len(*out))

}

func TestLinuxUtils_LsblkFail(t *testing.T) {
	e1 := mocks.EmptyExecutorSuccess{}
	l := NewLinuxUtils(e1)

	out, err := l.Lsblk(DriveTypeDisk)
	assert.Nil(t, out)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "invalid character")

	e2 := mocks.EmptyExecutorFail{}
	l.SetExecutor(e2)
	out, err = l.Lsblk(DriveTypeDisk)
	assert.Nil(t, out)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "error")

	e3 := mocks.NewMockExecutor(map[string]mocks.CmdOut{LsblkCmd: mocks.NoLsblkKey})
	l.SetExecutor(e3)
	out, err = l.Lsblk(DriveTypeDisk)
	assert.Nil(t, out)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unexpected lsblk output format")
}