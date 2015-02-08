package main

// single-line comment
/* multi-line comment: 
 * ... */

import ( "flag" ; "fmt" ; "os" ; "io/ioutil" ; "regexp" ; "encoding/json"
	log "github.com/alecthomas/log4go" ; "github.com/unknwon/goconfig"
	"github.com/go-yaml/yaml"
	lib "bitbucket.org/thebridge0491/introgo/intromain"
)

type OptsRecord struct {
	name string
}

func runIntromain(name string) () {
	matchQuit, err := regexp.MatchString("(?i)^quit$", name)
	
	if nil != err {
        fmt.Fprintf(os.Stderr, "Regexp MatchString error\n")
        os.Exit(1)
    }
	if matchQuit { fmt.Printf("Good match: %s to %s\n", name, "\"^quit$\"")
	} else { fmt.Printf("Does not match: %s to %s\n", name, "\"^quit$\"") }
}

func parseCmdopts(opts *OptsRecord) {
	flag.StringVar(&opts.name, "u", opts.name, "user name")
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
	
	opts := OptsRecord{name: "World"}
	
    parseCmdopts(&opts)
    
    var rowsArr = make([][]string, 0)
    iniCfg, err := goconfig.LoadConfigFile(rsrcPath + "/prac.conf")
    if nil != err {
    	fmt.Fprintf(os.Stderr, "goconfig.LoadConfigFile data error: %s\n", err)
    	//os.Exit(1)
    	rowsArr = append(rowsArr, []string{"????\n", "???", "???"})
    } else {
		rowsArr = append(rowsArr, []string{lib.IniCfgToStr(iniCfg), 
			iniCfg.MustValue("default", "domain"),
			iniCfg.MustValue("user1", "name")})
	}
    
    jsonStr, err := ioutil.ReadFile(rsrcPath + "/prac.json")
    if nil != err {
		fmt.Fprintf(os.Stderr, "ioutil.ReadFile data error: %s\n", err)
    	//os.Exit(1)
    	rowsArr = append(rowsArr, []string{"????\n", "???", "???"})
	}
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &jsonMap)
	if nil != jsonMap["user1"] {
		rowsArr = append(rowsArr, []string{fmt.Sprintf("%s", jsonMap),
			fmt.Sprintf("%s", jsonMap["domain"]),
			fmt.Sprintf("%s", jsonMap["user1"].(map[string]interface{})["name"])})
	}
	
	yamlStr, err := ioutil.ReadFile(rsrcPath + "/prac.yaml")
	if nil != err {
		fmt.Fprintf(os.Stderr, "ioutil.ReadFile data error: %s\n", err)
    	//os.Exit(1)
    	rowsArr = append(rowsArr, []string{"????\n", "???", "???"})
	}
	var yamlMap map[string]interface{}
	yaml.Unmarshal([]byte(yamlStr), &yamlMap)
	if nil != yamlMap["user1"] {
		rowsArr = append(rowsArr, []string{fmt.Sprintf("%s", yamlMap),
			fmt.Sprintf("%s", yamlMap["domain"]),
			fmt.Sprintf("%s", yamlMap["user1"].(map[interface{}]interface{})["name"])})
	}
	
    //sectList := iniCfg.GetSectionList()
    //fmt.Printf("%s\n", sectList)
    for _, row := range rowsArr {
		fmt.Printf("\nconfig: %s", row[0])
		fmt.Printf("\ndomain: %s", row[1])
		fmt.Printf("\nuser1Name: %s\n", row[2])
	}
    
    runIntromain(opts.name)
    
    log.Debug("exiting main()")
}
