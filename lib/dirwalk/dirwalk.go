package dirwalk

import (
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/svenfinke/gls/lib/format"
	custom_sort "github.com/svenfinke/gls/lib/sort"
	"github.com/svenfinke/gls/lib/types"
)

func Walk(options types.Options) {
	var files []os.DirEntry

	filepath.WalkDir(options.Args.File, func(p string, info os.DirEntry, err error) error {
		if strings.Count(options.Args.File, string(os.PathSeparator))+1 <= strings.Count(p, string(os.PathSeparator)) {
			return nil
		}

		if !options.All && info.Name() == "." {
			return nil
		}

		if (!options.All && !options.AlmostAll) && strings.HasPrefix(info.Name(), ".") {
			return nil
		}

		if options.IgnoreBackups && strings.HasSuffix(info.Name(), "~") {
			return nil
		}

		fileInfo, _ := info.Info()
		fileSize := format.FormatFileSize(options, fileInfo.Size())
		if options.SizeMaxLength < len(strconv.FormatInt(fileSize, 10)) {
			options.SizeMaxLength = len(strconv.FormatInt(fileSize, 10))
		}
		files = append(files, info)
		return err
	})

	if options.Sort {
		sort.Sort(custom_sort.ByCtime(files))
	}

}
