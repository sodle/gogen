package tests

import (
	"os"
	"path/filepath"
	"testing"

	config "github.com/coccyx/gogen/internal"
	"github.com/coccyx/gogen/run"
	"github.com/stretchr/testify/assert"
)

func TestReplay(t *testing.T) {
	// Setup environment
	config.ResetConfig()
	os.Setenv("GOGEN_HOME", "..")
	os.Setenv("GOGEN_ALWAYS_REFRESH", "")
	home := ".."
	os.Setenv("GOGEN_FULLCONFIG", filepath.Join(home, "tests", "replay", "fullreplay.yml"))

	c := config.NewConfig()
	run.Run(c)

	assert.Equal(t, `2001-10-20T12:00:00
2001-10-20T12:00:13
2001-10-20T12:00:14
2001-10-20T12:00:19
2001-10-20T12:00:29
`, c.Buf.String())
}
