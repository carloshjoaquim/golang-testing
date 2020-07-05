package test

import (
	"github.com/carloshjoaquim/golang-testing/src/api/app"
	"github.com/federicoleon/golang-restclient/rest"
	"os"
	"testing"
)

func TestMain (m *testing.M) {
	rest.StartMockupServer()
	go app.StartApp()
	os.Exit(m.Run())
}
