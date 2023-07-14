package template

import "fmt"

// for version only,do not change it

func Version(bitVer, cliVer, buildDate string) string {
	return fmt.Sprintf(`BitSrunLoginGo by Mmx233 ver.%s
BitSrunLoginGo-Cli by Leave_Time ver.%s
Last Build date:%s
`, bitVer, cliVer, buildDate)
}
