package generate

import (
	"github.com/yesilin/go-cutting/tools"
	"strings"
)

//生成取得管理员所有权
func TakeOwnership() {
	var reg = strings.Builder{}

	reg.WriteString("Windows Registry Editor Version 5.00\r\n")
	reg.WriteString(";取得文件修改权限\r\n")
	reg.WriteString("[HKEY_CLASSES_ROOT\\*\\shell\\runas]\r\n")
	reg.WriteString("@=\"管理员取得所有权\"\r\n")
	reg.WriteString("\"Icon\"=\"C:\\\\Windows\\\\System32\\\\imageres.dll,73\"\r\n")
	reg.WriteString("\"NoWorkingDirectory\"=\"\"\r\n")
	reg.WriteString("[HKEY_CLASSES_ROOT\\*\\shell\\runas\\command]\r\n")
	reg.WriteString("@=\"cmd.exe /c takeown /f \\\"%1\\\" && icacls \\\"%1\\\" /grant administrators:F\"\r\n")
	reg.WriteString("\"IsolatedCommand\"=\"cmd.exe /c takeown /f \\\"%1\\\" && icacls \\\"%1\\\" /grant administrators:F\"\r\n")
	reg.WriteString("[HKEY_CLASSES_ROOT\\exefile\\shell\\runas2]\r\n")
	reg.WriteString("@=\"管理员取得所有权\"\r\n")
	reg.WriteString("\"Icon\"=\"C:\\\\Windows\\\\System32\\\\imageres.dll,73\"\r\n")
	reg.WriteString("\"NoWorkingDirectory\"=\"\"\r\n")
	reg.WriteString("[HKEY_CLASSES_ROOT\\exefile\\shell\\runas2\\command]\r\n")
	reg.WriteString("@=\"cmd.exe /c takeown /f \\\"%1\\\" && icacls \\\"%1\\\" /grant administrators:F\"\r\n")
	reg.WriteString("\"IsolatedCommand\"=\"cmd.exe /c takeown /f \\\"%1\\\" && icacls \\\"%1\\\" /grant administrators:F\"\r\n")
	reg.WriteString("[HKEY_CLASSES_ROOT\\Directory\\shell\\runas]\r\n")
	reg.WriteString("@=\"管理员取得所有权\"\r\n")
	reg.WriteString("\"Icon\"=\"C:\\\\Windows\\\\System32\\\\imageres.dll,73\"\r\n")
	reg.WriteString("\"NoWorkingDirectory\"=\"\"\r\n")
	reg.WriteString("[HKEY_CLASSES_ROOT\\Directory\\shell\\runas\\command]\r\n")
	reg.WriteString("@=\"cmd.exe /c takeown /f \\\"%1\\\" /r /d y && icacls \\\"%1\\\" /grant administrators:F /t\"\r\n")
	reg.WriteString("\"IsolatedCommand\"=\"cmd.exe /c takeown /f \\\"%1\\\" /r /d y && icacls \\\"%1\\\" /grant administrators:F /t\"")

	// 转成字符串格式
	regStr := reg.String()
	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("Config\\takeOwnership.reg", regStr)
}
