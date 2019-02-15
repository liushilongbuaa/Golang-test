package main

import (
	"fmt"
	"io/ioutil"
	"syscall"
	"time"
)

const (
	ftyp uint32 = 0x66747970
)

func main() {
	fileInfos, err := ioutil.ReadDir("c:/Users/liushilong/Downloads")
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, v := range fileInfos {
		if v.Name() != "d10f30157912f2f18882ef8c1248ef24.mp4" {
			continue
		}
		tem := v.Sys().(*syscall.Win32FileAttributeData)
		date := time.Duration(time.Duration(tem.CreationTime.Nanoseconds()) * time.Nanosecond)
		fmt.Println(v.Name(), v.Size(), date.Hours())
	}
	fmt.Println("###########################################")

	mp4file, err := ioutil.ReadFile("c:/Users/liushilong/Downloads/d10f30157912f2f18882ef8c1248ef24.mp4")
	if err != nil {
		fmt.Println("ReadFile" + err.Error())
		return
	}
	var boxHeaders []BoxHeader
	var boxlength int
	for index := 0; index < len(mp4file); index += boxlength {
		if mp4file[index]|mp4file[index+1]|mp4file[index+2]|mp4file[index+3] == 0x01 {
			fmt.Println("length exceed.")
			return
		}
		box := BoxHeader{}
		box.Type = string(mp4file[index+4 : index+8])
		box.Size = uint32(mp4file[index])<<24 + uint32(mp4file[index+1])<<16 + uint32(mp4file[index+2])<<8 + uint32(mp4file[index+3])
		boxHeaders = append(boxHeaders, box)
		boxlength = int(box.Size)
	}
	fmt.Println(boxHeaders)
	var offset int
	for i := 0; i < len(boxHeaders)-1; i++ {
		offset += int(boxHeaders[i].Size)
	}
	moov := mp4file[offset+8:]
	var _boxHeaders []BoxHeader
	boxlength = 0
	for index := 0; index < len(moov); index += boxlength {
		if moov[index]|moov[index+1]|mp4file[index+2]|mp4file[index+3] == 0x01 {
			fmt.Println("length exceed.")
			return
		}
		box := BoxHeader{}
		box.Type = string(moov[index+4 : index+8])
		box.Size = uint32(moov[index])<<24 + uint32(moov[index+1])<<16 + uint32(moov[index+2])<<8 + uint32(moov[index+3])
		_boxHeaders = append(_boxHeaders, box)
		boxlength = int(box.Size)
	}

	fmt.Println(_boxHeaders)
}

//func printBox(in []byte) {
//	var boxlen []uint32
//	var boxtype []uint32
//	for i := 0; i < len(tem); {
//		var sum uint32
//		sum += uint32(tem[i]) << 24
//		sum += uint32(tem[i+1]) << 16
//		sum += uint32(tem[i+2]) << 8
//		sum += uint32(tem[i+3])
//		boxlen = append(boxlen, sum)

//		var type32 uint32
//		type32 += uint32(tem[i+4]) << 24
//		type32 += uint32(tem[i+5]) << 16
//		type32 += uint32(tem[i+6]) << 8
//		type32 += uint32(tem[i+7])
//		boxtype = append(boxtype, type32)

//		i += int(sum)
//		sum = 0
//	}
//	fmt.Println(boxlen)
//	fmt.Printf("%x\n", boxtype)
//}
