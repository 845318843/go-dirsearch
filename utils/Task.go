package utils

import "godir/common"

func ScanTask(g *Golimit, bar *Bar) {

	bar.AddCurOne()

	uri := ReadFromChan()
	DoRequest(common.Url, uri)
	defer g.Done()
	bar.Play()

}
