// generated by stringer -type=displayModeEnum; DO NOT EDIT

package kml

import "fmt"

const _displayModeEnum_name = "Defaulthide"

var _displayModeEnum_index = [...]uint8{0, 7, 11}

func (i displayModeEnum) String() string {
	if i < 0 || i+1 >= displayModeEnum(len(_displayModeEnum_index)) {
		return fmt.Sprintf("displayModeEnum(%d)", i)
	}
	return _displayModeEnum_name[_displayModeEnum_index[i]:_displayModeEnum_index[i+1]]
}