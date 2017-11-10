package main

var zndkPattern = [5]func(string) bool{
	CreateZndkValidFunc("zun"),
	CreateZndkValidFunc("doko"),
	CreateZndkValidFunc("ki"),
	CreateZndkValidFunc("yo"),
	CreateZndkValidFunc("shi"),
}

// ValidProofHash is a function to check the proof is valid
func ValidProofHash(proofHashes [8]ProofHash) (bool, int) {

	zndk := [8]func(string) bool{
		zndkPattern[0],
		zndkPattern[0],
		zndkPattern[0],
		zndkPattern[0],
		zndkPattern[1],
		zndkPattern[2],
		zndkPattern[3],
		zndkPattern[4],
	}

	i := 0
	for i < 8 {
		if zndk[i](proofHashes[7].Encode()) {
			break
		} else if i == 0 { // if index is 0 but not matched, go to 4
			i = 4
			continue
		}
		i++
	}
	if i == 8 {
		return false, -1
	}

	// Search Shi and Zun
	j := -1

	if zndk[0](proofHashes[0].Encode()) && zndk[7](proofHashes[7].Encode()) {
		j = 0
	} else {
		for i = 0; i < 7; i++ {
			if zndk[7](proofHashes[i].Encode()) &&
				zndk[0](proofHashes[i+1].Encode()) {
				j = i + 1
			}
		}
	}

	if j < 0 {
		return false, -1
	}

	for k := 0; k < 8; k++ {
		if !zndk[k](proofHashes[(j+k)%8].Encode()) {
			return false, -1
		}
	}

	return true, i
}
