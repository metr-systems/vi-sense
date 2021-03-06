package api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	. "github.com/vi-sense/vi-sense/app/api"
	. "github.com/vi-sense/vi-sense/app/model"
)

// This is a hack way to add test database for each case, as whole test will just share one database.
// You can read TestWithoutAuth's comment to know how to not share database each case.
func TestMain(m *testing.M) {
	SetupTestDatabase()
	CreateMockData("../../../sample-data", 3)
	exitVal := m.Run()
	DeleteTestDatabase()
	os.Exit(exitVal)
}

func TestPingRoute(t *testing.T) {
	r := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
