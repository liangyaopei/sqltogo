package sql_to_go

import (
	"fmt"
	"strings"

	"github.com/xwb1989/sqlparser"
)

// SqlToGo converts a sql create statement to Go struct
// sqlStmt for sql create statement, outputPkg for output directory
func SqlToGo(sqlStmt string, outputPkg string) (string, error) {
	statement, err := sqlparser.ParseStrictDDL(sqlStmt)
	if err != nil {
		return "", err
	}
	stmt, ok := statement.(*sqlparser.DDL)
	if !ok {
		return "", fmt.Errorf("input sql is not a create statment")
	}
	// convert to Go struct
	tableName := stmt.NewName.Name.String()
	res, err := stmtToGo(stmt, tableName, outputPkg)
	if err != nil {
		return "", err
	}
	return res, nil
}

func stmtToGo(stmt *sqlparser.DDL, tableName string, pkgName string) (string, error) {
	builder := strings.Builder{}

	header := fmt.Sprintf("package %s\n", pkgName)
	// import time package
	headerPkg := "import (\n" +
		"\t\"time\"\n" +
		")\n\n"
	importTime := false

	structName := snakeCaseToCamel(tableName)
	structStart := fmt.Sprintf("type %s struct { \n", structName)
	builder.WriteString(structStart)
	for _, col := range stmt.TableSpec.Columns {
		columnType := col.Type.Type
		if col.Type.Unsigned {
			columnType += " unsigned"
		}

		goType := sqlTypeMap[columnType]
		if goType == "time.Time" {
			importTime = true
		}

		field := snakeCaseToCamel(col.Name.String())
		comment := col.Type.Comment
		if comment == nil {
			builder.WriteString(fmt.Sprintf("\t%s\t%s\t\n", field, goType))
		} else {
			builder.WriteString(fmt.Sprintf("\t%s\t%s\t`comment:\"%s\"` \n",
				field, goType, string(comment.Val)))
		}
	}
	builder.WriteString("}\n")

	if importTime {
		return header + headerPkg + builder.String(), nil
	}
	return header + builder.String(), nil
}

// In sql, table name often is snake_case
// In Go, struct name often is camel case
func snakeCaseToCamel(str string) string {
	builder := strings.Builder{}
	index := 0
	if str[0] >= 'a' && str[0] <= 'z' {
		builder.WriteByte(str[0] - ('a' - 'A'))
		index = 1
	}
	for i := index; i < len(str); i++ {
		if str[i] == '_' && i+1 < len(str) {
			if str[i+1] >= 'a' && str[i+1] <= 'z' {
				builder.WriteByte(str[i+1] - ('a' - 'A'))
				i++
				continue
			}
		}
		builder.WriteByte(str[i])
	}
	return builder.String()
}
