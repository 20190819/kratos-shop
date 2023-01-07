package server

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.
// 由于此服务只对外提供 http 服务
var ProviderSet = wire.NewSet(NewHTTPServer)
