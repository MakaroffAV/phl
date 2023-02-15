package phl

import (
	"errors"
	"reflect"
	"regexp"
)

//	+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++	//

var (

	//	No found string submatch
	errNoFss error = errors.New("noMatchesFoundWithRegex")
)

//	+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++	//

// Parsing value from
// string with regex pattern
func parseField(r string, s string) (string, error) {

	e, eErr := regexp.Compile(r)
	if eErr != nil {
		return "", eErr
	}

	if m := e.FindAllStringSubmatch(s, -1); m != nil {
		return m[0][0], nil
	} else {
		return "", errNoFss
	}

}

//	+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Pass struct with regex expressions
// in tags and fill struct with parsed values
func FillStruct(d interface{}, s string, t string) (bool, error) {

	dV := reflect.ValueOf(d).Elem()
	dK := dV.Type()

	for i := 0; i < dV.NumField(); i++ {

		if m, mErr := parseField(dK.Field(i).Tag.Get(t), s); mErr != nil {
			return false, mErr
		} else {
			dV.Field(i).SetString(m)
		}

	}

	return true, nil

}

//	+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
