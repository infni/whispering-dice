package log

type AddlInfo map[string]interface{}

func (a AddlInfo) ToJson() map[string]interface{} {
	local := a
	j := make(map[string]interface{}, len(local))
	for k, v := range local {
		if addl, ok := v.(AddlInfo); ok {
			newV := addl.ToJson() // <--- RECURSION
			j[k] = newV
		} else {
			j[k] = v
		}
	}
	return j
}
