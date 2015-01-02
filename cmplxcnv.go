package cmplxcnv

import (
	"errors"
	//"fmt"
	//"regexp"
	"strconv"
	"strings"
)

// isImagString returns true if s is a properly formed represenation of an imaginary number
func isImagStr(s string) bool {
	return strings.Contains(s, "i") && s[len(s)-1] == 'i'
}

// imagStr returns the floating point portion of an imaginary string
// This should not be called unless isImagStr returns true
func imagStr(s string) string {
	return s[:len(s)-1]
}

// splitComplex takes a string represenation of a complex number, and
// splits it correctly into real and imaginary parts.
func splitComplex(s string) (rstr string, istr string, err error) {
	/*eReg := regexp.MustCompile("[eE]\\+")    // for searching for an 'e'  next to a '+'
	eSplit := regexp.MustCompile("[^eE]\\+") // for splitting without throwing away 'e'
	split := regexp.MustCompile("\\+")       // for splitting on the +
	var sarray []string
	fmt.Println(s)

	if eReg.FindString(s) != "" {
		sarray = eSplit.Split(s, 2)
	} else {
		sarray = split.Split(s, 2)
	}
	fmt.Println(sarray)*/

	// sarray := make([]string, 2)  // slice of strings to hold parsed substrings

	s1, s2 := s, "" // substrings vars, unknown real/img order
	for i := range s {
		if s[i] == '+' && s[i-1] != 'e' && s[i-1] != 'E' { // ignore '+' in exponents
			if len(s) == i { // '+' is at the end of the string, error
				err = errors.New("malformed complex number")
				return "", "", err
			} else {
				s1, s2 = s[:i], s[i+1:] // remove '+' and return substrings
			}
		}
	}

	//fmt.Println(s1, s2)

	switch { // check substrings, identify imag/real parts, there can be only one imag
	case isImagStr(s1) && !isImagStr(s2): // is the first string imag and the second real
		istr = imagStr(s1)
		rstr = s2
	case isImagStr(s2) && !isImagStr(s1): // is the second string imag and the first real
		istr = imagStr(s2)
		rstr = s1
	case isImagStr(s1): // can be only one well-formed string, is it imag
		istr = imagStr(s1)
	case !isImagStr(s1): // if not, it must be real
		rstr = s
	default: // all other cases indicate bad input
		err = errors.New("malformed complex number")
		return "", "", err
	}

	/* if sarray[1] != "" {
		if s := sarray[0]; !strings.Contains(s, "i") {
			rstr = s
		} else {
			err = errors.New("malformed complex number")
			return
		}
		if s := sarray[1]; strings.Contains(s, "i") {
			istr = s
		} else {
			err = errors.New("malformed complex number")
			return
		}
	} else if s := sarray[0]; strings.Contains(s, "i") {
		istr = s
	} else {
		rstr = s
	}

	fmt.Println(rstr, istr)

	if strings.Contains(istr, "i") {
		if istr[len(istr)-1] == 'i' {
			istr = istr[:len(istr)-1]
		}
	} else {
		//err = errors.New("no imaginary component")
	}*/

	return rstr, istr, nil
}

// ParseComplex takes a string representation of a complex number and
// returns the complex number.
func ParseComplex(s string, bitSize int) (c complex128, err error) {
	r, i := 0.0, 0.0
	rstr, istr, err := splitComplex(s) // first split the input string into parts
	if err != nil {
		return c, err
	}

	//fmt.Printf("real part: %v, img part: %v\n", rstr, istr)

	//if rstr != "" {
	r, err = strconv.ParseFloat(rstr, bitSize)
	if err != nil && rstr != "" { // ignore ParseFloat error from "", 0.0 value is correct
		//fmt.Println(err)
		//fmt.Println("real part error")
		return c, err
	}
	//}

	//if istr != "" {
	i, err = strconv.ParseFloat(istr, bitSize)
	if err != nil && istr != "" { // ignore error from "", 0.0 value is correct
		//fmt.Println(err)
		//fmt.Println("img part error")
		return c, err
	}
	//}
	//fmt.Println(r, i)
	//c, err = complex(r, i), nil

	return complex(r, i), nil
}
