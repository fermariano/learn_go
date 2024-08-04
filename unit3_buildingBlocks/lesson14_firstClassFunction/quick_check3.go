package main

/*
Quick check 14.3 Rewrite the following function signature to use a function type:
func drawTable(rows int, getRow func(row int) (string, string))
*/

func drawTable(rows int, getRow func(row int) (string, string))

type getRow func(rows int) (string, string)

func drawTableFn(rows int, g getRow)
