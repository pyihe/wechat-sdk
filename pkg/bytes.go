package pkg

func BytesContain(ele byte, b []byte) bool {
	for _, v := range b {
		if v == ele {
			return true
		}
	}
	return false
}
