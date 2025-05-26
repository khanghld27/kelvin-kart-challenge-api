package registry

import (
	"github.com/google/wire"
	"github.com/khanghld27/kelvin-kart-challenge-api/pkg/gormer"
)

var (
	singletonSet = wire.NewSet(
		gormer.GetDB,
	)
)
