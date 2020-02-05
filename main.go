package oauth2

import (
	"net/http"

	"github.com/nektro/go-util/types"
	"github.com/rakyll/statik/fs"

	_ "github.com/nektro/go.oauth2/statik"
)

var (
	mfs = new(types.MultiplexFileSystem)
)

func init() {
	statikFS, err := fs.New()
	if err != nil {
		return
	}
	mfs.Add(http.FileSystem(statikFS))
}
