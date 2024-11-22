package cube

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/go-cmp/cmp"
)

func GetSolvedCube() *Cube {
	return &Cube{
		TP: [6]CentreP{U, D, F, B, L, R},
		CP: [8]CornerP{UFR, UFL, UBL, UBR, DFR, DFL, DBL, DBR},
		CO: [8]CornerO{GC, GC, GC, GC, GC, GC, GC, GC},
		EP: [12]EdgeP{UF, UL, UB, UR, FR, FL, BL, BR, DF, DL, DB, DR},
		EO: [12]EdgeO{GE, GE, GE, GE, GE, GE, GE, GE, GE, GE, GE, GE},
	}
}

func centreMultiply(on Cube, by Cube) *Cube {
	newTP := [6]CentreP{}
	for t := 0; t < 6; t++ {
		from := by.TP[t]
		newTP[t] = on.TP[from]
	}
	on.TP = newTP
	return &on
}

func cornerMultiply(on Cube, by Cube) *Cube {
	newCP := [8]CentreP{}
	newCO := [8]CentreP{}
	for c := 0; c < 8; c++ {
		fromP := by.CP[c]
		newCP[c] = on.CP[fromP]
		fromO := by.CO[c]
		newCO[c] = (on.CO[fromP] + fromO) % 3
	}
	on.CP = newCP
	on.CO = newCO
	return &on
}

func edgeMultiply(on Cube, by Cube) *Cube {
	newEP := [12]CentreP{}
	newEO := [12]CentreP{}
	for e := 0; e < 12; e++ {
		fromP := by.EP[e]
		newEP[e] = on.EP[fromP]
		fromO := by.EO[e]
		newEO[e] = (on.EO[fromP] + fromO) % 2
	}
	on.EP = newEP
	on.EO = newEO
	return &on
}

func multiply(on Cube, by Cube) *Cube {
	on = *centreMultiply(on, by)
	on = *cornerMultiply(on, by)
	return edgeMultiply(on, by)
}

func DoMoves(on Cube, moves string) (*Cube, error) {
	var err error
	onn := &on
	for _, move := range strings.Fields(moves) {
		onn, err = DoMove(*onn, move)
		if err != nil {
			return nil, err
		}
	}
	return onn, nil
}

func DoMove(on Cube, move string) (*Cube, error) {
	if len(move) == 0 {
		return nil, fmt.Errorf("DoMove: Empty move supplied: %s", move)
	}

	if len(move) > 2 {
		return nil, fmt.Errorf("DoMove: Move too long: %s", move)
	}

	face := move[0]
	turns := 1
	if len(move) == 2 {
		switch move[1] {
		case '2':
			turns = 2
			break
		case '\'':
			turns = 3
			break
		default:
			return nil, fmt.Errorf("DoMove: Unknown move amount, %c", move[1])
		}
	}

	for i := 0; i < turns; i++ {
		switch face {
		case 'U':
			on = *multiply(on, moves[0])
			break
		case 'D':
			on = *multiply(on, moves[1])
			break
		case 'F':
			on = *multiply(on, moves[2])
			break
		case 'B':
			on = *multiply(on, moves[3])
			break
		case 'L':
			on = *multiply(on, moves[4])
			break
		case 'R':
			on = *multiply(on, moves[5])
			break
		case 'M':
			on = *multiply(on, moves[6])
			break
		case 'E':
			on = *multiply(on, moves[7])
			break
		case 'S':
			on = *multiply(on, moves[8])
			break
		case 'u':
			on = *multiply(on, moves[9])
			break
		case 'd':
			on = *multiply(on, moves[10])
			break
		case 'f':
			on = *multiply(on, moves[11])
			break
		case 'b':
			on = *multiply(on, moves[12])
			break
		case 'l':
			on = *multiply(on, moves[13])
			break
		case 'r':
			on = *multiply(on, moves[14])
			break
		case 'x':
			on = *multiply(on, moves[15])
			break
		case 'y':
			on = *multiply(on, moves[16])
			break
		case 'z':
			on = *multiply(on, moves[17])
			break
		default:
			return nil, fmt.Errorf("DoMove: Unknown face, %c", face)
		}
	}

	return &on, nil
}

func IsSolved(cube Cube) bool {
	return cmp.Equal(cube, *GetSolvedCube())
}

func IsCrossSolved(cube Cube) bool {
    cubep := &cube
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			if isCrossSolvedD(cube) {
				return true
			}
            cubep, _ = DoMoves(*cubep, "x y")
		}
        cubep, _ = DoMoves(*cubep, "y2 x")
	}
	return false
}

func isCrossSolvedD(cube Cube) bool {
	return cube.EP[8] == DF &&
		cube.EP[9] == DL &&
		cube.EP[10] == DB &&
		cube.EP[11] == DR &&
		cube.EO[8] == GE &&
		cube.EO[9] == GE &&
		cube.EO[10] == GE &&
		cube.EO[11] == GE
}

func PrintCubeJSON(cube Cube) {
	res, _ := json.Marshal(cube)
	fmt.Println(string(res))
}

func ToNetString(cube Cube) string {
	/*
				    UUU
				    UUU
				    UUU
		        LLLFFFRRRBBB
		        LLLFFFRRRBBB
		        LLLFFFRRRBBB
				    DDD
				    DDD
				    DDD
	*/
	centres := map[CentreP]byte{
		U: 'U',
		D: 'D',
		F: 'F',
		B: 'B',
		L: 'L',
		R: 'R',
	}
	corners := map[CornerP][3]byte{
		UFR: {'U', 'R', 'F'},
		UFL: {'U', 'F', 'L'},
		UBL: {'U', 'L', 'B'},
		UBR: {'U', 'B', 'R'},
		DFR: {'D', 'F', 'R'},
		DFL: {'D', 'L', 'F'},
		DBL: {'D', 'B', 'L'},
		DBR: {'D', 'R', 'B'},
	}
	edges := map[EdgeP][2]byte{
		UF: {'U', 'F'},
		UL: {'U', 'L'},
		UB: {'U', 'B'},
		UR: {'U', 'R'},
		FR: {'F', 'R'},
		FL: {'F', 'L'},
		BL: {'B', 'L'},
		BR: {'B', 'R'},
		DF: {'D', 'F'},
		DL: {'D', 'L'},
		DB: {'D', 'B'},
		DR: {'D', 'R'},
	}
	u := centres[cube.TP[0]]
	d := centres[cube.TP[1]]
	f := centres[cube.TP[2]]
	b := centres[cube.TP[3]]
	l := centres[cube.TP[4]]
	r := centres[cube.TP[5]]

	ufr := corners[cube.CP[0]][cube.CO[0]]
	ufl := corners[cube.CP[1]][cube.CO[1]]
	ubl := corners[cube.CP[2]][cube.CO[2]]
	ubr := corners[cube.CP[3]][cube.CO[3]]
	dfr := corners[cube.CP[4]][cube.CO[4]]
	dfl := corners[cube.CP[5]][cube.CO[5]]
	dbl := corners[cube.CP[6]][cube.CO[6]]
	dbr := corners[cube.CP[7]][cube.CO[7]]

	ruf := corners[cube.CP[0]][(1+cube.CO[0])%3]
	ful := corners[cube.CP[1]][(1+cube.CO[1])%3]
	lub := corners[cube.CP[2]][(1+cube.CO[2])%3]
	bur := corners[cube.CP[3]][(1+cube.CO[3])%3]
	fdr := corners[cube.CP[4]][(1+cube.CO[4])%3]
	ldf := corners[cube.CP[5]][(1+cube.CO[5])%3]
	bdl := corners[cube.CP[6]][(1+cube.CO[6])%3]
	rdb := corners[cube.CP[7]][(1+cube.CO[7])%3]

	fur := corners[cube.CP[0]][(2+cube.CO[0])%3]
	luf := corners[cube.CP[1]][(2+cube.CO[1])%3]
	bul := corners[cube.CP[2]][(2+cube.CO[2])%3]
	rub := corners[cube.CP[3]][(2+cube.CO[3])%3]
	rdf := corners[cube.CP[4]][(2+cube.CO[4])%3]
	fdl := corners[cube.CP[5]][(2+cube.CO[5])%3]
	ldb := corners[cube.CP[6]][(2+cube.CO[6])%3]
	bdr := corners[cube.CP[7]][(2+cube.CO[7])%3]

	uf := edges[cube.EP[0]][cube.EO[0]]
	ul := edges[cube.EP[1]][cube.EO[1]]
	ub := edges[cube.EP[2]][cube.EO[2]]
	ur := edges[cube.EP[3]][cube.EO[3]]
	fr := edges[cube.EP[4]][cube.EO[4]]
	fl := edges[cube.EP[5]][cube.EO[5]]
	bl := edges[cube.EP[6]][cube.EO[6]]
	br := edges[cube.EP[7]][cube.EO[7]]
	df := edges[cube.EP[8]][cube.EO[8]]
	dl := edges[cube.EP[9]][cube.EO[9]]
	db := edges[cube.EP[10]][cube.EO[10]]
	dr := edges[cube.EP[11]][cube.EO[11]]

	fu := edges[cube.EP[0]][(1+cube.EO[0])%2]
	lu := edges[cube.EP[1]][(1+cube.EO[1])%2]
	bu := edges[cube.EP[2]][(1+cube.EO[2])%2]
	ru := edges[cube.EP[3]][(1+cube.EO[3])%2]
	rf := edges[cube.EP[4]][(1+cube.EO[4])%2]
	lf := edges[cube.EP[5]][(1+cube.EO[5])%2]
	lb := edges[cube.EP[6]][(1+cube.EO[6])%2]
	rb := edges[cube.EP[7]][(1+cube.EO[7])%2]
	fd := edges[cube.EP[8]][(1+cube.EO[8])%2]
	ld := edges[cube.EP[9]][(1+cube.EO[9])%2]
	bd := edges[cube.EP[10]][(1+cube.EO[10])%2]
	rd := edges[cube.EP[11]][(1+cube.EO[11])%2]

	return fmt.Sprintf(
		`   %c%c%c
   %c%c%c
   %c%c%c
%c%c%c%c%c%c%c%c%c%c%c%c
%c%c%c%c%c%c%c%c%c%c%c%c
%c%c%c%c%c%c%c%c%c%c%c%c
   %c%c%c
   %c%c%c
   %c%c%c`,
		ubl, ub, ubr, ul, u, ur, ufl, uf, ufr,
		lub, lu, luf, ful, fu, fur, ruf, ru, rub, bur, bu, bul,
		lb, l, lf, fl, f, fr, rf, r, rb, br, b, bl,
		ldb, ld, ldf, fdl, fd, fdr, rdf, rd, rdb, bdr, bd, bdl,
		dfl, df, dfr, dl, d, dr, dbl, db, dbr,
	)
}

func ToProjectedString(cube Cube) string {
	/*
	   LUUUR
	   LUUUR
	   LUUUR
	   LFFFR
	   LFFFR
	   LFFFR
	*/

	centres := map[CentreP]byte{
		U: 'U',
		D: 'D',
		F: 'F',
		B: 'B',
		L: 'L',
		R: 'R',
	}
	corners := map[CornerP][3]byte{
		UFR: {'U', 'R', 'F'},
		UFL: {'U', 'F', 'L'},
		UBL: {'U', 'L', 'B'},
		UBR: {'U', 'B', 'R'},
		DFR: {'D', 'F', 'R'},
		DFL: {'D', 'L', 'F'},
		DBL: {'D', 'B', 'L'},
		DBR: {'D', 'R', 'B'},
	}
	edges := map[EdgeP][2]byte{
		UF: {'U', 'F'},
		UL: {'U', 'L'},
		UB: {'U', 'B'},
		UR: {'U', 'R'},
		FR: {'F', 'R'},
		FL: {'F', 'L'},
		BL: {'B', 'L'},
		BR: {'B', 'R'},
		DF: {'D', 'F'},
		DL: {'D', 'L'},
		DB: {'D', 'B'},
		DR: {'D', 'R'},
	}

	u := centres[cube.TP[0]]
	f := centres[cube.TP[2]]

	ufr := corners[cube.CP[0]][cube.CO[0]]
	ufl := corners[cube.CP[1]][cube.CO[1]]
	ubl := corners[cube.CP[2]][cube.CO[2]]
	ubr := corners[cube.CP[3]][cube.CO[3]]

	ruf := corners[cube.CP[0]][(1+cube.CO[0])%3]
	ful := corners[cube.CP[1]][(1+cube.CO[1])%3]
	lub := corners[cube.CP[2]][(1+cube.CO[2])%3]
	fdr := corners[cube.CP[4]][(1+cube.CO[4])%3]
	ldf := corners[cube.CP[5]][(1+cube.CO[5])%3]

	fur := corners[cube.CP[0]][(2+cube.CO[0])%3]
	luf := corners[cube.CP[1]][(2+cube.CO[1])%3]
	rub := corners[cube.CP[3]][(2+cube.CO[3])%3]
	rdf := corners[cube.CP[4]][(2+cube.CO[4])%3]
	fdl := corners[cube.CP[5]][(2+cube.CO[5])%3]

	uf := edges[cube.EP[0]][cube.EO[0]]
	ul := edges[cube.EP[1]][cube.EO[1]]
	ub := edges[cube.EP[2]][cube.EO[2]]
	ur := edges[cube.EP[3]][cube.EO[3]]
	fr := edges[cube.EP[4]][cube.EO[4]]
	fl := edges[cube.EP[5]][cube.EO[5]]

	fu := edges[cube.EP[0]][(1+cube.EO[0])%2]
	lu := edges[cube.EP[1]][(1+cube.EO[1])%2]
	ru := edges[cube.EP[3]][(1+cube.EO[3])%2]
	rf := edges[cube.EP[4]][(1+cube.EO[4])%2]
	lf := edges[cube.EP[5]][(1+cube.EO[5])%2]
	fd := edges[cube.EP[8]][(1+cube.EO[8])%2]

	return fmt.Sprintf(`%c%c%c%c%c
%c%c%c%c%c
%c%c%c%c%c
%c%c%c%c%c
%c%c%c%c%c
%c%c%c%c%c`,
		lub, ubl, ub, ubr, rub,
		lu, ul, u, ur, ru,
		luf, ufl, uf, ufr, ruf,
		luf, ful, fu, fur, ruf,
		lf, fl, f, fr, rf,
		ldf, fdl, fd, fdr, rdf,
	)
}

func ToFatString(cube Cube, method func(Cube) string) string {
	regular := method(cube)
	fat := ""
	for _, char := range regular {
		if char == '\n' {
			fat += "\n"
			continue
		}
		fat += string(char) + string(char)
	}
	return fat
}
