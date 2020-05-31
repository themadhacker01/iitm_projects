# This is a step-by-step guide of implementing the hash table and eBF :

**NOTE : ** All commands are to be executed in the terminal window open in pdsautomation directory as the file paths used have been hard-coded.

### Integrating the code

We need to integrate the above written code with the go-ethereum-1.8.27 source code. To do this we take the following steps :

* Add the packages `AMEY/edits` and `AMEY/persist` to `go/src`
* Add this to the end of `WriteBodyRLP()` function in `go-ethereum-1.8.27/core/rawdb/accessors_chain.go`
```go
// AMEY
// number == 0 is for the Genesis block to make sure array is created only once at the beginning
// len = 256 : a hash has 64 hexadecimal characters, each can be represented by 4 bits {0, 1}
// CreateReference() creates Approxbc/bool_file.json
if number == 0 {
	Edits.CreateArray(256)
	Edits.CreateReference()
}
```
* Add this just before the `return` statement of `submitTransaction()` function in `go-ethereum-1.8.27/internal/ethapi/api.go`
```go
// AMEY
// tx.Hash().Hex() returns the hash as a hexadecimal string
// The first 2 characters in the string are are 0, x --- we need to ignore these
// For now, we use offset index = 0
// AppendReference() writes offset index & string value to Approxbc/golden_reference.txt
offset, strVal := Edits.HashToBinary(tx.Hash().Hex()[2:])
Edits.WriteValue(offset, strVal)
Edits.AppendReference(offset, strVal)
```

### Explanation

There are 2 packages `edits` and `persist`.

**Description for** `edits/hash_table.go` :

* `writeToFile` : Writes array to file using persist package

* `readFromFile`
    1. Reads data from `./bool_file.json`
    2. Converts it into an array
    3. Returns array to calling function

* `CreateArray` - Creates an array of size n and writes to file

* `WriteValue`
    1. Converts string input into boolean
    2. Stores these values from given offset index
    3. Wrap around is enabled
    4. Writes to `./bool_file.json`

* `ReadValue`
    1. Reads a given no. of values from an offset index
    2. Wrap around is enabled
    
* `HashToBinary`
    1. Takes txn hash as hexadecimal string
    2. Converts it into binary format
    3. Initialises the array with this
    
**Description for** `persist/persist.go` :

The `persist` package is to convert a struct instance written in golang to JSON object. We use the `Save()` and `Load()` functions.
- The `Save()` function creates a file with .json extension and saves this object into the file.
- `Load()` reads a JSON object from the file and returns a golang struct.

This link has been diligently followed for this purpose : https://medium.com/@matryer/golang-advent-calendar-day-eleven-persisting-go-objects-to-disk-7caf1ee3d11d

**Description for** `validate.go` :

For a bitwise-or accumulator, usefulness is computed as the percentage of bits in the response (i.e., accumulated evidence) which are NOT overwritten with 1.

* `createUsePercent` : Creates the file use_percent.txt

* `appendUsePercent` : Appends the index & use percentage to the file use_percent.txt

* `main` : Calculates the use percentage and calls `createUsePercent()` and `appendUsePercent()` functions
    1. Reads the offset index and actual hash values from `golden_reference.txt` file and stores it in `index` and `gref`
    2. Reads stored hash value (last 248 bits) in `bool_file.json` from the offset index (first 8 bits), and stores this value in `valArr`
    3. Extracts each boolean value from `gref`, `valArr` and stores it in `g`, `v`
    4. Compares `g` and `v` using simple logic
    
    |    g    |    v    |    res    |
    |:-------:|:-------:|:-------:|
    |    0    |    0    |    1    |
    |    0    |    1    |    0    |
    |    1    |    0    |    x    |
    |    1    |    1    |    1    |
    
    This gives the boolean logic : `res = g || (!v)` which we use to check the no. of bits not overwritten, and hence, the usefulness of our solution