package level2

import (
	"fmt"
	"strconv"
)

// IkhwanOrAkhwat is a type to differentiate between Ikhwan and Akhwat.
// Ikwan has a prefix of ARN, while Akhwat has a prefix of ART.
type IkhwanOrAkhwat int

const (
	IKHWAN IkhwanOrAkhwat = iota // IKHWAN gives ARN prefix.
	AKHWAT                       // AKHWAT gives ART prefix.
)

// Prefix returns the prefix for the IkhwanOrAkhwat type.
func (i IkhwanOrAkhwat) Prefix() string {
	return [...]string{"ARN", "ART"}[i]
}

// CreateNIP creates a single NIP based on the given parameters.
func CreateNIP(ikhwanOrAkhwat IkhwanOrAkhwat, year string, month, id int) string {
	// We don't do validation yet, but we should.
	batch := year[2:]

	// The default semester is first.
	semester := 1
	// If the month is between 7 and 12, then it's the second semester.
	if month > 6 && month <= 12 {
		semester = 2
	}

	// The NIP format is: <prefix><batch><semester>-<id>
	return fmt.Sprintf("%s%s%d-%05d", ikhwanOrAkhwat.Prefix(), batch, semester, id)
}

// GenerateNIPs generates multiple NIPs based on the given parameters.
func GenerateNIPs(
	ikhwanOrAkhwat IkhwanOrAkhwat,
	year string,
	month, count, start int) []string {
	// Initialize the slice with the length of 0 and the capacity of count.
	// Please see https://blog.golang.org/slices-intro for more information.
	nips := make([]string, 0, count)

	// Loop from start to start+count.
	for i := start; i < start+count; i++ {
		nips = append(nips, CreateNIP(ikhwanOrAkhwat, year, month, i))
	}

	// Return the slice.
	return nips
}

// CreateNextNIP creates the next NIP based on the given NIP and the next number.
// "next" gives the "spaces" of the next NIP. For example, if the NIP is ARN191-00001,
// and the next is 3, then the next NIP will be ARN191-00004.
func CreateNextNIP(nip string, next int) string {
	// There are many ways of extracting data, including using regex.
	// But for now, we'll use the simple way.
	// Also, we don't do validation yet, but we should.
	prefix := nip[:3]
	batch := nip[3:5]
	semester := nip[5:6]

	id := nip[7:]
	// This can return an error, but we'll ignore it for now.
	currentID, _ := strconv.Atoi(id)

	// The NIP format is: <prefix><batch><semester>-<id>
	return fmt.Sprintf("%s%s%s-%05d", prefix, batch, semester, currentID+next)
}

// GenerateNextNIPs generates the next NIPs based on the given NIP and the count.
func GenerateNextNIPs(nip string, count int) []string {
	// Initialize the slice with the length of 0 and the capacity of count.
	nips := make([]string, 0, count)

	// Loop from 1 to count.
	for i := 1; i <= count; i++ {
		nips = append(nips, CreateNextNIP(nip, i))
	}

	// Return the slice.
	return nips
}
