package test_case

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/lukasjarosch/go-docx"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

func (s *TestCaseApiApi) ReadWordTemplate(c *gin.Context) {
	id := c.Query("id")
	data := testCaseService.SelectCaseTest(id)

	doc, _ := docx.Open("uploads/caseTemplate/template.docx")

	// 创建一个新的替换映射，将data的值与对应的占位符关联起来
	replaceMap := docx.PlaceholderMap{
		"name":           data.Name,
		"caseName":       data.CaseName,
		"prerequisites":  data.Prerequisites,
		"testProcedure":  data.TestProcedure,
		"expectedResult": data.ExpectedResult,
		"remark":         data.Remark,
	}

	// 将data的值替换到文档中
	err := doc.ReplaceAll(replaceMap)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": err,
		})
	}
	//// 复制第一页内容并替换数据
	//for i := 1; i < total; i++ {
	//	sourcePage, err := doc.GetPage(1)
	//	if err != nil {
	//		panic(err)
	//	}
	//	err = doc.AddPage(sourcePage)
	//	if err != nil {
	//		panic(err)
	//	}
	//}

	_ = doc.WriteToFile("uploads/caseTemplate/replaced.docx")

}

type Blob struct {
	ContentType string
	Data        []byte
}

func (s *TestCaseApiApi) FileBlob(c *gin.Context) {
	file, err := os.Open("uploads/caseTemplate/replaced.docx")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	contentType := mime.TypeByExtension(filepath.Ext("path/to/file"))

	blob := &Blob{
		ContentType: contentType,
		Data:        data,
	}

	c.DataFromReader(http.StatusOK, int64(len(blob.Data)), blob.ContentType, bytes.NewReader(blob.Data), map[string]string{})
}
