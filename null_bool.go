package be

import "database/sql"

type NullBool struct {
	value *sql.NullBool
}

func beNullBool(v any) {

}
