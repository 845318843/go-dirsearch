package common

import (
	"flag"
)

//显示logo信息
func Banner() {
	banner := `
	#   #       #                             ########           #       
	#    ##    ###      ##############   #         ##       #    ##  #   
	#     ##  ##        ##         ##    ##     ## #     ##### ########  
	#  #############    ##      #  ##     ##  #  ##  #     ##    ## ##   
	#       ##          ## ####### ##     #   #########    ## ########## 
	#   ###########     ##   ##    ##      #  ## ## ##    ##     ## ##   
	#       ##          ##   ## #  ##   ##### ########    # #  #######   
	# ###############   ## ####### ##     ##  ## ## ##   #####   ## #    
	#       ##          ##   ###   ##     ##  ## ## ##     ##    ##      
	#       ##          ##   ####  ##     ##  ########     ## ########   
	#  #############    ##   ## ## ##     ##  ## ## ##   # ##    ##      
	#      ## #         ##   ## #  ##     ##  ## ## ##    ### #########  
	#     ##   ##       #############     ### ## # ###     ###   ##      
	#    ##     ###     ##         ##    ## ###     #     ## ### #       
	#   ##       ####   #############   ##   ##########   #   ########## 
	# ##           #    #          #     #    ########   #      #######  
                     
							      网信安检查工具-目录扫描
`
	Colban.Println(banner)
}

//读取命令行参数
func Flag() {
	Banner()
	flag.StringVar(&Url, "u", "", "Target URL")
	flag.StringVar(&Extention, "e", "php, aspx, jsp, html, js", "Extension list separated by commas (e.g. php,asp)")
	flag.StringVar(&OutFile, "o", "", "Output File(default  url.result.txt)")
	flag.StringVar(&UrlFile, "uf", "", "Target URL File")
	flag.StringVar(&UA, "ua", "", "Set User-Agent(default  random)")
	flag.StringVar(&WordList, "wd", "dict/dict.txt", "Brute Dict(default  dict/dict.txt)")
	flag.StringVar(&ReqMethod, "m", "GET", "Request Method(default GET)")
	flag.IntVar(&ThreadNum, "t", 10, "Number of threads(default 10)")
	flag.IntVar(&Timeout, "to", 3, "Number of timeout(default 3)")
	flag.Parse()

}
