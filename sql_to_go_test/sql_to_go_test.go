package sql_to_go_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/liangyaopei/sqltogo"
)

func TestSqlToGo(t *testing.T) {
	inputFile := "./input.sql"
	output := "./_output.go"
	pkgName := "sql_to_go_test"
	// read input data
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		t.Errorf("read file err:%s", err.Error())
		return
	}
	sqlStmt := string(data)
	res, err := sqltogo.SqlToGo(sqlStmt, pkgName)
	if err != nil {
		t.Errorf("read file err:%s", err.Error())
		return
	}
	// print result
	t.Logf("\n%s", res)
	// save result to a file
	f, err := os.Create(output)
	if err != nil {
		fmt.Printf("create file err:%s", err.Error())
		return
	}
	_, err = f.WriteString(res)
	if err != nil {
		fmt.Printf("write file:%s", err.Error())
		return
	}
	// gofmt the output
	cmd := exec.Command("gofmt", "-w", output)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("go fmt:%s", err.Error())
		return
	}
}
