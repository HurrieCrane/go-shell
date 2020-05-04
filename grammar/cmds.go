package grammar

// CMD is an alias for int
// and represents commands
type CMD int

const (
	// Echo s
	Echo CMD = iota
	// TotalCmdCount s
	TotalCmdCount

	InvalidToken
)

var reverseLookupTable map[string]CMD = generateReverseLookupTable()

// generateReverseLookupTable creates a reverse
// lookup table to search the value of a string
// against an enumeration value
func generateReverseLookupTable() map[string]CMD {
	rlut := make(map[string]CMD)
	for i := 0; i < int(TotalCmdCount-1); i = i + 1 {
		rlut[CMD(i).String()] = CMD(i)
	}
	return rlut
}

// ReverseLookupString will return an enumeration
// value for a string in O(n) time
func ReverseLookupString(val *string) CMD {
	for i := 0; i < int(TotalCmdCount); i = i + 1 {
		if val, ok := reverseLookupTable[*val]; ok {
			return val
		}
	}
	return InvalidToken
}

//go:generate stringer -type=CMD
