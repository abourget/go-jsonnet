// Generated by: main
// TypeWriter: stringer
// Directive: +gen on astCompKind

package jsonnet

import (
	"fmt"
)

const _astCompKind_name = "astCompForastCompIf"

var _astCompKind_index = [...]uint8{0, 10, 19}

func (i CompKind) String() string {
	if i < 0 || i+1 >= CompKind(len(_astCompKind_index)) {
		return fmt.Sprintf("astCompKind(%d)", i)
	}
	return _astCompKind_name[_astCompKind_index[i]:_astCompKind_index[i+1]]
}
