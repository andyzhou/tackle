package file

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/andyzhou/pond"
)

//base face
type Base struct {
}

//gen md5
func (f *Base) GenMd5(dataBytes []byte) string {
	m := md5.New()
	m.Write(dataBytes)
	return hex.EncodeToString(m.Sum(nil))
}

//init sub pond instance
func (f *Base) InitPond(dataPath string) (*pond.Pond, error) {
	//init sub pond obj
	pond := pond.NewPond()

	//gen new config
	cfg := pond.GenConfig()
	cfg.DataPath = dataPath
	cfg.CheckSame = true
	cfg.WriteLazy = false
	cfg.FixedBlockSize = true
	cfg.UseMemoryMap = true

	//set config
	err := pond.SetConfig(cfg)
	if err != nil {
		return nil, err
	}
	return pond, nil
}