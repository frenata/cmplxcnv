ParseComplex, given a string representation of a complex number, returns a complex number.

- should handle exponents in both imaginary and real portions
- given only a real representation, will return a complex number with 0i as the imag portion
- currently only parses numbers with at most 1 real portion and at most 1 imag portion, seperated by a '+' sign
- bitSize is required for now to match ParseFloat from strconv package
