package files

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"test-case-gin/global"
	"test-case-gin/model/test-tools/Files"
	"test-case-gin/utils/errmsg"
	"time"

	"github.com/gin-gonic/gin"
)

type FilesApiApi struct{}

func (filesApiApi *FilesApiApi) GetFiles(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data, total := filesService.GetAllFiles(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// DelFile 删除文件
func (filesApiApi *FilesApiApi) DelFile(c *gin.Context) {
	//id, _ := strconv.Atoi(c.Query("id"))
	//uploadType := c.Query("uploadType")
	var params Files.Files
	_ = c.ShouldBindJSON(&params)

	if params.UploadType != "" {
		tableName := "DS_" + params.FileName
		query := fmt.Sprintf("DROP TABLE `%s`", tableName)
		err := global.GVA_DB.Exec(query).Error
		if err != nil {
			panic(err.Error())
		}
	}

	file, code := filesService.DeleteFile(params.ID)

	err := os.Remove(file.FileUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to delete file",
		})
		return
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// GetTypeFiles 根据类型获取文件列表
func (filesApiApi *FilesApiApi) GetTypeFiles(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	uploadType := c.Query("uploadType")
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data, total := filesService.GetUploadTypeFiles(pageSize, pageNum, uploadType)
	code := errmsg.SUCCESS
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

func (filesApiApi *FilesApiApi) UploadFile(c *gin.Context) {
	var files Files.Files
	// 从请求中获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.String(errmsg.BADREQUEST, "Bad Request")
		return
	}

	fileExt := filepath.Ext(file.Filename)
	//if fileExt != ".xlsx" && fileExt != ".xls" {
	//	c.JSON(400, gin.H{"error": "上传的不是excel文件"})
	//	return
	//}
	files.FileNameZh = file.Filename
	file.Filename = files.FileNameZh + "_" + strconv.FormatInt(time.Now().Unix(), 10)
	files.FileName = file.Filename
	files.FileType = fileExt
	files.FileUrl = "uploads/excel/" + files.FileName + fileExt
	// 将文件保存到指定路径
	if err := c.SaveUploadedFile(file, files.FileUrl); err != nil {
		c.JSON(500, gin.H{"error": "Failed to save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "文件上传成功"})

	// 特殊处理 - 入口
	uploadType := c.Query("uploadType")
	if uploadType != "" {
		files.UploadType = uploadType
		files.FileStatus = "0"
		// 读取文件头部第一列，gorm创建数据库表 表名为 file.Filename 列名为 文件列名 类型varchar(100)
		HandleCSV(c, files)
	}
	filesService.CreateFiles(&files)
}

func HandleCSV(c *gin.Context, files Files.Files) {
	// 打开CSV文件
	csvFile, err := os.Open(files.FileUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}

	// 创建CSV Reader
	reader := csv.NewReader(csvFile)
	reader.LazyQuotes = true

	// 读取第一行数据
	columnValues, err := reader.Read()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read CSV file"})
		return
	}

	// 创建数据库表
	tableName := "DS_" + files.FileName
	columns := make([]string, len(columnValues))
	for i, value := range columnValues {
		columns[i] = fmt.Sprintf("`%s` varchar(100)", value)
	}
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (%s);", tableName, strings.Join(columns, ","))
	err = global.GVA_DB.Exec(query).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create table"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "CSV file imported successfully"})
	defer func(csvFile *os.File) {
		err := csvFile.Close()
		if err != nil {

		}
	}(csvFile)
}
