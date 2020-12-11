package modules

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestDependentModuleImpl_DoPractice(t *testing.T) {
	testMainModule := mockMainModule{RunFunc: func(index int) {
		assert.Equal(t, index, 5)
	}}
	dependentModule := NewDependentInterface(&testMainModule)
	dependentModule.DoPractice("test")
}

type mockMainModule struct {
	RunFunc func(index int)
}

func (d *mockMainModule) RunMainModule(index int) {
	d.RunFunc(index)
}

