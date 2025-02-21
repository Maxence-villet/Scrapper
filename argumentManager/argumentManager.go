package argumentmanager

import (
	"fmt"
	"os"
	"strings"
)

func GetArguments() []string {
	return os.Args
}

func FilterArguments(args []string) ([]string, []string) {
	var keyWord []string
	var blacklist []string
	for i, arg := range args {
		if args[i] != args[0] {
			if arg == "-h" || arg == "--help" {
				fmt.Println("--------------------------------")
				fmt.Println("\n		Détail des commandes : \n")
				fmt.Println("	-k ou --key 		\"[OPTION1,OPTION2,...]\"")
				fmt.Println("	-b ou --blacklist 	\"[OPTION1,OPTION2,...]\"")
				fmt.Println()
				fmt.Println("--------------------------------")
				break
			}
			if i < len(args)-1 {
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
				if arg == "-b" || arg == "--blacklist" {
					if args[i+1] != "" && strings.Contains(args[i+1], "[") && strings.Contains(args[i+1], "]") {
						clean := strings.Replace(args[i+1], "[", "", -1)
						clean = strings.Replace(clean, "]", "", -1)
						var blacklistTmp []string = strings.Split(clean, ",")
						for _, b := range blacklistTmp {
							blacklist = append(blacklist, b)
						}
					}
				}
			}
		}
	}

	return keyWord, blacklist
}
