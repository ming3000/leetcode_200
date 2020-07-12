package shopee

import "testing"

func TestGenerateKeyList(t *testing.T) {
	//m := map[int64]string{1: "a", 2: "b", 3: "c"}
	//GenerateKeyList(m)
}

func TestFixStr(t *testing.T) {
	t.Log(FixStr2(" 0 "))
	t.Log(FixStr2("     1 "))
	t.Log(FixStr2("      "))
	t.Log(FixStr2(""))
	t.Log(FixStr2("A"))
}

func TestFixStr2(t *testing.T) {

}
