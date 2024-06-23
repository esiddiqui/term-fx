package format

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/esiddiqui/tfx/text"
)

// pad or trim the supplied to be max_len,
// if the string is longer than max_len, it
// will be updated with ellipses (...)
func pad(sb *strings.Builder, str string, max_len int) {
	if len(str) >= max_len {
		sb.WriteString(str[0 : max_len-4])
		sb.WriteString("... ")
		//return str[0:max_len-4] + "..." + " "
	} else {
		sb.WriteString(str)
		sb.WriteString(strings.Repeat(" ", max_len-len(str)))
		// return str + strings.Repeat(" ", max_len-len(str))
	}
}

const (
	tagKey1 string = "tfxcli"
)

// CliOut formats & displays the results array as a simple CLI
// output table in text format.
//
// T must be a struct & each field tagged with the key `tfxcli` will
// be displayed in a column.
//
// The ordering of the columns in the output is controlled by
// the
// The title & width of each column is controlled by the tag
// value in the format
//
//	field1 `tfxcli:"TITLE,50"`
func CliOutputTable[T any](results []T) error {

	if results == nil {
		fmt.Printf("0 items\n")
	}

	typeOfArr := reflect.TypeOf(results)
	kindOfArr := typeOfArr.Kind()
	if kindOfArr != reflect.Slice && kindOfArr != reflect.String {
		return fmt.Errorf("invalid type supplied as parameter, only array or slice of structs allowed")
	}

	typeOfElem := typeOfArr.Elem() // element type of the array
	kindOFElem := typeOfElem.Kind()
	if kindOFElem != reflect.Struct {
		return fmt.Errorf("invalid type supplied as parameter, only array or slice of structs allowed")
	}
	fields := make([]string, 0)
	titles := make([]string, 0)
	widths := make([]int64, 0)

	// inspect the struct
	for i := 0; i < typeOfElem.NumField(); i++ {
		f := typeOfElem.Field(i)
		val := f.Tag.Get(tagKey1)
		if len(val) != 0 {
			lstrFields := strings.Split(val, ",")
			titles = append(titles, lstrFields[0])
			if len(lstrFields) > 1 {
				if w, err := strconv.ParseInt(lstrFields[1], 10, 64); err != nil {
					widths = append(widths, 25)
				} else {
					widths = append(widths, w)
				}
			}
			fields = append(fields, f.Name)
		}
	}

	// top margin
	fmt.Println()

	// print title
	var sb strings.Builder
	for n, title := range titles {
		pad(&sb, title, int(widths[n]))
	}
	fmt.Println(text.Ul(sb.String()))

	// print values
	for _, res := range results {
		var sb strings.Builder
		rVal := reflect.ValueOf(&res)
		for n, field := range fields {
			rrval := rVal.Elem().FieldByName(field)
			switch rrval.Kind() {
			case reflect.Bool:
				v := fmt.Sprintf("%v", rrval.Bool())
				// fmt.Print()
				pad(&sb, v, int(widths[n]))
			case reflect.String:
				// fmt.Print()
				pad(&sb, rrval.String(), int(widths[n]))
			default:
				// fmt.Print()
				pad(&sb, "Not supported", int(widths[n]))
			}
		}
		fmt.Println(sb.String())
	}
	fmt.Printf("%v items\n", len(results))

	return nil
}
