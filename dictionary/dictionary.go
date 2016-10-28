package dictionary

import (
	"bufio"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/wmccracken/crackit/util"
	"github.com/wmccracken/stringutil"
)

func readFile(path string) (f *os.File) {
	f, err := os.Open(path)
	util.Check(err)

	return
}

func CheckAgainstDictionary(filePath string, word string) (passwordFound bool) {
	log.Info("Performing dictionary tests.")
	passwordFound = false
	dictFile := readFile(filePath)
	defer func() {
		err := dictFile.Close()
		util.Check(err)
	}()

	scanner := bufio.NewScanner(dictFile)

	for scanner.Scan() {
		dictWord := scanner.Text()

		if word == dictWord {
			passwordFound = true
			break
		} else if strings.ToUpper(dictWord) == word {
			passwordFound = true
			break
		} else if stringutil.Capitalize(dictWord) == word {
			passwordFound = true
			break
		}
	}
	return
}

func CheckMultipleAgainstDictionary(filePath string, word string, wordCount int) (passwordFound bool) {
	log.Info("Performing multi-word dictionary tests.")
	passwordFound = false
	dictFile := readFile(filePath)
	defer func() {
		err := dictFile.Close()
		util.Check(err)
	}()

	// bring the dictionary into an array
	scanner := bufio.NewScanner(dictFile)
	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	// how many words are in the dictionary?
	dictSize := len(words)

	// list of characters to join words together with
	//X will be used to directly join words (no special character)
	testChars := " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~X"

	// start with 2 words
	numWords := 2

	for passwordFound == false && numWords <= wordCount {
		placeCounter := make([]int, numWords) // these are zeroed by default

		// stay in the loop if on of the placeCounter values is
		// under the dictSize
		for passwordFound == false && util.ContainsLessThan(placeCounter, dictSize-1) {
			log.Info(placeCounter)
			var testString = ""
			for testCharIndex := 0; testCharIndex < len(testChars); testCharIndex++ {
				joinChar := ""
				if testChars[testCharIndex] != 'X' {
					joinChar = string(testChars[testCharIndex])
				}
				var parts []string
				for _, a := range placeCounter {
					parts = append(parts, words[a])
				}
				testString = strings.Join(parts, joinChar)
				log.Debug("Test String: " + testString)

				if word == testString {
					passwordFound = true
					break
				}
			}
			if passwordFound == true {
				break
			} else {
				// increment the placeCounters
				curPlace := len(placeCounter) - 1
				increment := true
				for increment == true && curPlace >= 0 {

					if placeCounter[curPlace] == dictSize-1 {
						if curPlace != 0 {
							placeCounter[curPlace] = 0
						} else {
							increment = false
						}
						// log.WithFields(log.Fields{"counter": curPlace, "value": placeCounter[curPlace]}).Info("Counter Rolling: ")
					} else {
						placeCounter[curPlace]++
						increment = false
					}
					curPlace--
				}
			}
		}

		if passwordFound {
			break
		}
		numWords++
	}
	return
}
