package magro

import (
	"testing"
)

func TestLine(t *testing.T) {
	v := Line()
	if v != 8 {
		t.Errorf("Line in test should return 8 but have %d\n", v)
	}
}

func TestFile(t *testing.T) {
	v := File()
	if "magro_test.go" != v {
		t.Errorf("File should return \"magro_test.go\" but have %s\n", v)
	}
}

func TestFunction(t *testing.T) {
	v := Function()
	if "github.com/Napat/magro.TestFunction" != v {
		t.Errorf("Function should return \"github.com/Napat/magro.TestFunction\" but have %s\n", v)
	}
}

func TestWhere(t *testing.T) {
	v := Where()
	if "File: magro_test.go Function: github.com/Napat/magro.TestWhere Line: 29" != v {
		t.Errorf("Where should return \"File: magro_test.go Function: github.com/Napat/magro.TestWhere Line: 29\" but have %s\n", v)
	}
}

func TestInfo(t *testing.T) {
	filename, funcName, line := Info()

	if line != 36 {
		t.Errorf("Line in test should return 36 but have %d\n", line)
	}
	if "magro_test.go" != filename {
		t.Errorf("File should return \"magro_test.go\" but have %s\n", filename)
	}
	if "github.com/Napat/magro.TestInfo" != funcName {
		t.Errorf("Function should return \"github.com/Napat/magro.TestInfo\" but have %s\n", funcName)
	}
}
