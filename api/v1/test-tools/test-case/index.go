package test_case

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"sync"
	"test-case-gin/model/test-tools/TestCase"
	"test-case-gin/utils/errmsg"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type TestCaseApiApi struct{}

func (s *TestCaseApiApi) AddTestCase(c *gin.Context) {
	var params TestCase.TestCase
	_ = c.ShouldBindJSON(&params)
	code := testCaseService.CreateTestCase(&params)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetTestCase 查询全部测试用例
func (s *TestCaseApiApi) GetTestCase(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id := c.Query("id")
	property := c.Query("property")

	switch {
	case pageSize >= 100:
		pageSize = 300
	case pageSize <= 0:
		pageSize = 10
	}
	//if pageNum == 0 {
	//	pageNum = 1
	//}
	var data []TestCase.TestCase
	var total int64
	if id == "" {
		data, total = testCaseService.GetAllTestCase(pageSize, pageNum)
	} else {
		data, total = testCaseService.GetAllTestCaseLimit(id, property, pageSize, pageNum)
	}

	// 开启信道
	var wg sync.WaitGroup
	for i := 0; i < len(data); i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			organizeInfo, _ := projectCaseService.GetFirstProject(data[index].ProjectId)
			data[index].ProjectName = organizeInfo.Name
		}(i)
	}
	wg.Wait()

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

type ID struct {
	Id int `json:"id"`
}
type IDs struct {
	Ids []int `json:"ids"`
}

// DelTestCase 删除测试实例
func (s *TestCaseApiApi) DelTestCase(c *gin.Context) {
	var params IDs
	_ = c.ShouldBindJSON(&params)
	var code int

	for _, id := range params.Ids {

		code = testCaseService.DeleteTestCase(id)
	}

	//code := model.DeleteTestCase(params.Id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// UpdateTestCase 修改测试实例
func (s *TestCaseApiApi) UpdateTestCase(c *gin.Context) {
	var params TestCase.TestCase
	_ = c.ShouldBindJSON(&params)
	code := testCaseService.UpdateTestCase(params)
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

type SearchTestCase struct {
	Id       string `json:"id"`
	Property string `json:"property"`
}

// LikeSelectTestCase 模糊查询菜单信息
func (s *TestCaseApiApi) LikeSelectTestCase(c *gin.Context) {
	var params SearchTestCase
	_ = c.ShouldBindJSON(&params)
	if params.Property == "" {
		data, total, code := testCaseService.SelectLikeTestCaseId(params.Id)
		c.JSON(
			http.StatusOK, gin.H{
				"status":  code,
				"data":    data,
				"total":   total,
				"message": errmsg.GetErrMsg(code),
			},
		)
	} else {
		data, total, code := testCaseService.SelectLikeTestCaseIdProperty(params.Id, params.Property)
		c.JSON(
			http.StatusOK, gin.H{
				"status":  code,
				"data":    data,
				"total":   total,
				"message": errmsg.GetErrMsg(code),
			},
		)
	}

}

func (s *TestCaseApiApi) ExportExcel(c *gin.Context) {

	id := c.Query("id")
	property := c.Query("property")
	var data []TestCase.TestCase
	var code int
	var err error
	if property == "" {
		data, _, code = testCaseService.SelectLikeTestCaseId(id)
	} else {
		data, _, code = testCaseService.SelectLikeTestCaseIdProperty(id, property)
	}
	if code == 500 {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}

	// 创建 Excel 文件对象
	f := excelize.NewFile()
	sheetName := "sheet1"

	// 获取结构体字段名并写入表头
	headers := getStructFieldNames(TestCase.ExportTestCase{})
	for i, header := range headers {
		col := 'A' + i
		cell := fmt.Sprintf("%c1", col)
		f.SetCellValue(sheetName, cell, header)

	}

	// 写入数据到单元格
	for i, item := range data {
		row := i + 2 // 数据行从第二行开始，因为第一行是表头
		fields := getStructFieldValues(item)
		for j, field := range fields[:10] {
			col := 'A' + j
			cell := fmt.Sprintf("%c%d", col, row)
			f.SetCellValue(sheetName, cell, field)
		}
	}

	// 设置 HTTP 响应头，告知浏览器返回文件流
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=export.xlsx")

	// 将 Excel 文件保存到响应体
	err = f.Write(c.Writer)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
}

func getStructFieldNames(s interface{}) []string {
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Struct {
		return nil
	}

	var headers []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		headers = append(headers, field.Name)
	}

	return headers
}

func getStructFieldValues(s interface{}) []interface{} {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Struct {
		return nil
	}

	var values []interface{}
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i).Interface()
		values = append(values, value)
	}

	return values
}
