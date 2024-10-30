package controllers

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/gin-gonic/gin"
	trash "github.com/hymkor/trash-go"
)

type SystemController struct {
}

func (sc *SystemController) Pong(c *gin.Context) {
	ReturnSuccessResponse(c, "OK")
}

// 打开文件管理器
// @platform: all
func (sc *SystemController) Explore(c *gin.Context) {
	param := make(map[string]interface{})
	err := c.BindJSON(&param)
	if err != nil {
		ReturnFailResponse(c, 400, "no param entry")
		return
	}

	entry, ok := param["entry"]
	if !ok {
		ReturnFailResponse(c, 400, "json has no field: entry")
	}

	entryStr, ok := entry.(string)
	if !ok {
		ReturnFailResponse(c, 400, "param entry should be string")
		return
	}

	_, err = os.Stat(entryStr)
	if os.IsNotExist(err) {
		ReturnFailResponse(c, 400, "file not exist")
		return
	}

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("explorer", entryStr)
	case "darwin":
		cmd = exec.Command("open", entryStr)
	case "linux":
		cmd = exec.Command("xdg-open", entryStr)
	default:
		ReturnFailResponse(c, 400, "unsupported platform")
		return
	}

	err = cmd.Run()
	if err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			ReturnFailResponse(nil, 400, "command run failed")
			return
		} else {
			ReturnSuccessResponse(c, "OK")
			return
		}
	} else {
		ReturnFailResponse(nil, 400, "command run failed")
		return
	}
}

// 将文件移动到回收站
// @platform: window
func (sc *SystemController) Trash(c *gin.Context) {
	param := make(map[string]interface{})
	err := c.BindJSON(&param)
	if err != nil {
		ReturnFailResponse(c, 400, "no param entry")
		return
	}

	entry, ok := param["entry"]
	if !ok {
		ReturnFailResponse(c, 400, "json has no field: entry")
	}

	entryStr, ok := entry.(string)
	if !ok {
		ReturnFailResponse(c, 400, "param entry should be string")
		return
	}

	err = trash.Throw(entryStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		ReturnFailResponse(nil, 400, "trash failed")
		return
	}
	ReturnSuccessResponse(c, nil)
}
