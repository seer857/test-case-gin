package projectDocument

import (
	"net/http"
	"test-case-gin/model/test-tools/ProjectDocument"
	"test-case-gin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

type ProjectDocumentApiApi struct{}

func buildTree(parentId string, data []ProjectDocument.ProjectDocument) []ProjectDocument.ProjectDocumentTree {
	subTree := make([]ProjectDocument.ProjectDocumentTree, 0)
	for _, m := range data {
		if m.ParentId == parentId {
			children := buildTree(m.ID, data)
			t := ProjectDocument.ProjectDocumentTree{
				ID:        m.ID,
				ParentId:  m.ParentId,
				Property:  m.Property,
				Name:      m.Name,
				ProjectId: m.ProjectId,
				Children:  children,
			}
			subTree = append(subTree, t)
		}
	}

	return subTree
}

func handleTree(data []ProjectDocument.ProjectDocument) []ProjectDocument.ProjectDocumentTree {
	return buildTree("", data)
}

// SelectDocumentTree 查询菜单树
func (p *ProjectDocumentApiApi) SelectDocumentTree(c *gin.Context) {
	id := c.Query("id")
	data, total := projectDocumentService.GetAllDocument(id)
	tree := handleTree(data)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    tree,
		"total":   total,
	})
}

// AddProjectDocument 新增文档目录
func (p *ProjectDocumentApiApi) AddProjectDocument(c *gin.Context) {
	var params ProjectDocument.ProjectDocument
	_ = c.ShouldBindJSON(&params)
	code := projectDocumentService.CreateProjectDocument(&params)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DelProjectDocument 删除项目文档
func (p *ProjectDocumentApiApi) DelProjectDocument(c *gin.Context) {
	id := c.Query("id")

	code := projectDocumentService.DeleteProjectDocument(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}
