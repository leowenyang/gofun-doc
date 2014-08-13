package hex4all

/**
 * @brief Asc2Hex
 * ASCII sting to Hex string
 *
 * @param srcASC : ASCII string
 *
 * @return dstHex : Hex string
 */

func Asc2Hex(srcASC [] byte) dstHex [] byte {
	for key, value := range(srcASC) {
		srcASC[key] = value - 0x30
	}
}

/**@brief Asc2Num
 * ASCII to Number
 *
 * @param srcASC : ASCII
 *
 * @return : Number
 */
func Asc2Num (srcASC byte) byte {
	// '0' - '9'
	if (srcASC >= 0x30) && (srcASC <= 0x39){
		return srcASC - 0x30
	} 

	// 'a' - 'f' 
	if (srcASC >= 0x61) && (srcASC <= 0x66){
		return srcASC - 0x57
	} 

	// 'A' - 'F'
	if (srcASC >= 0x41) && (srcASC <= 0x46){
		return srcASC - 0x37
	} 
}

/**
 * 
 */
func DAscTo1Num(srcASC byte, srcASC2 byte) byte {
	return (Asc2Num(srcASC2) << 4) | Asc2Num(srcASC1)
}