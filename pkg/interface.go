package pkg

// Interface2Map 将嵌套的map[string]interface全部转换成一层
func Interface2Map(data interface{}) (result map[string]interface{}) {
	result = make(map[string]interface{})
	m, ok := data.(map[string]interface{})
	if !ok {
		return
	}
	for k, iface := range m {
		switch v := iface.(type) {
		case map[string]interface{}:
			for i, u := range v {
				result[i] = u
			}
		default:
			result[k] = v
		}
	}
	return
}
