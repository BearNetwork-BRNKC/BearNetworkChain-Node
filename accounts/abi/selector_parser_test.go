

package abi

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestParseSelector(t *testing.T) {
	t.Parallel()
	mkType := func(types ...interface{}) []ArgumentMarshaling {
		var result []ArgumentMarshaling
		for i, typeOrComponents := range types {
			name := fmt.Sprintf("name%d", i)
			if typeName, ok := typeOrComponents.(string); ok {
				result = append(result, ArgumentMarshaling{name, typeName, typeName, nil, false})
			} else if components, ok := typeOrComponents.([]ArgumentMarshaling); ok {
				result = append(result, ArgumentMarshaling{name, "tuple", "tuple", components, false})
			} else if components, ok := typeOrComponents.([][]ArgumentMarshaling); ok {
				result = append(result, ArgumentMarshaling{name, "tuple[]", "tuple[]", components[0], false})
			} else {
				log.Fatalf("unexpected type %T", typeOrComponents)
			}
		}
		return result
	}
	tests := []struct {
		input string
		name  string
		args  []ArgumentMarshaling
	}{
		{"noargs()", "noargs", []ArgumentMarshaling{}},
		{"simple(uint256,uint256,uint256)", "simple", mkType("uint256", "uint256", "uint256")},
		{"other(uint256,address)", "other", mkType("uint256", "address")},
		{"withArray(uint256[],address[2],uint8[4][][5])", "withArray", mkType("uint256[]", "address[2]", "uint8[4][][5]")},
		{"singleNest(bytes32,uint8,(uint256,uint256),address)", "singleNest", mkType("bytes32", "uint8", mkType("uint256", "uint256"), "address")},
		{"multiNest(address,(uint256[],uint256),((address,bytes32),uint256))", "multiNest",
			mkType("address", mkType("uint256[]", "uint256"), mkType(mkType("address", "bytes32"), "uint256"))},
		{"arrayNest((uint256,uint256)[],bytes32)", "arrayNest", mkType([][]ArgumentMarshaling{mkType("uint256", "uint256")}, "bytes32")},
		{"multiArrayNest((uint256,uint256)[],(uint256,uint256)[])", "multiArrayNest",
			mkType([][]ArgumentMarshaling{mkType("uint256", "uint256")}, [][]ArgumentMarshaling{mkType("uint256", "uint256")})},
		{"singleArrayNestAndArray((uint256,uint256)[],bytes32[])", "singleArrayNestAndArray",
			mkType([][]ArgumentMarshaling{mkType("uint256", "uint256")}, "bytes32[]")},
		{"singleArrayNestWithArrayAndArray((uint256[],address[2],uint8[4][][5])[],bytes32[])", "singleArrayNestWithArrayAndArray",
			mkType([][]ArgumentMarshaling{mkType("uint256[]", "address[2]", "uint8[4][][5]")}, "bytes32[]")},
	}
	for i, tt := range tests {
		selector, err := ParseSelector(tt.input)
		if err != nil {
			t.Errorf("test %d: failed to parse selector '%v': %v", i, tt.input, err)
		}
		if selector.Name != tt.name {
			t.Errorf("test %d: unexpected function name: '%s' != '%s'", i, selector.Name, tt.name)
		}

		if selector.Type != "function" {
			t.Errorf("test %d: unexpected type: '%s' != '%s'", i, selector.Type, "function")
		}
		if !reflect.DeepEqual(selector.Inputs, tt.args) {
			t.Errorf("test %d: unexpected args: '%v' != '%v'", i, selector.Inputs, tt.args)
		}
	}
}
