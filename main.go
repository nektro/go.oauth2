package oauth2

import (
	"net/http"

	"github.com/nektro/go-util/types"
	"github.com/nektro/go-util/vflag"
	"github.com/rakyll/statik/fs"

	_ "github.com/nektro/go.oauth2/statik"
)

var (
	mfs = new(types.MultiplexFileSystem)
	doa string // default auth
)

func init() {
	statikFS, err := fs.New()
	if err != nil {
		return
	}
	mfs.Add(http.FileSystem(statikFS))

	vflag.StringVar(&doa, "oauth2-default-auth", "", "A default auth to use when multiple appconf's are enabled.")
}
