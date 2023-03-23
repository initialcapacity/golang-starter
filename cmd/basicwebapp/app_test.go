package main

import (
	"os"
	"testing"

	"github.com/initialcapacity/golang-starter/pkg/testsupport"
	"github.com/initialcapacity/golang-starter/pkg/websupport"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	_ = os.Setenv("PORT", "0")
	go main()
}

func TestApp(t *testing.T) {
	app, listener := newApp("localhost:8888")
	go websupport.Start(app, listener)
	testsupport.WaitForHealthy(app, "health")
	websupport.Stop(app)
}

func TestNewApp(t *testing.T) {
	_ = os.Setenv("PORT", "0")
	app, _ := newApp("localhost:0")
	assert.NotNil(t, app)
}
