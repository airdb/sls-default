package cmd

import (
	"github.com/MakeNowJust/heredoc"
)

var createShellExample = heredoc.Doc(`
$ go run cmd/cli/main.go  shell  -c ls --comment "xx" -s "ls -'altr'" -t "ls"

$ go run cmd/cli/main.go  shell  -c ls --comment "xx" -s "\"ls -'altr'\"" -t "ls"
	`)

var createDailyEnglishExample = heredoc.Doc(`
$ go run cmd/cli/main.go  shell  -c ls --comment "xx" -s "ls -'altr'" -t "ls"

$ go run cmd/cli/main.go  shell  -c ls --comment "xx" -s "\"ls -'altr'\"" -t "ls"
	`)
