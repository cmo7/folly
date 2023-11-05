package factories

import (
	"github.com/brianvoe/gofakeit/v6"
)

// faker is a global instance of gofakeit, scoped to this package.
var faker = gofakeit.NewCrypto()
