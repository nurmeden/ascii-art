package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	words := os.Args[1:]
	if len(words) == 3 {
		option := words[0]
		word := words[1]
		banner := words[2]
		readFiles := []string{"shadow", "standard", "thinkertoy"}
		if CheckInLatin(word) && (strings.HasPrefix(option, "--output=") && strings.HasSuffix(option, ".txt")) && indexOf(string(banner), readFiles) >= 0 {
			reading_file := string(banner) + ".txt"
			creating_file := option[9:]
			if indexOf(creating_file[:len(creating_file)-4], readFiles) < 0 {
				readFile, err := os.Open(reading_file)
				ErrorHandling(err)
				creatFile, err2 := os.Create(creating_file)
				ErrorHandling(err2)

				defer creatFile.Close()

				fileScanner := bufio.NewScanner(readFile)
				fileScanner.Split(bufio.ScanLines)
				var fileLines []string

				for fileScanner.Scan() {
					fileLines = append(fileLines, fileScanner.Text())
				}

				defer readFile.Close()
				// flag := true
				data := make(map[int]string)
				symbol := 32
				for i, line := range fileLines {
					if line == "" {
						data[i+1] = string(rune(symbol))
						symbol++
					}
				}
				fmt.Println(data)

				lines := []string{"lines", "lines2", "lines3", "lines4", "lines5", "lines6", "lines7", "lines8"}
				words_replace_all := strings.ReplaceAll(word, "\\n", "\n")
				splitted_words := strings.Split(words_replace_all, "\n")
				linesMap := make(map[string]string)
				index_newline := []int{}
				for _, words := range splitted_words {
					count := 0
					for _, lettersInWords := range string(words) {
						for keysIndx, lettersInData := range data {
							if string(lettersInWords) == lettersInData {
								for i := 0; i < len(lines); i++ {
									linesMap[lines[i]] += fileLines[keysIndx+i]
									if i == 0 {
										count += len(fileLines[keysIndx+i])
									}
								}
							}
						}
					}
					index_newline = append(index_newline, count)
				}
				counter_for_zero := 0
				for c := 0; c < len(index_newline); c++ {
					if index_newline[c] == 0 {
						counter_for_zero++
					}
				}
				flag := true
				if counter_for_zero == len(index_newline) {
					index_newline = index_newline[:len(index_newline)-1]
					if counter_for_zero == 3 {
						flag = false
					}
				}
				result := ""
				if counter_for_zero == 3 && !flag {
					result += "\n" + "\n"
				} else {
					start := 0
					if len(index_newline) == 2 && index_newline[0] == 0 && index_newline[1] == 0 {
						result += ""
					} else {
						for z := 0; z < len(index_newline); z++ {
							end := index_newline[z]
							if end == 0 {
								result += "\n"
							} else {
								for i := 0; i < len(lines); i++ {
									if i == len(lines)-1 && z == len(index_newline)-1 {
										result += linesMap[lines[i]][start:start+end] + "\n" + "\n"
									} else {
										result += linesMap[lines[i]][start:start+end] + "\n"
									}
								}
							}
							start = start + end
						}
					}
				}
				_, err3 := creatFile.WriteString(result)
				ErrorHandling(err3)
			} else {
				fmt.Println("The file does not open")
			}

		} else {
			fmt.Println("Incorrect input arguments")
		}
		if !CheckInLatin(word) {
			fmt.Println("Incorrect symbols")
		}
	} else {
		fmt.Println("Incorrect input arguments")
	}
}

func CheckInLatin(word string) bool {
	flagger := 0
	for _, letters := range word {
		if (letters >= 32 && letters <= 126) || letters == 10 {
			flagger++
		}
	}
	if flagger == len(word) {
		return true
	} else {
		return false
	}
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func ErrorHandling(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
