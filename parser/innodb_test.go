package parser_test

import (
	"io/ioutil"
	"testing"

	"github.com/mylxsw/asteria/log"
	"github.com/mylxsw/mysql-guard/parser"
)

func TestParseInnoDBDeadlocks(t *testing.T) {
	data, err := ioutil.ReadFile("./testdata/innodb_status.txt")
	if err != nil {
		panic(err)
	}

	log.WithFields(log.Fields{"deadlock": parser.ParseInnoDBDeadlocks(string(data))}).Debug("dead locks")
}
