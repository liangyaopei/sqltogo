package sql_to_go

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

func TestSqlToGo(t *testing.T) {
	inputFile := "./data/input.sql"
	output := "./data/output.go"
	pkgName := "data"
	// read input data
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		t.Errorf("read file err:%s", err.Error())
		return
	}
	sqlStmt := string(data)
	res, err := SqlToGo(sqlStmt, pkgName)
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
