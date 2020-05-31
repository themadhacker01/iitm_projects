package main

import (
	"log"
	"os"
	"strconv"

	Persist "persist"
)

type Arr struct {
	Size   int
	NewArr []bool
}

/*
======================================================================
writeToFile
- Writes array to file
- Uses persist package
======================================================================
*/

func writeToFile(arr []bool) {

	obj := &Arr{Size: len(arr), NewArr: arr}

	if _, err := os.Stat("./Approxbc"); os.IsNotExist(err) {
		os.Mkdir("./Approxbc", 0777)
	}

	if err := Persist.Save("./Approxbc/bool_file.json", obj); err != nil {
		log.Fatalln(err)
	}
}

/*
======================================================================
readFromFile
- Reads data from bool_file.json
- Converts it into an array
- Returns array to calling function
======================================================================
*/

func readFromFile() []bool {
	var obj2 Arr
	err := Persist.Load("./Approxbc/bool_file.json", &obj2)
	if err != nil {
		log.Fatalln("ERROR", err)
	}
	return obj2.NewArr
}

/*
======================================================================
CREATE : Creates an array of size n and writes to file
======================================================================
*/

func CreateArray(n int) []bool {
	arr := make([]bool, n, n)
	writeToFile(arr)
	return arr
}

/*
======================================================================
WRITE
- Converts string input into boolean
- Stores these values from given offset index
- Wrap around is enabled
- Writes to bool_file.json
======================================================================
*/

func WriteValue(i int, strVal string) {
	var arr = readFromFile()
	count := 0
	revCount := len(strVal) + i - len(arr)
	if revCount < 0 {
		revCount = 0
	}

	// SIMPLE write
	for k := 0; k < (len(strVal) - revCount); k++ {
		count++
		var val bool
		if (strVal[k] - 48) == 0 {
			val = false || arr[i+k]
		}
		if (strVal[k] - 48) == 1 {
			val = true || arr[i+k]
		}
		arr[i+k] = val
	}

	// WRAP AROUND and write
	for k := 0; k < revCount; k++ {
		var val bool
		if (strVal[count+k] - 48) == 0 {
			val = false || arr[k]
		}
		if (strVal[count+k] - 48) == 1 {
			val = true || arr[k]
		}
		arr[k] = val
	}

	writeToFile(arr)
}

/*
======================================================================
READ
- Reads a given no. of values from an offset index
- Wrap around is enabled
======================================================================
*/

func ReadValue(i int, numOfEle int) []bool {
	var readArr = make([]bool, numOfEle, numOfEle)
	var arr = readFromFile()
	count := 0
	revCount := numOfEle + i - len(arr)
	if revCount < 0 {
		revCount = 0
	}

	// SIMPLE read
	for k := 0; k < (numOfEle - revCount); k++ {
		count++
		readArr[k] = arr[k+i]
	}

	// WRAP AROUND and read
	for k := 0; k < revCount; k++ {
		readArr[k+count] = arr[k]
	}
	// fmt.Println("\n", readArr)
	return readArr
}

/*
======================================================================
HASH (hex) to BINARY
- Takes txn hash as hexadecimal string
- Converts it into binary format
- Initialises the array with this
======================================================================
*/

func HashToBinary(val string) (int, string) {
	var input string
	var char uint64

	// Splitting string into offset and array strings
	offStr := val[:2]
	arrStr := val[2:]

	// Converts offStr to integer value for offset index
	index, _ := strconv.ParseInt(offStr, 16, 64)

	// Taking txn hash in hexadecimal and converting it into binary format
	for k := 0; k < len(arrStr); k++ {

		// If hex char is a digit
		// Unicode value : 48-57
		// Convert to : 0-9
		if (arrStr[k] >= 48) && (arrStr[k] <= 57) {
			char = uint64(arrStr[k]) - 48
		}

		// If hex char is a letter
		// Unicode value : 97-122
		// Convert to : 10-15 (only a-f)
		if (arrStr[k] >= 97) && (arrStr[k] <= 122) {
			char = (uint64(arrStr[k]) - 97) + 10
		}

		str := strconv.FormatUint(char, 2)
		strLen := len(str)

		// Appending "0" to the binary string
		// To ensure that each binary string is of uniform length 4
		// Length of resulting bool array = 256
		if strLen < 4 {
			var app string
			for count := 0; count < 4-strLen; count++ {
				app += "0"
			}
			str = app + str
		}
		input += str
	}
	// int64 -> int requires expicit conversion
	return int(index), input
}

/*
======================================================================
FILE HANDLING

CreateReference() - creates golden_reference.txt
WriteReference() - writes (int, string) to golden_reference.txt
======================================================================
*/

func CreateReference() {
	var (
		newFile *os.File
		err     error
	)

	grefPath := "./Approxbc/golden_reference.txt"

	if _, err := os.Stat("./Approxbc"); os.IsNotExist(err) {
		os.Mkdir("./Approxbc", 0777)
	}

	newFile, err = os.Create(grefPath)
	if err != nil {
		log.Fatal("ERROR", err)
	}
	newFile.Close()
}

func AppendReference(i int, val string) {
	var (
		file *os.File
		err  error
	)

	grefPath := "./Approxbc/golden_reference.txt"

	val = strconv.Itoa(i) + " " + val + "\n"

	file, err = os.OpenFile(grefPath, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Fatal(err)
	}

	if _, err = file.WriteString(val); err != nil {
		log.Fatal(err)
	}

	file.Close()
}

/*
======================================================================
MAIN : First function to be executed
======================================================================
*/

// func main() {
// CreateReference()
// AppendReference(3, "ndsjic")

// fmt.Println(HashToBinary("0a9aff"))

// 	// Declare variables
// 	var n, i, numOfEle int
// 	var strVal string

// 	fmt.Println("\n================== CREATE ==================")

// 	// Input array size from user
// 	fmt.Print("\nSize of array : ")
// 	fmt.Scanln(&n)
// 	CreateArray(n)

// 	var arr = readFromFile()

// 	fmt.Println("\n=================== WRITE ===================")

// 	// To input offset index from user
// 	// Repeat loop until valid values are entered by user
// 	for {
// 		fmt.Print("\noffset index i : ")
// 		fmt.Scanln(&i)

// 		if (i >= 0) && (i < len(arr)) {
// 			break
// 		} else {
// 			fmt.Println("\nERROR. Invalid index. Please enter a valid index")
// 		}
// 	}

// 	// To input boolean string from user
// 	// Assume that user only enters 0/1 as string characters
// 	for {
// 		fmt.Print("\nChoice of boolean values :\n1 : true\n0 : false\nEnter string value : ")
// 		fmt.Scanln(&strVal)

// 		if len(strVal) <= len(arr) {
// 			break
// 		} else {
// 			fmt.Println("\nERROR. Invalid string. Please enter a valid string")
// 		}
// 	}
// 	WriteValue(i, strVal)

// 	fmt.Println("\n==================== READ ====================")

// 	// To input offset index from user
// 	// Repeat loop until valid values are entered by user
// 	for {
// 		fmt.Print("\noffset index i : ")
// 		fmt.Scanln(&i)

// 		if (i >= 0) && (i < len(arr)) {
// 			break
// 		} else {
// 			fmt.Println("\nERROR. Invalid index. Please enter a valid index")
// 		}
// 	}

// 	// To input no. of elements index from user
// 	// Repeat loop until valid values are entered by user
// 	for {
// 		fmt.Print("\nNo. of elements to read : ")
// 		fmt.Scanln(&numOfEle)

// 		if numOfEle <= len(arr) {
// 			break
// 		} else {
// 			fmt.Println("\nERROR. Invalid input. Please try again")
// 		}
// 	}

// 	ReadValue(i, numOfEle)
// }
