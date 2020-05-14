package audacity

import "encoding/json"

// Response is an output response from Audacity sent through the pipe from commands like GetInfo.
type Response struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	URL    string  `json:"url"`
	Params []Param `json:"params"`
	Tip    string  `json:"tip"`
}

// Param is a parameter in a response.
type Param struct {
	Key     string   `json:"key"`
	Type    string   `json:"type"`
	Default string   `json:"default"`
	Enum    []string `json:"enum"`
}

func (r Response) String() string {
	o, _ := json.Marshal(r)
	return string(o)
	// ps := ""
	// k := len(r.Params) - 1
	// for i, p := range r.Params {
	// 	if k == i {
	// 		ps += fmt.Sprintf(" %s ", p)
	// 	} else {
	// 		ps += fmt.Sprintf(" %s,", p)
	// 	}
	// }

	// return fmt.Sprintf(`{"id": "%s", "name": "%s", "url": "%s", "params": [%s], "tip": "%s"}`, r.ID, r.Name, r.URL, ps, r.Tip)
}

// func (p param) String() string {
// 	s := fmt.Sprintf(`{ "Key": "%s", "Type": "%s", "Default": "%s"`, p.Key, p.Type, p.Default)

// 	if k := len(p.Enum) - 1; k > -1 {

// 		es := ""
// 		for i, e := range p.Enum {
// 			if k == i {
// 				es += fmt.Sprintf(` "%s" `, e)
// 			} else {
// 				es += fmt.Sprintf(` "%s",`, e)
// 			}
// 		}

// 		s += fmt.Sprintf(`, "Enum": [%s]`, es)
// 	}
// 	return s + " }"
// }
