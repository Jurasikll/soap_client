package bpm

type BPM_actions struct {
	Soap_actions map[string]Soap_actions `toml:"soap_actions"`
}
type Soap_actions struct {
	Act      string
	Body_ptr string
	Headers  [][]string
}

func (sa *Soap_actions) Def_headers_map() map[string]string {

	headers := map[string]string{}
	header_count := len(sa.Headers[0])
	for i := 0; i < header_count; i++ {
		headers[sa.Headers[0][i]] = sa.Headers[1][i]

	}
	return headers

}
