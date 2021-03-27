//convert dna to rna
package strand

//convert DNA (gcta) to RNA (cgau) and return RNA string
func ToRNA(dna string) string {
	rna := ""

	for i := 0; i < len(dna); i++ {
		switch dna[i] {
		case 'C':
			rna += "G"
		case 'G':
			rna += "C"
		case 'A':
			rna += "U"
		case 'T':
			rna += "A"
		}
	}

	return rna
}
