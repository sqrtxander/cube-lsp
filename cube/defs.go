package cube

type (
	CentreP = int
	CornerP = int
	CornerO = int
	EdgeP   = int
	EdgeO   = int
)

const (
	U CentreP = iota
	D
	F
	B
	L
	R
)

const (
	UFR CornerP = iota
	UFL
	UBL
	UBR
	DFR
	DFL
	DBL
	DBR
)

const (
	GC CornerO = iota
	CCWC
	CWC
)

const (
	UF EdgeP = iota
	UL
	UB
	UR
	FR
	FL
	BL
	BR
	DF
	DL
	DB
	DR
)

const (
	GE EdgeO = iota
	BE
)

type Cube struct {
	TP [6]CentreP
	CP [8]CornerP
	CO [8]CornerO
	EP [12]EdgeP
	EO [12]EdgeO
}

var (
	centrePs = [6]CentreP{U, D, F, B, L, R}
	cornerPs = [8]CornerP{UFR, UFL, UBL, UBR, DFR, DFL, DBL, DBR}
	cornerOs = [8]CornerO{GC, GC, GC, GC, GC, GC, GC, GC}
	edgePs   = [12]EdgeP{UF, UL, UB, UR, FR, FL, BL, BR, DF, DL, DB, DR}
	edgeOs   = [12]EdgeO{GE, GE, GE, GE, GE, GE, GE, GE, GE, GE, GE, GE}
	moves    = []Cube{
		// U
		{
			TP: [6]CentreP{U, D, F, B, L, R},
			CP: [8]CornerP{UBR, UFR, UFL, UBL, DFR, DFL, DBL, DBR},
			CO: [8]CornerO{GC, GC, GC, GC, GC, GC, GC, GC},
			EP: [12]EdgeP{UR, UF, UL, UB, FR, FL, BL, BR, DF, DL, DB, DR},
			EO: [12]EdgeO{GE, GE, GE, GE, GE, GE, GE, GE, GE, GE, GE, GE},
		},

		// D
		{
			TP: [6]CentreP{U, D, F, B, L, R},
			CP: [8]CornerP{UFR, UFL, UBL, UBR, DFL, DBL, DBR, DFR},
			CO: [8]CornerO{GC, GC, GC, GC, GC, GC, GC, GC},
			EP: [12]EdgeP{UF, UL, UB, UR, FR, FL, BL, BR, DL, DB, DR, DF},
			EO: [12]EdgeO{GE, GE, GE, GE, GE, GE, GE, GE, GE, GE, GE, GE},
		},

		// F
		{
			TP: [6]CentreP{U, D, F, B, L, R},
			CP: [8]CornerP{UFL, DFL, UBL, UBR, UFR, DFR, DBL, DBR},
			CO: [8]CornerO{CWC, CCWC, GC, GC, CCWC, CWC, GC, GC},
			EP: [12]EdgeP{FL, UL, UB, UR, UF, DF, BL, BR, FR, DL, DB, DR},
			EO: [12]EdgeO{BE, GE, GE, GE, BE, BE, GE, GE, BE, GE, GE, GE},
		},

		// B
		{
			TP: [6]CentreP{U, D, F, B, L, R},
			CP: [8]CornerP{UFR, UFL, UBR, DBR, DFR, DFL, UBL, DBL},
			CO: [8]CornerO{GC, GC, CWC, CCWC, GC, GC, CCWC, CWC},
			EP: [12]EdgeP{UF, UL, BR, UR, FR, FL, UB, DB, DF, DL, BL, DR},
			EO: [12]EdgeO{GE, GE, BE, GE, GE, GE, BE, BE, GE, GE, BE, GE},
		},

		// L
		{
			TP: [6]CentreP{U, D, F, B, L, R},
			CP: [8]CornerP{UFR, UBL, DBL, UBR, DFR, UFL, DFL, DBR},
			CO: [8]CornerO{GC, CWC, CCWC, GC, GC, CCWC, CWC, GC},
			EP: [12]EdgeP{UF, BL, UB, UR, FR, UL, DL, BR, DF, FL, DB, DR},
			EO: [12]EdgeO{GE, GE, GE, GE, GE, GE, GE, GE, GE, GE, GE, GE},
		},

		// R
		{
			TP: [6]CentreP{U, D, F, B, L, R},
			CP: [8]CornerP{DFR, UFL, UBL, UFR, DBR, DFL, DBL, UBR},
			CO: [8]CornerO{CCWC, GC, GC, CWC, CWC, GC, GC, CCWC},
			EP: [12]EdgeP{UF, UL, UB, FR, DR, FL, BL, UR, DF, DL, DB, BR},
			EO: [12]EdgeO{GE, GE, GE, GE, GE, GE, GE, GE, GE, GE, GE, GE},
		},

		// M
		{
			TP: [6]CentreP{B, F, U, D, L, R},
			CP: [8]CornerP{UFR, UFL, UBL, UBR, DFR, DFL, DBL, DBR},
			CO: [8]CornerO{GC, GC, GC, GC, GC, GC, GC, GC},
			EP: [12]EdgeP{UB, UL, DB, UR, FR, FL, BL, BR, UF, DL, DF, DR},
			EO: [12]EdgeO{BE, GE, BE, GE, GE, GE, GE, GE, BE, GE, BE, GE},
		},

		// E
		{
			TP: [6]CentreP{U, D, L, R, B, F},
			CP: [8]CornerP{UFR, UFL, UBL, UBR, DFR, DFL, DBL, DBR},
			CO: [8]CornerO{GC, GC, GC, GC, GC, GC, GC, GC},
			EP: [12]EdgeP{UF, UL, UB, UR, FL, BL, BR, FR, DF, DL, DB, DR},
			EO: [12]EdgeO{GE, GE, GE, GE, BE, BE, BE, BE, GE, GE, GE, GE},
		},

		// S
		{
			TP: [6]CentreP{L, R, F, B, D, U},
			CP: [8]CornerP{UFR, UFL, UBL, UBR, DFR, DFL, DBL, DBR},
			CO: [8]CornerO{GC, GC, GC, GC, GC, GC, GC, GC},
			EP: [12]EdgeP{UF, DL, UB, UL, FR, FL, BL, BR, DF, DR, DB, UR},
			EO: [12]EdgeO{GE, BE, GE, BE, GE, GE, GE, GE, GE, BE, GE, BE},
		},
	}
)

func init() {
	// u
	on := *GetSolvedCube()
	on = *multiply(on, moves[0])
	on = *multiply(on, moves[7])
	on = *multiply(on, moves[7])
	on = *multiply(on, moves[7])
	moves = append(moves, on)

	// d
	on = *GetSolvedCube()
	on = *multiply(on, moves[1])
	on = *multiply(on, moves[7])
	moves = append(moves, on)

	// f
	on = *GetSolvedCube()
	on = *multiply(on, moves[2])
	on = *multiply(on, moves[8])
	moves = append(moves, on)

	// b
	on = *GetSolvedCube()
	on = *multiply(on, moves[3])
	on = *multiply(on, moves[8])
	on = *multiply(on, moves[8])
	on = *multiply(on, moves[8])
	moves = append(moves, on)

	// l
	on = *GetSolvedCube()
	on = *multiply(on, moves[4])
	on = *multiply(on, moves[6])
	moves = append(moves, on)

	// r
	on = *GetSolvedCube()
	on = *multiply(on, moves[5])
	on = *multiply(on, moves[6])
	on = *multiply(on, moves[6])
	on = *multiply(on, moves[6])
	moves = append(moves, on)

	// x
	on = *GetSolvedCube()
	on = *multiply(on, moves[4])
	on = *multiply(on, moves[4])
	on = *multiply(on, moves[4])
	on = *multiply(on, moves[6])
	on = *multiply(on, moves[6])
	on = *multiply(on, moves[6])
	on = *multiply(on, moves[5])
	moves = append(moves, on)

	// y
	on = *GetSolvedCube()
	on = *multiply(on, moves[0])
	on = *multiply(on, moves[7])
	on = *multiply(on, moves[7])
	on = *multiply(on, moves[7])
	on = *multiply(on, moves[1])
	on = *multiply(on, moves[1])
	on = *multiply(on, moves[1])
	moves = append(moves, on)

	// z
	on = *GetSolvedCube()
	on = *multiply(on, moves[2])
	on = *multiply(on, moves[8])
	on = *multiply(on, moves[3])
	moves = append(moves, on)
}
