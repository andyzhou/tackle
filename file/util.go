package file

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

/*
 * util face
 */

//face info
type Util struct {
}

//convert seconds to time string format
func (f *Util) Seconds2TimeStr(seconds int) string {
	var (
		hourStr, minuteStr, secondStr string
	)

	if seconds <= 0 {
		return ""
	}

	hourInt := seconds / 3600
	minuteInt := (seconds - hourInt * 3600) / 60
	secondInt := seconds - hourInt * 3600 - minuteInt * 60

	if hourInt > 0 {
		if hourInt > 9 {
			hourStr = fmt.Sprintf("%d:", hourInt)
		}else{
			hourStr = fmt.Sprintf("0%d:", hourInt)
		}
	}

	if minuteInt > 9 {
		minuteStr = fmt.Sprintf("%d", minuteInt)
	}else{
		minuteStr = fmt.Sprintf("0%d", minuteInt)
	}

	if secondInt > 9 {
		secondStr = fmt.Sprintf("%d", secondInt)
	}else{
		secondStr = fmt.Sprintf("0%d", secondInt)
	}

	//format time string
	timeStr := fmt.Sprintf("%s%s:%s", hourStr, minuteStr, secondStr)
	return timeStr
}

//execute 3rd command
func (f *Util) ExecCommand(
	command, args string) (string, error) {
	var (
		result string
	)
	//check
	if args == "" {
		return result, errors.New("invalid parameter")
	}

	//get third command path
	argSlice := strings.Split(args, " ")

	//exec command
	cmd := exec.Command(command, argSlice...)
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Printf("BaseFace::execCommand args:%v failed, err:%v\n", args, err.Error())
		return result, err
	}

	//init result
	result = string(out)
	return result, nil
}