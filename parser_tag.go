package vrule

import (
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/text/gregex"
)

// parseTagValue parses one sequence tag to field, rule and error message.
// The sequence tag is like: [alias@]rule[...#msg...]
func parseTagValue(tag string) (field, rule, msg string) {
	// Complete sequence tag.
	// Example: name@required|length:2,20|password3|same:password1#||密码强度不足|两次密码不一致
	match, _ := gregex.MatchString(`\s*((\w+)\s*@){0,1}\s*([^#]+)\s*(#\s*(.*)){0,1}\s*`, tag)
	if len(match) > 5 {
		msg = strings.TrimSpace(match[5])
		rule = strings.TrimSpace(match[3])
		field = strings.TrimSpace(match[2])
	} else {
		panic(fmt.Errorf(`invalid validation tag value: %s`, tag))
	}
	return
}
