package sort

import "io/fs"

type ByCtime []fs.DirEntry

func (a ByCtime) Len() int {
	return len(a)
}
func (a ByCtime) Less(i, j int) bool {
	iInfo, _ := a[i].Info()
	jInfo, _ := a[j].Info()

	return iInfo.ModTime().Unix() > jInfo.ModTime().Unix()
}
func (a ByCtime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
