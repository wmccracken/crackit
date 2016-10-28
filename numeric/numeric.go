package numeric

import (
	"regexp"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/wmccracken/crackit/util"
)

func CheckAgainstDigits(password string, maxlength int) (passwordFound bool) {
	match, err := regexp.MatchString("\\D", password)
	util.Check(err)

	if match {
		log.Info("Skipping digit check - password is not all digits")
		return false
	}
	passwordFound = false
	log.Info("Performing numeric tests.")
	i := 0

	for passwordFound != true && len(strconv.Itoa(i)) < maxlength {
		testString := strconv.Itoa(i)
		if testString == password {
			log.Info("Found an exact numeric match!")
			passwordFound = true
			break
		} else {
			i++
		}
	}

	i = 0

	for passwordFound != true && len(strconv.Itoa(i)) < maxlength {
		testString := strconv.Itoa(i)

		for passwordFound != true && len(testString) <= maxlength {
			testString = "0" + testString
			log.WithField("password", testString).Debug("Testing Password")
			if testString == password {
				log.Info("Found a padded numeric match!")
				passwordFound = true
				break
			}
		}
		if passwordFound == true {
			break
		}
		i++
	}

	return
}
