package intromain

import ( "fmt" ; "os" ; "time" ; "strings" //; "bufio"  
	log "github.com/alecthomas/log4go" )

// Greeting ...
func Greeting(greetPath string, name string) (string, error) {
	var ( res string = "" ; err error = nil )
	log.Info("Greeting()")
	//fIn, err := os.OpenFile(greet_path, os.O_RDONLY, 0666)
	fIn, err := os.Open(greetPath)
	if nil != err { return res, err }
	
	defer fIn.Close()
	
	bArr := make([]byte, 80)
	_, err = fIn.Read(bArr)
	if nil != err { return res, err }
	
	buf := string(bArr)
	res = strings.Replace(buf, "\n", "", -1) + name
	
	return res, err
}

// DelayChar ...
func DelayChar(millisecs float64) byte {
	var ( ch byte = 0 ; discard string )
	
	for {
		time.Sleep(time.Duration(millisecs) * time.Millisecond)
		fmt.Println("Type any character when ready.")
		//rdr := bufio.NewReader(os.Stdin)
		//ch, _ = rdr.ReadByte()
		fmt.Scanf("%c", &ch)
		fmt.Scanln(&discard)
		
		if 0 == ch || '\n' == ch { continue
		} else { break }
	}
	return ch
}

