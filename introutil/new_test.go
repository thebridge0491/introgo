package introutil_test

import ( "testing"
	lib "bitbucket.org/thebridge0491/introgo/introutil"
)

func TestMethod(t *testing.T) {
    if 4 != 2 * 2 { t.FailNow() }
}

func TestFloatMethod(t *testing.T) {
    //if 4.0 != 2.0 * 2.0 { t.FailNow() }
    if !lib.InEpsilon(epsilon * 4.0, 4.0, 2.0 * 2.0) { t.FailNow() }
}

func TestStrMethod(t *testing.T) {
    if "Hello" != "Hello" { t.FailNow() }
}

func TestFailedMethod(t *testing.T) {
    if 5 != 2 * 2 { t.Error(nil) }
}

func TestLogMethod(t *testing.T) {
    if true { t.Log(nil) }
}

func TestSkipMethod(t *testing.T) {
    t.SkipNow() // t.Skip(nil)
    if true { t.Error(nil) }
}
