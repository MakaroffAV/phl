package phl

import (
	"errors"
	"reflect"
	"regexp"
)

//	####################################################################### //

//	Package doc:
//

//	####################################################################### //

var (
	mErr error = errors.New("no matches")
)

//	+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++	//

// Desc:	parse value with regex
// Pass:	fdsjhfklsj
func pvwr(r string, s string) (*string, error) {

	p, pErr := regexp.Compile(r)
	if pErr != nil {
		return nil, pErr
	}

	if m := p.FindAllStringSubmatch(s, -1); m != nil {
		return &m, nil
	} else {
		return nil, mErr
	}

}

//	+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Desc:	Parsing values
// Pass:	Fill struct by regex
func Fsbr(d interface{}, s string, t string) (bool, error) {

	v := reflect.ValueOf(d).Elem()
	k := v.Type()

	for i := 0; i < v.NumField(); i++ {

		if p, err := pvwr(k.Field(i).Tag.Get(t), s); err == nil {
			v.Field(i).SetString(*p)
		} else {
			return false, err
		}

	}

	return true, nil

}

//	+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
