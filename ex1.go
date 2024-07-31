package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where my program begins")
}

func practical(P person) struct {
	sno   int
	first string
	last  string
} {
	return struct {
		sno   int
		first string
		last  string
	}{
		sno:   1,
		first: P.first,
		last:  P.last,
	}
}

func variadicFunction(i ...int) {
	for _, v := range i {
		fmt.Println(v)
	}
}

type person struct {
	first    string
	last     string
	iceCream string
}

type secretAgent struct {
	person
	sno int
}

func (s secretAgent) speak() {
	fmt.Println(s.person.first, s.person.last, s.person.iceCream, s.sno)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func arrayfunc(a [10]int) {
	a[0] = 11
	//fmt.Println(a)
}

func main() {

	fmt.Println("Hii")

	a := [...]int{42, 43, 44}
	fmt.Println(a[0])

	var b = [3]int{1, 2, 3}
	fmt.Println(b[0])

	for i, v := range b {
		fmt.Println(v)
		fmt.Println(b[i])
	}

	c := rand.Intn(251)
	fmt.Println("The value of c is", c)

	if c >= 0 && c <= 100 {
		fmt.Println("The value of c is between 0 to 100")
	} else if c > 100 && c <= 200 {
		fmt.Println("The value of c is between 100 and 200")
	} else {
		fmt.Println("The value of c is between 201 and 250")
	}

	switch {
	case c >= 0 && c <= 100:
		fmt.Println("The value of c is greater than 100 by switch statement")
	case c > 100 && c <= 200:
		fmt.Println("The value of c is between 100 and 200 by switch statement")
	default:
		fmt.Println("The value of c is between 201 and 250 by switch statement")
	}

	fmt.Println(true && true)

	var sli = [][]int{{42, 43, 44, 45, 46},
		{47, 48, 49, 50, 51},
		{44, 45, 46, 47, 48},
		{43, 44, 45, 46, 47},
	}
	fmt.Println(sli)
	for i, row := range sli {
		for j, value := range row {
			fmt.Printf("sli[%d][%d] = %d\n", i, j, value)
		}
		fmt.Println()
	}

	sl1 := []int{1, 2, 3}
	sl2 := []int{4, 5, 6}
	sl1 = append(sl1, sl2...)
	fmt.Println(sl1)

	sl3 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(sl3)
	sl3 = append(sl3[:3], sl3[6:]...)
	fmt.Println(sl3)

	sl4 := make([]string, 0, 10)
	fmt.Println(sl4)
	sl4 = append(sl4, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `
Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `
Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `
Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
		` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
		` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `
Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,
	)
	fmt.Println("Length - ", len(sl4), "Capacity - ", cap(sl4))
	for i := 0; i < len(sl4); i++ {
		fmt.Println(sl4[i])
	}

	//m := map[string][]string{
	//        "bond_james": {"shaken, not stirred", "martinis", "fast cars"},
	//		  "money_penny_jenny" : {"intelligence", "literature", "computer science"}
	//		  "no_dr" : {"cats", "ice cream", "sunsets"}
	//    }

	var m = make(map[string][]string, 9)
	m["bond_james"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	m["money_penny_jenny"] = []string{"intelligence", "literature", "computer science"}
	m["no_dr"] = []string{"cats", "ice cream", "sunsets"}
	fmt.Println(m)

	for k, ik := range m {
		fmt.Println(k)
		for _, i := range ik {
			fmt.Println(i)
		}
	}

	delete(m, "bond_james")
	fmt.Println(m)

	book := []string{"in", "my", "younger", "and", "more", "vulnerable", "years", "my", "father", "gave", "me", "some",
		"advice", "that", "i’ve", "been", "turning", "over", "in", "my", "mind", "ever", "since.", "whenever",
		"you", "feel", "like", "criticizing", "anyone,", "he", "told", "me,", "just", "remember", "that", "all",
		"the", "people", "in", "this", "world", "haven’t", "had", "the", "advantages", "that", "you’ve",
		"had.", "he", "didn’t", "say", "any", "more,", "but", "we’ve", "always", "been", "unusually",
		"communicative", "in", "a", "reserved", "way,", "and", "i", "understood", "that", "he", "meant",
		"a", "great", "deal", "more", "than", "that.", "in", "consequence,", "i’m", "inclined", "to",
		"reserve", "all", "judgements,", "a", "habit", "that", "has", "opened", "up", "many", "curious",
		"natures", "to", "me", "and", "also", "made", "me", "the", "victim", "of", "not", "a", "few",
		"veteran", "bores.", "the", "abnormal", "mind", "is", "quick", "to", "detect", "and", "attach",
		"itself", "to", "this", "quality", "when", "it", "appears", "in", "a", "normal", "person,", "and", "so",
		"it", "came", "about", "that", "in", "college", "i", "was", "unjustly", "accused", "of", "being", "a",
		"politician,", "because", "i", "was", "privy", "to", "the", "secret", "griefs", "of", "wild,", "unknown",
		"men.", "most", "of", "the", "confidences", "were", "unsought—frequently", "i", "have",
		"feigned", "sleep,", "preoccupation,", "or", "a", "hostile", "levity", "when", "i", "realized", "by",
		"some", "unmistakable", "sign", "that", "an", "intimate", "revelation", "was", "quivering", "on",
		"the", "horizon;", "for", "the", "intimate", "revelations", "of", "young", "men,", "or", "at", "least",
		"the", "terms", "in", "which", "they", "express", "them,", "are", "usually", "plagiaristic", "and",
		"marred", "by", "obvious", "suppressions.", "reserving", "judgements", "is", "a", "matter", "of",
		"infinite", "hope.", "i", "am", "still", "a", "little", "afraid", "of", "missing", "something", "if", "i",
		"forget", "that,", "as", "my", "father", "snobbishly", "suggested,", "and", "i", "snobbishly",
		"repeat,", "a", "sense", "of", "the", "fundamental", "decencies", "is", "parcelled", "out",
		"unequally", "at", "birth.", "and,", "after", "boasting", "this", "way", "of", "my", "tolerance,", "i",
		"come", "to", "the", "admission", "that", "it", "has", "a", "limit.", "conduct", "may", "be",
		"founded", "on", "the", "hard", "rock", "or", "the", "wet", "marshes,", "but", "after", "a", "certain",
		"point", "i", "don’t", "care", "what", "it’s", "founded", "on.", "when", "i", "came", "back", "from",
		"the", "east", "last", "autumn", "i", "felt", "that", "i", "wanted", "the", "world", "to", "be", "in",
		"uniform", "and", "at", "a", "sort", "of", "moral", "attention", "forever;", "i", "wanted", "no",
		"more", "riotous", "excursions", "with", "privileged", "glimpses", "into", "the", "human", "heart.",
		"only", "gatsby,", "the", "man", "who", "gives", "his", "name", "to", "this", "book,", "was",
		"exempt", "from", "my", "reaction—gatsby,", "who", "represented", "everything", "for", "which",
		"i", "have", "an", "unaffected", "scorn.", "if", "personality", "is", "an", "unbroken", "series", "of",
		"successful", "gestures,", "then", "there", "was", "something", "gorgeous", "about", "him,",
		"some", "heightened", "sensitivity", "to", "the", "promises", "of", "life,", "as", "if", "he", "were",
		"related", "to", "one", "of", "those", "intricate", "machines", "that", "register", "earthquakes",
		"ten", "thousand", "miles", "away.", "this", "responsiveness", "had", "nothing", "to", "do", "with",
		"that", "flabby", "impressionability", "which", "is", "dignified", "under", "the", "name", "of", "the",
		"creative", "temperament—it", "was", "an", "extraordinary", "gift", "for", "hope,", "a", "romantic",
		"readiness", "such", "as", "i", "have", "never", "found", "in", "any", "other", "person", "and",
		"which", "it", "is", "not", "likely", "i", "shall", "ever", "find", "again.", "no—gatsby", "turned", "out",
		"all", "right", "at", "the", "end;", "it", "is", "what", "preyed", "on", "gatsby,", "what", "foul", "dust",
		"floated", "in", "the", "wake", "of", "his", "dreams", "that", "temporarily", "closed", "out", "my",
		"interest", "in", "the", "abortive", "sorrows", "and", "short-winded", "elations", "of", "men.", "my",
		"family", "have", "been", "prominent,", "well-to-do", "people", "in", "this", "middle", "western",
		"city", "for", "three", "generations.", "the", "carraways", "are", "something", "of", "a", "clan,",
		"and", "we", "have", "a", "tradition", "that", "we’re", "descended", "from", "the", "dukes", "of",
		"buccleuch,", "but", "the", "actual", "founder", "of", "my", "line", "was", "my", "grandfather’s", "brother,", "who", "came", "here", "in", "fifty-one,", "sent", "a", "substitute", "to", "the", "civil",
		"war,", "and", "started", "the", "wholesale", "hardware", "business", "that", "my", "father",
		"carries", "on", "today.", "i", "never", "saw", "this", "great-uncle,", "but", "i’m", "supposed", "to",
		"look", "like", "him—with", "special", "reference", "to", "the", "rather", "hard-boiled", "painting",
		"that", "hangs", "in", "father’s", "office.", "i", "graduated", "from", "new", "haven", "in", "1915,",
		"just", "a", "quarter", "of", "a", "century", "after", "my", "father,", "and", "a", "little", "later", "i",
		"participated", "in", "that", "delayed", "teutonic", "migration", "known", "as", "the", "great",
		"war.", "i", "enjoyed", "the", "counter-raid", "so", "thoroughly", "that", "i", "came", "back",
		"restless.", "instead", "of", "being", "the", "warm", "centre", "of", "the", "world,", "the", "middle",
		"west", "now", "seemed", "like", "the", "ragged", "edge", "of", "the", "universe—so", "i",
		"decided", "to", "go", "east", "and", "learn", "the", "bond", "business.", "everybody", "i", "knew",
		"was", "in", "the", "bond", "business,", "so", "i", "supposed", "it", "could", "support", "one",
		"more", "single", "man.", "all", "my", "aunts", "and", "uncles", "talked", "it", "over", "as", "if",
		"they", "were", "choosing", "a", "prep", "school", "for", "me,", "and", "finally", "said,",
		"why—ye-es,", "with", "very", "grave,", "hesitant", "faces.", "father", "agreed", "to", "finance",
		"me", "for", "a", "year,", "and", "after", "various", "delays", "i", "came", "east,", "permanently,",
		"i", "thought,", "in", "the", "spring", "of", "twenty-two.", "the", "practical", "thing", "was", "to",
		"find", "rooms", "in", "the", "city,", "but", "it", "was", "a", "warm", "season,", "and", "i", "had",
		"just", "left", "a", "country", "of", "wide", "lawns", "and", "friendly", "trees,", "so", "when", "a",
		"young", "man", "at", "the", "office", "suggested", "that", "we", "take", "a", "house", "together",
		"in", "a", "commuting", "town,", "it", "sounded", "like", "a", "great", "idea.", "he", "found", "the",
		"house,", "a", "weather-beaten", "cardboard", "bungalow", "at", "eighty", "a", "month,", "but",
		"at", "the", "last", "minute", "the", "firm", "ordered", "him", "to", "washington,", "and", "i", "went",
		"out", "to", "the", "country", "alone.", "i", "had", "a", "dog—at", "least", "i", "had", "him", "for", "a",
		"few", "days", "until", "he", "ran", "away—and", "an", "old", "dodge", "and", "a", "finnish",
		"woman,", "who", "made", "my", "bed", "and", "cooked", "breakfast", "and", "muttered",
		"finnish", "wisdom", "to", "herself", "over", "the", "electric", "stove.", "it", "was", "lonely", "for",
		"a", "day", "or", "so", "until", "one", "morning", "some", "man,", "more", "recently", "arrived",
		"than", "i,", "stopped", "me", "on", "the", "road.", "how", "do", "you", "get", "to", "west", "egg",
		"village?", "he", "asked", "helplessly.", "i", "told", "him.", "and", "as", "i", "walked", "on", "i",
		"was", "lonely", "no", "longer.", "i", "was", "a", "guide,", "a", "pathfinder,", "an", "original",
		"settler.", "he", "had", "casually", "conferred", "on", "me", "the", "freedom", "of", "the",
		"neighbourhood.", "and", "so", "with", "the", "sunshine", "and", "the", "great", "bursts", "of",
		"leaves", "growing", "on", "the", "trees,", "just", "as", "things", "grow", "in", "fast", "movies,", "i",
		"had", "that", "familiar", "conviction", "that", "life", "was", "beginning", "over", "again", "with",
		"the", "summer.", "there", "was", "so", "much", "to", "read,", "for", "one", "thing,", "and", "so",
		"much", "fine", "health", "to", "be", "pulled", "down", "out", "of", "the", "young", "breath-giving",
		"air.", "i", "bought", "a", "dozen", "volumes", "on", "banking", "and", "credit", "and",
		"investment", "securities,", "and", "they", "stood", "on", "my", "shelf", "in", "red"}

	var bookMap = map[string]int{}
	for _, v := range book {
		bookMap[v]++
	}
	for k, v := range bookMap {
		fmt.Print("Key : ", k)
		fmt.Println("\tValue : ", v)
	}

	p := []person{
		{
			first:    "Aman",
			last:     "Kawadia",
			iceCream: "Chocolate",
		},
		{
			first:    "Pravin",
			last:     "Jalodiya",
			iceCream: "French Vanilla",
		},
	}
	fmt.Println(p)

	strMap := map[string]person{
		p[0].first: p[0],
		p[1].first: p[1],
	}
	for _, v := range strMap {
		fmt.Println(v)
	}

	fmt.Println(strMap["Aman"])

	fmt.Println(p)

	sl5 := [...][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
	}
	fmt.Println(sl5)

	sl6 := make([]int, 10, 10)
	fmt.Println(sl6)
	sl6 = append(sl6[:0], 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	fmt.Println(sl6)
	fmt.Println(len(sl6))
	fmt.Println(cap(sl6))
	sl6 = append(sl6, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	fmt.Println(sl6)
	fmt.Println(len(sl6))
	fmt.Println(cap(sl6))

	abc := func() {
		fmt.Println("abc")
	}
	abc()

	//var ptr *int  //nil pointer dereferencing
	//fmt.Println(*ptr)

	var practicalVar struct {
		sno   int
		first string
		last  string
	}
	practicalVar = practical(p[0])
	fmt.Println(practicalVar)

	variadicFunction(1, 2, 3, 4, 5, 6)
	variadicFunction()

	sec := secretAgent{
		person: p[0],
		sno:    1,
	}
	fmt.Println(sec)
	sec.speak()

	saySomething(sec)

	exampleVar := practicalVar
	fmt.Println(exampleVar)

	exampleVar.first = "Pravin"
	exampleVar.last = "Jalodiya"
	fmt.Println(exampleVar)
	fmt.Println(practicalVar)

	str1 := "My ⠧ name is the flash and i am the fastest man alive. When i was a child, my mother was killed by something impossible. My name is the flash and i am the fastest man alive. When i was a child, my mother was killed by something impossible. My name is the flash and i am the fastest man alive. When i was a child, my mother was killed by something impossible"
	fmt.Println(str1)
	fmt.Println(len(str1))

	fmt.Printf("%v\n", str1[3])

	var a12 byte = 191
	fmt.Printf("%c\n", a12)

	var ex rune = '¿'
	fmt.Printf("%c, %T, %#v\n", ex, ex, ex)

	var exa int32 = 2221
	fmt.Printf("%c, %T, %#v\n", exa, exa, exa)

	exa = ex
	fmt.Printf("%c, %T, %#v\n", exa, exa, exa)

	bb := []byte{'h', 'e', 'l', 'l', 'o', 121, 222}
	ss := string(bb)
	fmt.Println(ss)

	//bb1 := []int{121, 222, 2222}
	//ss1 := []bytes(bb1)
	//fmt.Println(ss)

	func() {
		fmt.Println("Ghost")
	}()

	var exArray [10]int
	exArray[0] = 1
	fmt.Println(exArray)

	arrayfunc(exArray)

	fmt.Println(exArray)

	per1 := person{
		"Aman",
		"Kawadia",
		"Chocolate",
	}
	per2 := per1
	per2.last = "Jain"
	fmt.Println(per1)
	fmt.Println(per2)
}
