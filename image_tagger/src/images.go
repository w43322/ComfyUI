package images

import (
	"fmt"
	"github.com/w43322/ComfyUI/image_tagger/biz/model/comfyui/image_tagger"
	"os"
	"strconv"
	"strings"
)

var imageList []*image_tagger.Image

func GetImages(tags []string) []*image_tagger.Image {
	return imageList
}

func init() {
	imageList = getImageListFromLocalFileSystem("/Users/bytedance/git/ComfyUI/output")
}

func getImageInfoFromFileName(fileName string) *image_tagger.Image {
	splits := strings.Split(fileName, "_")
	seed, _ := strconv.Atoi(splits[len(splits)-2])
	batchIndex := int64(seed & 0xf)
	provinceId := int64(seed >> 28)
	ownerTag := fmt.Sprintf("%c%c%c", seed>>20&0xff, seed>>12&0xff, seed>>4&0xff)

	return &image_tagger.Image{
		URL:    fmt.Sprintf("http://127.0.0.1:6789/output/%s", fileName),
		Width:  2400,
		Height: 1200,
		Tags:   nil,
		ProvinceMeta: &image_tagger.ProvinceMeta{
			ProvinceID: provinceId,
			OwnerTag:   ownerTag,
		},
		BatchIndex: batchIndex,
	}
}

func getImageListFromLocalFileSystem(path string) []*image_tagger.Image {
	dir, err := os.Open(path)
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return nil
	}

	ret := make([]*image_tagger.Image, 0, len(fileInfos)/2)
	for _, fileInfo := range fileInfos {
		fileName := fileInfo.Name()
		if !strings.HasSuffix(fileName, "_upscaled.png") {
			continue
		}
		image := getImageInfoFromFileName(fileName)
		ret = append(ret, image)
	}

	return ret
}
