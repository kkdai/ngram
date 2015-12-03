package ngram

import "sort"

type Ngram uint32

type docList []int

func (d docList) Len() int           { return len(d) }
func (d docList) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d docList) Less(i, j int) bool { return d[i] < d[j] }

type NgramType int

const ( // iota is reset to 0
	Twogram  NgramType = 2 // 0
	Trigram  NgramType = 3 // 1
	Fourgram NgramType = 4 // 2
)

//The trigram indexing result include all Document IDs and its Frequence in that document
type IndexResult struct {
	//Save all trigram mapping docID
	DocIDs map[int]bool

	//Save all trigram appear time for trigram deletion
	Freq map[int]int
}

// Extract one string to ngram list
// Note the Ngram is a uint32 for ascii code
func ExtractStringToNgram(str string, nType NgramType) []Ngram {
	if len(str) == 0 {
		return nil
	}

	nTypeInt := int(nType)
	var result []Ngram
	for i := 0; i < len(str)-(nTypeInt-1); i++ {
		var ngram Ngram
		for j := 0; j < nTypeInt; j++ {
			//ngram = Ngram(uint32(str[i])<<16 | uint32(str[i+1])<<8 | uint32(str[i+2]))
			shift := uint(j * 8)
			ngram = ngram + Ngram(uint32(str[i+(nTypeInt-1-j)])<<shift)
		}
		result = append(result, ngram)
	}

	return result
}

type NgramIndex struct {
	//To store all current trigram indexing result
	TrigramMap map[Ngram]IndexResult

	//it represent and document incremental index
	maxDocID int

	//it include currently all the doc list, it will be used when query string length less than 3
	docIDsMap map[int]bool

	//Ngram type
	ngramType NgramType
}

//Create a new trigram indexing
func NewNgramIndex(nType NgramType) *NgramIndex {
	t := new(NgramIndex)
	t.TrigramMap = make(map[Ngram]IndexResult)
	t.docIDsMap = make(map[int]bool)
	t.ngramType = nType
	return t
}

//Add new document into this trigram index
func (t *NgramIndex) Add(doc string) int {
	newDocID := t.maxDocID + 1
	trigrams := ExtractStringToNgram(doc, t.ngramType)
	for _, tg := range trigrams {
		var mapRet IndexResult
		var exist bool
		if mapRet, exist = t.TrigramMap[tg]; !exist {
			//New doc ID handle
			mapRet = IndexResult{}
			mapRet.DocIDs = make(map[int]bool)
			mapRet.Freq = make(map[int]int)
			mapRet.DocIDs[newDocID] = true
			mapRet.Freq[newDocID] = 1
		} else {
			//trigram already exist on this doc
			if _, docExist := mapRet.DocIDs[newDocID]; docExist {
				mapRet.Freq[newDocID] = mapRet.Freq[newDocID] + 1
			} else {
				//tg eixist but new doc id is not exist, add it
				mapRet.DocIDs[newDocID] = true
				mapRet.Freq[newDocID] = 1
			}
		}
		//Store or Add  result
		t.TrigramMap[tg] = mapRet
	}

	t.maxDocID = newDocID
	t.docIDsMap[newDocID] = true
	return newDocID
}

//Delete a doc from this trigram indexing
func (t *NgramIndex) Delete(doc string, docID int) {
	trigrams := ExtractStringToNgram(doc, t.ngramType)
	for _, tg := range trigrams {
		if obj, exist := t.TrigramMap[tg]; exist {
			if freq, docExist := obj.Freq[docID]; docExist && freq > 1 {
				obj.Freq[docID] = obj.Freq[docID] - 1
			} else {
				//need remove trigram from such docID
				delete(obj.Freq, docID)
				delete(obj.DocIDs, docID)
			}

			if len(obj.DocIDs) == 0 {
				//this object become empty remove this.
				delete(t.TrigramMap, tg)
				//TODO check if some doc id has no tg remove
			} else {
				//update back since there still other doc id exist
				t.TrigramMap[tg] = obj
			}
		} else {
			//trigram not exist in map, leave
			return
		}
	}
}

//This function help you to intersect two map
func IntersectTwoMap(IDsA, IDsB map[int]bool) map[int]bool {
	var retIDs map[int]bool   //for traversal it is smaller one
	var checkIDs map[int]bool //for checking it is bigger one
	if len(IDsA) >= len(IDsB) {
		retIDs = IDsB
		checkIDs = IDsA

	} else {
		retIDs = IDsA
		checkIDs = IDsB
	}

	for id, _ := range retIDs {
		if _, exist := checkIDs[id]; !exist {
			delete(retIDs, id)
		}
	}
	return retIDs
}

//Query a target string to return the doc ID
func (t *NgramIndex) Query(doc string) docList {
	trigrams := ExtractStringToNgram(doc, t.ngramType)
	if len(trigrams) == 0 {
		return t.getAllDocIDs()
	}

	//Find first trigram as base for intersect
	retObj, exist := t.TrigramMap[trigrams[0]]
	if !exist {
		return nil
	}
	retIDs := retObj.DocIDs

	//Remove first one and do intersect with other trigram
	trigrams = trigrams[1:]
	for _, tg := range trigrams {
		checkObj, exist := t.TrigramMap[tg]
		if !exist {
			return nil
		}
		checkIDs := checkObj.DocIDs
		retIDs = IntersectTwoMap(retIDs, checkIDs)
	}

	return getMapToSlice(retIDs)

}

//Transfer map to slice for return result
func getMapToSlice(inMap map[int]bool) docList {
	var retSlice docList
	for k, _ := range inMap {
		retSlice = append(retSlice, k)
	}
	sort.Sort(retSlice)
	return retSlice
}

func (t *NgramIndex) getAllDocIDs() docList {
	return getMapToSlice(t.docIDsMap)
}
