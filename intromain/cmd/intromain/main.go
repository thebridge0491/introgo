package main

// single-line comment
/* multi-line comment: 
 * ... */

import ( "flag" ; "fmt" ; "os" ; "regexp" ; "time" ; "math/rand" ; "reflect"
	"strings" ; "sort" ; log "github.com/alecthomas/log4go"
	"github.com/unknwon/goconfig" //; "io/ioutil" //; "encoding/json"
	//"github.com/BurntSushi/toml" ; "github.com/go-yaml/yaml"
	util "bitbucket.org/thebridge0491/introgo/introutil"
	practice "bitbucket.org/thebridge0491/introgo/intropractice"
	lib "bitbucket.org/thebridge0491/introgo/intromain"
)

type OptsRecord struct {
	name string
	num    int
	isExpt2 bool
}

// User type ...
type User struct {
	name   string
	num    int
	timeIn int64
}

type char byte

const PI float64 = 3.14

const (
	ZERO = iota
	//ONE
	NUMZ = 26
)

/*func deserializeStr(dataBytes []byte, dataFmt string, err error, errPfx string) map[string]interface{} {
	blankMap := make(map[string]interface{})
	blankMap["fmt"] = dataFmt
	if nil != err {
		fmt.Fprintf(os.Stderr, "%s: %s\n", errPfx, err)
    	//os.Exit(1)
		return blankMap
	}
	if "yaml" == dataFmt || "json" == dataFmt {
		yaml.Unmarshal([]byte(dataBytes), &blankMap)
	} else if "toml" == dataFmt {
		toml.Unmarshal([]byte(dataBytes), &blankMap)
	} //else if "json" == dataFmt {
	//	json.Unmarshal([]byte(dataBytes), &blankMap)
	//}
	return blankMap
}*/

func runIntromain(rsrcPath string, opts *OptsRecord) () {
	t1 := time.Now()
	timeIn := t1.Unix()
	
	// basic datatypes
	var isDone bool = false
	var ( numI int = 0 ; arrLen int = ZERO )
	var numIU uint = 100
	var seedp int64 = timeIn
	var ( delayMsecs float64 = 2.5e3 ; timeDiff float64 = 0.0 )
	var ( numF1 float64 = 100.0 ; numF2 float64 = 1.0e9 )
	var ch byte = 0
	
	rand.Seed(seedp)
	
	// pointers
	var numPtr *int = &numI
	//var int64Ptr *int64 = new(int64)
	
	// strings & arrays/slices
	var ( greetPath string = rsrcPath + "/greet.txt" ; dateBuf string)
	//var str1 [64]char
	//var numArr [4]int = [4]int{9, 011, 0x9, 9} //{bin, oct, hex, dec}
	numArr := [4]int{9, 011, 0x9, 9} //{bin, oct, hex, dec}
	var numSlice []int //= numArr[:]
	
	// composites
	user1 := User{name: opts.name}
	
	var userPtr *User = &user1
	var pers *lib.Person
    
    (*userPtr).num = rand.Intn(17) + 2
    userPtr.timeIn = timeIn

	arrLen = len(numArr)
	
	for i := 0; arrLen > i; i++ {
	    numI += numArr[i]
	}
	for _, val := range numArr { // idx, val := range numArr
	    numSlice = append(numSlice, val)
	}
	
	if *numPtr != (len(numSlice) * numSlice[0]) {
	    panic("not equal: *numPtr == (len(numSlice) * numSlice[0])")
	}
	
	ch = lib.DelayChar(delayMsecs)
	if 0 == ch { print("DelayChar error\n") }
	
	for !isDone {
	    numF1 += numF2
	    numIU += uint(numI)
	    isDone = true
	}
	
	
	matchQuit, err := regexp.MatchString("(?i)^quit$", opts.name)
	
	if nil != err {
        fmt.Fprintf(os.Stderr, "Regexp MatchString error\n")
        os.Exit(1)
    }
	if matchQuit { fmt.Printf("Good match: %s to %s\n", opts.name, "\"^quit$\"")
	} else { fmt.Printf("Does not match: %s to %s\n", opts.name, "\"^quit$\"") }
	
    
	greetBuf, err := lib.Greeting(greetPath, user1.name)
	if nil != err {
		fmt.Fprintf(os.Stderr, "%s.\n", err)
		os.Exit(1)
	}
    dateBuf = t1.Format(time.UnixDate)
	fmt.Printf("%s\n%s!\n", dateBuf, greetBuf)
    
    timeDiff = float64(time.Since(t1).Seconds())
	fmt.Printf("(program %s) Took %.1f seconds.\n", os.Args[0], timeDiff)
	fmt.Println(strings.Repeat("-", 40))
	
	var arrInts = []int{2, 0, 1, 4, 3}
	var iarrInts = util.IfcArrFromInts(arrInts)
	
	if opts.isExpt2 {
		fmt.Printf("Expt(%.1f, %.1f) = %.1f\n", 2.0, float64(user1.num),
			practice.ExptLp(2.0, float64(user1.num)))
		
		var iarrTmp = practice.CopyOf(iarrInts)
		practice.ReverseLp(iarrTmp)
		fmt.Printf("Reverse(%v): %v\n", iarrInts, iarrTmp)
		
		fmt.Printf("sort.Sort(sort.IntSlice(%v)): ", arrInts)
		sort.Sort(sort.IntSlice(arrInts))
		fmt.Println(arrInts)
	} else {
		fmt.Printf("Fact(%d) = %d\n", user1.num, practice.FactLp(user1.num))
		
		fmt.Printf("FindIndex(3, %v): %d\n", iarrInts,
			practice.FindIndexLp(3, iarrInts))
		
		fmt.Printf("append(%v, %d): %v\n", arrInts, 50, append(arrInts, 50))
	}
	fmt.Println(strings.Repeat("-", 40))
	
	pers = lib.NewPerson("I.M. Computer", 32)
	
	if reflect.TypeOf(pers) != reflect.TypeOf(&lib.Person{}) {
		panic("error: Type mismatch")
	}
	fmt.Println(pers.ToString())
	pers.SetAge(33)
	fmt.Printf("person.SetAge(%d): \n", 33)
	fmt.Println(pers.ToString())
	
	fmt.Println(strings.Repeat("-", 40))
	
	
	time.Sleep(100 * time.Millisecond)
}

func parseCmdopts(opts *OptsRecord) {
	flag.StringVar(&opts.name, "u", opts.name, "user name")
	flag.IntVar(&opts.num, "n", opts.num, "number n")
	flag.BoolVar(&opts.isExpt2, "2", opts.isExpt2, "expt 2 n")
	log.Info("parseCmdopts")
    flag.Parse()
	//fmt.Fprintln(os.Stderr, flag.Args())
}

func recoverMain() {
	if r := recover(); nil != r {
		fmt.Println("Recovered in main ---", r)
	}
}

// main - entry point (DocComment)
func main() {
	/*defer func() {
		if r := recover(); nil != r {
			fmt.Println("Recovered in main ---", r)
		}
	}()*/
	defer recoverMain()
	var rsrcPath string
	envRsrcPath, issetEnvRsrcPath := os.LookupEnv("RSRC_PATH")
	if issetEnvRsrcPath { rsrcPath = envRsrcPath
	} else { rsrcPath = "resources" }
	
	_, err := os.Stat(rsrcPath + "/log4go.xml")
	if os.IsNotExist(err) {
    	//panic(fmt.Sprintf("os.Stat file error: %s\n", err))
		fmt.Fprintf(os.Stderr, "os.Stat file error: %s\n", err)
    	//os.Exit(1)
    } else {
		log.LoadConfiguration(rsrcPath + "/log4go.xml")
	}
	
	opts := OptsRecord{name: "World", num: 0, isExpt2: false}
	
    parseCmdopts(&opts)
    
    var rowsArr = make([][]string, 0)
    iniCfg, err := goconfig.LoadConfigFile(rsrcPath + "/prac.conf")
    if nil != err {
    	fmt.Fprintf(os.Stderr, "goconfig.LoadConfigFile data error: %s\n", err)
    	//os.Exit(1)
    	rowsArr = append(rowsArr, []string{"????\n", "???", "???"})
    } else {
		rowsArr = append(rowsArr, []string{util.IniCfgToStr(iniCfg), 
			iniCfg.MustValue("default", "domain"),
			iniCfg.MustValue("user1", "name")})
	}
    
    
    /*jsonStr, jsonErr := ioutil.ReadFile(rsrcPath + "/prac.json")
	tomlStr, tomlErr := ioutil.ReadFile(rsrcPath + "/prac.toml")
	yamlStr, yamlErr := ioutil.ReadFile(rsrcPath + "/prac.yaml")
	
	var jsonMap = deserializeStr(jsonStr, "json", jsonErr,
		"ioutil.ReadFile data error")
	var tomlMap = deserializeStr(tomlStr, "toml", tomlErr,
		"ioutil.ReadFile data error")
	var yamlMap = deserializeStr(yamlStr, "yaml", yamlErr,
		"ioutil.ReadFile data error")
    
    if nil != jsonErr {
    	rowsArr = append(rowsArr, []string{"????\n", "???", "???"})
	} else if nil != jsonMap["user1"] {
		rowsArr = append(rowsArr, []string{fmt.Sprintf("%s", jsonMap),
			fmt.Sprintf("%s", jsonMap["domain"]),
			//fmt.Sprintf("%s", jsonMap["user1"].(map[string]interface{})["name"]),
			fmt.Sprintf("%s", jsonMap["user1"].(map[interface{}]interface{})["name"]),
			})
	}
	if nil != tomlErr {
    	rowsArr = append(rowsArr, []string{"????\n", "???", "???"})
	} else if nil != tomlMap["user1"] {
		rowsArr = append(rowsArr, []string{fmt.Sprintf("%s", tomlMap),
			fmt.Sprintf("%s", tomlMap["domain"]),
			fmt.Sprintf("%s", tomlMap["user1"].(map[string]interface{})["name"])})
	}
	if nil != yamlErr {
    	rowsArr = append(rowsArr, []string{"????\n", "???", "???"})
	} else if nil != yamlMap["user1"] {
		rowsArr = append(rowsArr, []string{fmt.Sprintf("%s", yamlMap),
			fmt.Sprintf("%s", yamlMap["domain"]),
			fmt.Sprintf("%s", yamlMap["user1"].(map[interface{}]interface{})["name"])})
	}*/
	
    //sectList := iniCfg.GetSectionList()
    //fmt.Printf("%s\n", sectList)
    for _, row := range rowsArr {
		fmt.Printf("\nconfig: %s", row[0])
		fmt.Printf("\ndomain: %s", row[1])
		fmt.Printf("\nuser1Name: %s\n", row[2])
	}
    
    runIntromain(rsrcPath, &opts)
    
    log.Debug("exiting main()")
}
