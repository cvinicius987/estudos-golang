package main

import "testing"

/*
- Na prática:
    - Arquivo: _test.go
    - BET: Testes, Exemplos e...
    - func BenchmarkFunc (b *testing.B) { ... b.N ... }
    - go test -bench . ← todos
    - go test -bench Func ← somente Func
- go help testflag
*/

func BenchmarkDivide(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Divide(50, 2)
	}
}
