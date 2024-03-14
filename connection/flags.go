package connection

import "strings"

func ParseFlags(args []string) map[string]string {
	flags := make(map[string]string)
	for i, arg := range args {
		if strings.HasPrefix(arg, "--") {
			if strings.Contains(arg, "=") {
				flag := strings.Split(arg, "=")
				flags[flag[0][2:]] = flag[1]
			} else {
				flags[args[i][2:]] = args[i+1]
			}
		}
	}

	return flags
}
