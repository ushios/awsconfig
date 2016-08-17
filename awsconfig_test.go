package awsconfig

import "testing"

func TestFileExists(t *testing.T) {
	test := func(filename string, e bool) {
		res := fileExists(filename)

		if res != e {
			t.Errorf("%s expected(%t) but (%t)", filename, e, res)
		}
	}

	test("./LICENSE", true)
	test("./nothing", false)
}
