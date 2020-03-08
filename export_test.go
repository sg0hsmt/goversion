package goversion

// ResetCache clear internal cache for testing.
func ResetCache() {
	mu.Lock()
	cache = nil
	mu.Unlock()
}

// SetExecCmd replace internal functions for testing.
func SetExecCmd(f func() ([]byte, error)) {
	execCmd = f
}
