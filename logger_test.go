//+build unit

package logger

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNew(t *testing.T) {
	//Act
	InitLogger(false)

	//Assert
	assert.Equal(t, reflect.TypeOf(*logStdOut), reflect.TypeOf(zap.Logger{}))
	assert.Equal(t, reflect.TypeOf(*logStdErr), reflect.TypeOf(zap.Logger{}))
}
