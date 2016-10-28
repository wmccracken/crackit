package main

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/wmccracken/crackit/ascii"
	"github.com/wmccracken/crackit/dictionary"
	"github.com/wmccracken/crackit/numeric"
)

var DICTFILE = "./english.txt"
var SHORTDICTFILE = "./short_english.txt"
var MAXPWLENGTH = 8
var WORDCOUNT = 2
var password = "science!science"

func main() {

	// log.SetLevel(log.DebugLevel)
	log.Info("Beginning test of password: " + password)
	// log.SetLevel(log.DebugLevel)
	start := time.Now()
	if dictionary.CheckAgainstDictionary(DICTFILE, password) {
		log.Info("Found a dictionary word match!")
	} else if dictionary.CheckMultipleAgainstDictionary(SHORTDICTFILE, password, WORDCOUNT) {
		log.Info("Found a multi-word dictionary match!")
	} else if numeric.CheckAgainstDigits(password, MAXPWLENGTH) {
		log.Info("Found a numeric match!")
	} else if ascii.CheckAscii(password, MAXPWLENGTH) {
		log.Info("Found an ascii match!")
	}

	elapsed := time.Since(start)
	log.WithField("elapsed", elapsed).Info("Finished cracking passwords.")

}
