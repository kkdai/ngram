package ngram_test

import (
	"testing"

	. "github.com/kkdai/ngram"
)

func TestTwogramlize(t *testing.T) {
	ret := ExtractStringToNgram("Cod", Twogram)
	if ret[0] != 17263 || ret[1] != 28516 {
		t.Errorf("Trigram failed, current:%v", ret)
	}
}

func TestTrigramlize(t *testing.T) {
	ret := ExtractStringToNgram("Cod", Trigram)
	if ret[0] != 4419428 {
		t.Errorf("Trigram failed, expect 4419428\n")
	}

	//string length longer than 3
	ret = ExtractStringToNgram("Code", Trigram)
	if ret[0] != 4419428 && ret[1] != 7300197 {
		t.Errorf("Trigram failed on longer string")
	}
}

func TestMapIntersect(t *testing.T) {
	mapA := make(map[int]bool)
	mapB := make(map[int]bool)

	mapA[1] = true
	mapA[2] = true
	mapB[1] = true

	ret := IntersectTwoMap(mapA, mapB)
	if len(ret) != 1 || ret[1] == false {
		t.Errorf("Map intersect error")
	}

	ret = IntersectTwoMap(mapB, mapA)
	if len(ret) != 1 || ret[1] == false {
		t.Errorf("Map intersect error")
	}

	mapA[3] = true
	mapB[3] = true
	mapA[4] = true

	ret = IntersectTwoMap(mapB, mapA)
	if len(ret) != 2 || ret[1] == false {
		t.Errorf("Map intersect error")
	}
}

func TestTrigramIndexBasicQuery(t *testing.T) {
	ti := NewNgramIndex(Trigram)
	ti.Add("Code is my life")
	ti.Add("Search")
	ti.Add("I write a lot of Codes")

	ret := ti.Query("Code")
	if ret[0] != 1 || ret[1] != 3 {
		t.Errorf("Basic query is failed.")
	}
}

func TestEmptyLessQuery(t *testing.T) {
	ti := NewNgramIndex(Trigram)
	ti.Add("Code is my life")
	ti.Add("Search")
	ti.Add("I write a lot of Codes")

	ret := ti.Query("te") //less than 3, should get all doc ID
	if len(ret) != 3 || ret[0] != 1 || ret[2] != 3 {
		t.Errorf("Error on less than 3 character query")
	}

	ret = ti.Query("")
	if len(ret) != 3 || ret[0] != 1 || ret[2] != 3 {
		t.Errorf("Error on empty character query")
	}
}

func TestEmptyLessQuery2(t *testing.T) {
	ti := NewNgramIndex(Fourgram)

	docs := []string{"Don't communicate by sharing memory, share memory by communicating.", "Channels orchestrate; mutexes serialize.", "The bigger the interface, the weaker the abstraction.", "Make the zero value useful", "interface{} says nothing"}

	for _, v := range docs {
		ti.Add(v)
	}

	ret := ti.Query("the") //less than 4, should get all doc ID
	if len(ret) != 5 || ret[0] != 1 || ret[2] != 3 {
		t.Errorf("Error on less than 4 character query")
	}

	ret = ti.Query("interface")
	if len(ret) != 2 || ret[0] != 3 || ret[1] != 5 {
		t.Errorf("Error on empty character query")
	}
}

func TestDelete(t *testing.T) {
	ti := NewNgramIndex(Trigram)
	ti.Add("Code is my life")

	ti.Delete("Code", 1)
	ret := ti.Query("Code")
	if len(ret) != 0 {
		t.Error("Basic delete failed", ret)
	}

	ret = ti.Query("life")
	if len(ret) != 1 || ret[0] != 1 {
		t.Error("Basic delete failed", ret)
	}
}

func BenchmarkAddTwogram(b *testing.B) {
	big := NewNgramIndex(Twogram)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.Add("1234567890")
	}
}
func BenchmarkAddTrigram(b *testing.B) {
	big := NewNgramIndex(Trigram)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.Add("1234567890")
	}
}

func BenchmarkAddFourgran(b *testing.B) {
	big := NewNgramIndex(Fourgram)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.Add("1234567890")
	}
}
func BenchmarkDeleteTwogram(b *testing.B) {

	big := NewNgramIndex(Twogram)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.Add("1234567890")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.Delete("1234567890", i)
	}
}

func BenchmarkDeleteTrigram(b *testing.B) {

	big := NewNgramIndex(Trigram)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.Add("1234567890")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.Delete("1234567890", i)
	}
}

func BenchmarkDeleteFourgram(b *testing.B) {

	big := NewNgramIndex(Fourgram)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.Add("1234567890")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.Delete("1234567890", i)
	}
}

func BenchmarkQueryTwogran(b *testing.B) {

	big := NewNgramIndex(Twogram)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.Add("1234567890")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.Query("1234567890")
	}
}

func BenchmarkQueryTrigran(b *testing.B) {

	big := NewNgramIndex(Trigram)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.Add("1234567890")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.Query("1234567890")
	}
}

func BenchmarkQueryFourgram(b *testing.B) {

	big := NewNgramIndex(Fourgram)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.Add("1234567890")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.Query("1234567890")
	}
}

func BenchmarkIntersectMap(b *testing.B) {

	DocA := make(map[int]bool)
	DocB := make(map[int]bool)
	for i := 0; i < 100000; i++ {
		DocA[i] = true
		DocB[i+1] = true
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IntersectTwoMap(DocA, DocB)
	}
}

func BenchmarkIntersectSlice(b *testing.B) {

	var DocA DocList
	var DocB DocList
	for i := 0; i < 100000; i++ {
		DocA = append(DocA, i)
		DocB = append(DocB, i+1)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IntersectTwoSlice(DocA, DocB)
	}
}
