package code

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImportPathForDir(t *testing.T) {
	wd, err := os.Getwd()
	require.NoError(t, err)

	assert.Equal(t, "github.com/DoNotPayHQ/gqlgen-fork/internal/code", ImportPathForDir(wd))
	assert.Equal(t, "github.com/DoNotPayHQ/gqlgen-fork/api", ImportPathForDir(filepath.Join(wd, "..", "..", "api")))

	// doesnt contain go code, but should still give a valid import path
	assert.Equal(t, "github.com/DoNotPayHQ/gqlgen-fork/docs", ImportPathForDir(filepath.Join(wd, "..", "..", "docs")))

	// directory does not exist
	assert.Equal(t, "github.com/DoNotPayHQ/gqlgen-fork/dos", ImportPathForDir(filepath.Join(wd, "..", "..", "dos")))

	if runtime.GOOS == "windows" {
		assert.Equal(t, "", ImportPathForDir("C:/doesnotexist"))
	} else {
		assert.Equal(t, "", ImportPathForDir("/doesnotexist"))
	}
}

func TestNameForPackage(t *testing.T) {
	assert.Equal(t, "api", NameForPackage("github.com/DoNotPayHQ/gqlgen-fork/api"))

	// does not contain go code, should still give a valid name
	assert.Equal(t, "docs", NameForPackage("github.com/DoNotPayHQ/gqlgen-fork/docs"))
	assert.Equal(t, "github_com", NameForPackage("github.com"))
}

func TestNameForDir(t *testing.T) {
	wd, err := os.Getwd()
	require.NoError(t, err)

	assert.Equal(t, "tmp", NameForDir("/tmp"))
	assert.Equal(t, "code", NameForDir(wd))
	assert.Equal(t, "internal", NameForDir(wd+"/.."))
	assert.Equal(t, "main", NameForDir(wd+"/../.."))
}
