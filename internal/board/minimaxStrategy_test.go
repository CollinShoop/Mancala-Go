package board

import "testing"

func BenchmarkBoard_GetMiniMaxMove_d4(b *testing.B) {
	GetMiniMaxMove_dn(b.N, 4, false)
}
func BenchmarkBoard_GetMiniMaxMove_d5(b *testing.B) {
	GetMiniMaxMove_dn(b.N, 5, false)
}
func BenchmarkBoard_GetMiniMaxMove_d6(b *testing.B) {
	GetMiniMaxMove_dn(b.N, 6, false)
}
func BenchmarkBoard_GetMiniMaxMove_d7(b *testing.B) {
	GetMiniMaxMove_dn(b.N, 7, false)
}
func BenchmarkBoard_GetMiniMaxMove_d4_mt(b *testing.B) {
	GetMiniMaxMove_dn(b.N, 4, true)
}
func BenchmarkBoard_GetMiniMaxMove_d5_mt(b *testing.B) {
	GetMiniMaxMove_dn(b.N, 5, true)
}
func BenchmarkBoard_GetMiniMaxMove_d6_mt(b *testing.B) {
	GetMiniMaxMove_dn(b.N, 6, true)
}
func BenchmarkBoard_GetMiniMaxMove_d7_mt(b *testing.B) {
	GetMiniMaxMove_dn(b.N, 7, true)
}

func GetMiniMaxMove_dn(N, depth int, multithreaded bool) {
	for n := 0; n < N; n++ {
		b := GetNewBoard()
		switch multithreaded {
		case true:
			minimax_mt(b, 0, 0, depth)
		case false:
			minimax(b, 0, 0, depth)
		}
	}
}
