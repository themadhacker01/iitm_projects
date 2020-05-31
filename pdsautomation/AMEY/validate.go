/*
For a bitwise-or accumulator, usefulness is computed as the percentage of bits in the response
(i.e., accumulated evidence) which are NOT overwritten with 1
*/

package main

import (
	"bufio"
	Edits "edits"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Creates the file use_percent.txt
func createUsePercent() {
	var (
		newFile *os.File
		err     error
	)
	// Creates use_percent to store usefulness percentage
	usePath := "./Approxbc/use_percent.txt"
	newFile, err = os.Create(usePath)
	if err != nil {
		log.Fatal("ERROR", err)
	}
	newFile.Close()
}

// To append the index, use percentage to the file use_percent.txt
func appendUsePercent(i int, percent float32) {
	var (
		file *os.File
		err  error
		f    string
	)

	grefPath := "./Approxbc/use_percent.txt"
	f = fmt.Sprintf("%f", percent)
	val := strconv.Itoa(i) + " " + f + "\n"

	file, err = os.OpenFile(grefPath, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Fatal(err)
	}

	if _, err = file.WriteString(val); err != nil {
		log.Fatal(err)
	}

	file.Close()
}

func main() {

	var (
		file  *os.File
		index int
		gref  string
		err   error
	)

	// Creates use_percent.txt
	createUsePercent()

	// Opens golden_reference file in read-only mode
	// To read the actual hash values for comparison
	path := "./Approxbc/golden_reference.txt"
	file, err = os.OpenFile(path, os.O_RDONLY, 0660)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		// Splits read lines into index and gref string
		str := scanner.Text()
		s := strings.Index(str, " ")
		index, _ = strconv.Atoi(str[0:s])
		gref = str[s+1:]

		// Reads value from bool_file.json
		// index : offset index
		// 248 = 256 (total bits in hash) - 8 (bits used in index)
		valArr := Edits.ReadValue(index, 248)

		// tot : total length of gref boolean string
		// count : no. of 0-value bits NOT overwritten by 1
		var tot, count int
		var percent float32
		tot = len(gref)

		// Looping over length of boolean string
		for k := 0; k < len(gref); k++ {

			var g, v, res bool

			// Converts each bit in gref string to boolean
			// Stores in g
			if (gref[k] - 48) == 1 {
				g = true
			} else {
				g = false
			}

			// Converts each bit in val string to boolean
			// Stores in v
			v = valArr[k]

			// Implementing logic for result
			res = (!v) || g
			if res == true {
				count++
			}
		}
		percent = float32(count) / float32(tot) * 100
		appendUsePercent(index, percent)
	}
}
