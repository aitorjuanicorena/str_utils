package str_utils

import (
	"errors"
)

func trimLastChars(str_in string, trim_pos int) (str_out string, err error) {

	err = errors.New("String size is lower than trim position")
	str_out = str_in
	size := len(str_in)

	if size > trim_pos {

		str_out = str_in[:trim_pos]

		err = nil

	}

	return

}
