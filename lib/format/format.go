package format

import (
	"math"

	"github.com/svenfinke/gls/lib/types"
)

func FormatFileSize(options types.Options, fileSize int64) (formattedFileSize int64) {
	var factor float64 = 1
	var factorBase float64 = 1024

	if options.Kibibytes {
		factorBase = 1000
	}

	factorBase = factorBase * 0.001
	switch options.BlockSize {
	case "K":
		factor = 0.001 * factorBase
	case "M":
		factor = 0.000001 * factorBase
	case "G":
		factor = 0.000000001 * factorBase
	case "T":
		factor = 0.000000000001 * factorBase
	}

	formattedFileSize = int64(math.Round(float64(fileSize) * factor))
	if formattedFileSize < 1 {
		formattedFileSize = 1
	}
	return
}
