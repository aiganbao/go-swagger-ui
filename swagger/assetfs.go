package swagger

import (
	"github.com/elazarl/go-bindata-assetfs"
	"net/http"
)

var SwaggerFileSystem = assetfs.AssetFS{
	Asset:     Asset,
	AssetDir:  AssetDir,
	AssetInfo: AssetInfo,
	Prefix:    "swagger-ui",
}

var SwaggerFs = func(dir string) http.FileSystem {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    dir,
	}
}

var SwaggerHttpHandle = http.StripPrefix("/swagger-ui/", http.FileServer(&SwaggerFileSystem))

var SwaggerPrefixHandle = func(prefix string) http.Handler {
	return http.StripPrefix(prefix, http.FileServer(&SwaggerFileSystem))
}
