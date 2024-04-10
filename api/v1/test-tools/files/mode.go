package files

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"test-case-gin/global"
	"test-case-gin/model/test-tools/Files"
	"test-case-gin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func (filesApiApi *FilesApiApi) DataInput(c *gin.Context) {
	var files Files.Files
	_ = c.ShouldBindJSON(&files)

	// 将 params 转换为 JSON 格式的字节数据
	jsonData, err := json.Marshal(files)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	url := "http://localhost:5000/data_input"
	// 发起POST请求
	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			// 处理关闭响应体时的错误
		}
	}()

	files.FileStatus = "2"
	tag := filesService.UpdateFiles(&files)

	// 读取响应体
	body, err := io.ReadAll(response.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": errmsg.GetErrMsg(200),
		"data":    string(body),
		"tag":     tag,
	})
}

// SelectInputData 对导入数据查看
func (filesApiApi *FilesApiApi) SelectInputData(c *gin.Context) {
	var files Files.Files
	_ = c.ShouldBindJSON(&files)
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	// 假设您要查询的表名为 "DS_" + files.FileName
	tableName := "DS_" + files.FileName

	var results []map[string]interface{}

	err := global.GVA_DB.Raw("SELECT * FROM `"+tableName+"` LIMIT ? OFFSET ?", pageSize, pageNum).Scan(&results).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
		return
	}
	var count int64
	global.GVA_DB.Raw("SELECT COUNT(*) FROM `" + tableName + "`").Scan(&count)

	c.JSON(http.StatusOK, gin.H{
		"data":  results,
		"total": count,
	})
}
