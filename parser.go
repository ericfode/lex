// Copyright (c) 2011 CZ.NIC z.s.p.o. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// blame: jnml, labs.nic.cz

// WARNING: If this file is parser.go then DO NOT EDIT.
// parser.go is generated by goyacc from parser.y (see the Makefile).

package lex

import (
	"fmt"
	"go/token"
)

var (
	defNamePos0 token.Position
)

type yySymType struct {
	yys  int
	str  string
	strs []string
}

const tSECTION_DIV = 57346
const tBLANKS = 57347
const tVERBATIM_OPEN = 57348
const tVERBATIM_CLOSE = 57349
const tSTARTS = 57350
const tSSTART = 57351
const tXSTART = 57352
const tYYT = 57353
const tYYB = 57354
const tYYC = 57355
const tYYN = 57356
const tDEF_NAME = 57357
const tDEFINITION = 57358
const tUNINDENTED_COMMENT = 57359
const tINDENTED_TEXT = 57360
const tVERBATIM_LINE = 57361
const tPATTERN_LINE = 57362
const tUSER_CODE_LINE = 57363
const tNAME = 57364
const tSTARTS_PATTERN_LINE = 57365

var yyToknames = []string{
	"tSECTION_DIV",
	"tBLANKS",
	"tVERBATIM_OPEN",
	"tVERBATIM_CLOSE",
	"tSTARTS",
	"tSSTART",
	"tXSTART",
	"tYYT",
	"tYYB",
	"tYYC",
	"tYYN",
	"tDEF_NAME",
	"tDEFINITION",
	"tUNINDENTED_COMMENT",
	"tINDENTED_TEXT",
	"tVERBATIM_LINE",
	"tPATTERN_LINE",
	"tUSER_CODE_LINE",
	"tNAME",
	"tSTARTS_PATTERN_LINE",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 47
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 67

var yyAct = []int{

	38, 6, 36, 11, 60, 61, 12, 13, 14, 15,
	16, 17, 8, 54, 9, 10, 63, 33, 55, 35,
	62, 7, 39, 66, 65, 57, 50, 43, 42, 32,
	59, 34, 47, 24, 51, 31, 46, 49, 37, 40,
	4, 64, 48, 45, 48, 52, 41, 44, 30, 26,
	27, 28, 19, 3, 23, 22, 21, 56, 20, 5,
	58, 25, 18, 2, 1, 53, 29,
}
var yyPact = []int{

	-1000, -1000, 36, -3, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 11,
	33, -1000, 34, 34, -1000, 33, -1000, -1000, -1000, 6,
	-1000, -1000, -1000, -1000, -1000, -1000, 20, -1000, 25, 32,
	4, 32, 18, -1000, -1000, -9, -1000, -1000, -1000, 3,
	-1000, -1000, 23, -21, -1000, -1000, -4, -1000, -8, -1000,
	-1000, 2, -1000, -1000, 0, -1000, -1000,
}
var yyPgo = []int{

	0, 0, 66, 33, 22, 65, 64, 63, 62, 2,
	61, 60, 59, 58, 57, 56, 55, 54, 53, 52,
	48, 47, 43, 41,
}
var yyR1 = []int{

	0, 6, 8, 6, 9, 9, 10, 11, 3, 12,
	13, 14, 12, 12, 12, 15, 12, 16, 12, 17,
	12, 12, 12, 12, 12, 18, 18, 7, 20, 20,
	21, 20, 20, 22, 23, 20, 19, 19, 4, 4,
	5, 5, 5, 2, 2, 1, 1,
}
var yyR2 = []int{

	0, 1, 0, 4, 0, 1, 0, 0, 5, 1,
	0, 0, 6, 1, 1, 0, 4, 0, 3, 0,
	3, 2, 2, 2, 2, 0, 2, 3, 1, 1,
	0, 4, 1, 0, 0, 6, 0, 2, 2, 3,
	1, 3, 1, 0, 2, 0, 2,
}
var yyChk = []int{

	-1000, -6, -7, -18, 4, -12, 4, 24, 15, 17,
	18, 6, 9, 10, 11, 12, 13, 14, -8, -19,
	-13, -15, -16, -17, -3, -10, -3, -3, -3, -2,
	-20, 24, 18, 6, 20, 8, -9, 5, -1, -4,
	5, -4, -9, 21, -21, -22, 16, 7, 19, 5,
	22, 16, -1, -5, 22, 27, -14, 22, -11, 7,
	25, 26, 24, 24, -23, 22, 23,
}
var yyDef = []int{

	25, -2, 1, 0, 2, 26, 36, 9, 10, 13,
	14, 15, 17, 19, 6, 6, 6, 6, 43, 27,
	4, 45, 0, 0, 21, 4, 22, 23, 24, 3,
	37, 28, 29, 30, 32, 33, 0, 5, 0, 18,
	0, 20, 0, 44, 45, 0, 11, 16, 46, 0,
	38, 7, 0, 0, 40, 42, 0, 39, 0, 31,
	34, 0, 12, 8, 0, 41, 35,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	24, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 27, 3, 26, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 25,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23,
}
var yyTok3 = []int{
	0,
}

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c > 0 && c <= len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return fmt.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return fmt.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		fmt.Printf("lex %U %s\n", uint(char), yyTokname(c))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		fmt.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
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
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
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
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
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
			if yyn < 0 || yyn == yychar {
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
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				fmt.Printf("%s", yyStatname(yystate))
				fmt.Printf("saw %s\n", yyTokname(yychar))
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

				/* the current p has no shift onn "error", pop stack */
				if yyDebug >= 2 {
					fmt.Printf("error recovery pops state %d, uncovers %d\n",
						yyS[yyp].yys, yyS[yyp-1].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				fmt.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		fmt.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
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

	case 2:

		{
			sc(yylex).Begin(_USER)
		}
	case 3:

		{
			usrCode = yyS[yypt-0].str
		}
	case 6:

		{
			sc(yylex).PushState(_DEF_NAME)
		}
	case 7:

		{
			sc(yylex).PopState()
		}
	case 8:

		{
			yyVAL.str = yyS[yypt-2].str
		}
	case 10:

		{
			defNamePos0 = sc(yylex).TokenStart()
			sc(yylex).PushState(_DEF_NAME)
		}
	case 11:

		{
			if _, ok := defs[yyS[yypt-3].str]; ok {
				logErr(fmt.Sprintf("%s:%q redefined", defNamePos0, yyS[yypt-3].str))
			}

			defs[yyS[yypt-3].str] = yyS[yypt-0].str
			defPos[yyS[yypt-3].str] = sc(yylex).TokenStart()
			sc(yylex).PopState()
		}
	case 13:

		{
			defCode = append(defCode, yyS[yypt-0].str+"\n")
		}
	case 14:

		{
			defCode = append(defCode, yyS[yypt-0].str+"\n")
		}
	case 15:

		{
			sc(yylex).PushState(_VERBATIM)
		}
	case 16:

		{
			defCode = append(defCode, yyS[yypt-1].str+"\n")
			sc(yylex).PopState()
		}
	case 17:

		{
			sc(yylex).PushState(_DEF_STARTS)
		}
	case 18:

		{
			sStarts = append(sStarts, yyS[yypt-0].strs...)
			sc(yylex).PopState()
		}
	case 19:

		{
			sc(yylex).PushState(_DEF_STARTS)
		}
	case 20:

		{
			xStarts = append(xStarts, yyS[yypt-0].strs...)
			for _, start := range yyS[yypt-0].strs {
				isXStart[start] = true
			}
			sc(yylex).PopState()
		}
	case 21:

		{
			_yyt = yyS[yypt-0].str
		}
	case 22:

		{
			_yyb = yyS[yypt-0].str
		}
	case 23:

		{
			_yyc = yyS[yypt-0].str
		}
	case 24:

		{
			_yyn = yyS[yypt-0].str
		}
	case 27:

		{
			for s := range unrefStarts {
				logErr(fmt.Sprintf("%s:start condition %q declared and not used", sc(yylex).TokenStart(), s))
			}
			if len(rules) == 1 {
				logErr(fmt.Sprintf("%s:no rules defined", sc(yylex).TokenStart()))
			}
		}
	case 29:

		{
			moreAction(yyS[yypt-0].str)
		}
	case 30:

		{
			sc(yylex).PushState(_VERBATIM)
		}
	case 31:

		{
			moreAction(yyS[yypt-1].str)
			sc(yylex).PopState()
		}
	case 32:

		{
			rulePos = append(rulePos, sc(yylex).TokenStart())
			pat, re, action, bol, eol := parsePattern(sc(yylex).TokenStart(), yyS[yypt-0].str, map[string]bool{})
			unreachableRules[len(rules)] = true
			rules = append(rules, rule{nil, pat, re, action, nil, nil, bol, eol})
		}
	case 33:

		{
			sc(yylex).PushState(_STARTS)
		}
	case 34:

		{
			sc(yylex).PopState()
		}
	case 35:

		{
			rulePos = append(rulePos, sc(yylex).TokenStart())
			pat, re, action, bol, eol := parsePattern(sc(yylex).TokenStart(), yyS[yypt-0].str, map[string]bool{})
			unreachableRules[len(rules)] = true
			rules = append(rules, rule{yyS[yypt-3].strs, pat, re, action, nil, nil, bol, eol})
		}
	case 36:

		{
			sc(yylex).PushState(_RULES)
		}
	case 38:

		{
			if !addStartSet(yyS[yypt-0].str) {
				logErr(fmt.Sprintf("%s:start condition %q redeclared", sc(yylex).TokenStart(), yyS[yypt-0].str))
			}
			yyVAL.strs = append(yyVAL.strs, yyS[yypt-0].str)
		}
	case 39:

		{
			if !addStartSet(yyS[yypt-0].str) {
				logErr(fmt.Sprintf("%s:start condition %q redeclared", sc(yylex).TokenStart(), yyS[yypt-0].str))
			}
			yyVAL.strs = append(yyVAL.strs, yyS[yypt-0].str)
		}
	case 40:

		{
			if _, ok := defStarts[yyS[yypt-0].str]; !ok {
				logErr(fmt.Sprintf("%s:start condition %q undefined", sc(yylex).TokenStart(), yyS[yypt-0].str))
			}
			delete(unrefStarts, yyS[yypt-0].str)
			yyVAL.strs = append(yyVAL.strs, yyS[yypt-0].str)
		}
	case 41:

		{
			if _, ok := defStarts[yyS[yypt-0].str]; !ok {
				logErr(fmt.Sprintf("%s:start condition %q undefined", sc(yylex).TokenStart(), yyS[yypt-0].str))
			}
			delete(unrefStarts, yyS[yypt-0].str)
			yyVAL.strs = append(yyVAL.strs, yyS[yypt-0].str)
		}
	case 42:

		{
			yyVAL.strs = append(yyVAL.strs, "*")
		}
	case 43:

		{
			yyVAL.str = ""
		}
	case 44:

		{
			yyVAL.str += yyS[yypt-0].str
		}
	case 45:

		{
			yyVAL.str = ""
		}
	case 46:

		{
			yyVAL.str += yyS[yypt-0].str
		}
	}
	goto yystack /* stack new state and value */
}
