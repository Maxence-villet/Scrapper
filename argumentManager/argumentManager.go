package argumentmanager

import (
	"fmt"
	"os"
	"strings"
)

func GetArguments() []string {
	return os.Args
}

func FilterArguments(args []string) []string {
	var keyWord []string
	for i, arg := range args {
		if args[i] != args[0] {
			if arg == "-h" || arg == "--help" {
				fmt.Println("Usage: scrap ")
			}
			if arg == "-k" || arg == "--key" {
				if args[i+1] != "" && strings.Contains(args[i+1], "[") && strings.Contains(args[i+1], "]") {
					clean := strings.Replace(args[i+1], "[", "", -1)
					clean = strings.Replace(clean, "]", "", -1)
					var keyWordTmp []string = strings.Split(clean, ",")
					for _, k := range keyWordTmp {
						keyWord = append(keyWord, k)
					}
				}
			}
		}
	}
	if keyWord == nil {
		keyWord = args
	}
	return keyWord
}
