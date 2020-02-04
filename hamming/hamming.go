package hamming

import "errors"
//This is a function to calculate the hamming distance between 2 DNA strands. Hamming distance is the difference in terms of amino acids in DNA strands. Can read more about hamming distance here: https://en.wikipedia.org/wiki/Hamming_distance
func Distance(a, b string) (int, error) {
	//length of the DNA strands must be same to calculate hamming distance.
	if len(a) != len(b) {
		return -1, errors.New("length of string A and B should be the same to calculate hamming distance")
	}
	hammingDistance := 0
	for i :=0; i<len(a) ;i++  {
		if a[i] != b[i] {
			hammingDistance++
		}
	}
	return hammingDistance, nil
}
