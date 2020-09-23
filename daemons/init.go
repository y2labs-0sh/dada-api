package daemons

var (
	daemons = make(map[string]Daemon)
)

func Get(name string) (Daemon, bool) {
	d, ok := daemons[name]
	if !ok {
		return nil, false
	}
	return d, true
}
