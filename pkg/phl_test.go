package phl

import (
	"regexp/syntax"
	"testing"
)

//	+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++	//

type (

	//	Data structure for
	//	describing the test case for "parseField" func
	testCaseParseField struct {
		arg testCaseParseFieldArg
		exp testCaseParseFieldExp
	}

	//	Data structure for
	//	describing the arg of test case for "parseField" func
	testCaseParseFieldArg struct {
		a1 string
		a2 string
	}

	//	Data structure for
	//	describing the exp of test case for "parseField" func
	testCaseParseFieldExp struct {
		r1 string
		r2 error
	}
)

//	---------------------------------------------------------------------	//

var (

	//	Test cases for "parseField" func
	testCasesParseField = []testCaseParseField{
		{
			arg: testCaseParseFieldArg{
				a1: "\\d+",
				a2: "sdfsdf23472askjsad",
			},
			exp: testCaseParseFieldExp{
				r1: "23472",
				r2: nil,
			},
		},
		{
			arg: testCaseParseFieldArg{
				a1: "iot",
				a2: "sdfsdf23472askjsad",
			},
			exp: testCaseParseFieldExp{
				r1: "",
				r2: errNoFss,
			},
		},
		{
			arg: testCaseParseFieldArg{
				a1: "?",
				a2: "sdfsdf",
			},
			exp: testCaseParseFieldExp{
				r1: "",
				//	TODO: describe error
				r2: &syntax.Error{Code: "missing argument to repetition operator", Expr: "?"},
			},
		},
	}
)

//	---------------------------------------------------------------------	//

// Test "parseField" func
func TestParseField(t *testing.T) {

	for i, c := range testCasesParseField {

		m, mErr := parseField(c.arg.a1, c.arg.a2)

		if m != c.exp.r1 || mErr != c.exp.r2 {
			t.Fatalf(
				`
					Test failed:	Test (%d) for "parseField" func
									(returned (%s, %s) != expected (%s, %s))
				`,
				i,
				m,
				mErr,
				c.exp.r1,
				c.exp.r2,
			)
		}

	}

}

//	+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++	//

type (

	//	Data structure for
	//	describing the test case for "FillStruct" func
	testCaseFillStruct struct {
		arg testCaseFillStructArg
		exp testCaseFillStructExp
	}

	//	Data structure for
	//	describing the arg of test case for "FillStruct" func
	testCaseFillStructArg struct {
		d interface{}
		s string
		t string
	}

	//	Data structure for
	//	describing the exp of test case for "FillStruct" func
	testCaseFillStructExp struct {
		r1 bool
		r2 error
	}

	//	Data structure for
	//	describing the #1 d-param arg of test case for "FillStruct" func
	testCaseFillStructArgD1 struct {
		M string `regex:"sit ([a-z]+)"`
	}

	//	Data structure for
	//	describing the #1 d-param arg of test case for "FillStruct" func
	testCaseFillStructArgD2 struct {
		M string `regex:"?"`
	}
)

//	---------------------------------------------------------------------	//

var (

	//	Test cases for "FillStruct" func
	testCasesFillStruct = []testCaseFillStruct{
		{
			arg: testCaseFillStructArg{
				d: new(testCaseFillStructArgD1),
				s: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Massa placerat duis ultricies lacus.",
				t: "regex",
			},
			exp: testCaseFillStructExp{
				r1: true,
				r2: nil,
			},
		},
		{
			arg: testCaseFillStructArg{
				d: new(testCaseFillStructArgD2),
				s: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Massa placerat duis ultricies lacus.",
				t: "regex",
			},
			exp: testCaseFillStructExp{
				r1: false,
				//	TODO: describe error
				r2: &syntax.Error{Code: "missing argument to repetition operator", Expr: "?"},
			},
		},
	}
)

//	---------------------------------------------------------------------	//

// Test for "FillStruct" func
func TestFillStruct(t *testing.T) {

	for i, c := range testCasesFillStruct {

		f, fErr := FillStruct(c.arg.d, c.arg.s, c.arg.t)

		if f != c.exp.r1 || fErr != c.exp.r2 {
			t.Fatalf(
				`
					Test failed:	Test (%d) for "FillStruct" func
									(returned (%t, %s) != expected (%t, %s))
				`,
				i,
				f,
				fErr,
				c.exp.r1,
				c.exp.r2,
			)
		}

	}

}

//	+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++	//
