package throt

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

func success(success bool) string {
	if success {
		return succeed
	}
	return failed
}
