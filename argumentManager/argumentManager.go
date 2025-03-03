package argumentmanager

import (
	"fmt"
	"os"
	"strings"

	readcsvfile "scrap.com/readCsvFile"
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
				fmt.Println("	-i ou --input 	example.csv")
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

				if arg == "-i" || arg == "--input" {
					if args[i+1] != "" {
						csv := readcsvfile.NewReadCsvFile(args[i+1])
						csv.SplitCsvData()

						str := ConvertArrayToString(csv.GetBlackList())
						clean := strings.Replace(str, " ", "", -1)
						var blacklistTmp []string = strings.Split(clean, ",")

						for i, b := range blacklistTmp {
							blacklistTmp[i] = strings.Replace(b, " ", "", -1)
						}

						for _, b := range blacklistTmp {
							blacklist = append(blacklist, b)
						}

						str = ConvertArrayToString(csv.GetWhiteList())
						clean = strings.Replace(str, " ", "", -1)
						var whitelistTmp []string = strings.Split(str, ",")

						for i, w := range whitelistTmp {
							whitelistTmp[i] = strings.Replace(w, " ", "", -1)
						}

						for _, k := range whitelistTmp {
							keyWord = append(keyWord, k)
						}

					} else {
						fmt.Println("Erreur : argument manquant ou invalide après -i/--input")
					}
				}
			}
		}
	}

	keyWord = trimSpaces(keyWord)
	blacklist = trimSpaces(blacklist)

	return keyWord, blacklist
}

func ConvertArrayToString(slice []string) string {
	return strings.Join(slice, ",")
}

// remove Space between elements
func trimSpaces(slice []string) []string {
	for i, v := range slice {
		slice[i] = strings.TrimSpace(v)
	}
	return slice
}
