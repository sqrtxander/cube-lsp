package cube

import "testing"

func TestSolved(t *testing.T) {
	alg := ""
	expected := `   UUU
   UUU
   UUU
LLLFFFRRRBBB
LLLFFFRRRBBB
LLLFFFRRRBBB
   DDD
   DDD
   DDD`
	cube, err := DoMoves(*GetSolvedCube(), alg)
	actual := ToNetString(*cube)
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	if expected != actual {
		t.Fatalf("TestSolved: expected %s but got %s", expected, actual)
	}
}

func TestCheckerboardSlice(t *testing.T) {
	alg := "M2 E2 S2"
	expected := `   UDU
   DUD
   UDU
LRLFBFRLRBFB
RLRBFBLRLFBF
LRLFBFRLRBFB
   DUD
   UDU
   DUD`
	cube, err := DoMoves(*GetSolvedCube(), alg)
	actual := ToNetString(*cube)
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	if expected != actual {
		t.Fatalf("TestCheckerboardSlice: expected %s but got %s", expected, actual)
	}
}

func TestJPerm(t *testing.T) {
	alg := "R U R' F' R U R' U' R' F R2 U' R'"
	expected := `   UUU
   UUU
   UUU
FRRBFFRBBLLL
LLLFFFRRRBBB
LLLFFFRRRBBB
   DDD
   DDD
   DDD`
	cube, err := DoMoves(*GetSolvedCube(), alg)
	actual := ToNetString(*cube)
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	if expected != actual {
		t.Fatalf("TestJPerm: expected %s but got %s", expected, actual)
	}
}

func TestScramble(t *testing.T) {
	alg := "D' B R D' B D F2 L U L2 F' U2 D2 B D2 B' R2 B' U2 L2 B'"
	expected := `   RFR
   DUR
   DLL
URFLBBDUUFRB
ULBUFFDRBRBF
LDDRULBBDFFF
   BLU
   LDD
   ULR`
	cube, err := DoMoves(*GetSolvedCube(), alg)
	actual := ToNetString(*cube)
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	if expected != actual {
		t.Fatalf("TestScramble: expected %s but got %s", expected, actual)
	}
}

func TestSuperflip(t *testing.T) {
	alg := "M' U' M' U' M' U' M' U' x y M' U' M' U' M' U' M' U' x y M' U' M' U' M' U' M' U' x y"
	expected := `   UBU
   LUR
   UFU
LULFUFRURBUB
BLFLFRFRBRBL
LDLFDFRDRBDB
   DFD
   LDR
   DBD`
	cube, err := DoMoves(*GetSolvedCube(), alg)
	actual := ToNetString(*cube)
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	if expected != actual {
		t.Fatalf("TestSuperflip: expected %s but got %s", expected, actual)
	}
}

func TestIsSolved(t *testing.T) {
	cube := GetSolvedCube()
	expected := true
	actual := IsSolved(*cube)
	if expected != actual {
		t.Fatalf("TestIsSolved: expected %t but got %t", expected, actual)
	}
	cube, err := DoMoves(*cube, "S")
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	expected = false
	actual = IsSolved(*cube)
	if expected != actual {
		t.Fatalf("TestIsSolved: expected %t but got %t", expected, actual)
	}
}
