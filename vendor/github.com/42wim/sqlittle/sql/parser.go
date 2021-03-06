// Code generated by goyacc -o parser.go parser.go.y. DO NOT EDIT.

//line parser.go.y:2
package sql

import __yyfmt__ "fmt"

//line parser.go.y:2

//line parser.go.y:5
type yySymType struct {
	yys                  int
	literal              string
	identifier           string
	signedNumber         int64
	statement            interface{}
	columnNameList       []string
	columnName           string
	columnDefList        []ColumnDef
	columnDef            ColumnDef
	indexedColumnList    []IndexedColumn
	indexedColumn        IndexedColumn
	name                 string
	withoutRowid         bool
	unique               bool
	bool                 bool
	collate              string
	sortOrder            SortOrder
	columnConstraint     columnConstraint
	columnConstraintList []columnConstraint
	tableConstraint      TableConstraint
	tableConstraintList  []TableConstraint
	foreignKeyClause     ForeignKeyClause
	triggerAction        TriggerAction
	trigger              Trigger
	triggerList          []Trigger
	where                Expression
	expr                 Expression
	exprList             []Expression
	float                float64
}

const ABORT = 57346
const ACTION = 57347
const AND = 57348
const ASC = 57349
const AUTOINCREMENT = 57350
const CASCADE = 57351
const CHECK = 57352
const COLLATE = 57353
const CONFLICT = 57354
const CONSTRAINT = 57355
const CREATE = 57356
const DEFAULT = 57357
const DEFERRABLE = 57358
const DEFERRED = 57359
const DELETE = 57360
const DESC = 57361
const FAIL = 57362
const FOREIGN = 57363
const FROM = 57364
const GLOB = 57365
const IGNORE = 57366
const IN = 57367
const INDEX = 57368
const INITIALLY = 57369
const IS = 57370
const KEY = 57371
const LIKE = 57372
const MATCH = 57373
const NO = 57374
const NOT = 57375
const NULL = 57376
const ON = 57377
const OR = 57378
const PRIMARY = 57379
const REFERENCES = 57380
const REGEXP = 57381
const REPLACE = 57382
const RESTRICT = 57383
const ROLLBACK = 57384
const ROWID = 57385
const SELECT = 57386
const SET = 57387
const TABLE = 57388
const UNIQUE = 57389
const UPDATE = 57390
const WHERE = 57391
const WITHOUT = 57392
const tBare = 57393
const tLiteral = 57394
const tIdentifier = 57395
const tOperator = 57396
const tSignedNumber = 57397
const tFloat = 57398

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ABORT",
	"ACTION",
	"AND",
	"ASC",
	"AUTOINCREMENT",
	"CASCADE",
	"CHECK",
	"COLLATE",
	"CONFLICT",
	"CONSTRAINT",
	"CREATE",
	"DEFAULT",
	"DEFERRABLE",
	"DEFERRED",
	"DELETE",
	"DESC",
	"FAIL",
	"FOREIGN",
	"FROM",
	"GLOB",
	"IGNORE",
	"IN",
	"INDEX",
	"INITIALLY",
	"IS",
	"KEY",
	"LIKE",
	"MATCH",
	"NO",
	"NOT",
	"NULL",
	"ON",
	"OR",
	"PRIMARY",
	"REFERENCES",
	"REGEXP",
	"REPLACE",
	"RESTRICT",
	"ROLLBACK",
	"ROWID",
	"SELECT",
	"SET",
	"TABLE",
	"UNIQUE",
	"UPDATE",
	"WHERE",
	"WITHOUT",
	"tBare",
	"tLiteral",
	"tIdentifier",
	"tOperator",
	"tSignedNumber",
	"tFloat",
	"'-'",
	"'+'",
	"','",
	"'('",
	"')'",
	"'*'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 88,
	60, 6,
	-2, 94,
	-1, 89,
	60, 7,
	-2, 95,
}

const yyPrivate = 57344

const yyLast = 229

var yyAct = [...]int{

	174, 84, 82, 51, 9, 129, 85, 11, 79, 69,
	86, 99, 151, 149, 80, 158, 109, 20, 155, 123,
	11, 23, 154, 25, 153, 83, 11, 121, 56, 28,
	34, 34, 36, 11, 117, 149, 28, 150, 109, 31,
	145, 60, 88, 87, 89, 106, 71, 93, 91, 92,
	95, 90, 76, 67, 77, 113, 68, 55, 116, 115,
	78, 113, 143, 114, 116, 115, 53, 63, 126, 114,
	102, 24, 109, 113, 110, 12, 116, 115, 134, 103,
	104, 114, 108, 13, 107, 14, 40, 98, 41, 70,
	97, 74, 75, 118, 10, 15, 17, 43, 103, 104,
	52, 166, 119, 120, 122, 19, 74, 75, 11, 157,
	71, 130, 72, 73, 37, 131, 137, 138, 139, 140,
	142, 30, 136, 135, 132, 11, 39, 96, 130, 146,
	144, 71, 93, 91, 92, 12, 71, 152, 72, 73,
	64, 172, 18, 13, 13, 14, 14, 29, 6, 59,
	12, 11, 64, 29, 161, 61, 12, 162, 13, 66,
	14, 176, 164, 168, 13, 57, 14, 48, 50, 42,
	181, 173, 49, 33, 179, 58, 27, 94, 5, 169,
	65, 148, 21, 170, 178, 160, 100, 8, 128, 180,
	46, 45, 163, 177, 44, 52, 39, 175, 101, 171,
	112, 167, 125, 35, 47, 54, 22, 182, 38, 156,
	147, 127, 141, 81, 133, 159, 165, 32, 124, 111,
	62, 16, 26, 7, 105, 4, 3, 2, 1,
}
var yyPact = [...]int{

	134, -1000, -1000, -1000, -1000, 32, 49, 83, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 93, 156, -1000, 32, 93,
	11, 93, -1000, -1000, 107, 86, -20, -1000, 93, 93,
	93, 113, 27, 157, 6, 157, -3, 128, -1000, 93,
	183, 17, 157, -1000, 151, -1000, 125, -1000, -7, 55,
	93, -1000, 93, 81, 157, -9, -1000, 148, -10, 98,
	-1000, 128, -1000, 44, -1000, 179, -1000, -9, -1000, -1000,
	-1000, -1000, 81, 81, -1000, -1000, -1000, -15, 23, 13,
	-1000, 189, 19, -1000, -26, -1000, -1000, -1000, -1000, -1000,
	-9, 76, 76, -1000, -33, -9, -41, -1000, -1000, 194,
	-1000, -1000, 7, -1000, -1000, 172, 92, -1000, 81, -9,
	29, 179, 40, -9, -9, -9, -9, -9, 1, -1000,
	-1000, -9, -21, 92, -1000, -1000, -1000, 154, -1000, -24,
	-1000, -49, -1000, -1000, -9, -1000, -1000, 19, 19, 19,
	19, -37, 19, -1000, -43, 74, -46, -1000, 168, 92,
	-1000, -1000, 19, -1000, -9, -1000, -1000, 180, 62, 66,
	-1000, -1000, 19, 159, -1000, -1000, 123, -1000, -1000, -1000,
	-1000, -1000, 152, 152, -1000, 155, -1000, -1000, 202, -1000,
	-1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 228, 227, 226, 225, 1, 9, 6, 10, 4,
	187, 5, 224, 223, 222, 176, 8, 14, 173, 114,
	221, 220, 219, 11, 218, 97, 169, 28, 217, 3,
	0, 216, 215, 214, 213, 2, 212, 211, 210, 209,
}
var yyR1 = [...]int{

	0, 1, 1, 1, 6, 6, 5, 5, 7, 7,
	7, 8, 8, 8, 9, 9, 11, 11, 12, 10,
	10, 13, 13, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 26, 26, 26, 27, 27, 27, 29,
	19, 19, 28, 28, 28, 24, 24, 14, 14, 15,
	15, 18, 18, 18, 18, 22, 22, 23, 23, 23,
	21, 21, 20, 20, 39, 39, 39, 39, 39, 39,
	16, 16, 34, 17, 30, 30, 30, 30, 30, 31,
	31, 32, 32, 37, 37, 38, 38, 33, 33, 35,
	35, 35, 35, 35, 35, 35, 35, 35, 35, 35,
	35, 36, 36, 36, 2, 3, 4,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	2, 1, 2, 2, 1, 1, 1, 3, 3, 1,
	1, 1, 3, 4, 1, 2, 1, 4, 2, 2,
	2, 2, 1, 0, 1, 2, 5, 5, 6, 6,
	0, 2, 0, 3, 4, 0, 1, 1, 3, 3,
	3, 0, 1, 4, 6, 0, 2, 0, 1, 1,
	0, 2, 0, 1, 0, 3, 3, 3, 3, 3,
	1, 3, 1, 3, 2, 2, 1, 1, 2, 3,
	3, 0, 2, 0, 1, 0, 2, 0, 2, 1,
	4, 1, 1, 1, 1, 1, 3, 3, 3, 3,
	3, 0, 1, 3, 4, 8, 10,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, 44, 14, -13, -10, -9,
	62, -5, 43, 51, 53, 46, -20, 47, 59, 22,
	-5, 26, -10, -5, 60, -5, -14, -15, -9, 40,
	35, 59, -28, -18, -5, -18, -5, -19, -15, 13,
	59, 61, -26, -25, 37, 34, 33, 47, 10, 15,
	11, -29, 38, 60, -26, 60, -27, 37, 47, 21,
	-5, -19, -21, 50, -25, 29, 34, 60, -7, -6,
	34, 55, 57, 58, 51, 52, -5, -5, -7, -16,
	-17, -34, -35, 34, -5, -7, -8, 52, 51, 53,
	60, 57, 58, 56, 29, 60, 29, -27, 43, -23,
	7, 19, -35, -7, -7, -12, 60, 61, 59, 59,
	61, -22, 11, 54, 62, 58, 57, 60, -35, -8,
	-8, 60, -16, 60, -24, 8, 61, -37, 16, -11,
	-9, -7, -17, -33, 49, -23, -6, -35, -35, -35,
	-35, -36, -35, 61, -16, 61, -11, -38, 27, 59,
	61, 61, -35, 61, 59, 61, -39, 35, 61, -32,
	17, -9, -35, 12, -29, -31, 35, 42, 4, 20,
	24, 40, 18, 48, -30, 45, 9, 41, 32, -30,
	34, 15, 5,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 0, 62, 0, 21, 19,
	20, 14, 15, 6, 7, 0, 0, 63, 0, 0,
	0, 0, 22, 104, 0, 0, 42, 47, 51, 51,
	0, 40, 0, 33, 52, 33, 0, 0, 48, 0,
	40, 60, 49, 34, 0, 24, 0, 26, 0, 0,
	0, 32, 0, 0, 50, 0, 43, 0, 0, 0,
	41, 0, 105, 0, 35, 57, 25, 0, 28, 29,
	30, 8, 0, 0, 4, 5, 31, 0, 0, 0,
	70, 55, 72, 89, 0, 91, 92, 93, -2, -2,
	0, 0, 0, 11, 0, 0, 0, 44, 61, 45,
	58, 59, 0, 9, 10, 83, 0, 53, 0, 0,
	87, 57, 0, 0, 0, 0, 0, 101, 0, 12,
	13, 0, 0, 0, 23, 46, 27, 85, 84, 0,
	16, 0, 71, 106, 0, 73, 56, 96, 97, 98,
	99, 0, 102, 100, 0, 64, 0, 81, 0, 0,
	18, 54, 88, 90, 0, 36, 37, 0, 0, 39,
	86, 17, 103, 0, 38, 82, 0, 65, 66, 67,
	68, 69, 0, 0, 79, 0, 76, 77, 0, 80,
	74, 75, 78,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	60, 61, 62, 58, 59, 57,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ??, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:131
		{
			yyVAL.literal = yyDollar[1].identifier
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:134
		{
			yyVAL.literal = yyDollar[1].identifier
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:139
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:142
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:147
		{
			yyVAL.signedNumber = yyDollar[1].signedNumber
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:150
		{
			yyVAL.signedNumber = -yyDollar[2].signedNumber
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:153
		{
			yyVAL.signedNumber = yyDollar[2].signedNumber
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:158
		{
			yyVAL.float = yyDollar[1].float
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:161
		{
			yyVAL.float = -yyDollar[2].float
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:164
		{
			yyVAL.float = yyDollar[2].float
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:169
		{
			yyVAL.columnName = yyDollar[1].identifier
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:172
		{
			yyVAL.columnName = "ROWID"
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:177
		{
			yyVAL.columnNameList = []string{yyDollar[1].columnName}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:180
		{
			yyVAL.columnNameList = append(yyDollar[1].columnNameList, yyDollar[3].columnName)
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:185
		{
			yyVAL.columnNameList = yyDollar[2].columnNameList
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:190
		{
			yyVAL.columnName = yyDollar[1].columnName
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:193
		{
			yyVAL.columnName = "*"
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:198
		{
			yyVAL.columnNameList = []string{yyDollar[1].columnName}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:201
		{
			yyVAL.columnNameList = append(yyDollar[1].columnNameList, yyDollar[3].columnName)
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:207
		{
			yyVAL.columnConstraint = ccPrimaryKey{yyDollar[3].sortOrder, yyDollar[4].bool}
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:210
		{
			yyVAL.columnConstraint = ccNull(true)
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:213
		{
			yyVAL.columnConstraint = ccNull(false)
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:216
		{
			yyVAL.columnConstraint = ccUnique(true)
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:219
		{
			yyVAL.columnConstraint = ccCheck{expr: yyDollar[3].expr}
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:222
		{
			yyVAL.columnConstraint = ccDefault(yyDollar[2].signedNumber)
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:225
		{
			yyVAL.columnConstraint = ccDefault(yyDollar[2].literal)
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:228
		{
			yyVAL.columnConstraint = ccDefault(nil)
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:231
		{
			yyVAL.columnConstraint = ccCollate(yyDollar[2].identifier)
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:234
		{
			yyVAL.columnConstraint = ccReferences(yyDollar[1].foreignKeyClause)
		}
	case 33:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:239
		{
			yyVAL.columnConstraintList = nil
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:242
		{
			yyVAL.columnConstraintList = []columnConstraint{yyDollar[1].columnConstraint}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:245
		{
			yyVAL.columnConstraintList = append(yyDollar[1].columnConstraintList, yyDollar[2].columnConstraint)
		}
	case 36:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:250
		{
			yyVAL.tableConstraint = TablePrimaryKey{yyDollar[4].indexedColumnList}
		}
	case 37:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:253
		{
			yyVAL.tableConstraint = TableUnique{
				IndexedColumns: yyDollar[3].indexedColumnList,
			}
		}
	case 38:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:258
		{
			yyVAL.tableConstraint = TableForeignKey{
				Columns: yyDollar[4].columnNameList,
				Clause:  yyDollar[6].foreignKeyClause,
			}
		}
	case 39:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:266
		{
			yyVAL.foreignKeyClause = ForeignKeyClause{
				ForeignTable:      yyDollar[2].identifier,
				ForeignColumns:    yyDollar[3].columnNameList,
				Deferrable:        yyDollar[4].bool,
				InitiallyDeferred: yyDollar[5].bool,
				Triggers:          yyDollar[6].triggerList,
			}
		}
	case 40:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:277
		{
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:278
		{
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:282
		{
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:283
		{
			yyVAL.tableConstraintList = []TableConstraint{yyDollar[3].tableConstraint}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:286
		{
			yyVAL.tableConstraintList = append(yyDollar[1].tableConstraintList, yyDollar[4].tableConstraint)
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:292
		{
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:293
		{
			yyVAL.bool = true
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:298
		{
			yyVAL.columnDefList = []ColumnDef{yyDollar[1].columnDef}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:301
		{
			yyVAL.columnDefList = append(yyDollar[1].columnDefList, yyDollar[3].columnDef)
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:306
		{
			yyVAL.columnDef = makeColumnDef(yyDollar[1].columnName, yyDollar[2].name, yyDollar[3].columnConstraintList)
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:309
		{
			yyVAL.columnDef = makeColumnDef("REPLACE", yyDollar[2].name, yyDollar[3].columnConstraintList)
		}
	case 51:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:314
		{
			yyVAL.name = ""
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:317
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 53:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:320
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 54:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:323
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 55:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:328
		{
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:329
		{
			yyVAL.collate = yyDollar[2].literal
		}
	case 57:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:334
		{
			yyVAL.sortOrder = Asc
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:337
		{
			yyVAL.sortOrder = Asc
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:340
		{
			yyVAL.sortOrder = Desc
		}
	case 60:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:345
		{
			yyVAL.withoutRowid = false
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:348
		{
			yyVAL.withoutRowid = true
		}
	case 62:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:353
		{
			yyVAL.unique = false
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:356
		{
			yyVAL.unique = true
		}
	case 64:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:361
		{
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:362
		{
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:364
		{
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:366
		{
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:368
		{
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:370
		{
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:374
		{
			yyVAL.indexedColumnList = []IndexedColumn{yyDollar[1].indexedColumn}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:377
		{
			yyVAL.indexedColumnList = append(yyDollar[1].indexedColumnList, yyDollar[3].indexedColumn)
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:382
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:387
		{
			yyVAL.indexedColumn = newIndexColumn(yyDollar[1].expr, yyDollar[2].collate, yyDollar[3].sortOrder)
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:392
		{
			yyVAL.triggerAction = ActionSetNull
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:395
		{
			yyVAL.triggerAction = ActionSetDefault
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:398
		{
			yyVAL.triggerAction = ActionCascade
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:401
		{
			yyVAL.triggerAction = ActionRestrict
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:404
		{
			yyVAL.triggerAction = ActionNoAction
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:409
		{
			yyVAL.trigger = TriggerOnDelete(yyDollar[3].triggerAction)
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:412
		{
			yyVAL.trigger = TriggerOnUpdate(yyDollar[3].triggerAction)
		}
	case 81:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:417
		{
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:418
		{
			yyVAL.triggerList = append(yyDollar[1].triggerList, yyDollar[2].trigger)
		}
	case 83:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:423
		{
			yyVAL.bool = false
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:426
		{
			yyVAL.bool = true
		}
	case 85:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:431
		{
			yyVAL.bool = false
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:434
		{
			yyVAL.bool = true
		}
	case 87:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:439
		{
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:440
		{
			yyVAL.where = yyDollar[2].expr
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:445
		{
			yyVAL.expr = nil
		}
	case 90:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:448
		{
			yyVAL.expr = ExFunction{yyDollar[1].identifier, yyDollar[3].exprList}
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:451
		{
			yyVAL.expr = yyDollar[1].signedNumber
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:454
		{
			yyVAL.expr = yyDollar[1].float
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:457
		{
			yyVAL.expr = yyDollar[1].identifier
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:460
		{
			yyVAL.expr = ExColumn(yyDollar[1].identifier)
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:463
		{
			yyVAL.expr = ExColumn(yyDollar[1].identifier)
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:466
		{
			yyVAL.expr = ExBinaryOp{yyDollar[2].identifier, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:469
		{
			yyVAL.expr = ExBinaryOp{"*", yyDollar[1].expr, yyDollar[3].expr}
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:472
		{
			yyVAL.expr = ExBinaryOp{"+", yyDollar[1].expr, yyDollar[3].expr}
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:475
		{
			yyVAL.expr = ExBinaryOp{"-", yyDollar[1].expr, yyDollar[3].expr}
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:478
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 101:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:483
		{
			yyVAL.exprList = nil
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:486
		{
			yyVAL.exprList = []Expression{yyDollar[1].expr}
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:489
		{
			yyVAL.exprList = append(yyDollar[1].exprList, yyDollar[3].expr)
		}
	case 104:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:494
		{
			yylex.(*lexer).result = SelectStmt{Columns: yyDollar[2].columnNameList, Table: yyDollar[4].identifier}
		}
	case 105:
		yyDollar = yyS[yypt-8 : yypt+1]
//line parser.go.y:499
		{
			yylex.(*lexer).result = CreateTableStmt{
				Table:        yyDollar[3].identifier,
				Columns:      yyDollar[5].columnDefList,
				Constraints:  yyDollar[6].tableConstraintList,
				WithoutRowid: yyDollar[8].withoutRowid,
			}
		}
	case 106:
		yyDollar = yyS[yypt-10 : yypt+1]
//line parser.go.y:509
		{
			yylex.(*lexer).result = CreateIndexStmt{
				Index:          yyDollar[4].identifier,
				Table:          yyDollar[6].identifier,
				Unique:         yyDollar[2].unique,
				IndexedColumns: yyDollar[8].indexedColumnList,
				Where:          yyDollar[10].where,
			}
		}
	}
	goto yystack /* stack new state and value */
}
