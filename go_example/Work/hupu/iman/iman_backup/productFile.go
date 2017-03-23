package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"strings"
)

func init() {
	productFile = new(Backup)
	productPaht := bakPath
	productFile.pathName = productPaht

	productFile.fileName = "zentaopms.tar.gz"
	productFile.cmd = "cd /var/www/; rm -rf zentaopms.tar.gz; tar zcvf zentaopms.tar.gz --exclude=zentaopms/backup --exclude=zentaopms/tmp zentaopms/"
}

func tarFile(bak *Backup) {
	if bak.fileName == "zentaopms.tar.gz" {
		file, err := os.Open(bak.pathName + "/zentaopms.tar.gz")
		defer file.Close()
		checkError(err)

		gr, err := gzip.NewReader(file)
		defer gr.Close()
		checkError(err)

		tr := tar.NewReader(gr)
		for {
			h, err := tr.Next()
			if err == io.EOF {
				break
			}
			checkError(err)

			log.Printf("正在解压 【%s】.", h.Name)

			info := h.FileInfo()
			var dstPath string
			if strings.Contains(h.Name, "/www/bbs/") {
				dstPath = bbs.pathName
			} else {
				dstPath = product.pathName
			}

			log.Println("-----------------正在解压文件,请耐心等待.------------------")
			if info.IsDir() {
				MkBakDir(dstPath + "/" + h.Name)
			} else {
				fw, err := os.OpenFile(dstPath+"/"+h.Name, os.O_CREATE|os.O_WRONLY, 0755)
				defer fw.Close()
				checkError(err)

				_, err = io.Copy(fw, tr)
				checkError(err)
			}
		}

		// 删除tar文件
		file.Close()
		os.RemoveAll(bak.pathName + "/zentaopms.tar.gz")
		log.Println("----------文件【zentaopms.tar.gz】解压成功.------------")
	}
}
