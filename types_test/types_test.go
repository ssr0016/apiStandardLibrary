package typestest

import (
	"fmt"
	"testing"

	"github.com/ssr0016/gobank/types"
	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	acc, err := types.NewAccount("a", "b", "123456")
	assert.Nil(t, err)

	fmt.Printf("%+v\n", acc)
}
