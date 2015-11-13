package go_map_vs_switch

import (
	"math/rand"
	"os"
	"testing"
)

var randInputs []int
var ascInputs []int

func TestMain(m *testing.M) {
	for i := 0; i < 4096; i++ {
		randInputs = append(randInputs, rand.Int())
	}

	for i := 0; i < 4096; i++ {
		ascInputs = append(ascInputs, i)
	}

	os.Exit(m.Run())
}

func BenchmarkOptimalCacheFriendlinessSwitchMinimalFunc4(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch i % 4 {
		case 0:
			n += Minimal0(i)
		case 1:
			n += Minimal1(i)
		case 2:
			n += Minimal2(i)
		case 3:
			n += Minimal3(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessMapMinimalFunc4(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += MinimalFuncs[i%4](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessSwitchMinimalFunc4(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch ascInputs[i%len(ascInputs)] % 4 {
		case 0:
			n += Minimal0(i)
		case 1:
			n += Minimal1(i)
		case 2:
			n += Minimal2(i)
		case 3:
			n += Minimal3(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessMapMinimalFunc4(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += MinimalFuncs[ascInputs[i%len(ascInputs)]%4](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessSwitchMinimalFunc4(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch randInputs[i%len(randInputs)] % 4 {
		case 0:
			n += Minimal0(i)
		case 1:
			n += Minimal1(i)
		case 2:
			n += Minimal2(i)
		case 3:
			n += Minimal3(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessMapMinimalFunc4(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += MinimalFuncs[randInputs[i%len(randInputs)]%4](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessSwitchNoInlineFunc4(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch i % 4 {
		case 0:
			n += NoInline0(i)
		case 1:
			n += NoInline1(i)
		case 2:
			n += NoInline2(i)
		case 3:
			n += NoInline3(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessMapNoInlineFunc4(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += NoInlineFuncs[i%4](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessSwitchNoInlineFunc4(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch ascInputs[i%len(ascInputs)] % 4 {
		case 0:
			n += NoInline0(i)
		case 1:
			n += NoInline1(i)
		case 2:
			n += NoInline2(i)
		case 3:
			n += NoInline3(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessMapNoInlineFunc4(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += NoInlineFuncs[ascInputs[i%len(ascInputs)]%4](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessSwitchNoInlineFunc4(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch randInputs[i%len(randInputs)] % 4 {
		case 0:
			n += NoInline0(i)
		case 1:
			n += NoInline1(i)
		case 2:
			n += NoInline2(i)
		case 3:
			n += NoInline3(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessMapNoInlineFunc4(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += NoInlineFuncs[randInputs[i%len(randInputs)]%4](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessSwitchMinimalFunc8(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch i % 8 {
		case 0:
			n += Minimal0(i)
		case 1:
			n += Minimal1(i)
		case 2:
			n += Minimal2(i)
		case 3:
			n += Minimal3(i)
		case 4:
			n += Minimal4(i)
		case 5:
			n += Minimal5(i)
		case 6:
			n += Minimal6(i)
		case 7:
			n += Minimal7(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessMapMinimalFunc8(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += MinimalFuncs[i%8](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessSwitchMinimalFunc8(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch ascInputs[i%len(ascInputs)] % 8 {
		case 0:
			n += Minimal0(i)
		case 1:
			n += Minimal1(i)
		case 2:
			n += Minimal2(i)
		case 3:
			n += Minimal3(i)
		case 4:
			n += Minimal4(i)
		case 5:
			n += Minimal5(i)
		case 6:
			n += Minimal6(i)
		case 7:
			n += Minimal7(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessMapMinimalFunc8(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += MinimalFuncs[ascInputs[i%len(ascInputs)]%8](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessSwitchMinimalFunc8(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch randInputs[i%len(randInputs)] % 8 {
		case 0:
			n += Minimal0(i)
		case 1:
			n += Minimal1(i)
		case 2:
			n += Minimal2(i)
		case 3:
			n += Minimal3(i)
		case 4:
			n += Minimal4(i)
		case 5:
			n += Minimal5(i)
		case 6:
			n += Minimal6(i)
		case 7:
			n += Minimal7(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessMapMinimalFunc8(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += MinimalFuncs[randInputs[i%len(randInputs)]%8](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessSwitchNoInlineFunc8(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch i % 8 {
		case 0:
			n += NoInline0(i)
		case 1:
			n += NoInline1(i)
		case 2:
			n += NoInline2(i)
		case 3:
			n += NoInline3(i)
		case 4:
			n += NoInline4(i)
		case 5:
			n += NoInline5(i)
		case 6:
			n += NoInline6(i)
		case 7:
			n += NoInline7(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessMapNoInlineFunc8(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += NoInlineFuncs[i%8](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessSwitchNoInlineFunc8(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch ascInputs[i%len(ascInputs)] % 8 {
		case 0:
			n += NoInline0(i)
		case 1:
			n += NoInline1(i)
		case 2:
			n += NoInline2(i)
		case 3:
			n += NoInline3(i)
		case 4:
			n += NoInline4(i)
		case 5:
			n += NoInline5(i)
		case 6:
			n += NoInline6(i)
		case 7:
			n += NoInline7(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessMapNoInlineFunc8(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += NoInlineFuncs[ascInputs[i%len(ascInputs)]%8](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessSwitchNoInlineFunc8(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch randInputs[i%len(randInputs)] % 8 {
		case 0:
			n += NoInline0(i)
		case 1:
			n += NoInline1(i)
		case 2:
			n += NoInline2(i)
		case 3:
			n += NoInline3(i)
		case 4:
			n += NoInline4(i)
		case 5:
			n += NoInline5(i)
		case 6:
			n += NoInline6(i)
		case 7:
			n += NoInline7(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessMapNoInlineFunc8(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += NoInlineFuncs[randInputs[i%len(randInputs)]%8](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessSwitchMinimalFunc16(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch i % 16 {
		case 0:
			n += Minimal0(i)
		case 1:
			n += Minimal1(i)
		case 2:
			n += Minimal2(i)
		case 3:
			n += Minimal3(i)
		case 4:
			n += Minimal4(i)
		case 5:
			n += Minimal5(i)
		case 6:
			n += Minimal6(i)
		case 7:
			n += Minimal7(i)
		case 8:
			n += Minimal8(i)
		case 9:
			n += Minimal9(i)
		case 10:
			n += Minimal10(i)
		case 11:
			n += Minimal11(i)
		case 12:
			n += Minimal12(i)
		case 13:
			n += Minimal13(i)
		case 14:
			n += Minimal14(i)
		case 15:
			n += Minimal15(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessMapMinimalFunc16(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += MinimalFuncs[i%16](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessSwitchMinimalFunc16(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch ascInputs[i%len(ascInputs)] % 16 {
		case 0:
			n += Minimal0(i)
		case 1:
			n += Minimal1(i)
		case 2:
			n += Minimal2(i)
		case 3:
			n += Minimal3(i)
		case 4:
			n += Minimal4(i)
		case 5:
			n += Minimal5(i)
		case 6:
			n += Minimal6(i)
		case 7:
			n += Minimal7(i)
		case 8:
			n += Minimal8(i)
		case 9:
			n += Minimal9(i)
		case 10:
			n += Minimal10(i)
		case 11:
			n += Minimal11(i)
		case 12:
			n += Minimal12(i)
		case 13:
			n += Minimal13(i)
		case 14:
			n += Minimal14(i)
		case 15:
			n += Minimal15(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessMapMinimalFunc16(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += MinimalFuncs[ascInputs[i%len(ascInputs)]%16](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessSwitchMinimalFunc16(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch randInputs[i%len(randInputs)] % 16 {
		case 0:
			n += Minimal0(i)
		case 1:
			n += Minimal1(i)
		case 2:
			n += Minimal2(i)
		case 3:
			n += Minimal3(i)
		case 4:
			n += Minimal4(i)
		case 5:
			n += Minimal5(i)
		case 6:
			n += Minimal6(i)
		case 7:
			n += Minimal7(i)
		case 8:
			n += Minimal8(i)
		case 9:
			n += Minimal9(i)
		case 10:
			n += Minimal10(i)
		case 11:
			n += Minimal11(i)
		case 12:
			n += Minimal12(i)
		case 13:
			n += Minimal13(i)
		case 14:
			n += Minimal14(i)
		case 15:
			n += Minimal15(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessMapMinimalFunc16(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += MinimalFuncs[randInputs[i%len(randInputs)]%16](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessSwitchNoInlineFunc16(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch i % 16 {
		case 0:
			n += NoInline0(i)
		case 1:
			n += NoInline1(i)
		case 2:
			n += NoInline2(i)
		case 3:
			n += NoInline3(i)
		case 4:
			n += NoInline4(i)
		case 5:
			n += NoInline5(i)
		case 6:
			n += NoInline6(i)
		case 7:
			n += NoInline7(i)
		case 8:
			n += NoInline8(i)
		case 9:
			n += NoInline9(i)
		case 10:
			n += NoInline10(i)
		case 11:
			n += NoInline11(i)
		case 12:
			n += NoInline12(i)
		case 13:
			n += NoInline13(i)
		case 14:
			n += NoInline14(i)
		case 15:
			n += NoInline15(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessMapNoInlineFunc16(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += NoInlineFuncs[i%16](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessSwitchNoInlineFunc16(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch ascInputs[i%len(ascInputs)] % 16 {
		case 0:
			n += NoInline0(i)
		case 1:
			n += NoInline1(i)
		case 2:
			n += NoInline2(i)
		case 3:
			n += NoInline3(i)
		case 4:
			n += NoInline4(i)
		case 5:
			n += NoInline5(i)
		case 6:
			n += NoInline6(i)
		case 7:
			n += NoInline7(i)
		case 8:
			n += NoInline8(i)
		case 9:
			n += NoInline9(i)
		case 10:
			n += NoInline10(i)
		case 11:
			n += NoInline11(i)
		case 12:
			n += NoInline12(i)
		case 13:
			n += NoInline13(i)
		case 14:
			n += NoInline14(i)
		case 15:
			n += NoInline15(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessMapNoInlineFunc16(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += NoInlineFuncs[ascInputs[i%len(ascInputs)]%16](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessSwitchNoInlineFunc16(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch randInputs[i%len(randInputs)] % 16 {
		case 0:
			n += NoInline0(i)
		case 1:
			n += NoInline1(i)
		case 2:
			n += NoInline2(i)
		case 3:
			n += NoInline3(i)
		case 4:
			n += NoInline4(i)
		case 5:
			n += NoInline5(i)
		case 6:
			n += NoInline6(i)
		case 7:
			n += NoInline7(i)
		case 8:
			n += NoInline8(i)
		case 9:
			n += NoInline9(i)
		case 10:
			n += NoInline10(i)
		case 11:
			n += NoInline11(i)
		case 12:
			n += NoInline12(i)
		case 13:
			n += NoInline13(i)
		case 14:
			n += NoInline14(i)
		case 15:
			n += NoInline15(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessMapNoInlineFunc16(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += NoInlineFuncs[randInputs[i%len(randInputs)]%16](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessSwitchMinimalFunc32(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch i % 32 {
		case 0:
			n += Minimal0(i)
		case 1:
			n += Minimal1(i)
		case 2:
			n += Minimal2(i)
		case 3:
			n += Minimal3(i)
		case 4:
			n += Minimal4(i)
		case 5:
			n += Minimal5(i)
		case 6:
			n += Minimal6(i)
		case 7:
			n += Minimal7(i)
		case 8:
			n += Minimal8(i)
		case 9:
			n += Minimal9(i)
		case 10:
			n += Minimal10(i)
		case 11:
			n += Minimal11(i)
		case 12:
			n += Minimal12(i)
		case 13:
			n += Minimal13(i)
		case 14:
			n += Minimal14(i)
		case 15:
			n += Minimal15(i)
		case 16:
			n += Minimal16(i)
		case 17:
			n += Minimal17(i)
		case 18:
			n += Minimal18(i)
		case 19:
			n += Minimal19(i)
		case 20:
			n += Minimal20(i)
		case 21:
			n += Minimal21(i)
		case 22:
			n += Minimal22(i)
		case 23:
			n += Minimal23(i)
		case 24:
			n += Minimal24(i)
		case 25:
			n += Minimal25(i)
		case 26:
			n += Minimal26(i)
		case 27:
			n += Minimal27(i)
		case 28:
			n += Minimal28(i)
		case 29:
			n += Minimal29(i)
		case 30:
			n += Minimal30(i)
		case 31:
			n += Minimal31(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessMapMinimalFunc32(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += MinimalFuncs[i%32](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessSwitchMinimalFunc32(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch ascInputs[i%len(ascInputs)] % 32 {
		case 0:
			n += Minimal0(i)
		case 1:
			n += Minimal1(i)
		case 2:
			n += Minimal2(i)
		case 3:
			n += Minimal3(i)
		case 4:
			n += Minimal4(i)
		case 5:
			n += Minimal5(i)
		case 6:
			n += Minimal6(i)
		case 7:
			n += Minimal7(i)
		case 8:
			n += Minimal8(i)
		case 9:
			n += Minimal9(i)
		case 10:
			n += Minimal10(i)
		case 11:
			n += Minimal11(i)
		case 12:
			n += Minimal12(i)
		case 13:
			n += Minimal13(i)
		case 14:
			n += Minimal14(i)
		case 15:
			n += Minimal15(i)
		case 16:
			n += Minimal16(i)
		case 17:
			n += Minimal17(i)
		case 18:
			n += Minimal18(i)
		case 19:
			n += Minimal19(i)
		case 20:
			n += Minimal20(i)
		case 21:
			n += Minimal21(i)
		case 22:
			n += Minimal22(i)
		case 23:
			n += Minimal23(i)
		case 24:
			n += Minimal24(i)
		case 25:
			n += Minimal25(i)
		case 26:
			n += Minimal26(i)
		case 27:
			n += Minimal27(i)
		case 28:
			n += Minimal28(i)
		case 29:
			n += Minimal29(i)
		case 30:
			n += Minimal30(i)
		case 31:
			n += Minimal31(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessMapMinimalFunc32(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += MinimalFuncs[ascInputs[i%len(ascInputs)]%32](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessSwitchMinimalFunc32(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch randInputs[i%len(randInputs)] % 32 {
		case 0:
			n += Minimal0(i)
		case 1:
			n += Minimal1(i)
		case 2:
			n += Minimal2(i)
		case 3:
			n += Minimal3(i)
		case 4:
			n += Minimal4(i)
		case 5:
			n += Minimal5(i)
		case 6:
			n += Minimal6(i)
		case 7:
			n += Minimal7(i)
		case 8:
			n += Minimal8(i)
		case 9:
			n += Minimal9(i)
		case 10:
			n += Minimal10(i)
		case 11:
			n += Minimal11(i)
		case 12:
			n += Minimal12(i)
		case 13:
			n += Minimal13(i)
		case 14:
			n += Minimal14(i)
		case 15:
			n += Minimal15(i)
		case 16:
			n += Minimal16(i)
		case 17:
			n += Minimal17(i)
		case 18:
			n += Minimal18(i)
		case 19:
			n += Minimal19(i)
		case 20:
			n += Minimal20(i)
		case 21:
			n += Minimal21(i)
		case 22:
			n += Minimal22(i)
		case 23:
			n += Minimal23(i)
		case 24:
			n += Minimal24(i)
		case 25:
			n += Minimal25(i)
		case 26:
			n += Minimal26(i)
		case 27:
			n += Minimal27(i)
		case 28:
			n += Minimal28(i)
		case 29:
			n += Minimal29(i)
		case 30:
			n += Minimal30(i)
		case 31:
			n += Minimal31(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessMapMinimalFunc32(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += MinimalFuncs[randInputs[i%len(randInputs)]%32](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessSwitchNoInlineFunc32(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch i % 32 {
		case 0:
			n += NoInline0(i)
		case 1:
			n += NoInline1(i)
		case 2:
			n += NoInline2(i)
		case 3:
			n += NoInline3(i)
		case 4:
			n += NoInline4(i)
		case 5:
			n += NoInline5(i)
		case 6:
			n += NoInline6(i)
		case 7:
			n += NoInline7(i)
		case 8:
			n += NoInline8(i)
		case 9:
			n += NoInline9(i)
		case 10:
			n += NoInline10(i)
		case 11:
			n += NoInline11(i)
		case 12:
			n += NoInline12(i)
		case 13:
			n += NoInline13(i)
		case 14:
			n += NoInline14(i)
		case 15:
			n += NoInline15(i)
		case 16:
			n += NoInline16(i)
		case 17:
			n += NoInline17(i)
		case 18:
			n += NoInline18(i)
		case 19:
			n += NoInline19(i)
		case 20:
			n += NoInline20(i)
		case 21:
			n += NoInline21(i)
		case 22:
			n += NoInline22(i)
		case 23:
			n += NoInline23(i)
		case 24:
			n += NoInline24(i)
		case 25:
			n += NoInline25(i)
		case 26:
			n += NoInline26(i)
		case 27:
			n += NoInline27(i)
		case 28:
			n += NoInline28(i)
		case 29:
			n += NoInline29(i)
		case 30:
			n += NoInline30(i)
		case 31:
			n += NoInline31(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessMapNoInlineFunc32(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += NoInlineFuncs[i%32](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessSwitchNoInlineFunc32(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch ascInputs[i%len(ascInputs)] % 32 {
		case 0:
			n += NoInline0(i)
		case 1:
			n += NoInline1(i)
		case 2:
			n += NoInline2(i)
		case 3:
			n += NoInline3(i)
		case 4:
			n += NoInline4(i)
		case 5:
			n += NoInline5(i)
		case 6:
			n += NoInline6(i)
		case 7:
			n += NoInline7(i)
		case 8:
			n += NoInline8(i)
		case 9:
			n += NoInline9(i)
		case 10:
			n += NoInline10(i)
		case 11:
			n += NoInline11(i)
		case 12:
			n += NoInline12(i)
		case 13:
			n += NoInline13(i)
		case 14:
			n += NoInline14(i)
		case 15:
			n += NoInline15(i)
		case 16:
			n += NoInline16(i)
		case 17:
			n += NoInline17(i)
		case 18:
			n += NoInline18(i)
		case 19:
			n += NoInline19(i)
		case 20:
			n += NoInline20(i)
		case 21:
			n += NoInline21(i)
		case 22:
			n += NoInline22(i)
		case 23:
			n += NoInline23(i)
		case 24:
			n += NoInline24(i)
		case 25:
			n += NoInline25(i)
		case 26:
			n += NoInline26(i)
		case 27:
			n += NoInline27(i)
		case 28:
			n += NoInline28(i)
		case 29:
			n += NoInline29(i)
		case 30:
			n += NoInline30(i)
		case 31:
			n += NoInline31(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessMapNoInlineFunc32(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += NoInlineFuncs[ascInputs[i%len(ascInputs)]%32](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessSwitchNoInlineFunc32(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch randInputs[i%len(randInputs)] % 32 {
		case 0:
			n += NoInline0(i)
		case 1:
			n += NoInline1(i)
		case 2:
			n += NoInline2(i)
		case 3:
			n += NoInline3(i)
		case 4:
			n += NoInline4(i)
		case 5:
			n += NoInline5(i)
		case 6:
			n += NoInline6(i)
		case 7:
			n += NoInline7(i)
		case 8:
			n += NoInline8(i)
		case 9:
			n += NoInline9(i)
		case 10:
			n += NoInline10(i)
		case 11:
			n += NoInline11(i)
		case 12:
			n += NoInline12(i)
		case 13:
			n += NoInline13(i)
		case 14:
			n += NoInline14(i)
		case 15:
			n += NoInline15(i)
		case 16:
			n += NoInline16(i)
		case 17:
			n += NoInline17(i)
		case 18:
			n += NoInline18(i)
		case 19:
			n += NoInline19(i)
		case 20:
			n += NoInline20(i)
		case 21:
			n += NoInline21(i)
		case 22:
			n += NoInline22(i)
		case 23:
			n += NoInline23(i)
		case 24:
			n += NoInline24(i)
		case 25:
			n += NoInline25(i)
		case 26:
			n += NoInline26(i)
		case 27:
			n += NoInline27(i)
		case 28:
			n += NoInline28(i)
		case 29:
			n += NoInline29(i)
		case 30:
			n += NoInline30(i)
		case 31:
			n += NoInline31(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessMapNoInlineFunc32(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += NoInlineFuncs[randInputs[i%len(randInputs)]%32](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessSwitchMinimalFunc64(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch i % 64 {
		case 0:
			n += Minimal0(i)
		case 1:
			n += Minimal1(i)
		case 2:
			n += Minimal2(i)
		case 3:
			n += Minimal3(i)
		case 4:
			n += Minimal4(i)
		case 5:
			n += Minimal5(i)
		case 6:
			n += Minimal6(i)
		case 7:
			n += Minimal7(i)
		case 8:
			n += Minimal8(i)
		case 9:
			n += Minimal9(i)
		case 10:
			n += Minimal10(i)
		case 11:
			n += Minimal11(i)
		case 12:
			n += Minimal12(i)
		case 13:
			n += Minimal13(i)
		case 14:
			n += Minimal14(i)
		case 15:
			n += Minimal15(i)
		case 16:
			n += Minimal16(i)
		case 17:
			n += Minimal17(i)
		case 18:
			n += Minimal18(i)
		case 19:
			n += Minimal19(i)
		case 20:
			n += Minimal20(i)
		case 21:
			n += Minimal21(i)
		case 22:
			n += Minimal22(i)
		case 23:
			n += Minimal23(i)
		case 24:
			n += Minimal24(i)
		case 25:
			n += Minimal25(i)
		case 26:
			n += Minimal26(i)
		case 27:
			n += Minimal27(i)
		case 28:
			n += Minimal28(i)
		case 29:
			n += Minimal29(i)
		case 30:
			n += Minimal30(i)
		case 31:
			n += Minimal31(i)
		case 32:
			n += Minimal32(i)
		case 33:
			n += Minimal33(i)
		case 34:
			n += Minimal34(i)
		case 35:
			n += Minimal35(i)
		case 36:
			n += Minimal36(i)
		case 37:
			n += Minimal37(i)
		case 38:
			n += Minimal38(i)
		case 39:
			n += Minimal39(i)
		case 40:
			n += Minimal40(i)
		case 41:
			n += Minimal41(i)
		case 42:
			n += Minimal42(i)
		case 43:
			n += Minimal43(i)
		case 44:
			n += Minimal44(i)
		case 45:
			n += Minimal45(i)
		case 46:
			n += Minimal46(i)
		case 47:
			n += Minimal47(i)
		case 48:
			n += Minimal48(i)
		case 49:
			n += Minimal49(i)
		case 50:
			n += Minimal50(i)
		case 51:
			n += Minimal51(i)
		case 52:
			n += Minimal52(i)
		case 53:
			n += Minimal53(i)
		case 54:
			n += Minimal54(i)
		case 55:
			n += Minimal55(i)
		case 56:
			n += Minimal56(i)
		case 57:
			n += Minimal57(i)
		case 58:
			n += Minimal58(i)
		case 59:
			n += Minimal59(i)
		case 60:
			n += Minimal60(i)
		case 61:
			n += Minimal61(i)
		case 62:
			n += Minimal62(i)
		case 63:
			n += Minimal63(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessMapMinimalFunc64(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += MinimalFuncs[i%64](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessSwitchMinimalFunc64(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch ascInputs[i%len(ascInputs)] % 64 {
		case 0:
			n += Minimal0(i)
		case 1:
			n += Minimal1(i)
		case 2:
			n += Minimal2(i)
		case 3:
			n += Minimal3(i)
		case 4:
			n += Minimal4(i)
		case 5:
			n += Minimal5(i)
		case 6:
			n += Minimal6(i)
		case 7:
			n += Minimal7(i)
		case 8:
			n += Minimal8(i)
		case 9:
			n += Minimal9(i)
		case 10:
			n += Minimal10(i)
		case 11:
			n += Minimal11(i)
		case 12:
			n += Minimal12(i)
		case 13:
			n += Minimal13(i)
		case 14:
			n += Minimal14(i)
		case 15:
			n += Minimal15(i)
		case 16:
			n += Minimal16(i)
		case 17:
			n += Minimal17(i)
		case 18:
			n += Minimal18(i)
		case 19:
			n += Minimal19(i)
		case 20:
			n += Minimal20(i)
		case 21:
			n += Minimal21(i)
		case 22:
			n += Minimal22(i)
		case 23:
			n += Minimal23(i)
		case 24:
			n += Minimal24(i)
		case 25:
			n += Minimal25(i)
		case 26:
			n += Minimal26(i)
		case 27:
			n += Minimal27(i)
		case 28:
			n += Minimal28(i)
		case 29:
			n += Minimal29(i)
		case 30:
			n += Minimal30(i)
		case 31:
			n += Minimal31(i)
		case 32:
			n += Minimal32(i)
		case 33:
			n += Minimal33(i)
		case 34:
			n += Minimal34(i)
		case 35:
			n += Minimal35(i)
		case 36:
			n += Minimal36(i)
		case 37:
			n += Minimal37(i)
		case 38:
			n += Minimal38(i)
		case 39:
			n += Minimal39(i)
		case 40:
			n += Minimal40(i)
		case 41:
			n += Minimal41(i)
		case 42:
			n += Minimal42(i)
		case 43:
			n += Minimal43(i)
		case 44:
			n += Minimal44(i)
		case 45:
			n += Minimal45(i)
		case 46:
			n += Minimal46(i)
		case 47:
			n += Minimal47(i)
		case 48:
			n += Minimal48(i)
		case 49:
			n += Minimal49(i)
		case 50:
			n += Minimal50(i)
		case 51:
			n += Minimal51(i)
		case 52:
			n += Minimal52(i)
		case 53:
			n += Minimal53(i)
		case 54:
			n += Minimal54(i)
		case 55:
			n += Minimal55(i)
		case 56:
			n += Minimal56(i)
		case 57:
			n += Minimal57(i)
		case 58:
			n += Minimal58(i)
		case 59:
			n += Minimal59(i)
		case 60:
			n += Minimal60(i)
		case 61:
			n += Minimal61(i)
		case 62:
			n += Minimal62(i)
		case 63:
			n += Minimal63(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessMapMinimalFunc64(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += MinimalFuncs[ascInputs[i%len(ascInputs)]%64](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessSwitchMinimalFunc64(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch randInputs[i%len(randInputs)] % 64 {
		case 0:
			n += Minimal0(i)
		case 1:
			n += Minimal1(i)
		case 2:
			n += Minimal2(i)
		case 3:
			n += Minimal3(i)
		case 4:
			n += Minimal4(i)
		case 5:
			n += Minimal5(i)
		case 6:
			n += Minimal6(i)
		case 7:
			n += Minimal7(i)
		case 8:
			n += Minimal8(i)
		case 9:
			n += Minimal9(i)
		case 10:
			n += Minimal10(i)
		case 11:
			n += Minimal11(i)
		case 12:
			n += Minimal12(i)
		case 13:
			n += Minimal13(i)
		case 14:
			n += Minimal14(i)
		case 15:
			n += Minimal15(i)
		case 16:
			n += Minimal16(i)
		case 17:
			n += Minimal17(i)
		case 18:
			n += Minimal18(i)
		case 19:
			n += Minimal19(i)
		case 20:
			n += Minimal20(i)
		case 21:
			n += Minimal21(i)
		case 22:
			n += Minimal22(i)
		case 23:
			n += Minimal23(i)
		case 24:
			n += Minimal24(i)
		case 25:
			n += Minimal25(i)
		case 26:
			n += Minimal26(i)
		case 27:
			n += Minimal27(i)
		case 28:
			n += Minimal28(i)
		case 29:
			n += Minimal29(i)
		case 30:
			n += Minimal30(i)
		case 31:
			n += Minimal31(i)
		case 32:
			n += Minimal32(i)
		case 33:
			n += Minimal33(i)
		case 34:
			n += Minimal34(i)
		case 35:
			n += Minimal35(i)
		case 36:
			n += Minimal36(i)
		case 37:
			n += Minimal37(i)
		case 38:
			n += Minimal38(i)
		case 39:
			n += Minimal39(i)
		case 40:
			n += Minimal40(i)
		case 41:
			n += Minimal41(i)
		case 42:
			n += Minimal42(i)
		case 43:
			n += Minimal43(i)
		case 44:
			n += Minimal44(i)
		case 45:
			n += Minimal45(i)
		case 46:
			n += Minimal46(i)
		case 47:
			n += Minimal47(i)
		case 48:
			n += Minimal48(i)
		case 49:
			n += Minimal49(i)
		case 50:
			n += Minimal50(i)
		case 51:
			n += Minimal51(i)
		case 52:
			n += Minimal52(i)
		case 53:
			n += Minimal53(i)
		case 54:
			n += Minimal54(i)
		case 55:
			n += Minimal55(i)
		case 56:
			n += Minimal56(i)
		case 57:
			n += Minimal57(i)
		case 58:
			n += Minimal58(i)
		case 59:
			n += Minimal59(i)
		case 60:
			n += Minimal60(i)
		case 61:
			n += Minimal61(i)
		case 62:
			n += Minimal62(i)
		case 63:
			n += Minimal63(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessMapMinimalFunc64(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += MinimalFuncs[randInputs[i%len(randInputs)]%64](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessSwitchNoInlineFunc64(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch i % 64 {
		case 0:
			n += NoInline0(i)
		case 1:
			n += NoInline1(i)
		case 2:
			n += NoInline2(i)
		case 3:
			n += NoInline3(i)
		case 4:
			n += NoInline4(i)
		case 5:
			n += NoInline5(i)
		case 6:
			n += NoInline6(i)
		case 7:
			n += NoInline7(i)
		case 8:
			n += NoInline8(i)
		case 9:
			n += NoInline9(i)
		case 10:
			n += NoInline10(i)
		case 11:
			n += NoInline11(i)
		case 12:
			n += NoInline12(i)
		case 13:
			n += NoInline13(i)
		case 14:
			n += NoInline14(i)
		case 15:
			n += NoInline15(i)
		case 16:
			n += NoInline16(i)
		case 17:
			n += NoInline17(i)
		case 18:
			n += NoInline18(i)
		case 19:
			n += NoInline19(i)
		case 20:
			n += NoInline20(i)
		case 21:
			n += NoInline21(i)
		case 22:
			n += NoInline22(i)
		case 23:
			n += NoInline23(i)
		case 24:
			n += NoInline24(i)
		case 25:
			n += NoInline25(i)
		case 26:
			n += NoInline26(i)
		case 27:
			n += NoInline27(i)
		case 28:
			n += NoInline28(i)
		case 29:
			n += NoInline29(i)
		case 30:
			n += NoInline30(i)
		case 31:
			n += NoInline31(i)
		case 32:
			n += NoInline32(i)
		case 33:
			n += NoInline33(i)
		case 34:
			n += NoInline34(i)
		case 35:
			n += NoInline35(i)
		case 36:
			n += NoInline36(i)
		case 37:
			n += NoInline37(i)
		case 38:
			n += NoInline38(i)
		case 39:
			n += NoInline39(i)
		case 40:
			n += NoInline40(i)
		case 41:
			n += NoInline41(i)
		case 42:
			n += NoInline42(i)
		case 43:
			n += NoInline43(i)
		case 44:
			n += NoInline44(i)
		case 45:
			n += NoInline45(i)
		case 46:
			n += NoInline46(i)
		case 47:
			n += NoInline47(i)
		case 48:
			n += NoInline48(i)
		case 49:
			n += NoInline49(i)
		case 50:
			n += NoInline50(i)
		case 51:
			n += NoInline51(i)
		case 52:
			n += NoInline52(i)
		case 53:
			n += NoInline53(i)
		case 54:
			n += NoInline54(i)
		case 55:
			n += NoInline55(i)
		case 56:
			n += NoInline56(i)
		case 57:
			n += NoInline57(i)
		case 58:
			n += NoInline58(i)
		case 59:
			n += NoInline59(i)
		case 60:
			n += NoInline60(i)
		case 61:
			n += NoInline61(i)
		case 62:
			n += NoInline62(i)
		case 63:
			n += NoInline63(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessMapNoInlineFunc64(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += NoInlineFuncs[i%64](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessSwitchNoInlineFunc64(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch ascInputs[i%len(ascInputs)] % 64 {
		case 0:
			n += NoInline0(i)
		case 1:
			n += NoInline1(i)
		case 2:
			n += NoInline2(i)
		case 3:
			n += NoInline3(i)
		case 4:
			n += NoInline4(i)
		case 5:
			n += NoInline5(i)
		case 6:
			n += NoInline6(i)
		case 7:
			n += NoInline7(i)
		case 8:
			n += NoInline8(i)
		case 9:
			n += NoInline9(i)
		case 10:
			n += NoInline10(i)
		case 11:
			n += NoInline11(i)
		case 12:
			n += NoInline12(i)
		case 13:
			n += NoInline13(i)
		case 14:
			n += NoInline14(i)
		case 15:
			n += NoInline15(i)
		case 16:
			n += NoInline16(i)
		case 17:
			n += NoInline17(i)
		case 18:
			n += NoInline18(i)
		case 19:
			n += NoInline19(i)
		case 20:
			n += NoInline20(i)
		case 21:
			n += NoInline21(i)
		case 22:
			n += NoInline22(i)
		case 23:
			n += NoInline23(i)
		case 24:
			n += NoInline24(i)
		case 25:
			n += NoInline25(i)
		case 26:
			n += NoInline26(i)
		case 27:
			n += NoInline27(i)
		case 28:
			n += NoInline28(i)
		case 29:
			n += NoInline29(i)
		case 30:
			n += NoInline30(i)
		case 31:
			n += NoInline31(i)
		case 32:
			n += NoInline32(i)
		case 33:
			n += NoInline33(i)
		case 34:
			n += NoInline34(i)
		case 35:
			n += NoInline35(i)
		case 36:
			n += NoInline36(i)
		case 37:
			n += NoInline37(i)
		case 38:
			n += NoInline38(i)
		case 39:
			n += NoInline39(i)
		case 40:
			n += NoInline40(i)
		case 41:
			n += NoInline41(i)
		case 42:
			n += NoInline42(i)
		case 43:
			n += NoInline43(i)
		case 44:
			n += NoInline44(i)
		case 45:
			n += NoInline45(i)
		case 46:
			n += NoInline46(i)
		case 47:
			n += NoInline47(i)
		case 48:
			n += NoInline48(i)
		case 49:
			n += NoInline49(i)
		case 50:
			n += NoInline50(i)
		case 51:
			n += NoInline51(i)
		case 52:
			n += NoInline52(i)
		case 53:
			n += NoInline53(i)
		case 54:
			n += NoInline54(i)
		case 55:
			n += NoInline55(i)
		case 56:
			n += NoInline56(i)
		case 57:
			n += NoInline57(i)
		case 58:
			n += NoInline58(i)
		case 59:
			n += NoInline59(i)
		case 60:
			n += NoInline60(i)
		case 61:
			n += NoInline61(i)
		case 62:
			n += NoInline62(i)
		case 63:
			n += NoInline63(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessMapNoInlineFunc64(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += NoInlineFuncs[ascInputs[i%len(ascInputs)]%64](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessSwitchNoInlineFunc64(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch randInputs[i%len(randInputs)] % 64 {
		case 0:
			n += NoInline0(i)
		case 1:
			n += NoInline1(i)
		case 2:
			n += NoInline2(i)
		case 3:
			n += NoInline3(i)
		case 4:
			n += NoInline4(i)
		case 5:
			n += NoInline5(i)
		case 6:
			n += NoInline6(i)
		case 7:
			n += NoInline7(i)
		case 8:
			n += NoInline8(i)
		case 9:
			n += NoInline9(i)
		case 10:
			n += NoInline10(i)
		case 11:
			n += NoInline11(i)
		case 12:
			n += NoInline12(i)
		case 13:
			n += NoInline13(i)
		case 14:
			n += NoInline14(i)
		case 15:
			n += NoInline15(i)
		case 16:
			n += NoInline16(i)
		case 17:
			n += NoInline17(i)
		case 18:
			n += NoInline18(i)
		case 19:
			n += NoInline19(i)
		case 20:
			n += NoInline20(i)
		case 21:
			n += NoInline21(i)
		case 22:
			n += NoInline22(i)
		case 23:
			n += NoInline23(i)
		case 24:
			n += NoInline24(i)
		case 25:
			n += NoInline25(i)
		case 26:
			n += NoInline26(i)
		case 27:
			n += NoInline27(i)
		case 28:
			n += NoInline28(i)
		case 29:
			n += NoInline29(i)
		case 30:
			n += NoInline30(i)
		case 31:
			n += NoInline31(i)
		case 32:
			n += NoInline32(i)
		case 33:
			n += NoInline33(i)
		case 34:
			n += NoInline34(i)
		case 35:
			n += NoInline35(i)
		case 36:
			n += NoInline36(i)
		case 37:
			n += NoInline37(i)
		case 38:
			n += NoInline38(i)
		case 39:
			n += NoInline39(i)
		case 40:
			n += NoInline40(i)
		case 41:
			n += NoInline41(i)
		case 42:
			n += NoInline42(i)
		case 43:
			n += NoInline43(i)
		case 44:
			n += NoInline44(i)
		case 45:
			n += NoInline45(i)
		case 46:
			n += NoInline46(i)
		case 47:
			n += NoInline47(i)
		case 48:
			n += NoInline48(i)
		case 49:
			n += NoInline49(i)
		case 50:
			n += NoInline50(i)
		case 51:
			n += NoInline51(i)
		case 52:
			n += NoInline52(i)
		case 53:
			n += NoInline53(i)
		case 54:
			n += NoInline54(i)
		case 55:
			n += NoInline55(i)
		case 56:
			n += NoInline56(i)
		case 57:
			n += NoInline57(i)
		case 58:
			n += NoInline58(i)
		case 59:
			n += NoInline59(i)
		case 60:
			n += NoInline60(i)
		case 61:
			n += NoInline61(i)
		case 62:
			n += NoInline62(i)
		case 63:
			n += NoInline63(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessMapNoInlineFunc64(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += NoInlineFuncs[randInputs[i%len(randInputs)]%64](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessSwitchMinimalFunc128(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch i % 128 {
		case 0:
			n += Minimal0(i)
		case 1:
			n += Minimal1(i)
		case 2:
			n += Minimal2(i)
		case 3:
			n += Minimal3(i)
		case 4:
			n += Minimal4(i)
		case 5:
			n += Minimal5(i)
		case 6:
			n += Minimal6(i)
		case 7:
			n += Minimal7(i)
		case 8:
			n += Minimal8(i)
		case 9:
			n += Minimal9(i)
		case 10:
			n += Minimal10(i)
		case 11:
			n += Minimal11(i)
		case 12:
			n += Minimal12(i)
		case 13:
			n += Minimal13(i)
		case 14:
			n += Minimal14(i)
		case 15:
			n += Minimal15(i)
		case 16:
			n += Minimal16(i)
		case 17:
			n += Minimal17(i)
		case 18:
			n += Minimal18(i)
		case 19:
			n += Minimal19(i)
		case 20:
			n += Minimal20(i)
		case 21:
			n += Minimal21(i)
		case 22:
			n += Minimal22(i)
		case 23:
			n += Minimal23(i)
		case 24:
			n += Minimal24(i)
		case 25:
			n += Minimal25(i)
		case 26:
			n += Minimal26(i)
		case 27:
			n += Minimal27(i)
		case 28:
			n += Minimal28(i)
		case 29:
			n += Minimal29(i)
		case 30:
			n += Minimal30(i)
		case 31:
			n += Minimal31(i)
		case 32:
			n += Minimal32(i)
		case 33:
			n += Minimal33(i)
		case 34:
			n += Minimal34(i)
		case 35:
			n += Minimal35(i)
		case 36:
			n += Minimal36(i)
		case 37:
			n += Minimal37(i)
		case 38:
			n += Minimal38(i)
		case 39:
			n += Minimal39(i)
		case 40:
			n += Minimal40(i)
		case 41:
			n += Minimal41(i)
		case 42:
			n += Minimal42(i)
		case 43:
			n += Minimal43(i)
		case 44:
			n += Minimal44(i)
		case 45:
			n += Minimal45(i)
		case 46:
			n += Minimal46(i)
		case 47:
			n += Minimal47(i)
		case 48:
			n += Minimal48(i)
		case 49:
			n += Minimal49(i)
		case 50:
			n += Minimal50(i)
		case 51:
			n += Minimal51(i)
		case 52:
			n += Minimal52(i)
		case 53:
			n += Minimal53(i)
		case 54:
			n += Minimal54(i)
		case 55:
			n += Minimal55(i)
		case 56:
			n += Minimal56(i)
		case 57:
			n += Minimal57(i)
		case 58:
			n += Minimal58(i)
		case 59:
			n += Minimal59(i)
		case 60:
			n += Minimal60(i)
		case 61:
			n += Minimal61(i)
		case 62:
			n += Minimal62(i)
		case 63:
			n += Minimal63(i)
		case 64:
			n += Minimal64(i)
		case 65:
			n += Minimal65(i)
		case 66:
			n += Minimal66(i)
		case 67:
			n += Minimal67(i)
		case 68:
			n += Minimal68(i)
		case 69:
			n += Minimal69(i)
		case 70:
			n += Minimal70(i)
		case 71:
			n += Minimal71(i)
		case 72:
			n += Minimal72(i)
		case 73:
			n += Minimal73(i)
		case 74:
			n += Minimal74(i)
		case 75:
			n += Minimal75(i)
		case 76:
			n += Minimal76(i)
		case 77:
			n += Minimal77(i)
		case 78:
			n += Minimal78(i)
		case 79:
			n += Minimal79(i)
		case 80:
			n += Minimal80(i)
		case 81:
			n += Minimal81(i)
		case 82:
			n += Minimal82(i)
		case 83:
			n += Minimal83(i)
		case 84:
			n += Minimal84(i)
		case 85:
			n += Minimal85(i)
		case 86:
			n += Minimal86(i)
		case 87:
			n += Minimal87(i)
		case 88:
			n += Minimal88(i)
		case 89:
			n += Minimal89(i)
		case 90:
			n += Minimal90(i)
		case 91:
			n += Minimal91(i)
		case 92:
			n += Minimal92(i)
		case 93:
			n += Minimal93(i)
		case 94:
			n += Minimal94(i)
		case 95:
			n += Minimal95(i)
		case 96:
			n += Minimal96(i)
		case 97:
			n += Minimal97(i)
		case 98:
			n += Minimal98(i)
		case 99:
			n += Minimal99(i)
		case 100:
			n += Minimal100(i)
		case 101:
			n += Minimal101(i)
		case 102:
			n += Minimal102(i)
		case 103:
			n += Minimal103(i)
		case 104:
			n += Minimal104(i)
		case 105:
			n += Minimal105(i)
		case 106:
			n += Minimal106(i)
		case 107:
			n += Minimal107(i)
		case 108:
			n += Minimal108(i)
		case 109:
			n += Minimal109(i)
		case 110:
			n += Minimal110(i)
		case 111:
			n += Minimal111(i)
		case 112:
			n += Minimal112(i)
		case 113:
			n += Minimal113(i)
		case 114:
			n += Minimal114(i)
		case 115:
			n += Minimal115(i)
		case 116:
			n += Minimal116(i)
		case 117:
			n += Minimal117(i)
		case 118:
			n += Minimal118(i)
		case 119:
			n += Minimal119(i)
		case 120:
			n += Minimal120(i)
		case 121:
			n += Minimal121(i)
		case 122:
			n += Minimal122(i)
		case 123:
			n += Minimal123(i)
		case 124:
			n += Minimal124(i)
		case 125:
			n += Minimal125(i)
		case 126:
			n += Minimal126(i)
		case 127:
			n += Minimal127(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessMapMinimalFunc128(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += MinimalFuncs[i%128](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessSwitchMinimalFunc128(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch ascInputs[i%len(ascInputs)] % 128 {
		case 0:
			n += Minimal0(i)
		case 1:
			n += Minimal1(i)
		case 2:
			n += Minimal2(i)
		case 3:
			n += Minimal3(i)
		case 4:
			n += Minimal4(i)
		case 5:
			n += Minimal5(i)
		case 6:
			n += Minimal6(i)
		case 7:
			n += Minimal7(i)
		case 8:
			n += Minimal8(i)
		case 9:
			n += Minimal9(i)
		case 10:
			n += Minimal10(i)
		case 11:
			n += Minimal11(i)
		case 12:
			n += Minimal12(i)
		case 13:
			n += Minimal13(i)
		case 14:
			n += Minimal14(i)
		case 15:
			n += Minimal15(i)
		case 16:
			n += Minimal16(i)
		case 17:
			n += Minimal17(i)
		case 18:
			n += Minimal18(i)
		case 19:
			n += Minimal19(i)
		case 20:
			n += Minimal20(i)
		case 21:
			n += Minimal21(i)
		case 22:
			n += Minimal22(i)
		case 23:
			n += Minimal23(i)
		case 24:
			n += Minimal24(i)
		case 25:
			n += Minimal25(i)
		case 26:
			n += Minimal26(i)
		case 27:
			n += Minimal27(i)
		case 28:
			n += Minimal28(i)
		case 29:
			n += Minimal29(i)
		case 30:
			n += Minimal30(i)
		case 31:
			n += Minimal31(i)
		case 32:
			n += Minimal32(i)
		case 33:
			n += Minimal33(i)
		case 34:
			n += Minimal34(i)
		case 35:
			n += Minimal35(i)
		case 36:
			n += Minimal36(i)
		case 37:
			n += Minimal37(i)
		case 38:
			n += Minimal38(i)
		case 39:
			n += Minimal39(i)
		case 40:
			n += Minimal40(i)
		case 41:
			n += Minimal41(i)
		case 42:
			n += Minimal42(i)
		case 43:
			n += Minimal43(i)
		case 44:
			n += Minimal44(i)
		case 45:
			n += Minimal45(i)
		case 46:
			n += Minimal46(i)
		case 47:
			n += Minimal47(i)
		case 48:
			n += Minimal48(i)
		case 49:
			n += Minimal49(i)
		case 50:
			n += Minimal50(i)
		case 51:
			n += Minimal51(i)
		case 52:
			n += Minimal52(i)
		case 53:
			n += Minimal53(i)
		case 54:
			n += Minimal54(i)
		case 55:
			n += Minimal55(i)
		case 56:
			n += Minimal56(i)
		case 57:
			n += Minimal57(i)
		case 58:
			n += Minimal58(i)
		case 59:
			n += Minimal59(i)
		case 60:
			n += Minimal60(i)
		case 61:
			n += Minimal61(i)
		case 62:
			n += Minimal62(i)
		case 63:
			n += Minimal63(i)
		case 64:
			n += Minimal64(i)
		case 65:
			n += Minimal65(i)
		case 66:
			n += Minimal66(i)
		case 67:
			n += Minimal67(i)
		case 68:
			n += Minimal68(i)
		case 69:
			n += Minimal69(i)
		case 70:
			n += Minimal70(i)
		case 71:
			n += Minimal71(i)
		case 72:
			n += Minimal72(i)
		case 73:
			n += Minimal73(i)
		case 74:
			n += Minimal74(i)
		case 75:
			n += Minimal75(i)
		case 76:
			n += Minimal76(i)
		case 77:
			n += Minimal77(i)
		case 78:
			n += Minimal78(i)
		case 79:
			n += Minimal79(i)
		case 80:
			n += Minimal80(i)
		case 81:
			n += Minimal81(i)
		case 82:
			n += Minimal82(i)
		case 83:
			n += Minimal83(i)
		case 84:
			n += Minimal84(i)
		case 85:
			n += Minimal85(i)
		case 86:
			n += Minimal86(i)
		case 87:
			n += Minimal87(i)
		case 88:
			n += Minimal88(i)
		case 89:
			n += Minimal89(i)
		case 90:
			n += Minimal90(i)
		case 91:
			n += Minimal91(i)
		case 92:
			n += Minimal92(i)
		case 93:
			n += Minimal93(i)
		case 94:
			n += Minimal94(i)
		case 95:
			n += Minimal95(i)
		case 96:
			n += Minimal96(i)
		case 97:
			n += Minimal97(i)
		case 98:
			n += Minimal98(i)
		case 99:
			n += Minimal99(i)
		case 100:
			n += Minimal100(i)
		case 101:
			n += Minimal101(i)
		case 102:
			n += Minimal102(i)
		case 103:
			n += Minimal103(i)
		case 104:
			n += Minimal104(i)
		case 105:
			n += Minimal105(i)
		case 106:
			n += Minimal106(i)
		case 107:
			n += Minimal107(i)
		case 108:
			n += Minimal108(i)
		case 109:
			n += Minimal109(i)
		case 110:
			n += Minimal110(i)
		case 111:
			n += Minimal111(i)
		case 112:
			n += Minimal112(i)
		case 113:
			n += Minimal113(i)
		case 114:
			n += Minimal114(i)
		case 115:
			n += Minimal115(i)
		case 116:
			n += Minimal116(i)
		case 117:
			n += Minimal117(i)
		case 118:
			n += Minimal118(i)
		case 119:
			n += Minimal119(i)
		case 120:
			n += Minimal120(i)
		case 121:
			n += Minimal121(i)
		case 122:
			n += Minimal122(i)
		case 123:
			n += Minimal123(i)
		case 124:
			n += Minimal124(i)
		case 125:
			n += Minimal125(i)
		case 126:
			n += Minimal126(i)
		case 127:
			n += Minimal127(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessMapMinimalFunc128(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += MinimalFuncs[ascInputs[i%len(ascInputs)]%128](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessSwitchMinimalFunc128(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch randInputs[i%len(randInputs)] % 128 {
		case 0:
			n += Minimal0(i)
		case 1:
			n += Minimal1(i)
		case 2:
			n += Minimal2(i)
		case 3:
			n += Minimal3(i)
		case 4:
			n += Minimal4(i)
		case 5:
			n += Minimal5(i)
		case 6:
			n += Minimal6(i)
		case 7:
			n += Minimal7(i)
		case 8:
			n += Minimal8(i)
		case 9:
			n += Minimal9(i)
		case 10:
			n += Minimal10(i)
		case 11:
			n += Minimal11(i)
		case 12:
			n += Minimal12(i)
		case 13:
			n += Minimal13(i)
		case 14:
			n += Minimal14(i)
		case 15:
			n += Minimal15(i)
		case 16:
			n += Minimal16(i)
		case 17:
			n += Minimal17(i)
		case 18:
			n += Minimal18(i)
		case 19:
			n += Minimal19(i)
		case 20:
			n += Minimal20(i)
		case 21:
			n += Minimal21(i)
		case 22:
			n += Minimal22(i)
		case 23:
			n += Minimal23(i)
		case 24:
			n += Minimal24(i)
		case 25:
			n += Minimal25(i)
		case 26:
			n += Minimal26(i)
		case 27:
			n += Minimal27(i)
		case 28:
			n += Minimal28(i)
		case 29:
			n += Minimal29(i)
		case 30:
			n += Minimal30(i)
		case 31:
			n += Minimal31(i)
		case 32:
			n += Minimal32(i)
		case 33:
			n += Minimal33(i)
		case 34:
			n += Minimal34(i)
		case 35:
			n += Minimal35(i)
		case 36:
			n += Minimal36(i)
		case 37:
			n += Minimal37(i)
		case 38:
			n += Minimal38(i)
		case 39:
			n += Minimal39(i)
		case 40:
			n += Minimal40(i)
		case 41:
			n += Minimal41(i)
		case 42:
			n += Minimal42(i)
		case 43:
			n += Minimal43(i)
		case 44:
			n += Minimal44(i)
		case 45:
			n += Minimal45(i)
		case 46:
			n += Minimal46(i)
		case 47:
			n += Minimal47(i)
		case 48:
			n += Minimal48(i)
		case 49:
			n += Minimal49(i)
		case 50:
			n += Minimal50(i)
		case 51:
			n += Minimal51(i)
		case 52:
			n += Minimal52(i)
		case 53:
			n += Minimal53(i)
		case 54:
			n += Minimal54(i)
		case 55:
			n += Minimal55(i)
		case 56:
			n += Minimal56(i)
		case 57:
			n += Minimal57(i)
		case 58:
			n += Minimal58(i)
		case 59:
			n += Minimal59(i)
		case 60:
			n += Minimal60(i)
		case 61:
			n += Minimal61(i)
		case 62:
			n += Minimal62(i)
		case 63:
			n += Minimal63(i)
		case 64:
			n += Minimal64(i)
		case 65:
			n += Minimal65(i)
		case 66:
			n += Minimal66(i)
		case 67:
			n += Minimal67(i)
		case 68:
			n += Minimal68(i)
		case 69:
			n += Minimal69(i)
		case 70:
			n += Minimal70(i)
		case 71:
			n += Minimal71(i)
		case 72:
			n += Minimal72(i)
		case 73:
			n += Minimal73(i)
		case 74:
			n += Minimal74(i)
		case 75:
			n += Minimal75(i)
		case 76:
			n += Minimal76(i)
		case 77:
			n += Minimal77(i)
		case 78:
			n += Minimal78(i)
		case 79:
			n += Minimal79(i)
		case 80:
			n += Minimal80(i)
		case 81:
			n += Minimal81(i)
		case 82:
			n += Minimal82(i)
		case 83:
			n += Minimal83(i)
		case 84:
			n += Minimal84(i)
		case 85:
			n += Minimal85(i)
		case 86:
			n += Minimal86(i)
		case 87:
			n += Minimal87(i)
		case 88:
			n += Minimal88(i)
		case 89:
			n += Minimal89(i)
		case 90:
			n += Minimal90(i)
		case 91:
			n += Minimal91(i)
		case 92:
			n += Minimal92(i)
		case 93:
			n += Minimal93(i)
		case 94:
			n += Minimal94(i)
		case 95:
			n += Minimal95(i)
		case 96:
			n += Minimal96(i)
		case 97:
			n += Minimal97(i)
		case 98:
			n += Minimal98(i)
		case 99:
			n += Minimal99(i)
		case 100:
			n += Minimal100(i)
		case 101:
			n += Minimal101(i)
		case 102:
			n += Minimal102(i)
		case 103:
			n += Minimal103(i)
		case 104:
			n += Minimal104(i)
		case 105:
			n += Minimal105(i)
		case 106:
			n += Minimal106(i)
		case 107:
			n += Minimal107(i)
		case 108:
			n += Minimal108(i)
		case 109:
			n += Minimal109(i)
		case 110:
			n += Minimal110(i)
		case 111:
			n += Minimal111(i)
		case 112:
			n += Minimal112(i)
		case 113:
			n += Minimal113(i)
		case 114:
			n += Minimal114(i)
		case 115:
			n += Minimal115(i)
		case 116:
			n += Minimal116(i)
		case 117:
			n += Minimal117(i)
		case 118:
			n += Minimal118(i)
		case 119:
			n += Minimal119(i)
		case 120:
			n += Minimal120(i)
		case 121:
			n += Minimal121(i)
		case 122:
			n += Minimal122(i)
		case 123:
			n += Minimal123(i)
		case 124:
			n += Minimal124(i)
		case 125:
			n += Minimal125(i)
		case 126:
			n += Minimal126(i)
		case 127:
			n += Minimal127(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessMapMinimalFunc128(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += MinimalFuncs[randInputs[i%len(randInputs)]%128](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessSwitchNoInlineFunc128(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch i % 128 {
		case 0:
			n += NoInline0(i)
		case 1:
			n += NoInline1(i)
		case 2:
			n += NoInline2(i)
		case 3:
			n += NoInline3(i)
		case 4:
			n += NoInline4(i)
		case 5:
			n += NoInline5(i)
		case 6:
			n += NoInline6(i)
		case 7:
			n += NoInline7(i)
		case 8:
			n += NoInline8(i)
		case 9:
			n += NoInline9(i)
		case 10:
			n += NoInline10(i)
		case 11:
			n += NoInline11(i)
		case 12:
			n += NoInline12(i)
		case 13:
			n += NoInline13(i)
		case 14:
			n += NoInline14(i)
		case 15:
			n += NoInline15(i)
		case 16:
			n += NoInline16(i)
		case 17:
			n += NoInline17(i)
		case 18:
			n += NoInline18(i)
		case 19:
			n += NoInline19(i)
		case 20:
			n += NoInline20(i)
		case 21:
			n += NoInline21(i)
		case 22:
			n += NoInline22(i)
		case 23:
			n += NoInline23(i)
		case 24:
			n += NoInline24(i)
		case 25:
			n += NoInline25(i)
		case 26:
			n += NoInline26(i)
		case 27:
			n += NoInline27(i)
		case 28:
			n += NoInline28(i)
		case 29:
			n += NoInline29(i)
		case 30:
			n += NoInline30(i)
		case 31:
			n += NoInline31(i)
		case 32:
			n += NoInline32(i)
		case 33:
			n += NoInline33(i)
		case 34:
			n += NoInline34(i)
		case 35:
			n += NoInline35(i)
		case 36:
			n += NoInline36(i)
		case 37:
			n += NoInline37(i)
		case 38:
			n += NoInline38(i)
		case 39:
			n += NoInline39(i)
		case 40:
			n += NoInline40(i)
		case 41:
			n += NoInline41(i)
		case 42:
			n += NoInline42(i)
		case 43:
			n += NoInline43(i)
		case 44:
			n += NoInline44(i)
		case 45:
			n += NoInline45(i)
		case 46:
			n += NoInline46(i)
		case 47:
			n += NoInline47(i)
		case 48:
			n += NoInline48(i)
		case 49:
			n += NoInline49(i)
		case 50:
			n += NoInline50(i)
		case 51:
			n += NoInline51(i)
		case 52:
			n += NoInline52(i)
		case 53:
			n += NoInline53(i)
		case 54:
			n += NoInline54(i)
		case 55:
			n += NoInline55(i)
		case 56:
			n += NoInline56(i)
		case 57:
			n += NoInline57(i)
		case 58:
			n += NoInline58(i)
		case 59:
			n += NoInline59(i)
		case 60:
			n += NoInline60(i)
		case 61:
			n += NoInline61(i)
		case 62:
			n += NoInline62(i)
		case 63:
			n += NoInline63(i)
		case 64:
			n += NoInline64(i)
		case 65:
			n += NoInline65(i)
		case 66:
			n += NoInline66(i)
		case 67:
			n += NoInline67(i)
		case 68:
			n += NoInline68(i)
		case 69:
			n += NoInline69(i)
		case 70:
			n += NoInline70(i)
		case 71:
			n += NoInline71(i)
		case 72:
			n += NoInline72(i)
		case 73:
			n += NoInline73(i)
		case 74:
			n += NoInline74(i)
		case 75:
			n += NoInline75(i)
		case 76:
			n += NoInline76(i)
		case 77:
			n += NoInline77(i)
		case 78:
			n += NoInline78(i)
		case 79:
			n += NoInline79(i)
		case 80:
			n += NoInline80(i)
		case 81:
			n += NoInline81(i)
		case 82:
			n += NoInline82(i)
		case 83:
			n += NoInline83(i)
		case 84:
			n += NoInline84(i)
		case 85:
			n += NoInline85(i)
		case 86:
			n += NoInline86(i)
		case 87:
			n += NoInline87(i)
		case 88:
			n += NoInline88(i)
		case 89:
			n += NoInline89(i)
		case 90:
			n += NoInline90(i)
		case 91:
			n += NoInline91(i)
		case 92:
			n += NoInline92(i)
		case 93:
			n += NoInline93(i)
		case 94:
			n += NoInline94(i)
		case 95:
			n += NoInline95(i)
		case 96:
			n += NoInline96(i)
		case 97:
			n += NoInline97(i)
		case 98:
			n += NoInline98(i)
		case 99:
			n += NoInline99(i)
		case 100:
			n += NoInline100(i)
		case 101:
			n += NoInline101(i)
		case 102:
			n += NoInline102(i)
		case 103:
			n += NoInline103(i)
		case 104:
			n += NoInline104(i)
		case 105:
			n += NoInline105(i)
		case 106:
			n += NoInline106(i)
		case 107:
			n += NoInline107(i)
		case 108:
			n += NoInline108(i)
		case 109:
			n += NoInline109(i)
		case 110:
			n += NoInline110(i)
		case 111:
			n += NoInline111(i)
		case 112:
			n += NoInline112(i)
		case 113:
			n += NoInline113(i)
		case 114:
			n += NoInline114(i)
		case 115:
			n += NoInline115(i)
		case 116:
			n += NoInline116(i)
		case 117:
			n += NoInline117(i)
		case 118:
			n += NoInline118(i)
		case 119:
			n += NoInline119(i)
		case 120:
			n += NoInline120(i)
		case 121:
			n += NoInline121(i)
		case 122:
			n += NoInline122(i)
		case 123:
			n += NoInline123(i)
		case 124:
			n += NoInline124(i)
		case 125:
			n += NoInline125(i)
		case 126:
			n += NoInline126(i)
		case 127:
			n += NoInline127(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessMapNoInlineFunc128(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += NoInlineFuncs[i%128](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessSwitchNoInlineFunc128(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch ascInputs[i%len(ascInputs)] % 128 {
		case 0:
			n += NoInline0(i)
		case 1:
			n += NoInline1(i)
		case 2:
			n += NoInline2(i)
		case 3:
			n += NoInline3(i)
		case 4:
			n += NoInline4(i)
		case 5:
			n += NoInline5(i)
		case 6:
			n += NoInline6(i)
		case 7:
			n += NoInline7(i)
		case 8:
			n += NoInline8(i)
		case 9:
			n += NoInline9(i)
		case 10:
			n += NoInline10(i)
		case 11:
			n += NoInline11(i)
		case 12:
			n += NoInline12(i)
		case 13:
			n += NoInline13(i)
		case 14:
			n += NoInline14(i)
		case 15:
			n += NoInline15(i)
		case 16:
			n += NoInline16(i)
		case 17:
			n += NoInline17(i)
		case 18:
			n += NoInline18(i)
		case 19:
			n += NoInline19(i)
		case 20:
			n += NoInline20(i)
		case 21:
			n += NoInline21(i)
		case 22:
			n += NoInline22(i)
		case 23:
			n += NoInline23(i)
		case 24:
			n += NoInline24(i)
		case 25:
			n += NoInline25(i)
		case 26:
			n += NoInline26(i)
		case 27:
			n += NoInline27(i)
		case 28:
			n += NoInline28(i)
		case 29:
			n += NoInline29(i)
		case 30:
			n += NoInline30(i)
		case 31:
			n += NoInline31(i)
		case 32:
			n += NoInline32(i)
		case 33:
			n += NoInline33(i)
		case 34:
			n += NoInline34(i)
		case 35:
			n += NoInline35(i)
		case 36:
			n += NoInline36(i)
		case 37:
			n += NoInline37(i)
		case 38:
			n += NoInline38(i)
		case 39:
			n += NoInline39(i)
		case 40:
			n += NoInline40(i)
		case 41:
			n += NoInline41(i)
		case 42:
			n += NoInline42(i)
		case 43:
			n += NoInline43(i)
		case 44:
			n += NoInline44(i)
		case 45:
			n += NoInline45(i)
		case 46:
			n += NoInline46(i)
		case 47:
			n += NoInline47(i)
		case 48:
			n += NoInline48(i)
		case 49:
			n += NoInline49(i)
		case 50:
			n += NoInline50(i)
		case 51:
			n += NoInline51(i)
		case 52:
			n += NoInline52(i)
		case 53:
			n += NoInline53(i)
		case 54:
			n += NoInline54(i)
		case 55:
			n += NoInline55(i)
		case 56:
			n += NoInline56(i)
		case 57:
			n += NoInline57(i)
		case 58:
			n += NoInline58(i)
		case 59:
			n += NoInline59(i)
		case 60:
			n += NoInline60(i)
		case 61:
			n += NoInline61(i)
		case 62:
			n += NoInline62(i)
		case 63:
			n += NoInline63(i)
		case 64:
			n += NoInline64(i)
		case 65:
			n += NoInline65(i)
		case 66:
			n += NoInline66(i)
		case 67:
			n += NoInline67(i)
		case 68:
			n += NoInline68(i)
		case 69:
			n += NoInline69(i)
		case 70:
			n += NoInline70(i)
		case 71:
			n += NoInline71(i)
		case 72:
			n += NoInline72(i)
		case 73:
			n += NoInline73(i)
		case 74:
			n += NoInline74(i)
		case 75:
			n += NoInline75(i)
		case 76:
			n += NoInline76(i)
		case 77:
			n += NoInline77(i)
		case 78:
			n += NoInline78(i)
		case 79:
			n += NoInline79(i)
		case 80:
			n += NoInline80(i)
		case 81:
			n += NoInline81(i)
		case 82:
			n += NoInline82(i)
		case 83:
			n += NoInline83(i)
		case 84:
			n += NoInline84(i)
		case 85:
			n += NoInline85(i)
		case 86:
			n += NoInline86(i)
		case 87:
			n += NoInline87(i)
		case 88:
			n += NoInline88(i)
		case 89:
			n += NoInline89(i)
		case 90:
			n += NoInline90(i)
		case 91:
			n += NoInline91(i)
		case 92:
			n += NoInline92(i)
		case 93:
			n += NoInline93(i)
		case 94:
			n += NoInline94(i)
		case 95:
			n += NoInline95(i)
		case 96:
			n += NoInline96(i)
		case 97:
			n += NoInline97(i)
		case 98:
			n += NoInline98(i)
		case 99:
			n += NoInline99(i)
		case 100:
			n += NoInline100(i)
		case 101:
			n += NoInline101(i)
		case 102:
			n += NoInline102(i)
		case 103:
			n += NoInline103(i)
		case 104:
			n += NoInline104(i)
		case 105:
			n += NoInline105(i)
		case 106:
			n += NoInline106(i)
		case 107:
			n += NoInline107(i)
		case 108:
			n += NoInline108(i)
		case 109:
			n += NoInline109(i)
		case 110:
			n += NoInline110(i)
		case 111:
			n += NoInline111(i)
		case 112:
			n += NoInline112(i)
		case 113:
			n += NoInline113(i)
		case 114:
			n += NoInline114(i)
		case 115:
			n += NoInline115(i)
		case 116:
			n += NoInline116(i)
		case 117:
			n += NoInline117(i)
		case 118:
			n += NoInline118(i)
		case 119:
			n += NoInline119(i)
		case 120:
			n += NoInline120(i)
		case 121:
			n += NoInline121(i)
		case 122:
			n += NoInline122(i)
		case 123:
			n += NoInline123(i)
		case 124:
			n += NoInline124(i)
		case 125:
			n += NoInline125(i)
		case 126:
			n += NoInline126(i)
		case 127:
			n += NoInline127(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessMapNoInlineFunc128(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += NoInlineFuncs[ascInputs[i%len(ascInputs)]%128](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessSwitchNoInlineFunc128(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch randInputs[i%len(randInputs)] % 128 {
		case 0:
			n += NoInline0(i)
		case 1:
			n += NoInline1(i)
		case 2:
			n += NoInline2(i)
		case 3:
			n += NoInline3(i)
		case 4:
			n += NoInline4(i)
		case 5:
			n += NoInline5(i)
		case 6:
			n += NoInline6(i)
		case 7:
			n += NoInline7(i)
		case 8:
			n += NoInline8(i)
		case 9:
			n += NoInline9(i)
		case 10:
			n += NoInline10(i)
		case 11:
			n += NoInline11(i)
		case 12:
			n += NoInline12(i)
		case 13:
			n += NoInline13(i)
		case 14:
			n += NoInline14(i)
		case 15:
			n += NoInline15(i)
		case 16:
			n += NoInline16(i)
		case 17:
			n += NoInline17(i)
		case 18:
			n += NoInline18(i)
		case 19:
			n += NoInline19(i)
		case 20:
			n += NoInline20(i)
		case 21:
			n += NoInline21(i)
		case 22:
			n += NoInline22(i)
		case 23:
			n += NoInline23(i)
		case 24:
			n += NoInline24(i)
		case 25:
			n += NoInline25(i)
		case 26:
			n += NoInline26(i)
		case 27:
			n += NoInline27(i)
		case 28:
			n += NoInline28(i)
		case 29:
			n += NoInline29(i)
		case 30:
			n += NoInline30(i)
		case 31:
			n += NoInline31(i)
		case 32:
			n += NoInline32(i)
		case 33:
			n += NoInline33(i)
		case 34:
			n += NoInline34(i)
		case 35:
			n += NoInline35(i)
		case 36:
			n += NoInline36(i)
		case 37:
			n += NoInline37(i)
		case 38:
			n += NoInline38(i)
		case 39:
			n += NoInline39(i)
		case 40:
			n += NoInline40(i)
		case 41:
			n += NoInline41(i)
		case 42:
			n += NoInline42(i)
		case 43:
			n += NoInline43(i)
		case 44:
			n += NoInline44(i)
		case 45:
			n += NoInline45(i)
		case 46:
			n += NoInline46(i)
		case 47:
			n += NoInline47(i)
		case 48:
			n += NoInline48(i)
		case 49:
			n += NoInline49(i)
		case 50:
			n += NoInline50(i)
		case 51:
			n += NoInline51(i)
		case 52:
			n += NoInline52(i)
		case 53:
			n += NoInline53(i)
		case 54:
			n += NoInline54(i)
		case 55:
			n += NoInline55(i)
		case 56:
			n += NoInline56(i)
		case 57:
			n += NoInline57(i)
		case 58:
			n += NoInline58(i)
		case 59:
			n += NoInline59(i)
		case 60:
			n += NoInline60(i)
		case 61:
			n += NoInline61(i)
		case 62:
			n += NoInline62(i)
		case 63:
			n += NoInline63(i)
		case 64:
			n += NoInline64(i)
		case 65:
			n += NoInline65(i)
		case 66:
			n += NoInline66(i)
		case 67:
			n += NoInline67(i)
		case 68:
			n += NoInline68(i)
		case 69:
			n += NoInline69(i)
		case 70:
			n += NoInline70(i)
		case 71:
			n += NoInline71(i)
		case 72:
			n += NoInline72(i)
		case 73:
			n += NoInline73(i)
		case 74:
			n += NoInline74(i)
		case 75:
			n += NoInline75(i)
		case 76:
			n += NoInline76(i)
		case 77:
			n += NoInline77(i)
		case 78:
			n += NoInline78(i)
		case 79:
			n += NoInline79(i)
		case 80:
			n += NoInline80(i)
		case 81:
			n += NoInline81(i)
		case 82:
			n += NoInline82(i)
		case 83:
			n += NoInline83(i)
		case 84:
			n += NoInline84(i)
		case 85:
			n += NoInline85(i)
		case 86:
			n += NoInline86(i)
		case 87:
			n += NoInline87(i)
		case 88:
			n += NoInline88(i)
		case 89:
			n += NoInline89(i)
		case 90:
			n += NoInline90(i)
		case 91:
			n += NoInline91(i)
		case 92:
			n += NoInline92(i)
		case 93:
			n += NoInline93(i)
		case 94:
			n += NoInline94(i)
		case 95:
			n += NoInline95(i)
		case 96:
			n += NoInline96(i)
		case 97:
			n += NoInline97(i)
		case 98:
			n += NoInline98(i)
		case 99:
			n += NoInline99(i)
		case 100:
			n += NoInline100(i)
		case 101:
			n += NoInline101(i)
		case 102:
			n += NoInline102(i)
		case 103:
			n += NoInline103(i)
		case 104:
			n += NoInline104(i)
		case 105:
			n += NoInline105(i)
		case 106:
			n += NoInline106(i)
		case 107:
			n += NoInline107(i)
		case 108:
			n += NoInline108(i)
		case 109:
			n += NoInline109(i)
		case 110:
			n += NoInline110(i)
		case 111:
			n += NoInline111(i)
		case 112:
			n += NoInline112(i)
		case 113:
			n += NoInline113(i)
		case 114:
			n += NoInline114(i)
		case 115:
			n += NoInline115(i)
		case 116:
			n += NoInline116(i)
		case 117:
			n += NoInline117(i)
		case 118:
			n += NoInline118(i)
		case 119:
			n += NoInline119(i)
		case 120:
			n += NoInline120(i)
		case 121:
			n += NoInline121(i)
		case 122:
			n += NoInline122(i)
		case 123:
			n += NoInline123(i)
		case 124:
			n += NoInline124(i)
		case 125:
			n += NoInline125(i)
		case 126:
			n += NoInline126(i)
		case 127:
			n += NoInline127(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessMapNoInlineFunc128(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += NoInlineFuncs[randInputs[i%len(randInputs)]%128](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessSwitchMinimalFunc256(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch i % 256 {
		case 0:
			n += Minimal0(i)
		case 1:
			n += Minimal1(i)
		case 2:
			n += Minimal2(i)
		case 3:
			n += Minimal3(i)
		case 4:
			n += Minimal4(i)
		case 5:
			n += Minimal5(i)
		case 6:
			n += Minimal6(i)
		case 7:
			n += Minimal7(i)
		case 8:
			n += Minimal8(i)
		case 9:
			n += Minimal9(i)
		case 10:
			n += Minimal10(i)
		case 11:
			n += Minimal11(i)
		case 12:
			n += Minimal12(i)
		case 13:
			n += Minimal13(i)
		case 14:
			n += Minimal14(i)
		case 15:
			n += Minimal15(i)
		case 16:
			n += Minimal16(i)
		case 17:
			n += Minimal17(i)
		case 18:
			n += Minimal18(i)
		case 19:
			n += Minimal19(i)
		case 20:
			n += Minimal20(i)
		case 21:
			n += Minimal21(i)
		case 22:
			n += Minimal22(i)
		case 23:
			n += Minimal23(i)
		case 24:
			n += Minimal24(i)
		case 25:
			n += Minimal25(i)
		case 26:
			n += Minimal26(i)
		case 27:
			n += Minimal27(i)
		case 28:
			n += Minimal28(i)
		case 29:
			n += Minimal29(i)
		case 30:
			n += Minimal30(i)
		case 31:
			n += Minimal31(i)
		case 32:
			n += Minimal32(i)
		case 33:
			n += Minimal33(i)
		case 34:
			n += Minimal34(i)
		case 35:
			n += Minimal35(i)
		case 36:
			n += Minimal36(i)
		case 37:
			n += Minimal37(i)
		case 38:
			n += Minimal38(i)
		case 39:
			n += Minimal39(i)
		case 40:
			n += Minimal40(i)
		case 41:
			n += Minimal41(i)
		case 42:
			n += Minimal42(i)
		case 43:
			n += Minimal43(i)
		case 44:
			n += Minimal44(i)
		case 45:
			n += Minimal45(i)
		case 46:
			n += Minimal46(i)
		case 47:
			n += Minimal47(i)
		case 48:
			n += Minimal48(i)
		case 49:
			n += Minimal49(i)
		case 50:
			n += Minimal50(i)
		case 51:
			n += Minimal51(i)
		case 52:
			n += Minimal52(i)
		case 53:
			n += Minimal53(i)
		case 54:
			n += Minimal54(i)
		case 55:
			n += Minimal55(i)
		case 56:
			n += Minimal56(i)
		case 57:
			n += Minimal57(i)
		case 58:
			n += Minimal58(i)
		case 59:
			n += Minimal59(i)
		case 60:
			n += Minimal60(i)
		case 61:
			n += Minimal61(i)
		case 62:
			n += Minimal62(i)
		case 63:
			n += Minimal63(i)
		case 64:
			n += Minimal64(i)
		case 65:
			n += Minimal65(i)
		case 66:
			n += Minimal66(i)
		case 67:
			n += Minimal67(i)
		case 68:
			n += Minimal68(i)
		case 69:
			n += Minimal69(i)
		case 70:
			n += Minimal70(i)
		case 71:
			n += Minimal71(i)
		case 72:
			n += Minimal72(i)
		case 73:
			n += Minimal73(i)
		case 74:
			n += Minimal74(i)
		case 75:
			n += Minimal75(i)
		case 76:
			n += Minimal76(i)
		case 77:
			n += Minimal77(i)
		case 78:
			n += Minimal78(i)
		case 79:
			n += Minimal79(i)
		case 80:
			n += Minimal80(i)
		case 81:
			n += Minimal81(i)
		case 82:
			n += Minimal82(i)
		case 83:
			n += Minimal83(i)
		case 84:
			n += Minimal84(i)
		case 85:
			n += Minimal85(i)
		case 86:
			n += Minimal86(i)
		case 87:
			n += Minimal87(i)
		case 88:
			n += Minimal88(i)
		case 89:
			n += Minimal89(i)
		case 90:
			n += Minimal90(i)
		case 91:
			n += Minimal91(i)
		case 92:
			n += Minimal92(i)
		case 93:
			n += Minimal93(i)
		case 94:
			n += Minimal94(i)
		case 95:
			n += Minimal95(i)
		case 96:
			n += Minimal96(i)
		case 97:
			n += Minimal97(i)
		case 98:
			n += Minimal98(i)
		case 99:
			n += Minimal99(i)
		case 100:
			n += Minimal100(i)
		case 101:
			n += Minimal101(i)
		case 102:
			n += Minimal102(i)
		case 103:
			n += Minimal103(i)
		case 104:
			n += Minimal104(i)
		case 105:
			n += Minimal105(i)
		case 106:
			n += Minimal106(i)
		case 107:
			n += Minimal107(i)
		case 108:
			n += Minimal108(i)
		case 109:
			n += Minimal109(i)
		case 110:
			n += Minimal110(i)
		case 111:
			n += Minimal111(i)
		case 112:
			n += Minimal112(i)
		case 113:
			n += Minimal113(i)
		case 114:
			n += Minimal114(i)
		case 115:
			n += Minimal115(i)
		case 116:
			n += Minimal116(i)
		case 117:
			n += Minimal117(i)
		case 118:
			n += Minimal118(i)
		case 119:
			n += Minimal119(i)
		case 120:
			n += Minimal120(i)
		case 121:
			n += Minimal121(i)
		case 122:
			n += Minimal122(i)
		case 123:
			n += Minimal123(i)
		case 124:
			n += Minimal124(i)
		case 125:
			n += Minimal125(i)
		case 126:
			n += Minimal126(i)
		case 127:
			n += Minimal127(i)
		case 128:
			n += Minimal128(i)
		case 129:
			n += Minimal129(i)
		case 130:
			n += Minimal130(i)
		case 131:
			n += Minimal131(i)
		case 132:
			n += Minimal132(i)
		case 133:
			n += Minimal133(i)
		case 134:
			n += Minimal134(i)
		case 135:
			n += Minimal135(i)
		case 136:
			n += Minimal136(i)
		case 137:
			n += Minimal137(i)
		case 138:
			n += Minimal138(i)
		case 139:
			n += Minimal139(i)
		case 140:
			n += Minimal140(i)
		case 141:
			n += Minimal141(i)
		case 142:
			n += Minimal142(i)
		case 143:
			n += Minimal143(i)
		case 144:
			n += Minimal144(i)
		case 145:
			n += Minimal145(i)
		case 146:
			n += Minimal146(i)
		case 147:
			n += Minimal147(i)
		case 148:
			n += Minimal148(i)
		case 149:
			n += Minimal149(i)
		case 150:
			n += Minimal150(i)
		case 151:
			n += Minimal151(i)
		case 152:
			n += Minimal152(i)
		case 153:
			n += Minimal153(i)
		case 154:
			n += Minimal154(i)
		case 155:
			n += Minimal155(i)
		case 156:
			n += Minimal156(i)
		case 157:
			n += Minimal157(i)
		case 158:
			n += Minimal158(i)
		case 159:
			n += Minimal159(i)
		case 160:
			n += Minimal160(i)
		case 161:
			n += Minimal161(i)
		case 162:
			n += Minimal162(i)
		case 163:
			n += Minimal163(i)
		case 164:
			n += Minimal164(i)
		case 165:
			n += Minimal165(i)
		case 166:
			n += Minimal166(i)
		case 167:
			n += Minimal167(i)
		case 168:
			n += Minimal168(i)
		case 169:
			n += Minimal169(i)
		case 170:
			n += Minimal170(i)
		case 171:
			n += Minimal171(i)
		case 172:
			n += Minimal172(i)
		case 173:
			n += Minimal173(i)
		case 174:
			n += Minimal174(i)
		case 175:
			n += Minimal175(i)
		case 176:
			n += Minimal176(i)
		case 177:
			n += Minimal177(i)
		case 178:
			n += Minimal178(i)
		case 179:
			n += Minimal179(i)
		case 180:
			n += Minimal180(i)
		case 181:
			n += Minimal181(i)
		case 182:
			n += Minimal182(i)
		case 183:
			n += Minimal183(i)
		case 184:
			n += Minimal184(i)
		case 185:
			n += Minimal185(i)
		case 186:
			n += Minimal186(i)
		case 187:
			n += Minimal187(i)
		case 188:
			n += Minimal188(i)
		case 189:
			n += Minimal189(i)
		case 190:
			n += Minimal190(i)
		case 191:
			n += Minimal191(i)
		case 192:
			n += Minimal192(i)
		case 193:
			n += Minimal193(i)
		case 194:
			n += Minimal194(i)
		case 195:
			n += Minimal195(i)
		case 196:
			n += Minimal196(i)
		case 197:
			n += Minimal197(i)
		case 198:
			n += Minimal198(i)
		case 199:
			n += Minimal199(i)
		case 200:
			n += Minimal200(i)
		case 201:
			n += Minimal201(i)
		case 202:
			n += Minimal202(i)
		case 203:
			n += Minimal203(i)
		case 204:
			n += Minimal204(i)
		case 205:
			n += Minimal205(i)
		case 206:
			n += Minimal206(i)
		case 207:
			n += Minimal207(i)
		case 208:
			n += Minimal208(i)
		case 209:
			n += Minimal209(i)
		case 210:
			n += Minimal210(i)
		case 211:
			n += Minimal211(i)
		case 212:
			n += Minimal212(i)
		case 213:
			n += Minimal213(i)
		case 214:
			n += Minimal214(i)
		case 215:
			n += Minimal215(i)
		case 216:
			n += Minimal216(i)
		case 217:
			n += Minimal217(i)
		case 218:
			n += Minimal218(i)
		case 219:
			n += Minimal219(i)
		case 220:
			n += Minimal220(i)
		case 221:
			n += Minimal221(i)
		case 222:
			n += Minimal222(i)
		case 223:
			n += Minimal223(i)
		case 224:
			n += Minimal224(i)
		case 225:
			n += Minimal225(i)
		case 226:
			n += Minimal226(i)
		case 227:
			n += Minimal227(i)
		case 228:
			n += Minimal228(i)
		case 229:
			n += Minimal229(i)
		case 230:
			n += Minimal230(i)
		case 231:
			n += Minimal231(i)
		case 232:
			n += Minimal232(i)
		case 233:
			n += Minimal233(i)
		case 234:
			n += Minimal234(i)
		case 235:
			n += Minimal235(i)
		case 236:
			n += Minimal236(i)
		case 237:
			n += Minimal237(i)
		case 238:
			n += Minimal238(i)
		case 239:
			n += Minimal239(i)
		case 240:
			n += Minimal240(i)
		case 241:
			n += Minimal241(i)
		case 242:
			n += Minimal242(i)
		case 243:
			n += Minimal243(i)
		case 244:
			n += Minimal244(i)
		case 245:
			n += Minimal245(i)
		case 246:
			n += Minimal246(i)
		case 247:
			n += Minimal247(i)
		case 248:
			n += Minimal248(i)
		case 249:
			n += Minimal249(i)
		case 250:
			n += Minimal250(i)
		case 251:
			n += Minimal251(i)
		case 252:
			n += Minimal252(i)
		case 253:
			n += Minimal253(i)
		case 254:
			n += Minimal254(i)
		case 255:
			n += Minimal255(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessMapMinimalFunc256(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += MinimalFuncs[i%256](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessSwitchMinimalFunc256(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch ascInputs[i%len(ascInputs)] % 256 {
		case 0:
			n += Minimal0(i)
		case 1:
			n += Minimal1(i)
		case 2:
			n += Minimal2(i)
		case 3:
			n += Minimal3(i)
		case 4:
			n += Minimal4(i)
		case 5:
			n += Minimal5(i)
		case 6:
			n += Minimal6(i)
		case 7:
			n += Minimal7(i)
		case 8:
			n += Minimal8(i)
		case 9:
			n += Minimal9(i)
		case 10:
			n += Minimal10(i)
		case 11:
			n += Minimal11(i)
		case 12:
			n += Minimal12(i)
		case 13:
			n += Minimal13(i)
		case 14:
			n += Minimal14(i)
		case 15:
			n += Minimal15(i)
		case 16:
			n += Minimal16(i)
		case 17:
			n += Minimal17(i)
		case 18:
			n += Minimal18(i)
		case 19:
			n += Minimal19(i)
		case 20:
			n += Minimal20(i)
		case 21:
			n += Minimal21(i)
		case 22:
			n += Minimal22(i)
		case 23:
			n += Minimal23(i)
		case 24:
			n += Minimal24(i)
		case 25:
			n += Minimal25(i)
		case 26:
			n += Minimal26(i)
		case 27:
			n += Minimal27(i)
		case 28:
			n += Minimal28(i)
		case 29:
			n += Minimal29(i)
		case 30:
			n += Minimal30(i)
		case 31:
			n += Minimal31(i)
		case 32:
			n += Minimal32(i)
		case 33:
			n += Minimal33(i)
		case 34:
			n += Minimal34(i)
		case 35:
			n += Minimal35(i)
		case 36:
			n += Minimal36(i)
		case 37:
			n += Minimal37(i)
		case 38:
			n += Minimal38(i)
		case 39:
			n += Minimal39(i)
		case 40:
			n += Minimal40(i)
		case 41:
			n += Minimal41(i)
		case 42:
			n += Minimal42(i)
		case 43:
			n += Minimal43(i)
		case 44:
			n += Minimal44(i)
		case 45:
			n += Minimal45(i)
		case 46:
			n += Minimal46(i)
		case 47:
			n += Minimal47(i)
		case 48:
			n += Minimal48(i)
		case 49:
			n += Minimal49(i)
		case 50:
			n += Minimal50(i)
		case 51:
			n += Minimal51(i)
		case 52:
			n += Minimal52(i)
		case 53:
			n += Minimal53(i)
		case 54:
			n += Minimal54(i)
		case 55:
			n += Minimal55(i)
		case 56:
			n += Minimal56(i)
		case 57:
			n += Minimal57(i)
		case 58:
			n += Minimal58(i)
		case 59:
			n += Minimal59(i)
		case 60:
			n += Minimal60(i)
		case 61:
			n += Minimal61(i)
		case 62:
			n += Minimal62(i)
		case 63:
			n += Minimal63(i)
		case 64:
			n += Minimal64(i)
		case 65:
			n += Minimal65(i)
		case 66:
			n += Minimal66(i)
		case 67:
			n += Minimal67(i)
		case 68:
			n += Minimal68(i)
		case 69:
			n += Minimal69(i)
		case 70:
			n += Minimal70(i)
		case 71:
			n += Minimal71(i)
		case 72:
			n += Minimal72(i)
		case 73:
			n += Minimal73(i)
		case 74:
			n += Minimal74(i)
		case 75:
			n += Minimal75(i)
		case 76:
			n += Minimal76(i)
		case 77:
			n += Minimal77(i)
		case 78:
			n += Minimal78(i)
		case 79:
			n += Minimal79(i)
		case 80:
			n += Minimal80(i)
		case 81:
			n += Minimal81(i)
		case 82:
			n += Minimal82(i)
		case 83:
			n += Minimal83(i)
		case 84:
			n += Minimal84(i)
		case 85:
			n += Minimal85(i)
		case 86:
			n += Minimal86(i)
		case 87:
			n += Minimal87(i)
		case 88:
			n += Minimal88(i)
		case 89:
			n += Minimal89(i)
		case 90:
			n += Minimal90(i)
		case 91:
			n += Minimal91(i)
		case 92:
			n += Minimal92(i)
		case 93:
			n += Minimal93(i)
		case 94:
			n += Minimal94(i)
		case 95:
			n += Minimal95(i)
		case 96:
			n += Minimal96(i)
		case 97:
			n += Minimal97(i)
		case 98:
			n += Minimal98(i)
		case 99:
			n += Minimal99(i)
		case 100:
			n += Minimal100(i)
		case 101:
			n += Minimal101(i)
		case 102:
			n += Minimal102(i)
		case 103:
			n += Minimal103(i)
		case 104:
			n += Minimal104(i)
		case 105:
			n += Minimal105(i)
		case 106:
			n += Minimal106(i)
		case 107:
			n += Minimal107(i)
		case 108:
			n += Minimal108(i)
		case 109:
			n += Minimal109(i)
		case 110:
			n += Minimal110(i)
		case 111:
			n += Minimal111(i)
		case 112:
			n += Minimal112(i)
		case 113:
			n += Minimal113(i)
		case 114:
			n += Minimal114(i)
		case 115:
			n += Minimal115(i)
		case 116:
			n += Minimal116(i)
		case 117:
			n += Minimal117(i)
		case 118:
			n += Minimal118(i)
		case 119:
			n += Minimal119(i)
		case 120:
			n += Minimal120(i)
		case 121:
			n += Minimal121(i)
		case 122:
			n += Minimal122(i)
		case 123:
			n += Minimal123(i)
		case 124:
			n += Minimal124(i)
		case 125:
			n += Minimal125(i)
		case 126:
			n += Minimal126(i)
		case 127:
			n += Minimal127(i)
		case 128:
			n += Minimal128(i)
		case 129:
			n += Minimal129(i)
		case 130:
			n += Minimal130(i)
		case 131:
			n += Minimal131(i)
		case 132:
			n += Minimal132(i)
		case 133:
			n += Minimal133(i)
		case 134:
			n += Minimal134(i)
		case 135:
			n += Minimal135(i)
		case 136:
			n += Minimal136(i)
		case 137:
			n += Minimal137(i)
		case 138:
			n += Minimal138(i)
		case 139:
			n += Minimal139(i)
		case 140:
			n += Minimal140(i)
		case 141:
			n += Minimal141(i)
		case 142:
			n += Minimal142(i)
		case 143:
			n += Minimal143(i)
		case 144:
			n += Minimal144(i)
		case 145:
			n += Minimal145(i)
		case 146:
			n += Minimal146(i)
		case 147:
			n += Minimal147(i)
		case 148:
			n += Minimal148(i)
		case 149:
			n += Minimal149(i)
		case 150:
			n += Minimal150(i)
		case 151:
			n += Minimal151(i)
		case 152:
			n += Minimal152(i)
		case 153:
			n += Minimal153(i)
		case 154:
			n += Minimal154(i)
		case 155:
			n += Minimal155(i)
		case 156:
			n += Minimal156(i)
		case 157:
			n += Minimal157(i)
		case 158:
			n += Minimal158(i)
		case 159:
			n += Minimal159(i)
		case 160:
			n += Minimal160(i)
		case 161:
			n += Minimal161(i)
		case 162:
			n += Minimal162(i)
		case 163:
			n += Minimal163(i)
		case 164:
			n += Minimal164(i)
		case 165:
			n += Minimal165(i)
		case 166:
			n += Minimal166(i)
		case 167:
			n += Minimal167(i)
		case 168:
			n += Minimal168(i)
		case 169:
			n += Minimal169(i)
		case 170:
			n += Minimal170(i)
		case 171:
			n += Minimal171(i)
		case 172:
			n += Minimal172(i)
		case 173:
			n += Minimal173(i)
		case 174:
			n += Minimal174(i)
		case 175:
			n += Minimal175(i)
		case 176:
			n += Minimal176(i)
		case 177:
			n += Minimal177(i)
		case 178:
			n += Minimal178(i)
		case 179:
			n += Minimal179(i)
		case 180:
			n += Minimal180(i)
		case 181:
			n += Minimal181(i)
		case 182:
			n += Minimal182(i)
		case 183:
			n += Minimal183(i)
		case 184:
			n += Minimal184(i)
		case 185:
			n += Minimal185(i)
		case 186:
			n += Minimal186(i)
		case 187:
			n += Minimal187(i)
		case 188:
			n += Minimal188(i)
		case 189:
			n += Minimal189(i)
		case 190:
			n += Minimal190(i)
		case 191:
			n += Minimal191(i)
		case 192:
			n += Minimal192(i)
		case 193:
			n += Minimal193(i)
		case 194:
			n += Minimal194(i)
		case 195:
			n += Minimal195(i)
		case 196:
			n += Minimal196(i)
		case 197:
			n += Minimal197(i)
		case 198:
			n += Minimal198(i)
		case 199:
			n += Minimal199(i)
		case 200:
			n += Minimal200(i)
		case 201:
			n += Minimal201(i)
		case 202:
			n += Minimal202(i)
		case 203:
			n += Minimal203(i)
		case 204:
			n += Minimal204(i)
		case 205:
			n += Minimal205(i)
		case 206:
			n += Minimal206(i)
		case 207:
			n += Minimal207(i)
		case 208:
			n += Minimal208(i)
		case 209:
			n += Minimal209(i)
		case 210:
			n += Minimal210(i)
		case 211:
			n += Minimal211(i)
		case 212:
			n += Minimal212(i)
		case 213:
			n += Minimal213(i)
		case 214:
			n += Minimal214(i)
		case 215:
			n += Minimal215(i)
		case 216:
			n += Minimal216(i)
		case 217:
			n += Minimal217(i)
		case 218:
			n += Minimal218(i)
		case 219:
			n += Minimal219(i)
		case 220:
			n += Minimal220(i)
		case 221:
			n += Minimal221(i)
		case 222:
			n += Minimal222(i)
		case 223:
			n += Minimal223(i)
		case 224:
			n += Minimal224(i)
		case 225:
			n += Minimal225(i)
		case 226:
			n += Minimal226(i)
		case 227:
			n += Minimal227(i)
		case 228:
			n += Minimal228(i)
		case 229:
			n += Minimal229(i)
		case 230:
			n += Minimal230(i)
		case 231:
			n += Minimal231(i)
		case 232:
			n += Minimal232(i)
		case 233:
			n += Minimal233(i)
		case 234:
			n += Minimal234(i)
		case 235:
			n += Minimal235(i)
		case 236:
			n += Minimal236(i)
		case 237:
			n += Minimal237(i)
		case 238:
			n += Minimal238(i)
		case 239:
			n += Minimal239(i)
		case 240:
			n += Minimal240(i)
		case 241:
			n += Minimal241(i)
		case 242:
			n += Minimal242(i)
		case 243:
			n += Minimal243(i)
		case 244:
			n += Minimal244(i)
		case 245:
			n += Minimal245(i)
		case 246:
			n += Minimal246(i)
		case 247:
			n += Minimal247(i)
		case 248:
			n += Minimal248(i)
		case 249:
			n += Minimal249(i)
		case 250:
			n += Minimal250(i)
		case 251:
			n += Minimal251(i)
		case 252:
			n += Minimal252(i)
		case 253:
			n += Minimal253(i)
		case 254:
			n += Minimal254(i)
		case 255:
			n += Minimal255(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessMapMinimalFunc256(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += MinimalFuncs[ascInputs[i%len(ascInputs)]%256](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessSwitchMinimalFunc256(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch randInputs[i%len(randInputs)] % 256 {
		case 0:
			n += Minimal0(i)
		case 1:
			n += Minimal1(i)
		case 2:
			n += Minimal2(i)
		case 3:
			n += Minimal3(i)
		case 4:
			n += Minimal4(i)
		case 5:
			n += Minimal5(i)
		case 6:
			n += Minimal6(i)
		case 7:
			n += Minimal7(i)
		case 8:
			n += Minimal8(i)
		case 9:
			n += Minimal9(i)
		case 10:
			n += Minimal10(i)
		case 11:
			n += Minimal11(i)
		case 12:
			n += Minimal12(i)
		case 13:
			n += Minimal13(i)
		case 14:
			n += Minimal14(i)
		case 15:
			n += Minimal15(i)
		case 16:
			n += Minimal16(i)
		case 17:
			n += Minimal17(i)
		case 18:
			n += Minimal18(i)
		case 19:
			n += Minimal19(i)
		case 20:
			n += Minimal20(i)
		case 21:
			n += Minimal21(i)
		case 22:
			n += Minimal22(i)
		case 23:
			n += Minimal23(i)
		case 24:
			n += Minimal24(i)
		case 25:
			n += Minimal25(i)
		case 26:
			n += Minimal26(i)
		case 27:
			n += Minimal27(i)
		case 28:
			n += Minimal28(i)
		case 29:
			n += Minimal29(i)
		case 30:
			n += Minimal30(i)
		case 31:
			n += Minimal31(i)
		case 32:
			n += Minimal32(i)
		case 33:
			n += Minimal33(i)
		case 34:
			n += Minimal34(i)
		case 35:
			n += Minimal35(i)
		case 36:
			n += Minimal36(i)
		case 37:
			n += Minimal37(i)
		case 38:
			n += Minimal38(i)
		case 39:
			n += Minimal39(i)
		case 40:
			n += Minimal40(i)
		case 41:
			n += Minimal41(i)
		case 42:
			n += Minimal42(i)
		case 43:
			n += Minimal43(i)
		case 44:
			n += Minimal44(i)
		case 45:
			n += Minimal45(i)
		case 46:
			n += Minimal46(i)
		case 47:
			n += Minimal47(i)
		case 48:
			n += Minimal48(i)
		case 49:
			n += Minimal49(i)
		case 50:
			n += Minimal50(i)
		case 51:
			n += Minimal51(i)
		case 52:
			n += Minimal52(i)
		case 53:
			n += Minimal53(i)
		case 54:
			n += Minimal54(i)
		case 55:
			n += Minimal55(i)
		case 56:
			n += Minimal56(i)
		case 57:
			n += Minimal57(i)
		case 58:
			n += Minimal58(i)
		case 59:
			n += Minimal59(i)
		case 60:
			n += Minimal60(i)
		case 61:
			n += Minimal61(i)
		case 62:
			n += Minimal62(i)
		case 63:
			n += Minimal63(i)
		case 64:
			n += Minimal64(i)
		case 65:
			n += Minimal65(i)
		case 66:
			n += Minimal66(i)
		case 67:
			n += Minimal67(i)
		case 68:
			n += Minimal68(i)
		case 69:
			n += Minimal69(i)
		case 70:
			n += Minimal70(i)
		case 71:
			n += Minimal71(i)
		case 72:
			n += Minimal72(i)
		case 73:
			n += Minimal73(i)
		case 74:
			n += Minimal74(i)
		case 75:
			n += Minimal75(i)
		case 76:
			n += Minimal76(i)
		case 77:
			n += Minimal77(i)
		case 78:
			n += Minimal78(i)
		case 79:
			n += Minimal79(i)
		case 80:
			n += Minimal80(i)
		case 81:
			n += Minimal81(i)
		case 82:
			n += Minimal82(i)
		case 83:
			n += Minimal83(i)
		case 84:
			n += Minimal84(i)
		case 85:
			n += Minimal85(i)
		case 86:
			n += Minimal86(i)
		case 87:
			n += Minimal87(i)
		case 88:
			n += Minimal88(i)
		case 89:
			n += Minimal89(i)
		case 90:
			n += Minimal90(i)
		case 91:
			n += Minimal91(i)
		case 92:
			n += Minimal92(i)
		case 93:
			n += Minimal93(i)
		case 94:
			n += Minimal94(i)
		case 95:
			n += Minimal95(i)
		case 96:
			n += Minimal96(i)
		case 97:
			n += Minimal97(i)
		case 98:
			n += Minimal98(i)
		case 99:
			n += Minimal99(i)
		case 100:
			n += Minimal100(i)
		case 101:
			n += Minimal101(i)
		case 102:
			n += Minimal102(i)
		case 103:
			n += Minimal103(i)
		case 104:
			n += Minimal104(i)
		case 105:
			n += Minimal105(i)
		case 106:
			n += Minimal106(i)
		case 107:
			n += Minimal107(i)
		case 108:
			n += Minimal108(i)
		case 109:
			n += Minimal109(i)
		case 110:
			n += Minimal110(i)
		case 111:
			n += Minimal111(i)
		case 112:
			n += Minimal112(i)
		case 113:
			n += Minimal113(i)
		case 114:
			n += Minimal114(i)
		case 115:
			n += Minimal115(i)
		case 116:
			n += Minimal116(i)
		case 117:
			n += Minimal117(i)
		case 118:
			n += Minimal118(i)
		case 119:
			n += Minimal119(i)
		case 120:
			n += Minimal120(i)
		case 121:
			n += Minimal121(i)
		case 122:
			n += Minimal122(i)
		case 123:
			n += Minimal123(i)
		case 124:
			n += Minimal124(i)
		case 125:
			n += Minimal125(i)
		case 126:
			n += Minimal126(i)
		case 127:
			n += Minimal127(i)
		case 128:
			n += Minimal128(i)
		case 129:
			n += Minimal129(i)
		case 130:
			n += Minimal130(i)
		case 131:
			n += Minimal131(i)
		case 132:
			n += Minimal132(i)
		case 133:
			n += Minimal133(i)
		case 134:
			n += Minimal134(i)
		case 135:
			n += Minimal135(i)
		case 136:
			n += Minimal136(i)
		case 137:
			n += Minimal137(i)
		case 138:
			n += Minimal138(i)
		case 139:
			n += Minimal139(i)
		case 140:
			n += Minimal140(i)
		case 141:
			n += Minimal141(i)
		case 142:
			n += Minimal142(i)
		case 143:
			n += Minimal143(i)
		case 144:
			n += Minimal144(i)
		case 145:
			n += Minimal145(i)
		case 146:
			n += Minimal146(i)
		case 147:
			n += Minimal147(i)
		case 148:
			n += Minimal148(i)
		case 149:
			n += Minimal149(i)
		case 150:
			n += Minimal150(i)
		case 151:
			n += Minimal151(i)
		case 152:
			n += Minimal152(i)
		case 153:
			n += Minimal153(i)
		case 154:
			n += Minimal154(i)
		case 155:
			n += Minimal155(i)
		case 156:
			n += Minimal156(i)
		case 157:
			n += Minimal157(i)
		case 158:
			n += Minimal158(i)
		case 159:
			n += Minimal159(i)
		case 160:
			n += Minimal160(i)
		case 161:
			n += Minimal161(i)
		case 162:
			n += Minimal162(i)
		case 163:
			n += Minimal163(i)
		case 164:
			n += Minimal164(i)
		case 165:
			n += Minimal165(i)
		case 166:
			n += Minimal166(i)
		case 167:
			n += Minimal167(i)
		case 168:
			n += Minimal168(i)
		case 169:
			n += Minimal169(i)
		case 170:
			n += Minimal170(i)
		case 171:
			n += Minimal171(i)
		case 172:
			n += Minimal172(i)
		case 173:
			n += Minimal173(i)
		case 174:
			n += Minimal174(i)
		case 175:
			n += Minimal175(i)
		case 176:
			n += Minimal176(i)
		case 177:
			n += Minimal177(i)
		case 178:
			n += Minimal178(i)
		case 179:
			n += Minimal179(i)
		case 180:
			n += Minimal180(i)
		case 181:
			n += Minimal181(i)
		case 182:
			n += Minimal182(i)
		case 183:
			n += Minimal183(i)
		case 184:
			n += Minimal184(i)
		case 185:
			n += Minimal185(i)
		case 186:
			n += Minimal186(i)
		case 187:
			n += Minimal187(i)
		case 188:
			n += Minimal188(i)
		case 189:
			n += Minimal189(i)
		case 190:
			n += Minimal190(i)
		case 191:
			n += Minimal191(i)
		case 192:
			n += Minimal192(i)
		case 193:
			n += Minimal193(i)
		case 194:
			n += Minimal194(i)
		case 195:
			n += Minimal195(i)
		case 196:
			n += Minimal196(i)
		case 197:
			n += Minimal197(i)
		case 198:
			n += Minimal198(i)
		case 199:
			n += Minimal199(i)
		case 200:
			n += Minimal200(i)
		case 201:
			n += Minimal201(i)
		case 202:
			n += Minimal202(i)
		case 203:
			n += Minimal203(i)
		case 204:
			n += Minimal204(i)
		case 205:
			n += Minimal205(i)
		case 206:
			n += Minimal206(i)
		case 207:
			n += Minimal207(i)
		case 208:
			n += Minimal208(i)
		case 209:
			n += Minimal209(i)
		case 210:
			n += Minimal210(i)
		case 211:
			n += Minimal211(i)
		case 212:
			n += Minimal212(i)
		case 213:
			n += Minimal213(i)
		case 214:
			n += Minimal214(i)
		case 215:
			n += Minimal215(i)
		case 216:
			n += Minimal216(i)
		case 217:
			n += Minimal217(i)
		case 218:
			n += Minimal218(i)
		case 219:
			n += Minimal219(i)
		case 220:
			n += Minimal220(i)
		case 221:
			n += Minimal221(i)
		case 222:
			n += Minimal222(i)
		case 223:
			n += Minimal223(i)
		case 224:
			n += Minimal224(i)
		case 225:
			n += Minimal225(i)
		case 226:
			n += Minimal226(i)
		case 227:
			n += Minimal227(i)
		case 228:
			n += Minimal228(i)
		case 229:
			n += Minimal229(i)
		case 230:
			n += Minimal230(i)
		case 231:
			n += Minimal231(i)
		case 232:
			n += Minimal232(i)
		case 233:
			n += Minimal233(i)
		case 234:
			n += Minimal234(i)
		case 235:
			n += Minimal235(i)
		case 236:
			n += Minimal236(i)
		case 237:
			n += Minimal237(i)
		case 238:
			n += Minimal238(i)
		case 239:
			n += Minimal239(i)
		case 240:
			n += Minimal240(i)
		case 241:
			n += Minimal241(i)
		case 242:
			n += Minimal242(i)
		case 243:
			n += Minimal243(i)
		case 244:
			n += Minimal244(i)
		case 245:
			n += Minimal245(i)
		case 246:
			n += Minimal246(i)
		case 247:
			n += Minimal247(i)
		case 248:
			n += Minimal248(i)
		case 249:
			n += Minimal249(i)
		case 250:
			n += Minimal250(i)
		case 251:
			n += Minimal251(i)
		case 252:
			n += Minimal252(i)
		case 253:
			n += Minimal253(i)
		case 254:
			n += Minimal254(i)
		case 255:
			n += Minimal255(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessMapMinimalFunc256(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += MinimalFuncs[randInputs[i%len(randInputs)]%256](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessSwitchNoInlineFunc256(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch i % 256 {
		case 0:
			n += NoInline0(i)
		case 1:
			n += NoInline1(i)
		case 2:
			n += NoInline2(i)
		case 3:
			n += NoInline3(i)
		case 4:
			n += NoInline4(i)
		case 5:
			n += NoInline5(i)
		case 6:
			n += NoInline6(i)
		case 7:
			n += NoInline7(i)
		case 8:
			n += NoInline8(i)
		case 9:
			n += NoInline9(i)
		case 10:
			n += NoInline10(i)
		case 11:
			n += NoInline11(i)
		case 12:
			n += NoInline12(i)
		case 13:
			n += NoInline13(i)
		case 14:
			n += NoInline14(i)
		case 15:
			n += NoInline15(i)
		case 16:
			n += NoInline16(i)
		case 17:
			n += NoInline17(i)
		case 18:
			n += NoInline18(i)
		case 19:
			n += NoInline19(i)
		case 20:
			n += NoInline20(i)
		case 21:
			n += NoInline21(i)
		case 22:
			n += NoInline22(i)
		case 23:
			n += NoInline23(i)
		case 24:
			n += NoInline24(i)
		case 25:
			n += NoInline25(i)
		case 26:
			n += NoInline26(i)
		case 27:
			n += NoInline27(i)
		case 28:
			n += NoInline28(i)
		case 29:
			n += NoInline29(i)
		case 30:
			n += NoInline30(i)
		case 31:
			n += NoInline31(i)
		case 32:
			n += NoInline32(i)
		case 33:
			n += NoInline33(i)
		case 34:
			n += NoInline34(i)
		case 35:
			n += NoInline35(i)
		case 36:
			n += NoInline36(i)
		case 37:
			n += NoInline37(i)
		case 38:
			n += NoInline38(i)
		case 39:
			n += NoInline39(i)
		case 40:
			n += NoInline40(i)
		case 41:
			n += NoInline41(i)
		case 42:
			n += NoInline42(i)
		case 43:
			n += NoInline43(i)
		case 44:
			n += NoInline44(i)
		case 45:
			n += NoInline45(i)
		case 46:
			n += NoInline46(i)
		case 47:
			n += NoInline47(i)
		case 48:
			n += NoInline48(i)
		case 49:
			n += NoInline49(i)
		case 50:
			n += NoInline50(i)
		case 51:
			n += NoInline51(i)
		case 52:
			n += NoInline52(i)
		case 53:
			n += NoInline53(i)
		case 54:
			n += NoInline54(i)
		case 55:
			n += NoInline55(i)
		case 56:
			n += NoInline56(i)
		case 57:
			n += NoInline57(i)
		case 58:
			n += NoInline58(i)
		case 59:
			n += NoInline59(i)
		case 60:
			n += NoInline60(i)
		case 61:
			n += NoInline61(i)
		case 62:
			n += NoInline62(i)
		case 63:
			n += NoInline63(i)
		case 64:
			n += NoInline64(i)
		case 65:
			n += NoInline65(i)
		case 66:
			n += NoInline66(i)
		case 67:
			n += NoInline67(i)
		case 68:
			n += NoInline68(i)
		case 69:
			n += NoInline69(i)
		case 70:
			n += NoInline70(i)
		case 71:
			n += NoInline71(i)
		case 72:
			n += NoInline72(i)
		case 73:
			n += NoInline73(i)
		case 74:
			n += NoInline74(i)
		case 75:
			n += NoInline75(i)
		case 76:
			n += NoInline76(i)
		case 77:
			n += NoInline77(i)
		case 78:
			n += NoInline78(i)
		case 79:
			n += NoInline79(i)
		case 80:
			n += NoInline80(i)
		case 81:
			n += NoInline81(i)
		case 82:
			n += NoInline82(i)
		case 83:
			n += NoInline83(i)
		case 84:
			n += NoInline84(i)
		case 85:
			n += NoInline85(i)
		case 86:
			n += NoInline86(i)
		case 87:
			n += NoInline87(i)
		case 88:
			n += NoInline88(i)
		case 89:
			n += NoInline89(i)
		case 90:
			n += NoInline90(i)
		case 91:
			n += NoInline91(i)
		case 92:
			n += NoInline92(i)
		case 93:
			n += NoInline93(i)
		case 94:
			n += NoInline94(i)
		case 95:
			n += NoInline95(i)
		case 96:
			n += NoInline96(i)
		case 97:
			n += NoInline97(i)
		case 98:
			n += NoInline98(i)
		case 99:
			n += NoInline99(i)
		case 100:
			n += NoInline100(i)
		case 101:
			n += NoInline101(i)
		case 102:
			n += NoInline102(i)
		case 103:
			n += NoInline103(i)
		case 104:
			n += NoInline104(i)
		case 105:
			n += NoInline105(i)
		case 106:
			n += NoInline106(i)
		case 107:
			n += NoInline107(i)
		case 108:
			n += NoInline108(i)
		case 109:
			n += NoInline109(i)
		case 110:
			n += NoInline110(i)
		case 111:
			n += NoInline111(i)
		case 112:
			n += NoInline112(i)
		case 113:
			n += NoInline113(i)
		case 114:
			n += NoInline114(i)
		case 115:
			n += NoInline115(i)
		case 116:
			n += NoInline116(i)
		case 117:
			n += NoInline117(i)
		case 118:
			n += NoInline118(i)
		case 119:
			n += NoInline119(i)
		case 120:
			n += NoInline120(i)
		case 121:
			n += NoInline121(i)
		case 122:
			n += NoInline122(i)
		case 123:
			n += NoInline123(i)
		case 124:
			n += NoInline124(i)
		case 125:
			n += NoInline125(i)
		case 126:
			n += NoInline126(i)
		case 127:
			n += NoInline127(i)
		case 128:
			n += NoInline128(i)
		case 129:
			n += NoInline129(i)
		case 130:
			n += NoInline130(i)
		case 131:
			n += NoInline131(i)
		case 132:
			n += NoInline132(i)
		case 133:
			n += NoInline133(i)
		case 134:
			n += NoInline134(i)
		case 135:
			n += NoInline135(i)
		case 136:
			n += NoInline136(i)
		case 137:
			n += NoInline137(i)
		case 138:
			n += NoInline138(i)
		case 139:
			n += NoInline139(i)
		case 140:
			n += NoInline140(i)
		case 141:
			n += NoInline141(i)
		case 142:
			n += NoInline142(i)
		case 143:
			n += NoInline143(i)
		case 144:
			n += NoInline144(i)
		case 145:
			n += NoInline145(i)
		case 146:
			n += NoInline146(i)
		case 147:
			n += NoInline147(i)
		case 148:
			n += NoInline148(i)
		case 149:
			n += NoInline149(i)
		case 150:
			n += NoInline150(i)
		case 151:
			n += NoInline151(i)
		case 152:
			n += NoInline152(i)
		case 153:
			n += NoInline153(i)
		case 154:
			n += NoInline154(i)
		case 155:
			n += NoInline155(i)
		case 156:
			n += NoInline156(i)
		case 157:
			n += NoInline157(i)
		case 158:
			n += NoInline158(i)
		case 159:
			n += NoInline159(i)
		case 160:
			n += NoInline160(i)
		case 161:
			n += NoInline161(i)
		case 162:
			n += NoInline162(i)
		case 163:
			n += NoInline163(i)
		case 164:
			n += NoInline164(i)
		case 165:
			n += NoInline165(i)
		case 166:
			n += NoInline166(i)
		case 167:
			n += NoInline167(i)
		case 168:
			n += NoInline168(i)
		case 169:
			n += NoInline169(i)
		case 170:
			n += NoInline170(i)
		case 171:
			n += NoInline171(i)
		case 172:
			n += NoInline172(i)
		case 173:
			n += NoInline173(i)
		case 174:
			n += NoInline174(i)
		case 175:
			n += NoInline175(i)
		case 176:
			n += NoInline176(i)
		case 177:
			n += NoInline177(i)
		case 178:
			n += NoInline178(i)
		case 179:
			n += NoInline179(i)
		case 180:
			n += NoInline180(i)
		case 181:
			n += NoInline181(i)
		case 182:
			n += NoInline182(i)
		case 183:
			n += NoInline183(i)
		case 184:
			n += NoInline184(i)
		case 185:
			n += NoInline185(i)
		case 186:
			n += NoInline186(i)
		case 187:
			n += NoInline187(i)
		case 188:
			n += NoInline188(i)
		case 189:
			n += NoInline189(i)
		case 190:
			n += NoInline190(i)
		case 191:
			n += NoInline191(i)
		case 192:
			n += NoInline192(i)
		case 193:
			n += NoInline193(i)
		case 194:
			n += NoInline194(i)
		case 195:
			n += NoInline195(i)
		case 196:
			n += NoInline196(i)
		case 197:
			n += NoInline197(i)
		case 198:
			n += NoInline198(i)
		case 199:
			n += NoInline199(i)
		case 200:
			n += NoInline200(i)
		case 201:
			n += NoInline201(i)
		case 202:
			n += NoInline202(i)
		case 203:
			n += NoInline203(i)
		case 204:
			n += NoInline204(i)
		case 205:
			n += NoInline205(i)
		case 206:
			n += NoInline206(i)
		case 207:
			n += NoInline207(i)
		case 208:
			n += NoInline208(i)
		case 209:
			n += NoInline209(i)
		case 210:
			n += NoInline210(i)
		case 211:
			n += NoInline211(i)
		case 212:
			n += NoInline212(i)
		case 213:
			n += NoInline213(i)
		case 214:
			n += NoInline214(i)
		case 215:
			n += NoInline215(i)
		case 216:
			n += NoInline216(i)
		case 217:
			n += NoInline217(i)
		case 218:
			n += NoInline218(i)
		case 219:
			n += NoInline219(i)
		case 220:
			n += NoInline220(i)
		case 221:
			n += NoInline221(i)
		case 222:
			n += NoInline222(i)
		case 223:
			n += NoInline223(i)
		case 224:
			n += NoInline224(i)
		case 225:
			n += NoInline225(i)
		case 226:
			n += NoInline226(i)
		case 227:
			n += NoInline227(i)
		case 228:
			n += NoInline228(i)
		case 229:
			n += NoInline229(i)
		case 230:
			n += NoInline230(i)
		case 231:
			n += NoInline231(i)
		case 232:
			n += NoInline232(i)
		case 233:
			n += NoInline233(i)
		case 234:
			n += NoInline234(i)
		case 235:
			n += NoInline235(i)
		case 236:
			n += NoInline236(i)
		case 237:
			n += NoInline237(i)
		case 238:
			n += NoInline238(i)
		case 239:
			n += NoInline239(i)
		case 240:
			n += NoInline240(i)
		case 241:
			n += NoInline241(i)
		case 242:
			n += NoInline242(i)
		case 243:
			n += NoInline243(i)
		case 244:
			n += NoInline244(i)
		case 245:
			n += NoInline245(i)
		case 246:
			n += NoInline246(i)
		case 247:
			n += NoInline247(i)
		case 248:
			n += NoInline248(i)
		case 249:
			n += NoInline249(i)
		case 250:
			n += NoInline250(i)
		case 251:
			n += NoInline251(i)
		case 252:
			n += NoInline252(i)
		case 253:
			n += NoInline253(i)
		case 254:
			n += NoInline254(i)
		case 255:
			n += NoInline255(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessMapNoInlineFunc256(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += NoInlineFuncs[i%256](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessSwitchNoInlineFunc256(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch ascInputs[i%len(ascInputs)] % 256 {
		case 0:
			n += NoInline0(i)
		case 1:
			n += NoInline1(i)
		case 2:
			n += NoInline2(i)
		case 3:
			n += NoInline3(i)
		case 4:
			n += NoInline4(i)
		case 5:
			n += NoInline5(i)
		case 6:
			n += NoInline6(i)
		case 7:
			n += NoInline7(i)
		case 8:
			n += NoInline8(i)
		case 9:
			n += NoInline9(i)
		case 10:
			n += NoInline10(i)
		case 11:
			n += NoInline11(i)
		case 12:
			n += NoInline12(i)
		case 13:
			n += NoInline13(i)
		case 14:
			n += NoInline14(i)
		case 15:
			n += NoInline15(i)
		case 16:
			n += NoInline16(i)
		case 17:
			n += NoInline17(i)
		case 18:
			n += NoInline18(i)
		case 19:
			n += NoInline19(i)
		case 20:
			n += NoInline20(i)
		case 21:
			n += NoInline21(i)
		case 22:
			n += NoInline22(i)
		case 23:
			n += NoInline23(i)
		case 24:
			n += NoInline24(i)
		case 25:
			n += NoInline25(i)
		case 26:
			n += NoInline26(i)
		case 27:
			n += NoInline27(i)
		case 28:
			n += NoInline28(i)
		case 29:
			n += NoInline29(i)
		case 30:
			n += NoInline30(i)
		case 31:
			n += NoInline31(i)
		case 32:
			n += NoInline32(i)
		case 33:
			n += NoInline33(i)
		case 34:
			n += NoInline34(i)
		case 35:
			n += NoInline35(i)
		case 36:
			n += NoInline36(i)
		case 37:
			n += NoInline37(i)
		case 38:
			n += NoInline38(i)
		case 39:
			n += NoInline39(i)
		case 40:
			n += NoInline40(i)
		case 41:
			n += NoInline41(i)
		case 42:
			n += NoInline42(i)
		case 43:
			n += NoInline43(i)
		case 44:
			n += NoInline44(i)
		case 45:
			n += NoInline45(i)
		case 46:
			n += NoInline46(i)
		case 47:
			n += NoInline47(i)
		case 48:
			n += NoInline48(i)
		case 49:
			n += NoInline49(i)
		case 50:
			n += NoInline50(i)
		case 51:
			n += NoInline51(i)
		case 52:
			n += NoInline52(i)
		case 53:
			n += NoInline53(i)
		case 54:
			n += NoInline54(i)
		case 55:
			n += NoInline55(i)
		case 56:
			n += NoInline56(i)
		case 57:
			n += NoInline57(i)
		case 58:
			n += NoInline58(i)
		case 59:
			n += NoInline59(i)
		case 60:
			n += NoInline60(i)
		case 61:
			n += NoInline61(i)
		case 62:
			n += NoInline62(i)
		case 63:
			n += NoInline63(i)
		case 64:
			n += NoInline64(i)
		case 65:
			n += NoInline65(i)
		case 66:
			n += NoInline66(i)
		case 67:
			n += NoInline67(i)
		case 68:
			n += NoInline68(i)
		case 69:
			n += NoInline69(i)
		case 70:
			n += NoInline70(i)
		case 71:
			n += NoInline71(i)
		case 72:
			n += NoInline72(i)
		case 73:
			n += NoInline73(i)
		case 74:
			n += NoInline74(i)
		case 75:
			n += NoInline75(i)
		case 76:
			n += NoInline76(i)
		case 77:
			n += NoInline77(i)
		case 78:
			n += NoInline78(i)
		case 79:
			n += NoInline79(i)
		case 80:
			n += NoInline80(i)
		case 81:
			n += NoInline81(i)
		case 82:
			n += NoInline82(i)
		case 83:
			n += NoInline83(i)
		case 84:
			n += NoInline84(i)
		case 85:
			n += NoInline85(i)
		case 86:
			n += NoInline86(i)
		case 87:
			n += NoInline87(i)
		case 88:
			n += NoInline88(i)
		case 89:
			n += NoInline89(i)
		case 90:
			n += NoInline90(i)
		case 91:
			n += NoInline91(i)
		case 92:
			n += NoInline92(i)
		case 93:
			n += NoInline93(i)
		case 94:
			n += NoInline94(i)
		case 95:
			n += NoInline95(i)
		case 96:
			n += NoInline96(i)
		case 97:
			n += NoInline97(i)
		case 98:
			n += NoInline98(i)
		case 99:
			n += NoInline99(i)
		case 100:
			n += NoInline100(i)
		case 101:
			n += NoInline101(i)
		case 102:
			n += NoInline102(i)
		case 103:
			n += NoInline103(i)
		case 104:
			n += NoInline104(i)
		case 105:
			n += NoInline105(i)
		case 106:
			n += NoInline106(i)
		case 107:
			n += NoInline107(i)
		case 108:
			n += NoInline108(i)
		case 109:
			n += NoInline109(i)
		case 110:
			n += NoInline110(i)
		case 111:
			n += NoInline111(i)
		case 112:
			n += NoInline112(i)
		case 113:
			n += NoInline113(i)
		case 114:
			n += NoInline114(i)
		case 115:
			n += NoInline115(i)
		case 116:
			n += NoInline116(i)
		case 117:
			n += NoInline117(i)
		case 118:
			n += NoInline118(i)
		case 119:
			n += NoInline119(i)
		case 120:
			n += NoInline120(i)
		case 121:
			n += NoInline121(i)
		case 122:
			n += NoInline122(i)
		case 123:
			n += NoInline123(i)
		case 124:
			n += NoInline124(i)
		case 125:
			n += NoInline125(i)
		case 126:
			n += NoInline126(i)
		case 127:
			n += NoInline127(i)
		case 128:
			n += NoInline128(i)
		case 129:
			n += NoInline129(i)
		case 130:
			n += NoInline130(i)
		case 131:
			n += NoInline131(i)
		case 132:
			n += NoInline132(i)
		case 133:
			n += NoInline133(i)
		case 134:
			n += NoInline134(i)
		case 135:
			n += NoInline135(i)
		case 136:
			n += NoInline136(i)
		case 137:
			n += NoInline137(i)
		case 138:
			n += NoInline138(i)
		case 139:
			n += NoInline139(i)
		case 140:
			n += NoInline140(i)
		case 141:
			n += NoInline141(i)
		case 142:
			n += NoInline142(i)
		case 143:
			n += NoInline143(i)
		case 144:
			n += NoInline144(i)
		case 145:
			n += NoInline145(i)
		case 146:
			n += NoInline146(i)
		case 147:
			n += NoInline147(i)
		case 148:
			n += NoInline148(i)
		case 149:
			n += NoInline149(i)
		case 150:
			n += NoInline150(i)
		case 151:
			n += NoInline151(i)
		case 152:
			n += NoInline152(i)
		case 153:
			n += NoInline153(i)
		case 154:
			n += NoInline154(i)
		case 155:
			n += NoInline155(i)
		case 156:
			n += NoInline156(i)
		case 157:
			n += NoInline157(i)
		case 158:
			n += NoInline158(i)
		case 159:
			n += NoInline159(i)
		case 160:
			n += NoInline160(i)
		case 161:
			n += NoInline161(i)
		case 162:
			n += NoInline162(i)
		case 163:
			n += NoInline163(i)
		case 164:
			n += NoInline164(i)
		case 165:
			n += NoInline165(i)
		case 166:
			n += NoInline166(i)
		case 167:
			n += NoInline167(i)
		case 168:
			n += NoInline168(i)
		case 169:
			n += NoInline169(i)
		case 170:
			n += NoInline170(i)
		case 171:
			n += NoInline171(i)
		case 172:
			n += NoInline172(i)
		case 173:
			n += NoInline173(i)
		case 174:
			n += NoInline174(i)
		case 175:
			n += NoInline175(i)
		case 176:
			n += NoInline176(i)
		case 177:
			n += NoInline177(i)
		case 178:
			n += NoInline178(i)
		case 179:
			n += NoInline179(i)
		case 180:
			n += NoInline180(i)
		case 181:
			n += NoInline181(i)
		case 182:
			n += NoInline182(i)
		case 183:
			n += NoInline183(i)
		case 184:
			n += NoInline184(i)
		case 185:
			n += NoInline185(i)
		case 186:
			n += NoInline186(i)
		case 187:
			n += NoInline187(i)
		case 188:
			n += NoInline188(i)
		case 189:
			n += NoInline189(i)
		case 190:
			n += NoInline190(i)
		case 191:
			n += NoInline191(i)
		case 192:
			n += NoInline192(i)
		case 193:
			n += NoInline193(i)
		case 194:
			n += NoInline194(i)
		case 195:
			n += NoInline195(i)
		case 196:
			n += NoInline196(i)
		case 197:
			n += NoInline197(i)
		case 198:
			n += NoInline198(i)
		case 199:
			n += NoInline199(i)
		case 200:
			n += NoInline200(i)
		case 201:
			n += NoInline201(i)
		case 202:
			n += NoInline202(i)
		case 203:
			n += NoInline203(i)
		case 204:
			n += NoInline204(i)
		case 205:
			n += NoInline205(i)
		case 206:
			n += NoInline206(i)
		case 207:
			n += NoInline207(i)
		case 208:
			n += NoInline208(i)
		case 209:
			n += NoInline209(i)
		case 210:
			n += NoInline210(i)
		case 211:
			n += NoInline211(i)
		case 212:
			n += NoInline212(i)
		case 213:
			n += NoInline213(i)
		case 214:
			n += NoInline214(i)
		case 215:
			n += NoInline215(i)
		case 216:
			n += NoInline216(i)
		case 217:
			n += NoInline217(i)
		case 218:
			n += NoInline218(i)
		case 219:
			n += NoInline219(i)
		case 220:
			n += NoInline220(i)
		case 221:
			n += NoInline221(i)
		case 222:
			n += NoInline222(i)
		case 223:
			n += NoInline223(i)
		case 224:
			n += NoInline224(i)
		case 225:
			n += NoInline225(i)
		case 226:
			n += NoInline226(i)
		case 227:
			n += NoInline227(i)
		case 228:
			n += NoInline228(i)
		case 229:
			n += NoInline229(i)
		case 230:
			n += NoInline230(i)
		case 231:
			n += NoInline231(i)
		case 232:
			n += NoInline232(i)
		case 233:
			n += NoInline233(i)
		case 234:
			n += NoInline234(i)
		case 235:
			n += NoInline235(i)
		case 236:
			n += NoInline236(i)
		case 237:
			n += NoInline237(i)
		case 238:
			n += NoInline238(i)
		case 239:
			n += NoInline239(i)
		case 240:
			n += NoInline240(i)
		case 241:
			n += NoInline241(i)
		case 242:
			n += NoInline242(i)
		case 243:
			n += NoInline243(i)
		case 244:
			n += NoInline244(i)
		case 245:
			n += NoInline245(i)
		case 246:
			n += NoInline246(i)
		case 247:
			n += NoInline247(i)
		case 248:
			n += NoInline248(i)
		case 249:
			n += NoInline249(i)
		case 250:
			n += NoInline250(i)
		case 251:
			n += NoInline251(i)
		case 252:
			n += NoInline252(i)
		case 253:
			n += NoInline253(i)
		case 254:
			n += NoInline254(i)
		case 255:
			n += NoInline255(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessMapNoInlineFunc256(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += NoInlineFuncs[ascInputs[i%len(ascInputs)]%256](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessSwitchNoInlineFunc256(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch randInputs[i%len(randInputs)] % 256 {
		case 0:
			n += NoInline0(i)
		case 1:
			n += NoInline1(i)
		case 2:
			n += NoInline2(i)
		case 3:
			n += NoInline3(i)
		case 4:
			n += NoInline4(i)
		case 5:
			n += NoInline5(i)
		case 6:
			n += NoInline6(i)
		case 7:
			n += NoInline7(i)
		case 8:
			n += NoInline8(i)
		case 9:
			n += NoInline9(i)
		case 10:
			n += NoInline10(i)
		case 11:
			n += NoInline11(i)
		case 12:
			n += NoInline12(i)
		case 13:
			n += NoInline13(i)
		case 14:
			n += NoInline14(i)
		case 15:
			n += NoInline15(i)
		case 16:
			n += NoInline16(i)
		case 17:
			n += NoInline17(i)
		case 18:
			n += NoInline18(i)
		case 19:
			n += NoInline19(i)
		case 20:
			n += NoInline20(i)
		case 21:
			n += NoInline21(i)
		case 22:
			n += NoInline22(i)
		case 23:
			n += NoInline23(i)
		case 24:
			n += NoInline24(i)
		case 25:
			n += NoInline25(i)
		case 26:
			n += NoInline26(i)
		case 27:
			n += NoInline27(i)
		case 28:
			n += NoInline28(i)
		case 29:
			n += NoInline29(i)
		case 30:
			n += NoInline30(i)
		case 31:
			n += NoInline31(i)
		case 32:
			n += NoInline32(i)
		case 33:
			n += NoInline33(i)
		case 34:
			n += NoInline34(i)
		case 35:
			n += NoInline35(i)
		case 36:
			n += NoInline36(i)
		case 37:
			n += NoInline37(i)
		case 38:
			n += NoInline38(i)
		case 39:
			n += NoInline39(i)
		case 40:
			n += NoInline40(i)
		case 41:
			n += NoInline41(i)
		case 42:
			n += NoInline42(i)
		case 43:
			n += NoInline43(i)
		case 44:
			n += NoInline44(i)
		case 45:
			n += NoInline45(i)
		case 46:
			n += NoInline46(i)
		case 47:
			n += NoInline47(i)
		case 48:
			n += NoInline48(i)
		case 49:
			n += NoInline49(i)
		case 50:
			n += NoInline50(i)
		case 51:
			n += NoInline51(i)
		case 52:
			n += NoInline52(i)
		case 53:
			n += NoInline53(i)
		case 54:
			n += NoInline54(i)
		case 55:
			n += NoInline55(i)
		case 56:
			n += NoInline56(i)
		case 57:
			n += NoInline57(i)
		case 58:
			n += NoInline58(i)
		case 59:
			n += NoInline59(i)
		case 60:
			n += NoInline60(i)
		case 61:
			n += NoInline61(i)
		case 62:
			n += NoInline62(i)
		case 63:
			n += NoInline63(i)
		case 64:
			n += NoInline64(i)
		case 65:
			n += NoInline65(i)
		case 66:
			n += NoInline66(i)
		case 67:
			n += NoInline67(i)
		case 68:
			n += NoInline68(i)
		case 69:
			n += NoInline69(i)
		case 70:
			n += NoInline70(i)
		case 71:
			n += NoInline71(i)
		case 72:
			n += NoInline72(i)
		case 73:
			n += NoInline73(i)
		case 74:
			n += NoInline74(i)
		case 75:
			n += NoInline75(i)
		case 76:
			n += NoInline76(i)
		case 77:
			n += NoInline77(i)
		case 78:
			n += NoInline78(i)
		case 79:
			n += NoInline79(i)
		case 80:
			n += NoInline80(i)
		case 81:
			n += NoInline81(i)
		case 82:
			n += NoInline82(i)
		case 83:
			n += NoInline83(i)
		case 84:
			n += NoInline84(i)
		case 85:
			n += NoInline85(i)
		case 86:
			n += NoInline86(i)
		case 87:
			n += NoInline87(i)
		case 88:
			n += NoInline88(i)
		case 89:
			n += NoInline89(i)
		case 90:
			n += NoInline90(i)
		case 91:
			n += NoInline91(i)
		case 92:
			n += NoInline92(i)
		case 93:
			n += NoInline93(i)
		case 94:
			n += NoInline94(i)
		case 95:
			n += NoInline95(i)
		case 96:
			n += NoInline96(i)
		case 97:
			n += NoInline97(i)
		case 98:
			n += NoInline98(i)
		case 99:
			n += NoInline99(i)
		case 100:
			n += NoInline100(i)
		case 101:
			n += NoInline101(i)
		case 102:
			n += NoInline102(i)
		case 103:
			n += NoInline103(i)
		case 104:
			n += NoInline104(i)
		case 105:
			n += NoInline105(i)
		case 106:
			n += NoInline106(i)
		case 107:
			n += NoInline107(i)
		case 108:
			n += NoInline108(i)
		case 109:
			n += NoInline109(i)
		case 110:
			n += NoInline110(i)
		case 111:
			n += NoInline111(i)
		case 112:
			n += NoInline112(i)
		case 113:
			n += NoInline113(i)
		case 114:
			n += NoInline114(i)
		case 115:
			n += NoInline115(i)
		case 116:
			n += NoInline116(i)
		case 117:
			n += NoInline117(i)
		case 118:
			n += NoInline118(i)
		case 119:
			n += NoInline119(i)
		case 120:
			n += NoInline120(i)
		case 121:
			n += NoInline121(i)
		case 122:
			n += NoInline122(i)
		case 123:
			n += NoInline123(i)
		case 124:
			n += NoInline124(i)
		case 125:
			n += NoInline125(i)
		case 126:
			n += NoInline126(i)
		case 127:
			n += NoInline127(i)
		case 128:
			n += NoInline128(i)
		case 129:
			n += NoInline129(i)
		case 130:
			n += NoInline130(i)
		case 131:
			n += NoInline131(i)
		case 132:
			n += NoInline132(i)
		case 133:
			n += NoInline133(i)
		case 134:
			n += NoInline134(i)
		case 135:
			n += NoInline135(i)
		case 136:
			n += NoInline136(i)
		case 137:
			n += NoInline137(i)
		case 138:
			n += NoInline138(i)
		case 139:
			n += NoInline139(i)
		case 140:
			n += NoInline140(i)
		case 141:
			n += NoInline141(i)
		case 142:
			n += NoInline142(i)
		case 143:
			n += NoInline143(i)
		case 144:
			n += NoInline144(i)
		case 145:
			n += NoInline145(i)
		case 146:
			n += NoInline146(i)
		case 147:
			n += NoInline147(i)
		case 148:
			n += NoInline148(i)
		case 149:
			n += NoInline149(i)
		case 150:
			n += NoInline150(i)
		case 151:
			n += NoInline151(i)
		case 152:
			n += NoInline152(i)
		case 153:
			n += NoInline153(i)
		case 154:
			n += NoInline154(i)
		case 155:
			n += NoInline155(i)
		case 156:
			n += NoInline156(i)
		case 157:
			n += NoInline157(i)
		case 158:
			n += NoInline158(i)
		case 159:
			n += NoInline159(i)
		case 160:
			n += NoInline160(i)
		case 161:
			n += NoInline161(i)
		case 162:
			n += NoInline162(i)
		case 163:
			n += NoInline163(i)
		case 164:
			n += NoInline164(i)
		case 165:
			n += NoInline165(i)
		case 166:
			n += NoInline166(i)
		case 167:
			n += NoInline167(i)
		case 168:
			n += NoInline168(i)
		case 169:
			n += NoInline169(i)
		case 170:
			n += NoInline170(i)
		case 171:
			n += NoInline171(i)
		case 172:
			n += NoInline172(i)
		case 173:
			n += NoInline173(i)
		case 174:
			n += NoInline174(i)
		case 175:
			n += NoInline175(i)
		case 176:
			n += NoInline176(i)
		case 177:
			n += NoInline177(i)
		case 178:
			n += NoInline178(i)
		case 179:
			n += NoInline179(i)
		case 180:
			n += NoInline180(i)
		case 181:
			n += NoInline181(i)
		case 182:
			n += NoInline182(i)
		case 183:
			n += NoInline183(i)
		case 184:
			n += NoInline184(i)
		case 185:
			n += NoInline185(i)
		case 186:
			n += NoInline186(i)
		case 187:
			n += NoInline187(i)
		case 188:
			n += NoInline188(i)
		case 189:
			n += NoInline189(i)
		case 190:
			n += NoInline190(i)
		case 191:
			n += NoInline191(i)
		case 192:
			n += NoInline192(i)
		case 193:
			n += NoInline193(i)
		case 194:
			n += NoInline194(i)
		case 195:
			n += NoInline195(i)
		case 196:
			n += NoInline196(i)
		case 197:
			n += NoInline197(i)
		case 198:
			n += NoInline198(i)
		case 199:
			n += NoInline199(i)
		case 200:
			n += NoInline200(i)
		case 201:
			n += NoInline201(i)
		case 202:
			n += NoInline202(i)
		case 203:
			n += NoInline203(i)
		case 204:
			n += NoInline204(i)
		case 205:
			n += NoInline205(i)
		case 206:
			n += NoInline206(i)
		case 207:
			n += NoInline207(i)
		case 208:
			n += NoInline208(i)
		case 209:
			n += NoInline209(i)
		case 210:
			n += NoInline210(i)
		case 211:
			n += NoInline211(i)
		case 212:
			n += NoInline212(i)
		case 213:
			n += NoInline213(i)
		case 214:
			n += NoInline214(i)
		case 215:
			n += NoInline215(i)
		case 216:
			n += NoInline216(i)
		case 217:
			n += NoInline217(i)
		case 218:
			n += NoInline218(i)
		case 219:
			n += NoInline219(i)
		case 220:
			n += NoInline220(i)
		case 221:
			n += NoInline221(i)
		case 222:
			n += NoInline222(i)
		case 223:
			n += NoInline223(i)
		case 224:
			n += NoInline224(i)
		case 225:
			n += NoInline225(i)
		case 226:
			n += NoInline226(i)
		case 227:
			n += NoInline227(i)
		case 228:
			n += NoInline228(i)
		case 229:
			n += NoInline229(i)
		case 230:
			n += NoInline230(i)
		case 231:
			n += NoInline231(i)
		case 232:
			n += NoInline232(i)
		case 233:
			n += NoInline233(i)
		case 234:
			n += NoInline234(i)
		case 235:
			n += NoInline235(i)
		case 236:
			n += NoInline236(i)
		case 237:
			n += NoInline237(i)
		case 238:
			n += NoInline238(i)
		case 239:
			n += NoInline239(i)
		case 240:
			n += NoInline240(i)
		case 241:
			n += NoInline241(i)
		case 242:
			n += NoInline242(i)
		case 243:
			n += NoInline243(i)
		case 244:
			n += NoInline244(i)
		case 245:
			n += NoInline245(i)
		case 246:
			n += NoInline246(i)
		case 247:
			n += NoInline247(i)
		case 248:
			n += NoInline248(i)
		case 249:
			n += NoInline249(i)
		case 250:
			n += NoInline250(i)
		case 251:
			n += NoInline251(i)
		case 252:
			n += NoInline252(i)
		case 253:
			n += NoInline253(i)
		case 254:
			n += NoInline254(i)
		case 255:
			n += NoInline255(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessMapNoInlineFunc256(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += NoInlineFuncs[randInputs[i%len(randInputs)]%256](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessSwitchMinimalFunc512(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch i % 512 {
		case 0:
			n += Minimal0(i)
		case 1:
			n += Minimal1(i)
		case 2:
			n += Minimal2(i)
		case 3:
			n += Minimal3(i)
		case 4:
			n += Minimal4(i)
		case 5:
			n += Minimal5(i)
		case 6:
			n += Minimal6(i)
		case 7:
			n += Minimal7(i)
		case 8:
			n += Minimal8(i)
		case 9:
			n += Minimal9(i)
		case 10:
			n += Minimal10(i)
		case 11:
			n += Minimal11(i)
		case 12:
			n += Minimal12(i)
		case 13:
			n += Minimal13(i)
		case 14:
			n += Minimal14(i)
		case 15:
			n += Minimal15(i)
		case 16:
			n += Minimal16(i)
		case 17:
			n += Minimal17(i)
		case 18:
			n += Minimal18(i)
		case 19:
			n += Minimal19(i)
		case 20:
			n += Minimal20(i)
		case 21:
			n += Minimal21(i)
		case 22:
			n += Minimal22(i)
		case 23:
			n += Minimal23(i)
		case 24:
			n += Minimal24(i)
		case 25:
			n += Minimal25(i)
		case 26:
			n += Minimal26(i)
		case 27:
			n += Minimal27(i)
		case 28:
			n += Minimal28(i)
		case 29:
			n += Minimal29(i)
		case 30:
			n += Minimal30(i)
		case 31:
			n += Minimal31(i)
		case 32:
			n += Minimal32(i)
		case 33:
			n += Minimal33(i)
		case 34:
			n += Minimal34(i)
		case 35:
			n += Minimal35(i)
		case 36:
			n += Minimal36(i)
		case 37:
			n += Minimal37(i)
		case 38:
			n += Minimal38(i)
		case 39:
			n += Minimal39(i)
		case 40:
			n += Minimal40(i)
		case 41:
			n += Minimal41(i)
		case 42:
			n += Minimal42(i)
		case 43:
			n += Minimal43(i)
		case 44:
			n += Minimal44(i)
		case 45:
			n += Minimal45(i)
		case 46:
			n += Minimal46(i)
		case 47:
			n += Minimal47(i)
		case 48:
			n += Minimal48(i)
		case 49:
			n += Minimal49(i)
		case 50:
			n += Minimal50(i)
		case 51:
			n += Minimal51(i)
		case 52:
			n += Minimal52(i)
		case 53:
			n += Minimal53(i)
		case 54:
			n += Minimal54(i)
		case 55:
			n += Minimal55(i)
		case 56:
			n += Minimal56(i)
		case 57:
			n += Minimal57(i)
		case 58:
			n += Minimal58(i)
		case 59:
			n += Minimal59(i)
		case 60:
			n += Minimal60(i)
		case 61:
			n += Minimal61(i)
		case 62:
			n += Minimal62(i)
		case 63:
			n += Minimal63(i)
		case 64:
			n += Minimal64(i)
		case 65:
			n += Minimal65(i)
		case 66:
			n += Minimal66(i)
		case 67:
			n += Minimal67(i)
		case 68:
			n += Minimal68(i)
		case 69:
			n += Minimal69(i)
		case 70:
			n += Minimal70(i)
		case 71:
			n += Minimal71(i)
		case 72:
			n += Minimal72(i)
		case 73:
			n += Minimal73(i)
		case 74:
			n += Minimal74(i)
		case 75:
			n += Minimal75(i)
		case 76:
			n += Minimal76(i)
		case 77:
			n += Minimal77(i)
		case 78:
			n += Minimal78(i)
		case 79:
			n += Minimal79(i)
		case 80:
			n += Minimal80(i)
		case 81:
			n += Minimal81(i)
		case 82:
			n += Minimal82(i)
		case 83:
			n += Minimal83(i)
		case 84:
			n += Minimal84(i)
		case 85:
			n += Minimal85(i)
		case 86:
			n += Minimal86(i)
		case 87:
			n += Minimal87(i)
		case 88:
			n += Minimal88(i)
		case 89:
			n += Minimal89(i)
		case 90:
			n += Minimal90(i)
		case 91:
			n += Minimal91(i)
		case 92:
			n += Minimal92(i)
		case 93:
			n += Minimal93(i)
		case 94:
			n += Minimal94(i)
		case 95:
			n += Minimal95(i)
		case 96:
			n += Minimal96(i)
		case 97:
			n += Minimal97(i)
		case 98:
			n += Minimal98(i)
		case 99:
			n += Minimal99(i)
		case 100:
			n += Minimal100(i)
		case 101:
			n += Minimal101(i)
		case 102:
			n += Minimal102(i)
		case 103:
			n += Minimal103(i)
		case 104:
			n += Minimal104(i)
		case 105:
			n += Minimal105(i)
		case 106:
			n += Minimal106(i)
		case 107:
			n += Minimal107(i)
		case 108:
			n += Minimal108(i)
		case 109:
			n += Minimal109(i)
		case 110:
			n += Minimal110(i)
		case 111:
			n += Minimal111(i)
		case 112:
			n += Minimal112(i)
		case 113:
			n += Minimal113(i)
		case 114:
			n += Minimal114(i)
		case 115:
			n += Minimal115(i)
		case 116:
			n += Minimal116(i)
		case 117:
			n += Minimal117(i)
		case 118:
			n += Minimal118(i)
		case 119:
			n += Minimal119(i)
		case 120:
			n += Minimal120(i)
		case 121:
			n += Minimal121(i)
		case 122:
			n += Minimal122(i)
		case 123:
			n += Minimal123(i)
		case 124:
			n += Minimal124(i)
		case 125:
			n += Minimal125(i)
		case 126:
			n += Minimal126(i)
		case 127:
			n += Minimal127(i)
		case 128:
			n += Minimal128(i)
		case 129:
			n += Minimal129(i)
		case 130:
			n += Minimal130(i)
		case 131:
			n += Minimal131(i)
		case 132:
			n += Minimal132(i)
		case 133:
			n += Minimal133(i)
		case 134:
			n += Minimal134(i)
		case 135:
			n += Minimal135(i)
		case 136:
			n += Minimal136(i)
		case 137:
			n += Minimal137(i)
		case 138:
			n += Minimal138(i)
		case 139:
			n += Minimal139(i)
		case 140:
			n += Minimal140(i)
		case 141:
			n += Minimal141(i)
		case 142:
			n += Minimal142(i)
		case 143:
			n += Minimal143(i)
		case 144:
			n += Minimal144(i)
		case 145:
			n += Minimal145(i)
		case 146:
			n += Minimal146(i)
		case 147:
			n += Minimal147(i)
		case 148:
			n += Minimal148(i)
		case 149:
			n += Minimal149(i)
		case 150:
			n += Minimal150(i)
		case 151:
			n += Minimal151(i)
		case 152:
			n += Minimal152(i)
		case 153:
			n += Minimal153(i)
		case 154:
			n += Minimal154(i)
		case 155:
			n += Minimal155(i)
		case 156:
			n += Minimal156(i)
		case 157:
			n += Minimal157(i)
		case 158:
			n += Minimal158(i)
		case 159:
			n += Minimal159(i)
		case 160:
			n += Minimal160(i)
		case 161:
			n += Minimal161(i)
		case 162:
			n += Minimal162(i)
		case 163:
			n += Minimal163(i)
		case 164:
			n += Minimal164(i)
		case 165:
			n += Minimal165(i)
		case 166:
			n += Minimal166(i)
		case 167:
			n += Minimal167(i)
		case 168:
			n += Minimal168(i)
		case 169:
			n += Minimal169(i)
		case 170:
			n += Minimal170(i)
		case 171:
			n += Minimal171(i)
		case 172:
			n += Minimal172(i)
		case 173:
			n += Minimal173(i)
		case 174:
			n += Minimal174(i)
		case 175:
			n += Minimal175(i)
		case 176:
			n += Minimal176(i)
		case 177:
			n += Minimal177(i)
		case 178:
			n += Minimal178(i)
		case 179:
			n += Minimal179(i)
		case 180:
			n += Minimal180(i)
		case 181:
			n += Minimal181(i)
		case 182:
			n += Minimal182(i)
		case 183:
			n += Minimal183(i)
		case 184:
			n += Minimal184(i)
		case 185:
			n += Minimal185(i)
		case 186:
			n += Minimal186(i)
		case 187:
			n += Minimal187(i)
		case 188:
			n += Minimal188(i)
		case 189:
			n += Minimal189(i)
		case 190:
			n += Minimal190(i)
		case 191:
			n += Minimal191(i)
		case 192:
			n += Minimal192(i)
		case 193:
			n += Minimal193(i)
		case 194:
			n += Minimal194(i)
		case 195:
			n += Minimal195(i)
		case 196:
			n += Minimal196(i)
		case 197:
			n += Minimal197(i)
		case 198:
			n += Minimal198(i)
		case 199:
			n += Minimal199(i)
		case 200:
			n += Minimal200(i)
		case 201:
			n += Minimal201(i)
		case 202:
			n += Minimal202(i)
		case 203:
			n += Minimal203(i)
		case 204:
			n += Minimal204(i)
		case 205:
			n += Minimal205(i)
		case 206:
			n += Minimal206(i)
		case 207:
			n += Minimal207(i)
		case 208:
			n += Minimal208(i)
		case 209:
			n += Minimal209(i)
		case 210:
			n += Minimal210(i)
		case 211:
			n += Minimal211(i)
		case 212:
			n += Minimal212(i)
		case 213:
			n += Minimal213(i)
		case 214:
			n += Minimal214(i)
		case 215:
			n += Minimal215(i)
		case 216:
			n += Minimal216(i)
		case 217:
			n += Minimal217(i)
		case 218:
			n += Minimal218(i)
		case 219:
			n += Minimal219(i)
		case 220:
			n += Minimal220(i)
		case 221:
			n += Minimal221(i)
		case 222:
			n += Minimal222(i)
		case 223:
			n += Minimal223(i)
		case 224:
			n += Minimal224(i)
		case 225:
			n += Minimal225(i)
		case 226:
			n += Minimal226(i)
		case 227:
			n += Minimal227(i)
		case 228:
			n += Minimal228(i)
		case 229:
			n += Minimal229(i)
		case 230:
			n += Minimal230(i)
		case 231:
			n += Minimal231(i)
		case 232:
			n += Minimal232(i)
		case 233:
			n += Minimal233(i)
		case 234:
			n += Minimal234(i)
		case 235:
			n += Minimal235(i)
		case 236:
			n += Minimal236(i)
		case 237:
			n += Minimal237(i)
		case 238:
			n += Minimal238(i)
		case 239:
			n += Minimal239(i)
		case 240:
			n += Minimal240(i)
		case 241:
			n += Minimal241(i)
		case 242:
			n += Minimal242(i)
		case 243:
			n += Minimal243(i)
		case 244:
			n += Minimal244(i)
		case 245:
			n += Minimal245(i)
		case 246:
			n += Minimal246(i)
		case 247:
			n += Minimal247(i)
		case 248:
			n += Minimal248(i)
		case 249:
			n += Minimal249(i)
		case 250:
			n += Minimal250(i)
		case 251:
			n += Minimal251(i)
		case 252:
			n += Minimal252(i)
		case 253:
			n += Minimal253(i)
		case 254:
			n += Minimal254(i)
		case 255:
			n += Minimal255(i)
		case 256:
			n += Minimal256(i)
		case 257:
			n += Minimal257(i)
		case 258:
			n += Minimal258(i)
		case 259:
			n += Minimal259(i)
		case 260:
			n += Minimal260(i)
		case 261:
			n += Minimal261(i)
		case 262:
			n += Minimal262(i)
		case 263:
			n += Minimal263(i)
		case 264:
			n += Minimal264(i)
		case 265:
			n += Minimal265(i)
		case 266:
			n += Minimal266(i)
		case 267:
			n += Minimal267(i)
		case 268:
			n += Minimal268(i)
		case 269:
			n += Minimal269(i)
		case 270:
			n += Minimal270(i)
		case 271:
			n += Minimal271(i)
		case 272:
			n += Minimal272(i)
		case 273:
			n += Minimal273(i)
		case 274:
			n += Minimal274(i)
		case 275:
			n += Minimal275(i)
		case 276:
			n += Minimal276(i)
		case 277:
			n += Minimal277(i)
		case 278:
			n += Minimal278(i)
		case 279:
			n += Minimal279(i)
		case 280:
			n += Minimal280(i)
		case 281:
			n += Minimal281(i)
		case 282:
			n += Minimal282(i)
		case 283:
			n += Minimal283(i)
		case 284:
			n += Minimal284(i)
		case 285:
			n += Minimal285(i)
		case 286:
			n += Minimal286(i)
		case 287:
			n += Minimal287(i)
		case 288:
			n += Minimal288(i)
		case 289:
			n += Minimal289(i)
		case 290:
			n += Minimal290(i)
		case 291:
			n += Minimal291(i)
		case 292:
			n += Minimal292(i)
		case 293:
			n += Minimal293(i)
		case 294:
			n += Minimal294(i)
		case 295:
			n += Minimal295(i)
		case 296:
			n += Minimal296(i)
		case 297:
			n += Minimal297(i)
		case 298:
			n += Minimal298(i)
		case 299:
			n += Minimal299(i)
		case 300:
			n += Minimal300(i)
		case 301:
			n += Minimal301(i)
		case 302:
			n += Minimal302(i)
		case 303:
			n += Minimal303(i)
		case 304:
			n += Minimal304(i)
		case 305:
			n += Minimal305(i)
		case 306:
			n += Minimal306(i)
		case 307:
			n += Minimal307(i)
		case 308:
			n += Minimal308(i)
		case 309:
			n += Minimal309(i)
		case 310:
			n += Minimal310(i)
		case 311:
			n += Minimal311(i)
		case 312:
			n += Minimal312(i)
		case 313:
			n += Minimal313(i)
		case 314:
			n += Minimal314(i)
		case 315:
			n += Minimal315(i)
		case 316:
			n += Minimal316(i)
		case 317:
			n += Minimal317(i)
		case 318:
			n += Minimal318(i)
		case 319:
			n += Minimal319(i)
		case 320:
			n += Minimal320(i)
		case 321:
			n += Minimal321(i)
		case 322:
			n += Minimal322(i)
		case 323:
			n += Minimal323(i)
		case 324:
			n += Minimal324(i)
		case 325:
			n += Minimal325(i)
		case 326:
			n += Minimal326(i)
		case 327:
			n += Minimal327(i)
		case 328:
			n += Minimal328(i)
		case 329:
			n += Minimal329(i)
		case 330:
			n += Minimal330(i)
		case 331:
			n += Minimal331(i)
		case 332:
			n += Minimal332(i)
		case 333:
			n += Minimal333(i)
		case 334:
			n += Minimal334(i)
		case 335:
			n += Minimal335(i)
		case 336:
			n += Minimal336(i)
		case 337:
			n += Minimal337(i)
		case 338:
			n += Minimal338(i)
		case 339:
			n += Minimal339(i)
		case 340:
			n += Minimal340(i)
		case 341:
			n += Minimal341(i)
		case 342:
			n += Minimal342(i)
		case 343:
			n += Minimal343(i)
		case 344:
			n += Minimal344(i)
		case 345:
			n += Minimal345(i)
		case 346:
			n += Minimal346(i)
		case 347:
			n += Minimal347(i)
		case 348:
			n += Minimal348(i)
		case 349:
			n += Minimal349(i)
		case 350:
			n += Minimal350(i)
		case 351:
			n += Minimal351(i)
		case 352:
			n += Minimal352(i)
		case 353:
			n += Minimal353(i)
		case 354:
			n += Minimal354(i)
		case 355:
			n += Minimal355(i)
		case 356:
			n += Minimal356(i)
		case 357:
			n += Minimal357(i)
		case 358:
			n += Minimal358(i)
		case 359:
			n += Minimal359(i)
		case 360:
			n += Minimal360(i)
		case 361:
			n += Minimal361(i)
		case 362:
			n += Minimal362(i)
		case 363:
			n += Minimal363(i)
		case 364:
			n += Minimal364(i)
		case 365:
			n += Minimal365(i)
		case 366:
			n += Minimal366(i)
		case 367:
			n += Minimal367(i)
		case 368:
			n += Minimal368(i)
		case 369:
			n += Minimal369(i)
		case 370:
			n += Minimal370(i)
		case 371:
			n += Minimal371(i)
		case 372:
			n += Minimal372(i)
		case 373:
			n += Minimal373(i)
		case 374:
			n += Minimal374(i)
		case 375:
			n += Minimal375(i)
		case 376:
			n += Minimal376(i)
		case 377:
			n += Minimal377(i)
		case 378:
			n += Minimal378(i)
		case 379:
			n += Minimal379(i)
		case 380:
			n += Minimal380(i)
		case 381:
			n += Minimal381(i)
		case 382:
			n += Minimal382(i)
		case 383:
			n += Minimal383(i)
		case 384:
			n += Minimal384(i)
		case 385:
			n += Minimal385(i)
		case 386:
			n += Minimal386(i)
		case 387:
			n += Minimal387(i)
		case 388:
			n += Minimal388(i)
		case 389:
			n += Minimal389(i)
		case 390:
			n += Minimal390(i)
		case 391:
			n += Minimal391(i)
		case 392:
			n += Minimal392(i)
		case 393:
			n += Minimal393(i)
		case 394:
			n += Minimal394(i)
		case 395:
			n += Minimal395(i)
		case 396:
			n += Minimal396(i)
		case 397:
			n += Minimal397(i)
		case 398:
			n += Minimal398(i)
		case 399:
			n += Minimal399(i)
		case 400:
			n += Minimal400(i)
		case 401:
			n += Minimal401(i)
		case 402:
			n += Minimal402(i)
		case 403:
			n += Minimal403(i)
		case 404:
			n += Minimal404(i)
		case 405:
			n += Minimal405(i)
		case 406:
			n += Minimal406(i)
		case 407:
			n += Minimal407(i)
		case 408:
			n += Minimal408(i)
		case 409:
			n += Minimal409(i)
		case 410:
			n += Minimal410(i)
		case 411:
			n += Minimal411(i)
		case 412:
			n += Minimal412(i)
		case 413:
			n += Minimal413(i)
		case 414:
			n += Minimal414(i)
		case 415:
			n += Minimal415(i)
		case 416:
			n += Minimal416(i)
		case 417:
			n += Minimal417(i)
		case 418:
			n += Minimal418(i)
		case 419:
			n += Minimal419(i)
		case 420:
			n += Minimal420(i)
		case 421:
			n += Minimal421(i)
		case 422:
			n += Minimal422(i)
		case 423:
			n += Minimal423(i)
		case 424:
			n += Minimal424(i)
		case 425:
			n += Minimal425(i)
		case 426:
			n += Minimal426(i)
		case 427:
			n += Minimal427(i)
		case 428:
			n += Minimal428(i)
		case 429:
			n += Minimal429(i)
		case 430:
			n += Minimal430(i)
		case 431:
			n += Minimal431(i)
		case 432:
			n += Minimal432(i)
		case 433:
			n += Minimal433(i)
		case 434:
			n += Minimal434(i)
		case 435:
			n += Minimal435(i)
		case 436:
			n += Minimal436(i)
		case 437:
			n += Minimal437(i)
		case 438:
			n += Minimal438(i)
		case 439:
			n += Minimal439(i)
		case 440:
			n += Minimal440(i)
		case 441:
			n += Minimal441(i)
		case 442:
			n += Minimal442(i)
		case 443:
			n += Minimal443(i)
		case 444:
			n += Minimal444(i)
		case 445:
			n += Minimal445(i)
		case 446:
			n += Minimal446(i)
		case 447:
			n += Minimal447(i)
		case 448:
			n += Minimal448(i)
		case 449:
			n += Minimal449(i)
		case 450:
			n += Minimal450(i)
		case 451:
			n += Minimal451(i)
		case 452:
			n += Minimal452(i)
		case 453:
			n += Minimal453(i)
		case 454:
			n += Minimal454(i)
		case 455:
			n += Minimal455(i)
		case 456:
			n += Minimal456(i)
		case 457:
			n += Minimal457(i)
		case 458:
			n += Minimal458(i)
		case 459:
			n += Minimal459(i)
		case 460:
			n += Minimal460(i)
		case 461:
			n += Minimal461(i)
		case 462:
			n += Minimal462(i)
		case 463:
			n += Minimal463(i)
		case 464:
			n += Minimal464(i)
		case 465:
			n += Minimal465(i)
		case 466:
			n += Minimal466(i)
		case 467:
			n += Minimal467(i)
		case 468:
			n += Minimal468(i)
		case 469:
			n += Minimal469(i)
		case 470:
			n += Minimal470(i)
		case 471:
			n += Minimal471(i)
		case 472:
			n += Minimal472(i)
		case 473:
			n += Minimal473(i)
		case 474:
			n += Minimal474(i)
		case 475:
			n += Minimal475(i)
		case 476:
			n += Minimal476(i)
		case 477:
			n += Minimal477(i)
		case 478:
			n += Minimal478(i)
		case 479:
			n += Minimal479(i)
		case 480:
			n += Minimal480(i)
		case 481:
			n += Minimal481(i)
		case 482:
			n += Minimal482(i)
		case 483:
			n += Minimal483(i)
		case 484:
			n += Minimal484(i)
		case 485:
			n += Minimal485(i)
		case 486:
			n += Minimal486(i)
		case 487:
			n += Minimal487(i)
		case 488:
			n += Minimal488(i)
		case 489:
			n += Minimal489(i)
		case 490:
			n += Minimal490(i)
		case 491:
			n += Minimal491(i)
		case 492:
			n += Minimal492(i)
		case 493:
			n += Minimal493(i)
		case 494:
			n += Minimal494(i)
		case 495:
			n += Minimal495(i)
		case 496:
			n += Minimal496(i)
		case 497:
			n += Minimal497(i)
		case 498:
			n += Minimal498(i)
		case 499:
			n += Minimal499(i)
		case 500:
			n += Minimal500(i)
		case 501:
			n += Minimal501(i)
		case 502:
			n += Minimal502(i)
		case 503:
			n += Minimal503(i)
		case 504:
			n += Minimal504(i)
		case 505:
			n += Minimal505(i)
		case 506:
			n += Minimal506(i)
		case 507:
			n += Minimal507(i)
		case 508:
			n += Minimal508(i)
		case 509:
			n += Minimal509(i)
		case 510:
			n += Minimal510(i)
		case 511:
			n += Minimal511(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessMapMinimalFunc512(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += MinimalFuncs[i%512](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessSwitchMinimalFunc512(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch ascInputs[i%len(ascInputs)] % 512 {
		case 0:
			n += Minimal0(i)
		case 1:
			n += Minimal1(i)
		case 2:
			n += Minimal2(i)
		case 3:
			n += Minimal3(i)
		case 4:
			n += Minimal4(i)
		case 5:
			n += Minimal5(i)
		case 6:
			n += Minimal6(i)
		case 7:
			n += Minimal7(i)
		case 8:
			n += Minimal8(i)
		case 9:
			n += Minimal9(i)
		case 10:
			n += Minimal10(i)
		case 11:
			n += Minimal11(i)
		case 12:
			n += Minimal12(i)
		case 13:
			n += Minimal13(i)
		case 14:
			n += Minimal14(i)
		case 15:
			n += Minimal15(i)
		case 16:
			n += Minimal16(i)
		case 17:
			n += Minimal17(i)
		case 18:
			n += Minimal18(i)
		case 19:
			n += Minimal19(i)
		case 20:
			n += Minimal20(i)
		case 21:
			n += Minimal21(i)
		case 22:
			n += Minimal22(i)
		case 23:
			n += Minimal23(i)
		case 24:
			n += Minimal24(i)
		case 25:
			n += Minimal25(i)
		case 26:
			n += Minimal26(i)
		case 27:
			n += Minimal27(i)
		case 28:
			n += Minimal28(i)
		case 29:
			n += Minimal29(i)
		case 30:
			n += Minimal30(i)
		case 31:
			n += Minimal31(i)
		case 32:
			n += Minimal32(i)
		case 33:
			n += Minimal33(i)
		case 34:
			n += Minimal34(i)
		case 35:
			n += Minimal35(i)
		case 36:
			n += Minimal36(i)
		case 37:
			n += Minimal37(i)
		case 38:
			n += Minimal38(i)
		case 39:
			n += Minimal39(i)
		case 40:
			n += Minimal40(i)
		case 41:
			n += Minimal41(i)
		case 42:
			n += Minimal42(i)
		case 43:
			n += Minimal43(i)
		case 44:
			n += Minimal44(i)
		case 45:
			n += Minimal45(i)
		case 46:
			n += Minimal46(i)
		case 47:
			n += Minimal47(i)
		case 48:
			n += Minimal48(i)
		case 49:
			n += Minimal49(i)
		case 50:
			n += Minimal50(i)
		case 51:
			n += Minimal51(i)
		case 52:
			n += Minimal52(i)
		case 53:
			n += Minimal53(i)
		case 54:
			n += Minimal54(i)
		case 55:
			n += Minimal55(i)
		case 56:
			n += Minimal56(i)
		case 57:
			n += Minimal57(i)
		case 58:
			n += Minimal58(i)
		case 59:
			n += Minimal59(i)
		case 60:
			n += Minimal60(i)
		case 61:
			n += Minimal61(i)
		case 62:
			n += Minimal62(i)
		case 63:
			n += Minimal63(i)
		case 64:
			n += Minimal64(i)
		case 65:
			n += Minimal65(i)
		case 66:
			n += Minimal66(i)
		case 67:
			n += Minimal67(i)
		case 68:
			n += Minimal68(i)
		case 69:
			n += Minimal69(i)
		case 70:
			n += Minimal70(i)
		case 71:
			n += Minimal71(i)
		case 72:
			n += Minimal72(i)
		case 73:
			n += Minimal73(i)
		case 74:
			n += Minimal74(i)
		case 75:
			n += Minimal75(i)
		case 76:
			n += Minimal76(i)
		case 77:
			n += Minimal77(i)
		case 78:
			n += Minimal78(i)
		case 79:
			n += Minimal79(i)
		case 80:
			n += Minimal80(i)
		case 81:
			n += Minimal81(i)
		case 82:
			n += Minimal82(i)
		case 83:
			n += Minimal83(i)
		case 84:
			n += Minimal84(i)
		case 85:
			n += Minimal85(i)
		case 86:
			n += Minimal86(i)
		case 87:
			n += Minimal87(i)
		case 88:
			n += Minimal88(i)
		case 89:
			n += Minimal89(i)
		case 90:
			n += Minimal90(i)
		case 91:
			n += Minimal91(i)
		case 92:
			n += Minimal92(i)
		case 93:
			n += Minimal93(i)
		case 94:
			n += Minimal94(i)
		case 95:
			n += Minimal95(i)
		case 96:
			n += Minimal96(i)
		case 97:
			n += Minimal97(i)
		case 98:
			n += Minimal98(i)
		case 99:
			n += Minimal99(i)
		case 100:
			n += Minimal100(i)
		case 101:
			n += Minimal101(i)
		case 102:
			n += Minimal102(i)
		case 103:
			n += Minimal103(i)
		case 104:
			n += Minimal104(i)
		case 105:
			n += Minimal105(i)
		case 106:
			n += Minimal106(i)
		case 107:
			n += Minimal107(i)
		case 108:
			n += Minimal108(i)
		case 109:
			n += Minimal109(i)
		case 110:
			n += Minimal110(i)
		case 111:
			n += Minimal111(i)
		case 112:
			n += Minimal112(i)
		case 113:
			n += Minimal113(i)
		case 114:
			n += Minimal114(i)
		case 115:
			n += Minimal115(i)
		case 116:
			n += Minimal116(i)
		case 117:
			n += Minimal117(i)
		case 118:
			n += Minimal118(i)
		case 119:
			n += Minimal119(i)
		case 120:
			n += Minimal120(i)
		case 121:
			n += Minimal121(i)
		case 122:
			n += Minimal122(i)
		case 123:
			n += Minimal123(i)
		case 124:
			n += Minimal124(i)
		case 125:
			n += Minimal125(i)
		case 126:
			n += Minimal126(i)
		case 127:
			n += Minimal127(i)
		case 128:
			n += Minimal128(i)
		case 129:
			n += Minimal129(i)
		case 130:
			n += Minimal130(i)
		case 131:
			n += Minimal131(i)
		case 132:
			n += Minimal132(i)
		case 133:
			n += Minimal133(i)
		case 134:
			n += Minimal134(i)
		case 135:
			n += Minimal135(i)
		case 136:
			n += Minimal136(i)
		case 137:
			n += Minimal137(i)
		case 138:
			n += Minimal138(i)
		case 139:
			n += Minimal139(i)
		case 140:
			n += Minimal140(i)
		case 141:
			n += Minimal141(i)
		case 142:
			n += Minimal142(i)
		case 143:
			n += Minimal143(i)
		case 144:
			n += Minimal144(i)
		case 145:
			n += Minimal145(i)
		case 146:
			n += Minimal146(i)
		case 147:
			n += Minimal147(i)
		case 148:
			n += Minimal148(i)
		case 149:
			n += Minimal149(i)
		case 150:
			n += Minimal150(i)
		case 151:
			n += Minimal151(i)
		case 152:
			n += Minimal152(i)
		case 153:
			n += Minimal153(i)
		case 154:
			n += Minimal154(i)
		case 155:
			n += Minimal155(i)
		case 156:
			n += Minimal156(i)
		case 157:
			n += Minimal157(i)
		case 158:
			n += Minimal158(i)
		case 159:
			n += Minimal159(i)
		case 160:
			n += Minimal160(i)
		case 161:
			n += Minimal161(i)
		case 162:
			n += Minimal162(i)
		case 163:
			n += Minimal163(i)
		case 164:
			n += Minimal164(i)
		case 165:
			n += Minimal165(i)
		case 166:
			n += Minimal166(i)
		case 167:
			n += Minimal167(i)
		case 168:
			n += Minimal168(i)
		case 169:
			n += Minimal169(i)
		case 170:
			n += Minimal170(i)
		case 171:
			n += Minimal171(i)
		case 172:
			n += Minimal172(i)
		case 173:
			n += Minimal173(i)
		case 174:
			n += Minimal174(i)
		case 175:
			n += Minimal175(i)
		case 176:
			n += Minimal176(i)
		case 177:
			n += Minimal177(i)
		case 178:
			n += Minimal178(i)
		case 179:
			n += Minimal179(i)
		case 180:
			n += Minimal180(i)
		case 181:
			n += Minimal181(i)
		case 182:
			n += Minimal182(i)
		case 183:
			n += Minimal183(i)
		case 184:
			n += Minimal184(i)
		case 185:
			n += Minimal185(i)
		case 186:
			n += Minimal186(i)
		case 187:
			n += Minimal187(i)
		case 188:
			n += Minimal188(i)
		case 189:
			n += Minimal189(i)
		case 190:
			n += Minimal190(i)
		case 191:
			n += Minimal191(i)
		case 192:
			n += Minimal192(i)
		case 193:
			n += Minimal193(i)
		case 194:
			n += Minimal194(i)
		case 195:
			n += Minimal195(i)
		case 196:
			n += Minimal196(i)
		case 197:
			n += Minimal197(i)
		case 198:
			n += Minimal198(i)
		case 199:
			n += Minimal199(i)
		case 200:
			n += Minimal200(i)
		case 201:
			n += Minimal201(i)
		case 202:
			n += Minimal202(i)
		case 203:
			n += Minimal203(i)
		case 204:
			n += Minimal204(i)
		case 205:
			n += Minimal205(i)
		case 206:
			n += Minimal206(i)
		case 207:
			n += Minimal207(i)
		case 208:
			n += Minimal208(i)
		case 209:
			n += Minimal209(i)
		case 210:
			n += Minimal210(i)
		case 211:
			n += Minimal211(i)
		case 212:
			n += Minimal212(i)
		case 213:
			n += Minimal213(i)
		case 214:
			n += Minimal214(i)
		case 215:
			n += Minimal215(i)
		case 216:
			n += Minimal216(i)
		case 217:
			n += Minimal217(i)
		case 218:
			n += Minimal218(i)
		case 219:
			n += Minimal219(i)
		case 220:
			n += Minimal220(i)
		case 221:
			n += Minimal221(i)
		case 222:
			n += Minimal222(i)
		case 223:
			n += Minimal223(i)
		case 224:
			n += Minimal224(i)
		case 225:
			n += Minimal225(i)
		case 226:
			n += Minimal226(i)
		case 227:
			n += Minimal227(i)
		case 228:
			n += Minimal228(i)
		case 229:
			n += Minimal229(i)
		case 230:
			n += Minimal230(i)
		case 231:
			n += Minimal231(i)
		case 232:
			n += Minimal232(i)
		case 233:
			n += Minimal233(i)
		case 234:
			n += Minimal234(i)
		case 235:
			n += Minimal235(i)
		case 236:
			n += Minimal236(i)
		case 237:
			n += Minimal237(i)
		case 238:
			n += Minimal238(i)
		case 239:
			n += Minimal239(i)
		case 240:
			n += Minimal240(i)
		case 241:
			n += Minimal241(i)
		case 242:
			n += Minimal242(i)
		case 243:
			n += Minimal243(i)
		case 244:
			n += Minimal244(i)
		case 245:
			n += Minimal245(i)
		case 246:
			n += Minimal246(i)
		case 247:
			n += Minimal247(i)
		case 248:
			n += Minimal248(i)
		case 249:
			n += Minimal249(i)
		case 250:
			n += Minimal250(i)
		case 251:
			n += Minimal251(i)
		case 252:
			n += Minimal252(i)
		case 253:
			n += Minimal253(i)
		case 254:
			n += Minimal254(i)
		case 255:
			n += Minimal255(i)
		case 256:
			n += Minimal256(i)
		case 257:
			n += Minimal257(i)
		case 258:
			n += Minimal258(i)
		case 259:
			n += Minimal259(i)
		case 260:
			n += Minimal260(i)
		case 261:
			n += Minimal261(i)
		case 262:
			n += Minimal262(i)
		case 263:
			n += Minimal263(i)
		case 264:
			n += Minimal264(i)
		case 265:
			n += Minimal265(i)
		case 266:
			n += Minimal266(i)
		case 267:
			n += Minimal267(i)
		case 268:
			n += Minimal268(i)
		case 269:
			n += Minimal269(i)
		case 270:
			n += Minimal270(i)
		case 271:
			n += Minimal271(i)
		case 272:
			n += Minimal272(i)
		case 273:
			n += Minimal273(i)
		case 274:
			n += Minimal274(i)
		case 275:
			n += Minimal275(i)
		case 276:
			n += Minimal276(i)
		case 277:
			n += Minimal277(i)
		case 278:
			n += Minimal278(i)
		case 279:
			n += Minimal279(i)
		case 280:
			n += Minimal280(i)
		case 281:
			n += Minimal281(i)
		case 282:
			n += Minimal282(i)
		case 283:
			n += Minimal283(i)
		case 284:
			n += Minimal284(i)
		case 285:
			n += Minimal285(i)
		case 286:
			n += Minimal286(i)
		case 287:
			n += Minimal287(i)
		case 288:
			n += Minimal288(i)
		case 289:
			n += Minimal289(i)
		case 290:
			n += Minimal290(i)
		case 291:
			n += Minimal291(i)
		case 292:
			n += Minimal292(i)
		case 293:
			n += Minimal293(i)
		case 294:
			n += Minimal294(i)
		case 295:
			n += Minimal295(i)
		case 296:
			n += Minimal296(i)
		case 297:
			n += Minimal297(i)
		case 298:
			n += Minimal298(i)
		case 299:
			n += Minimal299(i)
		case 300:
			n += Minimal300(i)
		case 301:
			n += Minimal301(i)
		case 302:
			n += Minimal302(i)
		case 303:
			n += Minimal303(i)
		case 304:
			n += Minimal304(i)
		case 305:
			n += Minimal305(i)
		case 306:
			n += Minimal306(i)
		case 307:
			n += Minimal307(i)
		case 308:
			n += Minimal308(i)
		case 309:
			n += Minimal309(i)
		case 310:
			n += Minimal310(i)
		case 311:
			n += Minimal311(i)
		case 312:
			n += Minimal312(i)
		case 313:
			n += Minimal313(i)
		case 314:
			n += Minimal314(i)
		case 315:
			n += Minimal315(i)
		case 316:
			n += Minimal316(i)
		case 317:
			n += Minimal317(i)
		case 318:
			n += Minimal318(i)
		case 319:
			n += Minimal319(i)
		case 320:
			n += Minimal320(i)
		case 321:
			n += Minimal321(i)
		case 322:
			n += Minimal322(i)
		case 323:
			n += Minimal323(i)
		case 324:
			n += Minimal324(i)
		case 325:
			n += Minimal325(i)
		case 326:
			n += Minimal326(i)
		case 327:
			n += Minimal327(i)
		case 328:
			n += Minimal328(i)
		case 329:
			n += Minimal329(i)
		case 330:
			n += Minimal330(i)
		case 331:
			n += Minimal331(i)
		case 332:
			n += Minimal332(i)
		case 333:
			n += Minimal333(i)
		case 334:
			n += Minimal334(i)
		case 335:
			n += Minimal335(i)
		case 336:
			n += Minimal336(i)
		case 337:
			n += Minimal337(i)
		case 338:
			n += Minimal338(i)
		case 339:
			n += Minimal339(i)
		case 340:
			n += Minimal340(i)
		case 341:
			n += Minimal341(i)
		case 342:
			n += Minimal342(i)
		case 343:
			n += Minimal343(i)
		case 344:
			n += Minimal344(i)
		case 345:
			n += Minimal345(i)
		case 346:
			n += Minimal346(i)
		case 347:
			n += Minimal347(i)
		case 348:
			n += Minimal348(i)
		case 349:
			n += Minimal349(i)
		case 350:
			n += Minimal350(i)
		case 351:
			n += Minimal351(i)
		case 352:
			n += Minimal352(i)
		case 353:
			n += Minimal353(i)
		case 354:
			n += Minimal354(i)
		case 355:
			n += Minimal355(i)
		case 356:
			n += Minimal356(i)
		case 357:
			n += Minimal357(i)
		case 358:
			n += Minimal358(i)
		case 359:
			n += Minimal359(i)
		case 360:
			n += Minimal360(i)
		case 361:
			n += Minimal361(i)
		case 362:
			n += Minimal362(i)
		case 363:
			n += Minimal363(i)
		case 364:
			n += Minimal364(i)
		case 365:
			n += Minimal365(i)
		case 366:
			n += Minimal366(i)
		case 367:
			n += Minimal367(i)
		case 368:
			n += Minimal368(i)
		case 369:
			n += Minimal369(i)
		case 370:
			n += Minimal370(i)
		case 371:
			n += Minimal371(i)
		case 372:
			n += Minimal372(i)
		case 373:
			n += Minimal373(i)
		case 374:
			n += Minimal374(i)
		case 375:
			n += Minimal375(i)
		case 376:
			n += Minimal376(i)
		case 377:
			n += Minimal377(i)
		case 378:
			n += Minimal378(i)
		case 379:
			n += Minimal379(i)
		case 380:
			n += Minimal380(i)
		case 381:
			n += Minimal381(i)
		case 382:
			n += Minimal382(i)
		case 383:
			n += Minimal383(i)
		case 384:
			n += Minimal384(i)
		case 385:
			n += Minimal385(i)
		case 386:
			n += Minimal386(i)
		case 387:
			n += Minimal387(i)
		case 388:
			n += Minimal388(i)
		case 389:
			n += Minimal389(i)
		case 390:
			n += Minimal390(i)
		case 391:
			n += Minimal391(i)
		case 392:
			n += Minimal392(i)
		case 393:
			n += Minimal393(i)
		case 394:
			n += Minimal394(i)
		case 395:
			n += Minimal395(i)
		case 396:
			n += Minimal396(i)
		case 397:
			n += Minimal397(i)
		case 398:
			n += Minimal398(i)
		case 399:
			n += Minimal399(i)
		case 400:
			n += Minimal400(i)
		case 401:
			n += Minimal401(i)
		case 402:
			n += Minimal402(i)
		case 403:
			n += Minimal403(i)
		case 404:
			n += Minimal404(i)
		case 405:
			n += Minimal405(i)
		case 406:
			n += Minimal406(i)
		case 407:
			n += Minimal407(i)
		case 408:
			n += Minimal408(i)
		case 409:
			n += Minimal409(i)
		case 410:
			n += Minimal410(i)
		case 411:
			n += Minimal411(i)
		case 412:
			n += Minimal412(i)
		case 413:
			n += Minimal413(i)
		case 414:
			n += Minimal414(i)
		case 415:
			n += Minimal415(i)
		case 416:
			n += Minimal416(i)
		case 417:
			n += Minimal417(i)
		case 418:
			n += Minimal418(i)
		case 419:
			n += Minimal419(i)
		case 420:
			n += Minimal420(i)
		case 421:
			n += Minimal421(i)
		case 422:
			n += Minimal422(i)
		case 423:
			n += Minimal423(i)
		case 424:
			n += Minimal424(i)
		case 425:
			n += Minimal425(i)
		case 426:
			n += Minimal426(i)
		case 427:
			n += Minimal427(i)
		case 428:
			n += Minimal428(i)
		case 429:
			n += Minimal429(i)
		case 430:
			n += Minimal430(i)
		case 431:
			n += Minimal431(i)
		case 432:
			n += Minimal432(i)
		case 433:
			n += Minimal433(i)
		case 434:
			n += Minimal434(i)
		case 435:
			n += Minimal435(i)
		case 436:
			n += Minimal436(i)
		case 437:
			n += Minimal437(i)
		case 438:
			n += Minimal438(i)
		case 439:
			n += Minimal439(i)
		case 440:
			n += Minimal440(i)
		case 441:
			n += Minimal441(i)
		case 442:
			n += Minimal442(i)
		case 443:
			n += Minimal443(i)
		case 444:
			n += Minimal444(i)
		case 445:
			n += Minimal445(i)
		case 446:
			n += Minimal446(i)
		case 447:
			n += Minimal447(i)
		case 448:
			n += Minimal448(i)
		case 449:
			n += Minimal449(i)
		case 450:
			n += Minimal450(i)
		case 451:
			n += Minimal451(i)
		case 452:
			n += Minimal452(i)
		case 453:
			n += Minimal453(i)
		case 454:
			n += Minimal454(i)
		case 455:
			n += Minimal455(i)
		case 456:
			n += Minimal456(i)
		case 457:
			n += Minimal457(i)
		case 458:
			n += Minimal458(i)
		case 459:
			n += Minimal459(i)
		case 460:
			n += Minimal460(i)
		case 461:
			n += Minimal461(i)
		case 462:
			n += Minimal462(i)
		case 463:
			n += Minimal463(i)
		case 464:
			n += Minimal464(i)
		case 465:
			n += Minimal465(i)
		case 466:
			n += Minimal466(i)
		case 467:
			n += Minimal467(i)
		case 468:
			n += Minimal468(i)
		case 469:
			n += Minimal469(i)
		case 470:
			n += Minimal470(i)
		case 471:
			n += Minimal471(i)
		case 472:
			n += Minimal472(i)
		case 473:
			n += Minimal473(i)
		case 474:
			n += Minimal474(i)
		case 475:
			n += Minimal475(i)
		case 476:
			n += Minimal476(i)
		case 477:
			n += Minimal477(i)
		case 478:
			n += Minimal478(i)
		case 479:
			n += Minimal479(i)
		case 480:
			n += Minimal480(i)
		case 481:
			n += Minimal481(i)
		case 482:
			n += Minimal482(i)
		case 483:
			n += Minimal483(i)
		case 484:
			n += Minimal484(i)
		case 485:
			n += Minimal485(i)
		case 486:
			n += Minimal486(i)
		case 487:
			n += Minimal487(i)
		case 488:
			n += Minimal488(i)
		case 489:
			n += Minimal489(i)
		case 490:
			n += Minimal490(i)
		case 491:
			n += Minimal491(i)
		case 492:
			n += Minimal492(i)
		case 493:
			n += Minimal493(i)
		case 494:
			n += Minimal494(i)
		case 495:
			n += Minimal495(i)
		case 496:
			n += Minimal496(i)
		case 497:
			n += Minimal497(i)
		case 498:
			n += Minimal498(i)
		case 499:
			n += Minimal499(i)
		case 500:
			n += Minimal500(i)
		case 501:
			n += Minimal501(i)
		case 502:
			n += Minimal502(i)
		case 503:
			n += Minimal503(i)
		case 504:
			n += Minimal504(i)
		case 505:
			n += Minimal505(i)
		case 506:
			n += Minimal506(i)
		case 507:
			n += Minimal507(i)
		case 508:
			n += Minimal508(i)
		case 509:
			n += Minimal509(i)
		case 510:
			n += Minimal510(i)
		case 511:
			n += Minimal511(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessMapMinimalFunc512(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += MinimalFuncs[ascInputs[i%len(ascInputs)]%512](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessSwitchMinimalFunc512(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch randInputs[i%len(randInputs)] % 512 {
		case 0:
			n += Minimal0(i)
		case 1:
			n += Minimal1(i)
		case 2:
			n += Minimal2(i)
		case 3:
			n += Minimal3(i)
		case 4:
			n += Minimal4(i)
		case 5:
			n += Minimal5(i)
		case 6:
			n += Minimal6(i)
		case 7:
			n += Minimal7(i)
		case 8:
			n += Minimal8(i)
		case 9:
			n += Minimal9(i)
		case 10:
			n += Minimal10(i)
		case 11:
			n += Minimal11(i)
		case 12:
			n += Minimal12(i)
		case 13:
			n += Minimal13(i)
		case 14:
			n += Minimal14(i)
		case 15:
			n += Minimal15(i)
		case 16:
			n += Minimal16(i)
		case 17:
			n += Minimal17(i)
		case 18:
			n += Minimal18(i)
		case 19:
			n += Minimal19(i)
		case 20:
			n += Minimal20(i)
		case 21:
			n += Minimal21(i)
		case 22:
			n += Minimal22(i)
		case 23:
			n += Minimal23(i)
		case 24:
			n += Minimal24(i)
		case 25:
			n += Minimal25(i)
		case 26:
			n += Minimal26(i)
		case 27:
			n += Minimal27(i)
		case 28:
			n += Minimal28(i)
		case 29:
			n += Minimal29(i)
		case 30:
			n += Minimal30(i)
		case 31:
			n += Minimal31(i)
		case 32:
			n += Minimal32(i)
		case 33:
			n += Minimal33(i)
		case 34:
			n += Minimal34(i)
		case 35:
			n += Minimal35(i)
		case 36:
			n += Minimal36(i)
		case 37:
			n += Minimal37(i)
		case 38:
			n += Minimal38(i)
		case 39:
			n += Minimal39(i)
		case 40:
			n += Minimal40(i)
		case 41:
			n += Minimal41(i)
		case 42:
			n += Minimal42(i)
		case 43:
			n += Minimal43(i)
		case 44:
			n += Minimal44(i)
		case 45:
			n += Minimal45(i)
		case 46:
			n += Minimal46(i)
		case 47:
			n += Minimal47(i)
		case 48:
			n += Minimal48(i)
		case 49:
			n += Minimal49(i)
		case 50:
			n += Minimal50(i)
		case 51:
			n += Minimal51(i)
		case 52:
			n += Minimal52(i)
		case 53:
			n += Minimal53(i)
		case 54:
			n += Minimal54(i)
		case 55:
			n += Minimal55(i)
		case 56:
			n += Minimal56(i)
		case 57:
			n += Minimal57(i)
		case 58:
			n += Minimal58(i)
		case 59:
			n += Minimal59(i)
		case 60:
			n += Minimal60(i)
		case 61:
			n += Minimal61(i)
		case 62:
			n += Minimal62(i)
		case 63:
			n += Minimal63(i)
		case 64:
			n += Minimal64(i)
		case 65:
			n += Minimal65(i)
		case 66:
			n += Minimal66(i)
		case 67:
			n += Minimal67(i)
		case 68:
			n += Minimal68(i)
		case 69:
			n += Minimal69(i)
		case 70:
			n += Minimal70(i)
		case 71:
			n += Minimal71(i)
		case 72:
			n += Minimal72(i)
		case 73:
			n += Minimal73(i)
		case 74:
			n += Minimal74(i)
		case 75:
			n += Minimal75(i)
		case 76:
			n += Minimal76(i)
		case 77:
			n += Minimal77(i)
		case 78:
			n += Minimal78(i)
		case 79:
			n += Minimal79(i)
		case 80:
			n += Minimal80(i)
		case 81:
			n += Minimal81(i)
		case 82:
			n += Minimal82(i)
		case 83:
			n += Minimal83(i)
		case 84:
			n += Minimal84(i)
		case 85:
			n += Minimal85(i)
		case 86:
			n += Minimal86(i)
		case 87:
			n += Minimal87(i)
		case 88:
			n += Minimal88(i)
		case 89:
			n += Minimal89(i)
		case 90:
			n += Minimal90(i)
		case 91:
			n += Minimal91(i)
		case 92:
			n += Minimal92(i)
		case 93:
			n += Minimal93(i)
		case 94:
			n += Minimal94(i)
		case 95:
			n += Minimal95(i)
		case 96:
			n += Minimal96(i)
		case 97:
			n += Minimal97(i)
		case 98:
			n += Minimal98(i)
		case 99:
			n += Minimal99(i)
		case 100:
			n += Minimal100(i)
		case 101:
			n += Minimal101(i)
		case 102:
			n += Minimal102(i)
		case 103:
			n += Minimal103(i)
		case 104:
			n += Minimal104(i)
		case 105:
			n += Minimal105(i)
		case 106:
			n += Minimal106(i)
		case 107:
			n += Minimal107(i)
		case 108:
			n += Minimal108(i)
		case 109:
			n += Minimal109(i)
		case 110:
			n += Minimal110(i)
		case 111:
			n += Minimal111(i)
		case 112:
			n += Minimal112(i)
		case 113:
			n += Minimal113(i)
		case 114:
			n += Minimal114(i)
		case 115:
			n += Minimal115(i)
		case 116:
			n += Minimal116(i)
		case 117:
			n += Minimal117(i)
		case 118:
			n += Minimal118(i)
		case 119:
			n += Minimal119(i)
		case 120:
			n += Minimal120(i)
		case 121:
			n += Minimal121(i)
		case 122:
			n += Minimal122(i)
		case 123:
			n += Minimal123(i)
		case 124:
			n += Minimal124(i)
		case 125:
			n += Minimal125(i)
		case 126:
			n += Minimal126(i)
		case 127:
			n += Minimal127(i)
		case 128:
			n += Minimal128(i)
		case 129:
			n += Minimal129(i)
		case 130:
			n += Minimal130(i)
		case 131:
			n += Minimal131(i)
		case 132:
			n += Minimal132(i)
		case 133:
			n += Minimal133(i)
		case 134:
			n += Minimal134(i)
		case 135:
			n += Minimal135(i)
		case 136:
			n += Minimal136(i)
		case 137:
			n += Minimal137(i)
		case 138:
			n += Minimal138(i)
		case 139:
			n += Minimal139(i)
		case 140:
			n += Minimal140(i)
		case 141:
			n += Minimal141(i)
		case 142:
			n += Minimal142(i)
		case 143:
			n += Minimal143(i)
		case 144:
			n += Minimal144(i)
		case 145:
			n += Minimal145(i)
		case 146:
			n += Minimal146(i)
		case 147:
			n += Minimal147(i)
		case 148:
			n += Minimal148(i)
		case 149:
			n += Minimal149(i)
		case 150:
			n += Minimal150(i)
		case 151:
			n += Minimal151(i)
		case 152:
			n += Minimal152(i)
		case 153:
			n += Minimal153(i)
		case 154:
			n += Minimal154(i)
		case 155:
			n += Minimal155(i)
		case 156:
			n += Minimal156(i)
		case 157:
			n += Minimal157(i)
		case 158:
			n += Minimal158(i)
		case 159:
			n += Minimal159(i)
		case 160:
			n += Minimal160(i)
		case 161:
			n += Minimal161(i)
		case 162:
			n += Minimal162(i)
		case 163:
			n += Minimal163(i)
		case 164:
			n += Minimal164(i)
		case 165:
			n += Minimal165(i)
		case 166:
			n += Minimal166(i)
		case 167:
			n += Minimal167(i)
		case 168:
			n += Minimal168(i)
		case 169:
			n += Minimal169(i)
		case 170:
			n += Minimal170(i)
		case 171:
			n += Minimal171(i)
		case 172:
			n += Minimal172(i)
		case 173:
			n += Minimal173(i)
		case 174:
			n += Minimal174(i)
		case 175:
			n += Minimal175(i)
		case 176:
			n += Minimal176(i)
		case 177:
			n += Minimal177(i)
		case 178:
			n += Minimal178(i)
		case 179:
			n += Minimal179(i)
		case 180:
			n += Minimal180(i)
		case 181:
			n += Minimal181(i)
		case 182:
			n += Minimal182(i)
		case 183:
			n += Minimal183(i)
		case 184:
			n += Minimal184(i)
		case 185:
			n += Minimal185(i)
		case 186:
			n += Minimal186(i)
		case 187:
			n += Minimal187(i)
		case 188:
			n += Minimal188(i)
		case 189:
			n += Minimal189(i)
		case 190:
			n += Minimal190(i)
		case 191:
			n += Minimal191(i)
		case 192:
			n += Minimal192(i)
		case 193:
			n += Minimal193(i)
		case 194:
			n += Minimal194(i)
		case 195:
			n += Minimal195(i)
		case 196:
			n += Minimal196(i)
		case 197:
			n += Minimal197(i)
		case 198:
			n += Minimal198(i)
		case 199:
			n += Minimal199(i)
		case 200:
			n += Minimal200(i)
		case 201:
			n += Minimal201(i)
		case 202:
			n += Minimal202(i)
		case 203:
			n += Minimal203(i)
		case 204:
			n += Minimal204(i)
		case 205:
			n += Minimal205(i)
		case 206:
			n += Minimal206(i)
		case 207:
			n += Minimal207(i)
		case 208:
			n += Minimal208(i)
		case 209:
			n += Minimal209(i)
		case 210:
			n += Minimal210(i)
		case 211:
			n += Minimal211(i)
		case 212:
			n += Minimal212(i)
		case 213:
			n += Minimal213(i)
		case 214:
			n += Minimal214(i)
		case 215:
			n += Minimal215(i)
		case 216:
			n += Minimal216(i)
		case 217:
			n += Minimal217(i)
		case 218:
			n += Minimal218(i)
		case 219:
			n += Minimal219(i)
		case 220:
			n += Minimal220(i)
		case 221:
			n += Minimal221(i)
		case 222:
			n += Minimal222(i)
		case 223:
			n += Minimal223(i)
		case 224:
			n += Minimal224(i)
		case 225:
			n += Minimal225(i)
		case 226:
			n += Minimal226(i)
		case 227:
			n += Minimal227(i)
		case 228:
			n += Minimal228(i)
		case 229:
			n += Minimal229(i)
		case 230:
			n += Minimal230(i)
		case 231:
			n += Minimal231(i)
		case 232:
			n += Minimal232(i)
		case 233:
			n += Minimal233(i)
		case 234:
			n += Minimal234(i)
		case 235:
			n += Minimal235(i)
		case 236:
			n += Minimal236(i)
		case 237:
			n += Minimal237(i)
		case 238:
			n += Minimal238(i)
		case 239:
			n += Minimal239(i)
		case 240:
			n += Minimal240(i)
		case 241:
			n += Minimal241(i)
		case 242:
			n += Minimal242(i)
		case 243:
			n += Minimal243(i)
		case 244:
			n += Minimal244(i)
		case 245:
			n += Minimal245(i)
		case 246:
			n += Minimal246(i)
		case 247:
			n += Minimal247(i)
		case 248:
			n += Minimal248(i)
		case 249:
			n += Minimal249(i)
		case 250:
			n += Minimal250(i)
		case 251:
			n += Minimal251(i)
		case 252:
			n += Minimal252(i)
		case 253:
			n += Minimal253(i)
		case 254:
			n += Minimal254(i)
		case 255:
			n += Minimal255(i)
		case 256:
			n += Minimal256(i)
		case 257:
			n += Minimal257(i)
		case 258:
			n += Minimal258(i)
		case 259:
			n += Minimal259(i)
		case 260:
			n += Minimal260(i)
		case 261:
			n += Minimal261(i)
		case 262:
			n += Minimal262(i)
		case 263:
			n += Minimal263(i)
		case 264:
			n += Minimal264(i)
		case 265:
			n += Minimal265(i)
		case 266:
			n += Minimal266(i)
		case 267:
			n += Minimal267(i)
		case 268:
			n += Minimal268(i)
		case 269:
			n += Minimal269(i)
		case 270:
			n += Minimal270(i)
		case 271:
			n += Minimal271(i)
		case 272:
			n += Minimal272(i)
		case 273:
			n += Minimal273(i)
		case 274:
			n += Minimal274(i)
		case 275:
			n += Minimal275(i)
		case 276:
			n += Minimal276(i)
		case 277:
			n += Minimal277(i)
		case 278:
			n += Minimal278(i)
		case 279:
			n += Minimal279(i)
		case 280:
			n += Minimal280(i)
		case 281:
			n += Minimal281(i)
		case 282:
			n += Minimal282(i)
		case 283:
			n += Minimal283(i)
		case 284:
			n += Minimal284(i)
		case 285:
			n += Minimal285(i)
		case 286:
			n += Minimal286(i)
		case 287:
			n += Minimal287(i)
		case 288:
			n += Minimal288(i)
		case 289:
			n += Minimal289(i)
		case 290:
			n += Minimal290(i)
		case 291:
			n += Minimal291(i)
		case 292:
			n += Minimal292(i)
		case 293:
			n += Minimal293(i)
		case 294:
			n += Minimal294(i)
		case 295:
			n += Minimal295(i)
		case 296:
			n += Minimal296(i)
		case 297:
			n += Minimal297(i)
		case 298:
			n += Minimal298(i)
		case 299:
			n += Minimal299(i)
		case 300:
			n += Minimal300(i)
		case 301:
			n += Minimal301(i)
		case 302:
			n += Minimal302(i)
		case 303:
			n += Minimal303(i)
		case 304:
			n += Minimal304(i)
		case 305:
			n += Minimal305(i)
		case 306:
			n += Minimal306(i)
		case 307:
			n += Minimal307(i)
		case 308:
			n += Minimal308(i)
		case 309:
			n += Minimal309(i)
		case 310:
			n += Minimal310(i)
		case 311:
			n += Minimal311(i)
		case 312:
			n += Minimal312(i)
		case 313:
			n += Minimal313(i)
		case 314:
			n += Minimal314(i)
		case 315:
			n += Minimal315(i)
		case 316:
			n += Minimal316(i)
		case 317:
			n += Minimal317(i)
		case 318:
			n += Minimal318(i)
		case 319:
			n += Minimal319(i)
		case 320:
			n += Minimal320(i)
		case 321:
			n += Minimal321(i)
		case 322:
			n += Minimal322(i)
		case 323:
			n += Minimal323(i)
		case 324:
			n += Minimal324(i)
		case 325:
			n += Minimal325(i)
		case 326:
			n += Minimal326(i)
		case 327:
			n += Minimal327(i)
		case 328:
			n += Minimal328(i)
		case 329:
			n += Minimal329(i)
		case 330:
			n += Minimal330(i)
		case 331:
			n += Minimal331(i)
		case 332:
			n += Minimal332(i)
		case 333:
			n += Minimal333(i)
		case 334:
			n += Minimal334(i)
		case 335:
			n += Minimal335(i)
		case 336:
			n += Minimal336(i)
		case 337:
			n += Minimal337(i)
		case 338:
			n += Minimal338(i)
		case 339:
			n += Minimal339(i)
		case 340:
			n += Minimal340(i)
		case 341:
			n += Minimal341(i)
		case 342:
			n += Minimal342(i)
		case 343:
			n += Minimal343(i)
		case 344:
			n += Minimal344(i)
		case 345:
			n += Minimal345(i)
		case 346:
			n += Minimal346(i)
		case 347:
			n += Minimal347(i)
		case 348:
			n += Minimal348(i)
		case 349:
			n += Minimal349(i)
		case 350:
			n += Minimal350(i)
		case 351:
			n += Minimal351(i)
		case 352:
			n += Minimal352(i)
		case 353:
			n += Minimal353(i)
		case 354:
			n += Minimal354(i)
		case 355:
			n += Minimal355(i)
		case 356:
			n += Minimal356(i)
		case 357:
			n += Minimal357(i)
		case 358:
			n += Minimal358(i)
		case 359:
			n += Minimal359(i)
		case 360:
			n += Minimal360(i)
		case 361:
			n += Minimal361(i)
		case 362:
			n += Minimal362(i)
		case 363:
			n += Minimal363(i)
		case 364:
			n += Minimal364(i)
		case 365:
			n += Minimal365(i)
		case 366:
			n += Minimal366(i)
		case 367:
			n += Minimal367(i)
		case 368:
			n += Minimal368(i)
		case 369:
			n += Minimal369(i)
		case 370:
			n += Minimal370(i)
		case 371:
			n += Minimal371(i)
		case 372:
			n += Minimal372(i)
		case 373:
			n += Minimal373(i)
		case 374:
			n += Minimal374(i)
		case 375:
			n += Minimal375(i)
		case 376:
			n += Minimal376(i)
		case 377:
			n += Minimal377(i)
		case 378:
			n += Minimal378(i)
		case 379:
			n += Minimal379(i)
		case 380:
			n += Minimal380(i)
		case 381:
			n += Minimal381(i)
		case 382:
			n += Minimal382(i)
		case 383:
			n += Minimal383(i)
		case 384:
			n += Minimal384(i)
		case 385:
			n += Minimal385(i)
		case 386:
			n += Minimal386(i)
		case 387:
			n += Minimal387(i)
		case 388:
			n += Minimal388(i)
		case 389:
			n += Minimal389(i)
		case 390:
			n += Minimal390(i)
		case 391:
			n += Minimal391(i)
		case 392:
			n += Minimal392(i)
		case 393:
			n += Minimal393(i)
		case 394:
			n += Minimal394(i)
		case 395:
			n += Minimal395(i)
		case 396:
			n += Minimal396(i)
		case 397:
			n += Minimal397(i)
		case 398:
			n += Minimal398(i)
		case 399:
			n += Minimal399(i)
		case 400:
			n += Minimal400(i)
		case 401:
			n += Minimal401(i)
		case 402:
			n += Minimal402(i)
		case 403:
			n += Minimal403(i)
		case 404:
			n += Minimal404(i)
		case 405:
			n += Minimal405(i)
		case 406:
			n += Minimal406(i)
		case 407:
			n += Minimal407(i)
		case 408:
			n += Minimal408(i)
		case 409:
			n += Minimal409(i)
		case 410:
			n += Minimal410(i)
		case 411:
			n += Minimal411(i)
		case 412:
			n += Minimal412(i)
		case 413:
			n += Minimal413(i)
		case 414:
			n += Minimal414(i)
		case 415:
			n += Minimal415(i)
		case 416:
			n += Minimal416(i)
		case 417:
			n += Minimal417(i)
		case 418:
			n += Minimal418(i)
		case 419:
			n += Minimal419(i)
		case 420:
			n += Minimal420(i)
		case 421:
			n += Minimal421(i)
		case 422:
			n += Minimal422(i)
		case 423:
			n += Minimal423(i)
		case 424:
			n += Minimal424(i)
		case 425:
			n += Minimal425(i)
		case 426:
			n += Minimal426(i)
		case 427:
			n += Minimal427(i)
		case 428:
			n += Minimal428(i)
		case 429:
			n += Minimal429(i)
		case 430:
			n += Minimal430(i)
		case 431:
			n += Minimal431(i)
		case 432:
			n += Minimal432(i)
		case 433:
			n += Minimal433(i)
		case 434:
			n += Minimal434(i)
		case 435:
			n += Minimal435(i)
		case 436:
			n += Minimal436(i)
		case 437:
			n += Minimal437(i)
		case 438:
			n += Minimal438(i)
		case 439:
			n += Minimal439(i)
		case 440:
			n += Minimal440(i)
		case 441:
			n += Minimal441(i)
		case 442:
			n += Minimal442(i)
		case 443:
			n += Minimal443(i)
		case 444:
			n += Minimal444(i)
		case 445:
			n += Minimal445(i)
		case 446:
			n += Minimal446(i)
		case 447:
			n += Minimal447(i)
		case 448:
			n += Minimal448(i)
		case 449:
			n += Minimal449(i)
		case 450:
			n += Minimal450(i)
		case 451:
			n += Minimal451(i)
		case 452:
			n += Minimal452(i)
		case 453:
			n += Minimal453(i)
		case 454:
			n += Minimal454(i)
		case 455:
			n += Minimal455(i)
		case 456:
			n += Minimal456(i)
		case 457:
			n += Minimal457(i)
		case 458:
			n += Minimal458(i)
		case 459:
			n += Minimal459(i)
		case 460:
			n += Minimal460(i)
		case 461:
			n += Minimal461(i)
		case 462:
			n += Minimal462(i)
		case 463:
			n += Minimal463(i)
		case 464:
			n += Minimal464(i)
		case 465:
			n += Minimal465(i)
		case 466:
			n += Minimal466(i)
		case 467:
			n += Minimal467(i)
		case 468:
			n += Minimal468(i)
		case 469:
			n += Minimal469(i)
		case 470:
			n += Minimal470(i)
		case 471:
			n += Minimal471(i)
		case 472:
			n += Minimal472(i)
		case 473:
			n += Minimal473(i)
		case 474:
			n += Minimal474(i)
		case 475:
			n += Minimal475(i)
		case 476:
			n += Minimal476(i)
		case 477:
			n += Minimal477(i)
		case 478:
			n += Minimal478(i)
		case 479:
			n += Minimal479(i)
		case 480:
			n += Minimal480(i)
		case 481:
			n += Minimal481(i)
		case 482:
			n += Minimal482(i)
		case 483:
			n += Minimal483(i)
		case 484:
			n += Minimal484(i)
		case 485:
			n += Minimal485(i)
		case 486:
			n += Minimal486(i)
		case 487:
			n += Minimal487(i)
		case 488:
			n += Minimal488(i)
		case 489:
			n += Minimal489(i)
		case 490:
			n += Minimal490(i)
		case 491:
			n += Minimal491(i)
		case 492:
			n += Minimal492(i)
		case 493:
			n += Minimal493(i)
		case 494:
			n += Minimal494(i)
		case 495:
			n += Minimal495(i)
		case 496:
			n += Minimal496(i)
		case 497:
			n += Minimal497(i)
		case 498:
			n += Minimal498(i)
		case 499:
			n += Minimal499(i)
		case 500:
			n += Minimal500(i)
		case 501:
			n += Minimal501(i)
		case 502:
			n += Minimal502(i)
		case 503:
			n += Minimal503(i)
		case 504:
			n += Minimal504(i)
		case 505:
			n += Minimal505(i)
		case 506:
			n += Minimal506(i)
		case 507:
			n += Minimal507(i)
		case 508:
			n += Minimal508(i)
		case 509:
			n += Minimal509(i)
		case 510:
			n += Minimal510(i)
		case 511:
			n += Minimal511(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessMapMinimalFunc512(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += MinimalFuncs[randInputs[i%len(randInputs)]%512](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessSwitchNoInlineFunc512(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch i % 512 {
		case 0:
			n += NoInline0(i)
		case 1:
			n += NoInline1(i)
		case 2:
			n += NoInline2(i)
		case 3:
			n += NoInline3(i)
		case 4:
			n += NoInline4(i)
		case 5:
			n += NoInline5(i)
		case 6:
			n += NoInline6(i)
		case 7:
			n += NoInline7(i)
		case 8:
			n += NoInline8(i)
		case 9:
			n += NoInline9(i)
		case 10:
			n += NoInline10(i)
		case 11:
			n += NoInline11(i)
		case 12:
			n += NoInline12(i)
		case 13:
			n += NoInline13(i)
		case 14:
			n += NoInline14(i)
		case 15:
			n += NoInline15(i)
		case 16:
			n += NoInline16(i)
		case 17:
			n += NoInline17(i)
		case 18:
			n += NoInline18(i)
		case 19:
			n += NoInline19(i)
		case 20:
			n += NoInline20(i)
		case 21:
			n += NoInline21(i)
		case 22:
			n += NoInline22(i)
		case 23:
			n += NoInline23(i)
		case 24:
			n += NoInline24(i)
		case 25:
			n += NoInline25(i)
		case 26:
			n += NoInline26(i)
		case 27:
			n += NoInline27(i)
		case 28:
			n += NoInline28(i)
		case 29:
			n += NoInline29(i)
		case 30:
			n += NoInline30(i)
		case 31:
			n += NoInline31(i)
		case 32:
			n += NoInline32(i)
		case 33:
			n += NoInline33(i)
		case 34:
			n += NoInline34(i)
		case 35:
			n += NoInline35(i)
		case 36:
			n += NoInline36(i)
		case 37:
			n += NoInline37(i)
		case 38:
			n += NoInline38(i)
		case 39:
			n += NoInline39(i)
		case 40:
			n += NoInline40(i)
		case 41:
			n += NoInline41(i)
		case 42:
			n += NoInline42(i)
		case 43:
			n += NoInline43(i)
		case 44:
			n += NoInline44(i)
		case 45:
			n += NoInline45(i)
		case 46:
			n += NoInline46(i)
		case 47:
			n += NoInline47(i)
		case 48:
			n += NoInline48(i)
		case 49:
			n += NoInline49(i)
		case 50:
			n += NoInline50(i)
		case 51:
			n += NoInline51(i)
		case 52:
			n += NoInline52(i)
		case 53:
			n += NoInline53(i)
		case 54:
			n += NoInline54(i)
		case 55:
			n += NoInline55(i)
		case 56:
			n += NoInline56(i)
		case 57:
			n += NoInline57(i)
		case 58:
			n += NoInline58(i)
		case 59:
			n += NoInline59(i)
		case 60:
			n += NoInline60(i)
		case 61:
			n += NoInline61(i)
		case 62:
			n += NoInline62(i)
		case 63:
			n += NoInline63(i)
		case 64:
			n += NoInline64(i)
		case 65:
			n += NoInline65(i)
		case 66:
			n += NoInline66(i)
		case 67:
			n += NoInline67(i)
		case 68:
			n += NoInline68(i)
		case 69:
			n += NoInline69(i)
		case 70:
			n += NoInline70(i)
		case 71:
			n += NoInline71(i)
		case 72:
			n += NoInline72(i)
		case 73:
			n += NoInline73(i)
		case 74:
			n += NoInline74(i)
		case 75:
			n += NoInline75(i)
		case 76:
			n += NoInline76(i)
		case 77:
			n += NoInline77(i)
		case 78:
			n += NoInline78(i)
		case 79:
			n += NoInline79(i)
		case 80:
			n += NoInline80(i)
		case 81:
			n += NoInline81(i)
		case 82:
			n += NoInline82(i)
		case 83:
			n += NoInline83(i)
		case 84:
			n += NoInline84(i)
		case 85:
			n += NoInline85(i)
		case 86:
			n += NoInline86(i)
		case 87:
			n += NoInline87(i)
		case 88:
			n += NoInline88(i)
		case 89:
			n += NoInline89(i)
		case 90:
			n += NoInline90(i)
		case 91:
			n += NoInline91(i)
		case 92:
			n += NoInline92(i)
		case 93:
			n += NoInline93(i)
		case 94:
			n += NoInline94(i)
		case 95:
			n += NoInline95(i)
		case 96:
			n += NoInline96(i)
		case 97:
			n += NoInline97(i)
		case 98:
			n += NoInline98(i)
		case 99:
			n += NoInline99(i)
		case 100:
			n += NoInline100(i)
		case 101:
			n += NoInline101(i)
		case 102:
			n += NoInline102(i)
		case 103:
			n += NoInline103(i)
		case 104:
			n += NoInline104(i)
		case 105:
			n += NoInline105(i)
		case 106:
			n += NoInline106(i)
		case 107:
			n += NoInline107(i)
		case 108:
			n += NoInline108(i)
		case 109:
			n += NoInline109(i)
		case 110:
			n += NoInline110(i)
		case 111:
			n += NoInline111(i)
		case 112:
			n += NoInline112(i)
		case 113:
			n += NoInline113(i)
		case 114:
			n += NoInline114(i)
		case 115:
			n += NoInline115(i)
		case 116:
			n += NoInline116(i)
		case 117:
			n += NoInline117(i)
		case 118:
			n += NoInline118(i)
		case 119:
			n += NoInline119(i)
		case 120:
			n += NoInline120(i)
		case 121:
			n += NoInline121(i)
		case 122:
			n += NoInline122(i)
		case 123:
			n += NoInline123(i)
		case 124:
			n += NoInline124(i)
		case 125:
			n += NoInline125(i)
		case 126:
			n += NoInline126(i)
		case 127:
			n += NoInline127(i)
		case 128:
			n += NoInline128(i)
		case 129:
			n += NoInline129(i)
		case 130:
			n += NoInline130(i)
		case 131:
			n += NoInline131(i)
		case 132:
			n += NoInline132(i)
		case 133:
			n += NoInline133(i)
		case 134:
			n += NoInline134(i)
		case 135:
			n += NoInline135(i)
		case 136:
			n += NoInline136(i)
		case 137:
			n += NoInline137(i)
		case 138:
			n += NoInline138(i)
		case 139:
			n += NoInline139(i)
		case 140:
			n += NoInline140(i)
		case 141:
			n += NoInline141(i)
		case 142:
			n += NoInline142(i)
		case 143:
			n += NoInline143(i)
		case 144:
			n += NoInline144(i)
		case 145:
			n += NoInline145(i)
		case 146:
			n += NoInline146(i)
		case 147:
			n += NoInline147(i)
		case 148:
			n += NoInline148(i)
		case 149:
			n += NoInline149(i)
		case 150:
			n += NoInline150(i)
		case 151:
			n += NoInline151(i)
		case 152:
			n += NoInline152(i)
		case 153:
			n += NoInline153(i)
		case 154:
			n += NoInline154(i)
		case 155:
			n += NoInline155(i)
		case 156:
			n += NoInline156(i)
		case 157:
			n += NoInline157(i)
		case 158:
			n += NoInline158(i)
		case 159:
			n += NoInline159(i)
		case 160:
			n += NoInline160(i)
		case 161:
			n += NoInline161(i)
		case 162:
			n += NoInline162(i)
		case 163:
			n += NoInline163(i)
		case 164:
			n += NoInline164(i)
		case 165:
			n += NoInline165(i)
		case 166:
			n += NoInline166(i)
		case 167:
			n += NoInline167(i)
		case 168:
			n += NoInline168(i)
		case 169:
			n += NoInline169(i)
		case 170:
			n += NoInline170(i)
		case 171:
			n += NoInline171(i)
		case 172:
			n += NoInline172(i)
		case 173:
			n += NoInline173(i)
		case 174:
			n += NoInline174(i)
		case 175:
			n += NoInline175(i)
		case 176:
			n += NoInline176(i)
		case 177:
			n += NoInline177(i)
		case 178:
			n += NoInline178(i)
		case 179:
			n += NoInline179(i)
		case 180:
			n += NoInline180(i)
		case 181:
			n += NoInline181(i)
		case 182:
			n += NoInline182(i)
		case 183:
			n += NoInline183(i)
		case 184:
			n += NoInline184(i)
		case 185:
			n += NoInline185(i)
		case 186:
			n += NoInline186(i)
		case 187:
			n += NoInline187(i)
		case 188:
			n += NoInline188(i)
		case 189:
			n += NoInline189(i)
		case 190:
			n += NoInline190(i)
		case 191:
			n += NoInline191(i)
		case 192:
			n += NoInline192(i)
		case 193:
			n += NoInline193(i)
		case 194:
			n += NoInline194(i)
		case 195:
			n += NoInline195(i)
		case 196:
			n += NoInline196(i)
		case 197:
			n += NoInline197(i)
		case 198:
			n += NoInline198(i)
		case 199:
			n += NoInline199(i)
		case 200:
			n += NoInline200(i)
		case 201:
			n += NoInline201(i)
		case 202:
			n += NoInline202(i)
		case 203:
			n += NoInline203(i)
		case 204:
			n += NoInline204(i)
		case 205:
			n += NoInline205(i)
		case 206:
			n += NoInline206(i)
		case 207:
			n += NoInline207(i)
		case 208:
			n += NoInline208(i)
		case 209:
			n += NoInline209(i)
		case 210:
			n += NoInline210(i)
		case 211:
			n += NoInline211(i)
		case 212:
			n += NoInline212(i)
		case 213:
			n += NoInline213(i)
		case 214:
			n += NoInline214(i)
		case 215:
			n += NoInline215(i)
		case 216:
			n += NoInline216(i)
		case 217:
			n += NoInline217(i)
		case 218:
			n += NoInline218(i)
		case 219:
			n += NoInline219(i)
		case 220:
			n += NoInline220(i)
		case 221:
			n += NoInline221(i)
		case 222:
			n += NoInline222(i)
		case 223:
			n += NoInline223(i)
		case 224:
			n += NoInline224(i)
		case 225:
			n += NoInline225(i)
		case 226:
			n += NoInline226(i)
		case 227:
			n += NoInline227(i)
		case 228:
			n += NoInline228(i)
		case 229:
			n += NoInline229(i)
		case 230:
			n += NoInline230(i)
		case 231:
			n += NoInline231(i)
		case 232:
			n += NoInline232(i)
		case 233:
			n += NoInline233(i)
		case 234:
			n += NoInline234(i)
		case 235:
			n += NoInline235(i)
		case 236:
			n += NoInline236(i)
		case 237:
			n += NoInline237(i)
		case 238:
			n += NoInline238(i)
		case 239:
			n += NoInline239(i)
		case 240:
			n += NoInline240(i)
		case 241:
			n += NoInline241(i)
		case 242:
			n += NoInline242(i)
		case 243:
			n += NoInline243(i)
		case 244:
			n += NoInline244(i)
		case 245:
			n += NoInline245(i)
		case 246:
			n += NoInline246(i)
		case 247:
			n += NoInline247(i)
		case 248:
			n += NoInline248(i)
		case 249:
			n += NoInline249(i)
		case 250:
			n += NoInline250(i)
		case 251:
			n += NoInline251(i)
		case 252:
			n += NoInline252(i)
		case 253:
			n += NoInline253(i)
		case 254:
			n += NoInline254(i)
		case 255:
			n += NoInline255(i)
		case 256:
			n += NoInline256(i)
		case 257:
			n += NoInline257(i)
		case 258:
			n += NoInline258(i)
		case 259:
			n += NoInline259(i)
		case 260:
			n += NoInline260(i)
		case 261:
			n += NoInline261(i)
		case 262:
			n += NoInline262(i)
		case 263:
			n += NoInline263(i)
		case 264:
			n += NoInline264(i)
		case 265:
			n += NoInline265(i)
		case 266:
			n += NoInline266(i)
		case 267:
			n += NoInline267(i)
		case 268:
			n += NoInline268(i)
		case 269:
			n += NoInline269(i)
		case 270:
			n += NoInline270(i)
		case 271:
			n += NoInline271(i)
		case 272:
			n += NoInline272(i)
		case 273:
			n += NoInline273(i)
		case 274:
			n += NoInline274(i)
		case 275:
			n += NoInline275(i)
		case 276:
			n += NoInline276(i)
		case 277:
			n += NoInline277(i)
		case 278:
			n += NoInline278(i)
		case 279:
			n += NoInline279(i)
		case 280:
			n += NoInline280(i)
		case 281:
			n += NoInline281(i)
		case 282:
			n += NoInline282(i)
		case 283:
			n += NoInline283(i)
		case 284:
			n += NoInline284(i)
		case 285:
			n += NoInline285(i)
		case 286:
			n += NoInline286(i)
		case 287:
			n += NoInline287(i)
		case 288:
			n += NoInline288(i)
		case 289:
			n += NoInline289(i)
		case 290:
			n += NoInline290(i)
		case 291:
			n += NoInline291(i)
		case 292:
			n += NoInline292(i)
		case 293:
			n += NoInline293(i)
		case 294:
			n += NoInline294(i)
		case 295:
			n += NoInline295(i)
		case 296:
			n += NoInline296(i)
		case 297:
			n += NoInline297(i)
		case 298:
			n += NoInline298(i)
		case 299:
			n += NoInline299(i)
		case 300:
			n += NoInline300(i)
		case 301:
			n += NoInline301(i)
		case 302:
			n += NoInline302(i)
		case 303:
			n += NoInline303(i)
		case 304:
			n += NoInline304(i)
		case 305:
			n += NoInline305(i)
		case 306:
			n += NoInline306(i)
		case 307:
			n += NoInline307(i)
		case 308:
			n += NoInline308(i)
		case 309:
			n += NoInline309(i)
		case 310:
			n += NoInline310(i)
		case 311:
			n += NoInline311(i)
		case 312:
			n += NoInline312(i)
		case 313:
			n += NoInline313(i)
		case 314:
			n += NoInline314(i)
		case 315:
			n += NoInline315(i)
		case 316:
			n += NoInline316(i)
		case 317:
			n += NoInline317(i)
		case 318:
			n += NoInline318(i)
		case 319:
			n += NoInline319(i)
		case 320:
			n += NoInline320(i)
		case 321:
			n += NoInline321(i)
		case 322:
			n += NoInline322(i)
		case 323:
			n += NoInline323(i)
		case 324:
			n += NoInline324(i)
		case 325:
			n += NoInline325(i)
		case 326:
			n += NoInline326(i)
		case 327:
			n += NoInline327(i)
		case 328:
			n += NoInline328(i)
		case 329:
			n += NoInline329(i)
		case 330:
			n += NoInline330(i)
		case 331:
			n += NoInline331(i)
		case 332:
			n += NoInline332(i)
		case 333:
			n += NoInline333(i)
		case 334:
			n += NoInline334(i)
		case 335:
			n += NoInline335(i)
		case 336:
			n += NoInline336(i)
		case 337:
			n += NoInline337(i)
		case 338:
			n += NoInline338(i)
		case 339:
			n += NoInline339(i)
		case 340:
			n += NoInline340(i)
		case 341:
			n += NoInline341(i)
		case 342:
			n += NoInline342(i)
		case 343:
			n += NoInline343(i)
		case 344:
			n += NoInline344(i)
		case 345:
			n += NoInline345(i)
		case 346:
			n += NoInline346(i)
		case 347:
			n += NoInline347(i)
		case 348:
			n += NoInline348(i)
		case 349:
			n += NoInline349(i)
		case 350:
			n += NoInline350(i)
		case 351:
			n += NoInline351(i)
		case 352:
			n += NoInline352(i)
		case 353:
			n += NoInline353(i)
		case 354:
			n += NoInline354(i)
		case 355:
			n += NoInline355(i)
		case 356:
			n += NoInline356(i)
		case 357:
			n += NoInline357(i)
		case 358:
			n += NoInline358(i)
		case 359:
			n += NoInline359(i)
		case 360:
			n += NoInline360(i)
		case 361:
			n += NoInline361(i)
		case 362:
			n += NoInline362(i)
		case 363:
			n += NoInline363(i)
		case 364:
			n += NoInline364(i)
		case 365:
			n += NoInline365(i)
		case 366:
			n += NoInline366(i)
		case 367:
			n += NoInline367(i)
		case 368:
			n += NoInline368(i)
		case 369:
			n += NoInline369(i)
		case 370:
			n += NoInline370(i)
		case 371:
			n += NoInline371(i)
		case 372:
			n += NoInline372(i)
		case 373:
			n += NoInline373(i)
		case 374:
			n += NoInline374(i)
		case 375:
			n += NoInline375(i)
		case 376:
			n += NoInline376(i)
		case 377:
			n += NoInline377(i)
		case 378:
			n += NoInline378(i)
		case 379:
			n += NoInline379(i)
		case 380:
			n += NoInline380(i)
		case 381:
			n += NoInline381(i)
		case 382:
			n += NoInline382(i)
		case 383:
			n += NoInline383(i)
		case 384:
			n += NoInline384(i)
		case 385:
			n += NoInline385(i)
		case 386:
			n += NoInline386(i)
		case 387:
			n += NoInline387(i)
		case 388:
			n += NoInline388(i)
		case 389:
			n += NoInline389(i)
		case 390:
			n += NoInline390(i)
		case 391:
			n += NoInline391(i)
		case 392:
			n += NoInline392(i)
		case 393:
			n += NoInline393(i)
		case 394:
			n += NoInline394(i)
		case 395:
			n += NoInline395(i)
		case 396:
			n += NoInline396(i)
		case 397:
			n += NoInline397(i)
		case 398:
			n += NoInline398(i)
		case 399:
			n += NoInline399(i)
		case 400:
			n += NoInline400(i)
		case 401:
			n += NoInline401(i)
		case 402:
			n += NoInline402(i)
		case 403:
			n += NoInline403(i)
		case 404:
			n += NoInline404(i)
		case 405:
			n += NoInline405(i)
		case 406:
			n += NoInline406(i)
		case 407:
			n += NoInline407(i)
		case 408:
			n += NoInline408(i)
		case 409:
			n += NoInline409(i)
		case 410:
			n += NoInline410(i)
		case 411:
			n += NoInline411(i)
		case 412:
			n += NoInline412(i)
		case 413:
			n += NoInline413(i)
		case 414:
			n += NoInline414(i)
		case 415:
			n += NoInline415(i)
		case 416:
			n += NoInline416(i)
		case 417:
			n += NoInline417(i)
		case 418:
			n += NoInline418(i)
		case 419:
			n += NoInline419(i)
		case 420:
			n += NoInline420(i)
		case 421:
			n += NoInline421(i)
		case 422:
			n += NoInline422(i)
		case 423:
			n += NoInline423(i)
		case 424:
			n += NoInline424(i)
		case 425:
			n += NoInline425(i)
		case 426:
			n += NoInline426(i)
		case 427:
			n += NoInline427(i)
		case 428:
			n += NoInline428(i)
		case 429:
			n += NoInline429(i)
		case 430:
			n += NoInline430(i)
		case 431:
			n += NoInline431(i)
		case 432:
			n += NoInline432(i)
		case 433:
			n += NoInline433(i)
		case 434:
			n += NoInline434(i)
		case 435:
			n += NoInline435(i)
		case 436:
			n += NoInline436(i)
		case 437:
			n += NoInline437(i)
		case 438:
			n += NoInline438(i)
		case 439:
			n += NoInline439(i)
		case 440:
			n += NoInline440(i)
		case 441:
			n += NoInline441(i)
		case 442:
			n += NoInline442(i)
		case 443:
			n += NoInline443(i)
		case 444:
			n += NoInline444(i)
		case 445:
			n += NoInline445(i)
		case 446:
			n += NoInline446(i)
		case 447:
			n += NoInline447(i)
		case 448:
			n += NoInline448(i)
		case 449:
			n += NoInline449(i)
		case 450:
			n += NoInline450(i)
		case 451:
			n += NoInline451(i)
		case 452:
			n += NoInline452(i)
		case 453:
			n += NoInline453(i)
		case 454:
			n += NoInline454(i)
		case 455:
			n += NoInline455(i)
		case 456:
			n += NoInline456(i)
		case 457:
			n += NoInline457(i)
		case 458:
			n += NoInline458(i)
		case 459:
			n += NoInline459(i)
		case 460:
			n += NoInline460(i)
		case 461:
			n += NoInline461(i)
		case 462:
			n += NoInline462(i)
		case 463:
			n += NoInline463(i)
		case 464:
			n += NoInline464(i)
		case 465:
			n += NoInline465(i)
		case 466:
			n += NoInline466(i)
		case 467:
			n += NoInline467(i)
		case 468:
			n += NoInline468(i)
		case 469:
			n += NoInline469(i)
		case 470:
			n += NoInline470(i)
		case 471:
			n += NoInline471(i)
		case 472:
			n += NoInline472(i)
		case 473:
			n += NoInline473(i)
		case 474:
			n += NoInline474(i)
		case 475:
			n += NoInline475(i)
		case 476:
			n += NoInline476(i)
		case 477:
			n += NoInline477(i)
		case 478:
			n += NoInline478(i)
		case 479:
			n += NoInline479(i)
		case 480:
			n += NoInline480(i)
		case 481:
			n += NoInline481(i)
		case 482:
			n += NoInline482(i)
		case 483:
			n += NoInline483(i)
		case 484:
			n += NoInline484(i)
		case 485:
			n += NoInline485(i)
		case 486:
			n += NoInline486(i)
		case 487:
			n += NoInline487(i)
		case 488:
			n += NoInline488(i)
		case 489:
			n += NoInline489(i)
		case 490:
			n += NoInline490(i)
		case 491:
			n += NoInline491(i)
		case 492:
			n += NoInline492(i)
		case 493:
			n += NoInline493(i)
		case 494:
			n += NoInline494(i)
		case 495:
			n += NoInline495(i)
		case 496:
			n += NoInline496(i)
		case 497:
			n += NoInline497(i)
		case 498:
			n += NoInline498(i)
		case 499:
			n += NoInline499(i)
		case 500:
			n += NoInline500(i)
		case 501:
			n += NoInline501(i)
		case 502:
			n += NoInline502(i)
		case 503:
			n += NoInline503(i)
		case 504:
			n += NoInline504(i)
		case 505:
			n += NoInline505(i)
		case 506:
			n += NoInline506(i)
		case 507:
			n += NoInline507(i)
		case 508:
			n += NoInline508(i)
		case 509:
			n += NoInline509(i)
		case 510:
			n += NoInline510(i)
		case 511:
			n += NoInline511(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkOptimalCacheFriendlinessMapNoInlineFunc512(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += NoInlineFuncs[i%512](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessSwitchNoInlineFunc512(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch ascInputs[i%len(ascInputs)] % 512 {
		case 0:
			n += NoInline0(i)
		case 1:
			n += NoInline1(i)
		case 2:
			n += NoInline2(i)
		case 3:
			n += NoInline3(i)
		case 4:
			n += NoInline4(i)
		case 5:
			n += NoInline5(i)
		case 6:
			n += NoInline6(i)
		case 7:
			n += NoInline7(i)
		case 8:
			n += NoInline8(i)
		case 9:
			n += NoInline9(i)
		case 10:
			n += NoInline10(i)
		case 11:
			n += NoInline11(i)
		case 12:
			n += NoInline12(i)
		case 13:
			n += NoInline13(i)
		case 14:
			n += NoInline14(i)
		case 15:
			n += NoInline15(i)
		case 16:
			n += NoInline16(i)
		case 17:
			n += NoInline17(i)
		case 18:
			n += NoInline18(i)
		case 19:
			n += NoInline19(i)
		case 20:
			n += NoInline20(i)
		case 21:
			n += NoInline21(i)
		case 22:
			n += NoInline22(i)
		case 23:
			n += NoInline23(i)
		case 24:
			n += NoInline24(i)
		case 25:
			n += NoInline25(i)
		case 26:
			n += NoInline26(i)
		case 27:
			n += NoInline27(i)
		case 28:
			n += NoInline28(i)
		case 29:
			n += NoInline29(i)
		case 30:
			n += NoInline30(i)
		case 31:
			n += NoInline31(i)
		case 32:
			n += NoInline32(i)
		case 33:
			n += NoInline33(i)
		case 34:
			n += NoInline34(i)
		case 35:
			n += NoInline35(i)
		case 36:
			n += NoInline36(i)
		case 37:
			n += NoInline37(i)
		case 38:
			n += NoInline38(i)
		case 39:
			n += NoInline39(i)
		case 40:
			n += NoInline40(i)
		case 41:
			n += NoInline41(i)
		case 42:
			n += NoInline42(i)
		case 43:
			n += NoInline43(i)
		case 44:
			n += NoInline44(i)
		case 45:
			n += NoInline45(i)
		case 46:
			n += NoInline46(i)
		case 47:
			n += NoInline47(i)
		case 48:
			n += NoInline48(i)
		case 49:
			n += NoInline49(i)
		case 50:
			n += NoInline50(i)
		case 51:
			n += NoInline51(i)
		case 52:
			n += NoInline52(i)
		case 53:
			n += NoInline53(i)
		case 54:
			n += NoInline54(i)
		case 55:
			n += NoInline55(i)
		case 56:
			n += NoInline56(i)
		case 57:
			n += NoInline57(i)
		case 58:
			n += NoInline58(i)
		case 59:
			n += NoInline59(i)
		case 60:
			n += NoInline60(i)
		case 61:
			n += NoInline61(i)
		case 62:
			n += NoInline62(i)
		case 63:
			n += NoInline63(i)
		case 64:
			n += NoInline64(i)
		case 65:
			n += NoInline65(i)
		case 66:
			n += NoInline66(i)
		case 67:
			n += NoInline67(i)
		case 68:
			n += NoInline68(i)
		case 69:
			n += NoInline69(i)
		case 70:
			n += NoInline70(i)
		case 71:
			n += NoInline71(i)
		case 72:
			n += NoInline72(i)
		case 73:
			n += NoInline73(i)
		case 74:
			n += NoInline74(i)
		case 75:
			n += NoInline75(i)
		case 76:
			n += NoInline76(i)
		case 77:
			n += NoInline77(i)
		case 78:
			n += NoInline78(i)
		case 79:
			n += NoInline79(i)
		case 80:
			n += NoInline80(i)
		case 81:
			n += NoInline81(i)
		case 82:
			n += NoInline82(i)
		case 83:
			n += NoInline83(i)
		case 84:
			n += NoInline84(i)
		case 85:
			n += NoInline85(i)
		case 86:
			n += NoInline86(i)
		case 87:
			n += NoInline87(i)
		case 88:
			n += NoInline88(i)
		case 89:
			n += NoInline89(i)
		case 90:
			n += NoInline90(i)
		case 91:
			n += NoInline91(i)
		case 92:
			n += NoInline92(i)
		case 93:
			n += NoInline93(i)
		case 94:
			n += NoInline94(i)
		case 95:
			n += NoInline95(i)
		case 96:
			n += NoInline96(i)
		case 97:
			n += NoInline97(i)
		case 98:
			n += NoInline98(i)
		case 99:
			n += NoInline99(i)
		case 100:
			n += NoInline100(i)
		case 101:
			n += NoInline101(i)
		case 102:
			n += NoInline102(i)
		case 103:
			n += NoInline103(i)
		case 104:
			n += NoInline104(i)
		case 105:
			n += NoInline105(i)
		case 106:
			n += NoInline106(i)
		case 107:
			n += NoInline107(i)
		case 108:
			n += NoInline108(i)
		case 109:
			n += NoInline109(i)
		case 110:
			n += NoInline110(i)
		case 111:
			n += NoInline111(i)
		case 112:
			n += NoInline112(i)
		case 113:
			n += NoInline113(i)
		case 114:
			n += NoInline114(i)
		case 115:
			n += NoInline115(i)
		case 116:
			n += NoInline116(i)
		case 117:
			n += NoInline117(i)
		case 118:
			n += NoInline118(i)
		case 119:
			n += NoInline119(i)
		case 120:
			n += NoInline120(i)
		case 121:
			n += NoInline121(i)
		case 122:
			n += NoInline122(i)
		case 123:
			n += NoInline123(i)
		case 124:
			n += NoInline124(i)
		case 125:
			n += NoInline125(i)
		case 126:
			n += NoInline126(i)
		case 127:
			n += NoInline127(i)
		case 128:
			n += NoInline128(i)
		case 129:
			n += NoInline129(i)
		case 130:
			n += NoInline130(i)
		case 131:
			n += NoInline131(i)
		case 132:
			n += NoInline132(i)
		case 133:
			n += NoInline133(i)
		case 134:
			n += NoInline134(i)
		case 135:
			n += NoInline135(i)
		case 136:
			n += NoInline136(i)
		case 137:
			n += NoInline137(i)
		case 138:
			n += NoInline138(i)
		case 139:
			n += NoInline139(i)
		case 140:
			n += NoInline140(i)
		case 141:
			n += NoInline141(i)
		case 142:
			n += NoInline142(i)
		case 143:
			n += NoInline143(i)
		case 144:
			n += NoInline144(i)
		case 145:
			n += NoInline145(i)
		case 146:
			n += NoInline146(i)
		case 147:
			n += NoInline147(i)
		case 148:
			n += NoInline148(i)
		case 149:
			n += NoInline149(i)
		case 150:
			n += NoInline150(i)
		case 151:
			n += NoInline151(i)
		case 152:
			n += NoInline152(i)
		case 153:
			n += NoInline153(i)
		case 154:
			n += NoInline154(i)
		case 155:
			n += NoInline155(i)
		case 156:
			n += NoInline156(i)
		case 157:
			n += NoInline157(i)
		case 158:
			n += NoInline158(i)
		case 159:
			n += NoInline159(i)
		case 160:
			n += NoInline160(i)
		case 161:
			n += NoInline161(i)
		case 162:
			n += NoInline162(i)
		case 163:
			n += NoInline163(i)
		case 164:
			n += NoInline164(i)
		case 165:
			n += NoInline165(i)
		case 166:
			n += NoInline166(i)
		case 167:
			n += NoInline167(i)
		case 168:
			n += NoInline168(i)
		case 169:
			n += NoInline169(i)
		case 170:
			n += NoInline170(i)
		case 171:
			n += NoInline171(i)
		case 172:
			n += NoInline172(i)
		case 173:
			n += NoInline173(i)
		case 174:
			n += NoInline174(i)
		case 175:
			n += NoInline175(i)
		case 176:
			n += NoInline176(i)
		case 177:
			n += NoInline177(i)
		case 178:
			n += NoInline178(i)
		case 179:
			n += NoInline179(i)
		case 180:
			n += NoInline180(i)
		case 181:
			n += NoInline181(i)
		case 182:
			n += NoInline182(i)
		case 183:
			n += NoInline183(i)
		case 184:
			n += NoInline184(i)
		case 185:
			n += NoInline185(i)
		case 186:
			n += NoInline186(i)
		case 187:
			n += NoInline187(i)
		case 188:
			n += NoInline188(i)
		case 189:
			n += NoInline189(i)
		case 190:
			n += NoInline190(i)
		case 191:
			n += NoInline191(i)
		case 192:
			n += NoInline192(i)
		case 193:
			n += NoInline193(i)
		case 194:
			n += NoInline194(i)
		case 195:
			n += NoInline195(i)
		case 196:
			n += NoInline196(i)
		case 197:
			n += NoInline197(i)
		case 198:
			n += NoInline198(i)
		case 199:
			n += NoInline199(i)
		case 200:
			n += NoInline200(i)
		case 201:
			n += NoInline201(i)
		case 202:
			n += NoInline202(i)
		case 203:
			n += NoInline203(i)
		case 204:
			n += NoInline204(i)
		case 205:
			n += NoInline205(i)
		case 206:
			n += NoInline206(i)
		case 207:
			n += NoInline207(i)
		case 208:
			n += NoInline208(i)
		case 209:
			n += NoInline209(i)
		case 210:
			n += NoInline210(i)
		case 211:
			n += NoInline211(i)
		case 212:
			n += NoInline212(i)
		case 213:
			n += NoInline213(i)
		case 214:
			n += NoInline214(i)
		case 215:
			n += NoInline215(i)
		case 216:
			n += NoInline216(i)
		case 217:
			n += NoInline217(i)
		case 218:
			n += NoInline218(i)
		case 219:
			n += NoInline219(i)
		case 220:
			n += NoInline220(i)
		case 221:
			n += NoInline221(i)
		case 222:
			n += NoInline222(i)
		case 223:
			n += NoInline223(i)
		case 224:
			n += NoInline224(i)
		case 225:
			n += NoInline225(i)
		case 226:
			n += NoInline226(i)
		case 227:
			n += NoInline227(i)
		case 228:
			n += NoInline228(i)
		case 229:
			n += NoInline229(i)
		case 230:
			n += NoInline230(i)
		case 231:
			n += NoInline231(i)
		case 232:
			n += NoInline232(i)
		case 233:
			n += NoInline233(i)
		case 234:
			n += NoInline234(i)
		case 235:
			n += NoInline235(i)
		case 236:
			n += NoInline236(i)
		case 237:
			n += NoInline237(i)
		case 238:
			n += NoInline238(i)
		case 239:
			n += NoInline239(i)
		case 240:
			n += NoInline240(i)
		case 241:
			n += NoInline241(i)
		case 242:
			n += NoInline242(i)
		case 243:
			n += NoInline243(i)
		case 244:
			n += NoInline244(i)
		case 245:
			n += NoInline245(i)
		case 246:
			n += NoInline246(i)
		case 247:
			n += NoInline247(i)
		case 248:
			n += NoInline248(i)
		case 249:
			n += NoInline249(i)
		case 250:
			n += NoInline250(i)
		case 251:
			n += NoInline251(i)
		case 252:
			n += NoInline252(i)
		case 253:
			n += NoInline253(i)
		case 254:
			n += NoInline254(i)
		case 255:
			n += NoInline255(i)
		case 256:
			n += NoInline256(i)
		case 257:
			n += NoInline257(i)
		case 258:
			n += NoInline258(i)
		case 259:
			n += NoInline259(i)
		case 260:
			n += NoInline260(i)
		case 261:
			n += NoInline261(i)
		case 262:
			n += NoInline262(i)
		case 263:
			n += NoInline263(i)
		case 264:
			n += NoInline264(i)
		case 265:
			n += NoInline265(i)
		case 266:
			n += NoInline266(i)
		case 267:
			n += NoInline267(i)
		case 268:
			n += NoInline268(i)
		case 269:
			n += NoInline269(i)
		case 270:
			n += NoInline270(i)
		case 271:
			n += NoInline271(i)
		case 272:
			n += NoInline272(i)
		case 273:
			n += NoInline273(i)
		case 274:
			n += NoInline274(i)
		case 275:
			n += NoInline275(i)
		case 276:
			n += NoInline276(i)
		case 277:
			n += NoInline277(i)
		case 278:
			n += NoInline278(i)
		case 279:
			n += NoInline279(i)
		case 280:
			n += NoInline280(i)
		case 281:
			n += NoInline281(i)
		case 282:
			n += NoInline282(i)
		case 283:
			n += NoInline283(i)
		case 284:
			n += NoInline284(i)
		case 285:
			n += NoInline285(i)
		case 286:
			n += NoInline286(i)
		case 287:
			n += NoInline287(i)
		case 288:
			n += NoInline288(i)
		case 289:
			n += NoInline289(i)
		case 290:
			n += NoInline290(i)
		case 291:
			n += NoInline291(i)
		case 292:
			n += NoInline292(i)
		case 293:
			n += NoInline293(i)
		case 294:
			n += NoInline294(i)
		case 295:
			n += NoInline295(i)
		case 296:
			n += NoInline296(i)
		case 297:
			n += NoInline297(i)
		case 298:
			n += NoInline298(i)
		case 299:
			n += NoInline299(i)
		case 300:
			n += NoInline300(i)
		case 301:
			n += NoInline301(i)
		case 302:
			n += NoInline302(i)
		case 303:
			n += NoInline303(i)
		case 304:
			n += NoInline304(i)
		case 305:
			n += NoInline305(i)
		case 306:
			n += NoInline306(i)
		case 307:
			n += NoInline307(i)
		case 308:
			n += NoInline308(i)
		case 309:
			n += NoInline309(i)
		case 310:
			n += NoInline310(i)
		case 311:
			n += NoInline311(i)
		case 312:
			n += NoInline312(i)
		case 313:
			n += NoInline313(i)
		case 314:
			n += NoInline314(i)
		case 315:
			n += NoInline315(i)
		case 316:
			n += NoInline316(i)
		case 317:
			n += NoInline317(i)
		case 318:
			n += NoInline318(i)
		case 319:
			n += NoInline319(i)
		case 320:
			n += NoInline320(i)
		case 321:
			n += NoInline321(i)
		case 322:
			n += NoInline322(i)
		case 323:
			n += NoInline323(i)
		case 324:
			n += NoInline324(i)
		case 325:
			n += NoInline325(i)
		case 326:
			n += NoInline326(i)
		case 327:
			n += NoInline327(i)
		case 328:
			n += NoInline328(i)
		case 329:
			n += NoInline329(i)
		case 330:
			n += NoInline330(i)
		case 331:
			n += NoInline331(i)
		case 332:
			n += NoInline332(i)
		case 333:
			n += NoInline333(i)
		case 334:
			n += NoInline334(i)
		case 335:
			n += NoInline335(i)
		case 336:
			n += NoInline336(i)
		case 337:
			n += NoInline337(i)
		case 338:
			n += NoInline338(i)
		case 339:
			n += NoInline339(i)
		case 340:
			n += NoInline340(i)
		case 341:
			n += NoInline341(i)
		case 342:
			n += NoInline342(i)
		case 343:
			n += NoInline343(i)
		case 344:
			n += NoInline344(i)
		case 345:
			n += NoInline345(i)
		case 346:
			n += NoInline346(i)
		case 347:
			n += NoInline347(i)
		case 348:
			n += NoInline348(i)
		case 349:
			n += NoInline349(i)
		case 350:
			n += NoInline350(i)
		case 351:
			n += NoInline351(i)
		case 352:
			n += NoInline352(i)
		case 353:
			n += NoInline353(i)
		case 354:
			n += NoInline354(i)
		case 355:
			n += NoInline355(i)
		case 356:
			n += NoInline356(i)
		case 357:
			n += NoInline357(i)
		case 358:
			n += NoInline358(i)
		case 359:
			n += NoInline359(i)
		case 360:
			n += NoInline360(i)
		case 361:
			n += NoInline361(i)
		case 362:
			n += NoInline362(i)
		case 363:
			n += NoInline363(i)
		case 364:
			n += NoInline364(i)
		case 365:
			n += NoInline365(i)
		case 366:
			n += NoInline366(i)
		case 367:
			n += NoInline367(i)
		case 368:
			n += NoInline368(i)
		case 369:
			n += NoInline369(i)
		case 370:
			n += NoInline370(i)
		case 371:
			n += NoInline371(i)
		case 372:
			n += NoInline372(i)
		case 373:
			n += NoInline373(i)
		case 374:
			n += NoInline374(i)
		case 375:
			n += NoInline375(i)
		case 376:
			n += NoInline376(i)
		case 377:
			n += NoInline377(i)
		case 378:
			n += NoInline378(i)
		case 379:
			n += NoInline379(i)
		case 380:
			n += NoInline380(i)
		case 381:
			n += NoInline381(i)
		case 382:
			n += NoInline382(i)
		case 383:
			n += NoInline383(i)
		case 384:
			n += NoInline384(i)
		case 385:
			n += NoInline385(i)
		case 386:
			n += NoInline386(i)
		case 387:
			n += NoInline387(i)
		case 388:
			n += NoInline388(i)
		case 389:
			n += NoInline389(i)
		case 390:
			n += NoInline390(i)
		case 391:
			n += NoInline391(i)
		case 392:
			n += NoInline392(i)
		case 393:
			n += NoInline393(i)
		case 394:
			n += NoInline394(i)
		case 395:
			n += NoInline395(i)
		case 396:
			n += NoInline396(i)
		case 397:
			n += NoInline397(i)
		case 398:
			n += NoInline398(i)
		case 399:
			n += NoInline399(i)
		case 400:
			n += NoInline400(i)
		case 401:
			n += NoInline401(i)
		case 402:
			n += NoInline402(i)
		case 403:
			n += NoInline403(i)
		case 404:
			n += NoInline404(i)
		case 405:
			n += NoInline405(i)
		case 406:
			n += NoInline406(i)
		case 407:
			n += NoInline407(i)
		case 408:
			n += NoInline408(i)
		case 409:
			n += NoInline409(i)
		case 410:
			n += NoInline410(i)
		case 411:
			n += NoInline411(i)
		case 412:
			n += NoInline412(i)
		case 413:
			n += NoInline413(i)
		case 414:
			n += NoInline414(i)
		case 415:
			n += NoInline415(i)
		case 416:
			n += NoInline416(i)
		case 417:
			n += NoInline417(i)
		case 418:
			n += NoInline418(i)
		case 419:
			n += NoInline419(i)
		case 420:
			n += NoInline420(i)
		case 421:
			n += NoInline421(i)
		case 422:
			n += NoInline422(i)
		case 423:
			n += NoInline423(i)
		case 424:
			n += NoInline424(i)
		case 425:
			n += NoInline425(i)
		case 426:
			n += NoInline426(i)
		case 427:
			n += NoInline427(i)
		case 428:
			n += NoInline428(i)
		case 429:
			n += NoInline429(i)
		case 430:
			n += NoInline430(i)
		case 431:
			n += NoInline431(i)
		case 432:
			n += NoInline432(i)
		case 433:
			n += NoInline433(i)
		case 434:
			n += NoInline434(i)
		case 435:
			n += NoInline435(i)
		case 436:
			n += NoInline436(i)
		case 437:
			n += NoInline437(i)
		case 438:
			n += NoInline438(i)
		case 439:
			n += NoInline439(i)
		case 440:
			n += NoInline440(i)
		case 441:
			n += NoInline441(i)
		case 442:
			n += NoInline442(i)
		case 443:
			n += NoInline443(i)
		case 444:
			n += NoInline444(i)
		case 445:
			n += NoInline445(i)
		case 446:
			n += NoInline446(i)
		case 447:
			n += NoInline447(i)
		case 448:
			n += NoInline448(i)
		case 449:
			n += NoInline449(i)
		case 450:
			n += NoInline450(i)
		case 451:
			n += NoInline451(i)
		case 452:
			n += NoInline452(i)
		case 453:
			n += NoInline453(i)
		case 454:
			n += NoInline454(i)
		case 455:
			n += NoInline455(i)
		case 456:
			n += NoInline456(i)
		case 457:
			n += NoInline457(i)
		case 458:
			n += NoInline458(i)
		case 459:
			n += NoInline459(i)
		case 460:
			n += NoInline460(i)
		case 461:
			n += NoInline461(i)
		case 462:
			n += NoInline462(i)
		case 463:
			n += NoInline463(i)
		case 464:
			n += NoInline464(i)
		case 465:
			n += NoInline465(i)
		case 466:
			n += NoInline466(i)
		case 467:
			n += NoInline467(i)
		case 468:
			n += NoInline468(i)
		case 469:
			n += NoInline469(i)
		case 470:
			n += NoInline470(i)
		case 471:
			n += NoInline471(i)
		case 472:
			n += NoInline472(i)
		case 473:
			n += NoInline473(i)
		case 474:
			n += NoInline474(i)
		case 475:
			n += NoInline475(i)
		case 476:
			n += NoInline476(i)
		case 477:
			n += NoInline477(i)
		case 478:
			n += NoInline478(i)
		case 479:
			n += NoInline479(i)
		case 480:
			n += NoInline480(i)
		case 481:
			n += NoInline481(i)
		case 482:
			n += NoInline482(i)
		case 483:
			n += NoInline483(i)
		case 484:
			n += NoInline484(i)
		case 485:
			n += NoInline485(i)
		case 486:
			n += NoInline486(i)
		case 487:
			n += NoInline487(i)
		case 488:
			n += NoInline488(i)
		case 489:
			n += NoInline489(i)
		case 490:
			n += NoInline490(i)
		case 491:
			n += NoInline491(i)
		case 492:
			n += NoInline492(i)
		case 493:
			n += NoInline493(i)
		case 494:
			n += NoInline494(i)
		case 495:
			n += NoInline495(i)
		case 496:
			n += NoInline496(i)
		case 497:
			n += NoInline497(i)
		case 498:
			n += NoInline498(i)
		case 499:
			n += NoInline499(i)
		case 500:
			n += NoInline500(i)
		case 501:
			n += NoInline501(i)
		case 502:
			n += NoInline502(i)
		case 503:
			n += NoInline503(i)
		case 504:
			n += NoInline504(i)
		case 505:
			n += NoInline505(i)
		case 506:
			n += NoInline506(i)
		case 507:
			n += NoInline507(i)
		case 508:
			n += NoInline508(i)
		case 509:
			n += NoInline509(i)
		case 510:
			n += NoInline510(i)
		case 511:
			n += NoInline511(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkModerateCacheFriendlinessMapNoInlineFunc512(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += NoInlineFuncs[ascInputs[i%len(ascInputs)]%512](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessSwitchNoInlineFunc512(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch randInputs[i%len(randInputs)] % 512 {
		case 0:
			n += NoInline0(i)
		case 1:
			n += NoInline1(i)
		case 2:
			n += NoInline2(i)
		case 3:
			n += NoInline3(i)
		case 4:
			n += NoInline4(i)
		case 5:
			n += NoInline5(i)
		case 6:
			n += NoInline6(i)
		case 7:
			n += NoInline7(i)
		case 8:
			n += NoInline8(i)
		case 9:
			n += NoInline9(i)
		case 10:
			n += NoInline10(i)
		case 11:
			n += NoInline11(i)
		case 12:
			n += NoInline12(i)
		case 13:
			n += NoInline13(i)
		case 14:
			n += NoInline14(i)
		case 15:
			n += NoInline15(i)
		case 16:
			n += NoInline16(i)
		case 17:
			n += NoInline17(i)
		case 18:
			n += NoInline18(i)
		case 19:
			n += NoInline19(i)
		case 20:
			n += NoInline20(i)
		case 21:
			n += NoInline21(i)
		case 22:
			n += NoInline22(i)
		case 23:
			n += NoInline23(i)
		case 24:
			n += NoInline24(i)
		case 25:
			n += NoInline25(i)
		case 26:
			n += NoInline26(i)
		case 27:
			n += NoInline27(i)
		case 28:
			n += NoInline28(i)
		case 29:
			n += NoInline29(i)
		case 30:
			n += NoInline30(i)
		case 31:
			n += NoInline31(i)
		case 32:
			n += NoInline32(i)
		case 33:
			n += NoInline33(i)
		case 34:
			n += NoInline34(i)
		case 35:
			n += NoInline35(i)
		case 36:
			n += NoInline36(i)
		case 37:
			n += NoInline37(i)
		case 38:
			n += NoInline38(i)
		case 39:
			n += NoInline39(i)
		case 40:
			n += NoInline40(i)
		case 41:
			n += NoInline41(i)
		case 42:
			n += NoInline42(i)
		case 43:
			n += NoInline43(i)
		case 44:
			n += NoInline44(i)
		case 45:
			n += NoInline45(i)
		case 46:
			n += NoInline46(i)
		case 47:
			n += NoInline47(i)
		case 48:
			n += NoInline48(i)
		case 49:
			n += NoInline49(i)
		case 50:
			n += NoInline50(i)
		case 51:
			n += NoInline51(i)
		case 52:
			n += NoInline52(i)
		case 53:
			n += NoInline53(i)
		case 54:
			n += NoInline54(i)
		case 55:
			n += NoInline55(i)
		case 56:
			n += NoInline56(i)
		case 57:
			n += NoInline57(i)
		case 58:
			n += NoInline58(i)
		case 59:
			n += NoInline59(i)
		case 60:
			n += NoInline60(i)
		case 61:
			n += NoInline61(i)
		case 62:
			n += NoInline62(i)
		case 63:
			n += NoInline63(i)
		case 64:
			n += NoInline64(i)
		case 65:
			n += NoInline65(i)
		case 66:
			n += NoInline66(i)
		case 67:
			n += NoInline67(i)
		case 68:
			n += NoInline68(i)
		case 69:
			n += NoInline69(i)
		case 70:
			n += NoInline70(i)
		case 71:
			n += NoInline71(i)
		case 72:
			n += NoInline72(i)
		case 73:
			n += NoInline73(i)
		case 74:
			n += NoInline74(i)
		case 75:
			n += NoInline75(i)
		case 76:
			n += NoInline76(i)
		case 77:
			n += NoInline77(i)
		case 78:
			n += NoInline78(i)
		case 79:
			n += NoInline79(i)
		case 80:
			n += NoInline80(i)
		case 81:
			n += NoInline81(i)
		case 82:
			n += NoInline82(i)
		case 83:
			n += NoInline83(i)
		case 84:
			n += NoInline84(i)
		case 85:
			n += NoInline85(i)
		case 86:
			n += NoInline86(i)
		case 87:
			n += NoInline87(i)
		case 88:
			n += NoInline88(i)
		case 89:
			n += NoInline89(i)
		case 90:
			n += NoInline90(i)
		case 91:
			n += NoInline91(i)
		case 92:
			n += NoInline92(i)
		case 93:
			n += NoInline93(i)
		case 94:
			n += NoInline94(i)
		case 95:
			n += NoInline95(i)
		case 96:
			n += NoInline96(i)
		case 97:
			n += NoInline97(i)
		case 98:
			n += NoInline98(i)
		case 99:
			n += NoInline99(i)
		case 100:
			n += NoInline100(i)
		case 101:
			n += NoInline101(i)
		case 102:
			n += NoInline102(i)
		case 103:
			n += NoInline103(i)
		case 104:
			n += NoInline104(i)
		case 105:
			n += NoInline105(i)
		case 106:
			n += NoInline106(i)
		case 107:
			n += NoInline107(i)
		case 108:
			n += NoInline108(i)
		case 109:
			n += NoInline109(i)
		case 110:
			n += NoInline110(i)
		case 111:
			n += NoInline111(i)
		case 112:
			n += NoInline112(i)
		case 113:
			n += NoInline113(i)
		case 114:
			n += NoInline114(i)
		case 115:
			n += NoInline115(i)
		case 116:
			n += NoInline116(i)
		case 117:
			n += NoInline117(i)
		case 118:
			n += NoInline118(i)
		case 119:
			n += NoInline119(i)
		case 120:
			n += NoInline120(i)
		case 121:
			n += NoInline121(i)
		case 122:
			n += NoInline122(i)
		case 123:
			n += NoInline123(i)
		case 124:
			n += NoInline124(i)
		case 125:
			n += NoInline125(i)
		case 126:
			n += NoInline126(i)
		case 127:
			n += NoInline127(i)
		case 128:
			n += NoInline128(i)
		case 129:
			n += NoInline129(i)
		case 130:
			n += NoInline130(i)
		case 131:
			n += NoInline131(i)
		case 132:
			n += NoInline132(i)
		case 133:
			n += NoInline133(i)
		case 134:
			n += NoInline134(i)
		case 135:
			n += NoInline135(i)
		case 136:
			n += NoInline136(i)
		case 137:
			n += NoInline137(i)
		case 138:
			n += NoInline138(i)
		case 139:
			n += NoInline139(i)
		case 140:
			n += NoInline140(i)
		case 141:
			n += NoInline141(i)
		case 142:
			n += NoInline142(i)
		case 143:
			n += NoInline143(i)
		case 144:
			n += NoInline144(i)
		case 145:
			n += NoInline145(i)
		case 146:
			n += NoInline146(i)
		case 147:
			n += NoInline147(i)
		case 148:
			n += NoInline148(i)
		case 149:
			n += NoInline149(i)
		case 150:
			n += NoInline150(i)
		case 151:
			n += NoInline151(i)
		case 152:
			n += NoInline152(i)
		case 153:
			n += NoInline153(i)
		case 154:
			n += NoInline154(i)
		case 155:
			n += NoInline155(i)
		case 156:
			n += NoInline156(i)
		case 157:
			n += NoInline157(i)
		case 158:
			n += NoInline158(i)
		case 159:
			n += NoInline159(i)
		case 160:
			n += NoInline160(i)
		case 161:
			n += NoInline161(i)
		case 162:
			n += NoInline162(i)
		case 163:
			n += NoInline163(i)
		case 164:
			n += NoInline164(i)
		case 165:
			n += NoInline165(i)
		case 166:
			n += NoInline166(i)
		case 167:
			n += NoInline167(i)
		case 168:
			n += NoInline168(i)
		case 169:
			n += NoInline169(i)
		case 170:
			n += NoInline170(i)
		case 171:
			n += NoInline171(i)
		case 172:
			n += NoInline172(i)
		case 173:
			n += NoInline173(i)
		case 174:
			n += NoInline174(i)
		case 175:
			n += NoInline175(i)
		case 176:
			n += NoInline176(i)
		case 177:
			n += NoInline177(i)
		case 178:
			n += NoInline178(i)
		case 179:
			n += NoInline179(i)
		case 180:
			n += NoInline180(i)
		case 181:
			n += NoInline181(i)
		case 182:
			n += NoInline182(i)
		case 183:
			n += NoInline183(i)
		case 184:
			n += NoInline184(i)
		case 185:
			n += NoInline185(i)
		case 186:
			n += NoInline186(i)
		case 187:
			n += NoInline187(i)
		case 188:
			n += NoInline188(i)
		case 189:
			n += NoInline189(i)
		case 190:
			n += NoInline190(i)
		case 191:
			n += NoInline191(i)
		case 192:
			n += NoInline192(i)
		case 193:
			n += NoInline193(i)
		case 194:
			n += NoInline194(i)
		case 195:
			n += NoInline195(i)
		case 196:
			n += NoInline196(i)
		case 197:
			n += NoInline197(i)
		case 198:
			n += NoInline198(i)
		case 199:
			n += NoInline199(i)
		case 200:
			n += NoInline200(i)
		case 201:
			n += NoInline201(i)
		case 202:
			n += NoInline202(i)
		case 203:
			n += NoInline203(i)
		case 204:
			n += NoInline204(i)
		case 205:
			n += NoInline205(i)
		case 206:
			n += NoInline206(i)
		case 207:
			n += NoInline207(i)
		case 208:
			n += NoInline208(i)
		case 209:
			n += NoInline209(i)
		case 210:
			n += NoInline210(i)
		case 211:
			n += NoInline211(i)
		case 212:
			n += NoInline212(i)
		case 213:
			n += NoInline213(i)
		case 214:
			n += NoInline214(i)
		case 215:
			n += NoInline215(i)
		case 216:
			n += NoInline216(i)
		case 217:
			n += NoInline217(i)
		case 218:
			n += NoInline218(i)
		case 219:
			n += NoInline219(i)
		case 220:
			n += NoInline220(i)
		case 221:
			n += NoInline221(i)
		case 222:
			n += NoInline222(i)
		case 223:
			n += NoInline223(i)
		case 224:
			n += NoInline224(i)
		case 225:
			n += NoInline225(i)
		case 226:
			n += NoInline226(i)
		case 227:
			n += NoInline227(i)
		case 228:
			n += NoInline228(i)
		case 229:
			n += NoInline229(i)
		case 230:
			n += NoInline230(i)
		case 231:
			n += NoInline231(i)
		case 232:
			n += NoInline232(i)
		case 233:
			n += NoInline233(i)
		case 234:
			n += NoInline234(i)
		case 235:
			n += NoInline235(i)
		case 236:
			n += NoInline236(i)
		case 237:
			n += NoInline237(i)
		case 238:
			n += NoInline238(i)
		case 239:
			n += NoInline239(i)
		case 240:
			n += NoInline240(i)
		case 241:
			n += NoInline241(i)
		case 242:
			n += NoInline242(i)
		case 243:
			n += NoInline243(i)
		case 244:
			n += NoInline244(i)
		case 245:
			n += NoInline245(i)
		case 246:
			n += NoInline246(i)
		case 247:
			n += NoInline247(i)
		case 248:
			n += NoInline248(i)
		case 249:
			n += NoInline249(i)
		case 250:
			n += NoInline250(i)
		case 251:
			n += NoInline251(i)
		case 252:
			n += NoInline252(i)
		case 253:
			n += NoInline253(i)
		case 254:
			n += NoInline254(i)
		case 255:
			n += NoInline255(i)
		case 256:
			n += NoInline256(i)
		case 257:
			n += NoInline257(i)
		case 258:
			n += NoInline258(i)
		case 259:
			n += NoInline259(i)
		case 260:
			n += NoInline260(i)
		case 261:
			n += NoInline261(i)
		case 262:
			n += NoInline262(i)
		case 263:
			n += NoInline263(i)
		case 264:
			n += NoInline264(i)
		case 265:
			n += NoInline265(i)
		case 266:
			n += NoInline266(i)
		case 267:
			n += NoInline267(i)
		case 268:
			n += NoInline268(i)
		case 269:
			n += NoInline269(i)
		case 270:
			n += NoInline270(i)
		case 271:
			n += NoInline271(i)
		case 272:
			n += NoInline272(i)
		case 273:
			n += NoInline273(i)
		case 274:
			n += NoInline274(i)
		case 275:
			n += NoInline275(i)
		case 276:
			n += NoInline276(i)
		case 277:
			n += NoInline277(i)
		case 278:
			n += NoInline278(i)
		case 279:
			n += NoInline279(i)
		case 280:
			n += NoInline280(i)
		case 281:
			n += NoInline281(i)
		case 282:
			n += NoInline282(i)
		case 283:
			n += NoInline283(i)
		case 284:
			n += NoInline284(i)
		case 285:
			n += NoInline285(i)
		case 286:
			n += NoInline286(i)
		case 287:
			n += NoInline287(i)
		case 288:
			n += NoInline288(i)
		case 289:
			n += NoInline289(i)
		case 290:
			n += NoInline290(i)
		case 291:
			n += NoInline291(i)
		case 292:
			n += NoInline292(i)
		case 293:
			n += NoInline293(i)
		case 294:
			n += NoInline294(i)
		case 295:
			n += NoInline295(i)
		case 296:
			n += NoInline296(i)
		case 297:
			n += NoInline297(i)
		case 298:
			n += NoInline298(i)
		case 299:
			n += NoInline299(i)
		case 300:
			n += NoInline300(i)
		case 301:
			n += NoInline301(i)
		case 302:
			n += NoInline302(i)
		case 303:
			n += NoInline303(i)
		case 304:
			n += NoInline304(i)
		case 305:
			n += NoInline305(i)
		case 306:
			n += NoInline306(i)
		case 307:
			n += NoInline307(i)
		case 308:
			n += NoInline308(i)
		case 309:
			n += NoInline309(i)
		case 310:
			n += NoInline310(i)
		case 311:
			n += NoInline311(i)
		case 312:
			n += NoInline312(i)
		case 313:
			n += NoInline313(i)
		case 314:
			n += NoInline314(i)
		case 315:
			n += NoInline315(i)
		case 316:
			n += NoInline316(i)
		case 317:
			n += NoInline317(i)
		case 318:
			n += NoInline318(i)
		case 319:
			n += NoInline319(i)
		case 320:
			n += NoInline320(i)
		case 321:
			n += NoInline321(i)
		case 322:
			n += NoInline322(i)
		case 323:
			n += NoInline323(i)
		case 324:
			n += NoInline324(i)
		case 325:
			n += NoInline325(i)
		case 326:
			n += NoInline326(i)
		case 327:
			n += NoInline327(i)
		case 328:
			n += NoInline328(i)
		case 329:
			n += NoInline329(i)
		case 330:
			n += NoInline330(i)
		case 331:
			n += NoInline331(i)
		case 332:
			n += NoInline332(i)
		case 333:
			n += NoInline333(i)
		case 334:
			n += NoInline334(i)
		case 335:
			n += NoInline335(i)
		case 336:
			n += NoInline336(i)
		case 337:
			n += NoInline337(i)
		case 338:
			n += NoInline338(i)
		case 339:
			n += NoInline339(i)
		case 340:
			n += NoInline340(i)
		case 341:
			n += NoInline341(i)
		case 342:
			n += NoInline342(i)
		case 343:
			n += NoInline343(i)
		case 344:
			n += NoInline344(i)
		case 345:
			n += NoInline345(i)
		case 346:
			n += NoInline346(i)
		case 347:
			n += NoInline347(i)
		case 348:
			n += NoInline348(i)
		case 349:
			n += NoInline349(i)
		case 350:
			n += NoInline350(i)
		case 351:
			n += NoInline351(i)
		case 352:
			n += NoInline352(i)
		case 353:
			n += NoInline353(i)
		case 354:
			n += NoInline354(i)
		case 355:
			n += NoInline355(i)
		case 356:
			n += NoInline356(i)
		case 357:
			n += NoInline357(i)
		case 358:
			n += NoInline358(i)
		case 359:
			n += NoInline359(i)
		case 360:
			n += NoInline360(i)
		case 361:
			n += NoInline361(i)
		case 362:
			n += NoInline362(i)
		case 363:
			n += NoInline363(i)
		case 364:
			n += NoInline364(i)
		case 365:
			n += NoInline365(i)
		case 366:
			n += NoInline366(i)
		case 367:
			n += NoInline367(i)
		case 368:
			n += NoInline368(i)
		case 369:
			n += NoInline369(i)
		case 370:
			n += NoInline370(i)
		case 371:
			n += NoInline371(i)
		case 372:
			n += NoInline372(i)
		case 373:
			n += NoInline373(i)
		case 374:
			n += NoInline374(i)
		case 375:
			n += NoInline375(i)
		case 376:
			n += NoInline376(i)
		case 377:
			n += NoInline377(i)
		case 378:
			n += NoInline378(i)
		case 379:
			n += NoInline379(i)
		case 380:
			n += NoInline380(i)
		case 381:
			n += NoInline381(i)
		case 382:
			n += NoInline382(i)
		case 383:
			n += NoInline383(i)
		case 384:
			n += NoInline384(i)
		case 385:
			n += NoInline385(i)
		case 386:
			n += NoInline386(i)
		case 387:
			n += NoInline387(i)
		case 388:
			n += NoInline388(i)
		case 389:
			n += NoInline389(i)
		case 390:
			n += NoInline390(i)
		case 391:
			n += NoInline391(i)
		case 392:
			n += NoInline392(i)
		case 393:
			n += NoInline393(i)
		case 394:
			n += NoInline394(i)
		case 395:
			n += NoInline395(i)
		case 396:
			n += NoInline396(i)
		case 397:
			n += NoInline397(i)
		case 398:
			n += NoInline398(i)
		case 399:
			n += NoInline399(i)
		case 400:
			n += NoInline400(i)
		case 401:
			n += NoInline401(i)
		case 402:
			n += NoInline402(i)
		case 403:
			n += NoInline403(i)
		case 404:
			n += NoInline404(i)
		case 405:
			n += NoInline405(i)
		case 406:
			n += NoInline406(i)
		case 407:
			n += NoInline407(i)
		case 408:
			n += NoInline408(i)
		case 409:
			n += NoInline409(i)
		case 410:
			n += NoInline410(i)
		case 411:
			n += NoInline411(i)
		case 412:
			n += NoInline412(i)
		case 413:
			n += NoInline413(i)
		case 414:
			n += NoInline414(i)
		case 415:
			n += NoInline415(i)
		case 416:
			n += NoInline416(i)
		case 417:
			n += NoInline417(i)
		case 418:
			n += NoInline418(i)
		case 419:
			n += NoInline419(i)
		case 420:
			n += NoInline420(i)
		case 421:
			n += NoInline421(i)
		case 422:
			n += NoInline422(i)
		case 423:
			n += NoInline423(i)
		case 424:
			n += NoInline424(i)
		case 425:
			n += NoInline425(i)
		case 426:
			n += NoInline426(i)
		case 427:
			n += NoInline427(i)
		case 428:
			n += NoInline428(i)
		case 429:
			n += NoInline429(i)
		case 430:
			n += NoInline430(i)
		case 431:
			n += NoInline431(i)
		case 432:
			n += NoInline432(i)
		case 433:
			n += NoInline433(i)
		case 434:
			n += NoInline434(i)
		case 435:
			n += NoInline435(i)
		case 436:
			n += NoInline436(i)
		case 437:
			n += NoInline437(i)
		case 438:
			n += NoInline438(i)
		case 439:
			n += NoInline439(i)
		case 440:
			n += NoInline440(i)
		case 441:
			n += NoInline441(i)
		case 442:
			n += NoInline442(i)
		case 443:
			n += NoInline443(i)
		case 444:
			n += NoInline444(i)
		case 445:
			n += NoInline445(i)
		case 446:
			n += NoInline446(i)
		case 447:
			n += NoInline447(i)
		case 448:
			n += NoInline448(i)
		case 449:
			n += NoInline449(i)
		case 450:
			n += NoInline450(i)
		case 451:
			n += NoInline451(i)
		case 452:
			n += NoInline452(i)
		case 453:
			n += NoInline453(i)
		case 454:
			n += NoInline454(i)
		case 455:
			n += NoInline455(i)
		case 456:
			n += NoInline456(i)
		case 457:
			n += NoInline457(i)
		case 458:
			n += NoInline458(i)
		case 459:
			n += NoInline459(i)
		case 460:
			n += NoInline460(i)
		case 461:
			n += NoInline461(i)
		case 462:
			n += NoInline462(i)
		case 463:
			n += NoInline463(i)
		case 464:
			n += NoInline464(i)
		case 465:
			n += NoInline465(i)
		case 466:
			n += NoInline466(i)
		case 467:
			n += NoInline467(i)
		case 468:
			n += NoInline468(i)
		case 469:
			n += NoInline469(i)
		case 470:
			n += NoInline470(i)
		case 471:
			n += NoInline471(i)
		case 472:
			n += NoInline472(i)
		case 473:
			n += NoInline473(i)
		case 474:
			n += NoInline474(i)
		case 475:
			n += NoInline475(i)
		case 476:
			n += NoInline476(i)
		case 477:
			n += NoInline477(i)
		case 478:
			n += NoInline478(i)
		case 479:
			n += NoInline479(i)
		case 480:
			n += NoInline480(i)
		case 481:
			n += NoInline481(i)
		case 482:
			n += NoInline482(i)
		case 483:
			n += NoInline483(i)
		case 484:
			n += NoInline484(i)
		case 485:
			n += NoInline485(i)
		case 486:
			n += NoInline486(i)
		case 487:
			n += NoInline487(i)
		case 488:
			n += NoInline488(i)
		case 489:
			n += NoInline489(i)
		case 490:
			n += NoInline490(i)
		case 491:
			n += NoInline491(i)
		case 492:
			n += NoInline492(i)
		case 493:
			n += NoInline493(i)
		case 494:
			n += NoInline494(i)
		case 495:
			n += NoInline495(i)
		case 496:
			n += NoInline496(i)
		case 497:
			n += NoInline497(i)
		case 498:
			n += NoInline498(i)
		case 499:
			n += NoInline499(i)
		case 500:
			n += NoInline500(i)
		case 501:
			n += NoInline501(i)
		case 502:
			n += NoInline502(i)
		case 503:
			n += NoInline503(i)
		case 504:
			n += NoInline504(i)
		case 505:
			n += NoInline505(i)
		case 506:
			n += NoInline506(i)
		case 507:
			n += NoInline507(i)
		case 508:
			n += NoInline508(i)
		case 509:
			n += NoInline509(i)
		case 510:
			n += NoInline510(i)
		case 511:
			n += NoInline511(i)
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func BenchmarkPoorCacheFriendlinessMapNoInlineFunc512(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		n += NoInlineFuncs[randInputs[i%len(randInputs)]%512](i)
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}
