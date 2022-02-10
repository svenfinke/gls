package print

import (
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/fatih/color"
	"github.com/svenfinke/gls/lib/format"
	"github.com/svenfinke/gls/lib/types"
)

func Print(options types.Options) {
	var files []os.DirEntry
	var sizeMaxLength int = 0

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
		if sizeMaxLength < len(strconv.FormatInt(fileSize, 10)) {
			sizeMaxLength = len(strconv.FormatInt(fileSize, 10))
		}
		files = append(files, info)
		return err
	})

	for _, info := range files {
		if options.List {
			printList(options, info, sizeMaxLength)
		} else {
			printDefault(options, info)
		}
	}
	fmt.Print("\n")
}

func printDefault(options types.Options, file fs.DirEntry) {
	printColoredFilename(file)
	fmt.Print("  ")
}

func printList(options types.Options, file fs.DirEntry, fileSizeMaxLength int) {
	info, _ := file.Info()
	sep := " "
	stat := info.Sys().(*syscall.Stat_t)
	fileUser, _ := user.LookupId(strconv.FormatUint(uint64(stat.Uid), 10))
	fileGroup, _ := user.LookupGroupId(strconv.FormatUint(uint64(stat.Gid), 10))
	fileSize := format.FormatFileSize(options, info.Size())
	filePermissions := info.Mode().Perm().String()

	if info.IsDir() {
		filePermissions = "d" + filePermissions[1:]
	}

	fmt.Printf("%s%s", filePermissions, sep)
	fmt.Printf("%v%s", stat.Nlink, sep)
	fmt.Printf("%v%s", fileUser.Username, sep) // Owner
	fmt.Printf("%v%s", fileGroup.Name, sep)    // Group
	if options.Author {
		fmt.Printf("%v%s", fileUser.Username, sep) // "Author" - also seems faked in ls. It changes with the owner
	}
	// Generate Format String with maximal found length of size
	fmt.Printf(fmt.Sprintf("%%%vd%%s%%s", fileSizeMaxLength), fileSize, options.BlockSize, sep)
	fmt.Printf("%s%s", info.ModTime().Format("Jan"), sep)
	fmt.Printf("%s%s", info.ModTime().Format("02"), sep)
	fmt.Printf("%s%s", info.ModTime().Format("15:04"), sep)
	printColoredFilename(file)
	fmt.Print("\n")
}

func printColoredFilename(file fs.DirEntry) {
	d := color.New(color.FgWhite)

	if file.IsDir() {
		d = color.New(color.FgHiBlue).Add(color.Bold)
	}
	d.Printf("%v", file.Name())
}
