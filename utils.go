package gokml

func boolToInt(b bool) int {
	if !b {
		return 0
	}
	return 1
}
