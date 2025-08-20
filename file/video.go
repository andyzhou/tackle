package file

import (
	"errors"
	"fmt"
	"github.com/andyzhou/tackle/conf"
	"github.com/andyzhou/tackle/define"
	"github.com/andyzhou/tackle/json"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
 * video process face
 */

//face info
type Video struct {
	Util
}

//construct
func NewVideo() *Video {
	this := &Video{}
	return this
}

//take a snap image
//return snap byte data, error
func (f *Video) TakeSnapImage(
	videoFilePath string,
	metaJson *json.VideoMetaJson,
	startSeconds ...int) ([]byte, error) {
	var (
		scalePara string
		startSecond int
	)

	//basic check
	if videoFilePath == "" || metaJson == nil {
		return nil, errors.New("invalid parameter")
	}
	if len(startSeconds) > 0 {
		startSecond = startSeconds[0]
	}

	//get main conf
	mainConf := conf.RunAppConfig.GetMainConf().GetConfInfo()
	privatePath := mainConf.PrivatePath
	tempPath := mainConf.TempPath
	cmdPath := mainConf.CommandPath
	snapScaleWidth := mainConf.AnimateScale

	//process duration info
	if startSecond <= 0 {
		startSecond = metaJson.Duration/2
	}
	durationStart := f.Seconds2TimeStr(startSecond)

	//init snap file path
	snapFile := fmt.Sprintf("%d_snap.jpg", time.Now().UnixNano())
	snapFilePath := fmt.Sprintf("%v/%v/%v", privatePath, tempPath, snapFile)

	//rotate check
	if metaJson.Rotate == 90 {
		scalePara = fmt.Sprintf("scale='-1:%d'", snapScaleWidth)
	}else{
		diff := float64(metaJson.Height - metaJson.Width)/float64(metaJson.Width)
		if diff >= define.ScaleMorePercent {
			snapScaleWidth  = int(float64(snapScaleWidth) * (1 - define.ScaleMorePercent))
		}
		scalePara = fmt.Sprintf("scale='%d:-1'", snapScaleWidth)
	}

	//format mpeg command
	cmdOfMpeg := fmt.Sprintf("%v/%v", cmdPath, define.CmdOfMpeg)

	//set args
	args := fmt.Sprintf("-ss %s -i %s -y -vframes 1 -filter:v %s -q:v 5 %s",
		durationStart, videoFilePath, scalePara, snapFilePath)

	//exec command
	_, err := f.ExecCommand(cmdOfMpeg, args)
	if err != nil {
		return nil, err
	}

	//defer opt
	defer func() {
		os.RemoveAll(snapFilePath)
	}()

	//read snap file
	snapByteData, subErr := ioutil.ReadFile(snapFilePath)
	return snapByteData, subErr
}

//generate animate gif
//return animate byte data, error
func (f *Video) GenAnimateGif(
	videoFilePath string,
	metaJson *json.VideoMetaJson,
	startSecond, endSecond int) ([]byte, error) {
	var (
		scale string
	)

	//basic check
	if videoFilePath == "" || metaJson == nil ||
		startSecond < 0 || endSecond > metaJson.Duration {
		return nil, errors.New("invalid parameter")
	}

	//get current timestamp
	now := time.Now().Unix()

	//get relate conf
	mainConf := conf.RunAppConfig.GetMainConf().GetConfInfo()
	cmdPath := mainConf.CommandPath
	privatePath := mainConf.PrivatePath
	tempPath := mainConf.TempPath
	snapFps := mainConf.SnapFps
	animateScale := mainConf.AnimateScale

	//init palette file
	paletteFile := fmt.Sprintf("%d_palette.png", now)
	palettePath := fmt.Sprintf("%v/%v/%s", privatePath, tempPath, paletteFile)

	//remove palette file
	defer func() {
		os.RemoveAll(palettePath)
	}()

	//init gif file
	gifName := fmt.Sprintf("%d.gif", now)
	gifPath := fmt.Sprintf("%v/%s/%s", privatePath, tempPath, gifName)

	//process duration info
	durationStart := f.Seconds2TimeStr(startSecond)

	//get animate scale and max seconds
	animateSeconds := endSecond - startSecond

	//mpeg command
	cmdOfMpeg := fmt.Sprintf("%v/%v", cmdPath, define.CmdOfMpeg)

	//create palette
	args := fmt.Sprintf("-ss %s -t %d -i %s -vf fps=%d,scale=%d:-1:flags=lanczos,palettegen -y %s",
		durationStart, animateSeconds, videoFilePath, snapFps, animateScale, palettePath)
	_, err := f.ExecCommand(cmdOfMpeg, args)
	if err != nil {
		return nil, err
	}

	//check rotate
	if metaJson.Rotate == 90 {
		scale = fmt.Sprintf("scale=-1:%d", animateScale)
	}else{
		diff := float64(metaJson.Height - metaJson.Width)/float64(metaJson.Width)
		if diff >= define.ScaleMorePercent {
			animateScale  = int(float64(animateScale) * (1 - define.ScaleMorePercent))
		}
		scale = fmt.Sprintf("scale=%d:-1", animateScale)
	}

	//create high quality animate gif
	args = fmt.Sprintf("-ss %s -t %d -i %s -i %s -filter_complex fps=%d,%s:flags=lanczos[x];[x][1:v]paletteuse -y %s",
		durationStart, animateSeconds, videoFilePath, palettePath, snapFps, scale, gifPath)
	_, err = f.ExecCommand(cmdOfMpeg, args)
	if err != nil {
		return nil, err
	}

	//defer remove gif file
	defer func() {
		os.RemoveAll(gifPath)
	}()

	//read gif byte data
	byteData, subErr := ioutil.ReadFile(gifPath)
	return byteData, subErr
}

//get video meta info
//include duration, width, height, rate
func (f *Video) GetMetaInfo(
	videoFilePath string) *json.VideoMetaJson {
	var (
		rotate int
		duration float64
	)

	//check
	if videoFilePath == "" {
		return nil
	}

	//setup command full path
	mainConf := conf.RunAppConfig.GetMainConf().GetConfInfo()
	commandPath := mainConf.CommandPath
	probeCmdPath := fmt.Sprintf("%v/%v", commandPath, define.CmdOfProbe)

	//set args
	args := fmt.Sprintf("-v error -select_streams v:0 -show_entries format=duration -show_entries stream=width,height -show_entries stream_tags=rotate -of default=nw=1:nk=1 %s", videoFilePath)
	info, err := f.ExecCommand(probeCmdPath, args)
	if err != nil {
		return nil
	}

	//split info by `\n`
	tempSlice := strings.Split(info, "\n")
	sliceLen := len(tempSlice)
	if tempSlice == nil || sliceLen < 4 {
		return nil
	}

	//get relate value
	width, _ := strconv.Atoi(tempSlice[0])
	height, _ := strconv.Atoi(tempSlice[1])

	if sliceLen > 4 {
		//get rotate info
		rotate, _ = strconv.Atoi(tempSlice[2])
		duration, _ = strconv.ParseFloat(tempSlice[3], 64)
	}else{
		duration, _ = strconv.ParseFloat(tempSlice[2], 64)
	}
	durationInt := int(math.Ceil(duration))

	//init meta json
	metaJson := json.NewVideoMetaJson()
	metaJson.Width = width
	metaJson.Height = height
	metaJson.Rotate = rotate
	metaJson.Duration = durationInt

	return metaJson
}
