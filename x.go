package crypt4ruts

func roR(b byte) byte {
	if (b % 2) == 0 {
		return b >> 1
	}

	return (b >> 1) + 128
}

func roL(b byte) byte {
	if b < 128 {
		return b << 1
	}

	return (b << 1) + 1
}

func encode8(d, k byte) byte {
	for i := 0; i < 8; i++ {
		if (k & 128) == 0 {
			d = roL(d) + 4
		} else {
			d = roL(d) - 3
		}
		k = k << 1
	}
	return d
}

func decode8(e, k byte) byte {
	for i := 0; i < 8; i++ {
		if (k & 1) == 0 {
			e = roR(e - 4)
		} else {
			e = roR(e + 3)
		}
		k = k >> 1
	}
	return e
}
