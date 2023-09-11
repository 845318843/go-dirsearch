package main

import (
	"godir/common"
	"godir/utils"

	"github.com/fatih/color"
)

func main() {
	common.Flag()
	common.Parse()
	lim := utils.NewLimit(common.ThreadNum) //线程数
	var bar *utils.Bar
	for _, value := range common.Urls {
		bar = utils.NewBarOption(0, common.FileLine) //任务条
		go utils.ReadFromFile()
		common.Url = value
		common.Colban.Printf("\nUrl: %s                                  \n", value)
		for i := 0; i < int(bar.GetTotal()); i++ {
			lim.Add()
			go utils.ScanTask(lim, bar)
		}
		lim.Wait()
	}
	lim.Wait()
	bar.Finish()
	color.New(color.FgYellow).Printf("\nTask Completed")
}
