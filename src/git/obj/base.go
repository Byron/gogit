// Implement basic functions shared by all object types
package obj

import (
		"strings";
		"bytes";
		"encoding/hex";
		"os";
		"fmt";
)

// Type representing a sha1 with 20 bytes 
type Sha1 [20]byte;

// String converts a Sha1 to a hexadecimal representation
func (s *Sha1) String() string {
	return hex.EncodeToString(s)
}

// Compare returns True if lhs == rhs
func (s *Sha1) Equals( rhs *Sha1 ) bool {
	return bytes.Compare(s, rhs) == 0;
}


// Indicates the passed in string had a lenght != 40 bytes
type InvalidLengthError uint8;
func (s InvalidLengthError) String() string { return fmt.Sprintf("Invalid Length: %i", uint8(s)) }


// FromHex creates a new sha from its hex representation string
func FromHex(s string) (*Sha1, os.Error ) {
	if len(s) != 40 {
		return nil, InvalidLengthError(len(s));
	}
	sha := new(Sha1);
	_, err := hex.Decode(sha, strings.Bytes(s));
	if err != nil {
		return nil, err;
	}
	return sha, nil;
}
