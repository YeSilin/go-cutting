package generate

import (
	"github.com/yesilin/go-cutting/tools"
	"strings"
)

//生成取得管理员所有权
func AcquisitionOfOwnershipReg() {
	var reg = strings.Builder{}

	reg.WriteString("Windows Registry Editor Version 5.00\n")
	reg.WriteString(";取得文件修改权限\n")
	reg.WriteString("[HKEY_CLASSES_ROOT\\*\\shell\\runas]\n")
	reg.WriteString("@=\"管理员取得所有权\"\n")
	reg.WriteString("\"Icon\"=\"C:\\\\Windows\\\\System32\\\\imageres.dll,73\"\n")
	reg.WriteString("\"NoWorkingDirectory\"=\"\"\n")
	reg.WriteString("[HKEY_CLASSES_ROOT\\*\\shell\\runas\\command]\n")
	reg.WriteString("@=\"cmd.exe /c takeown /f \\\"%1\\\" && icacls \\\"%1\\\" /grant administrators:F\"\n")
	reg.WriteString("\"IsolatedCommand\"=\"cmd.exe /c takeown /f \\\"%1\\\" && icacls \\\"%1\\\" /grant administrators:F\"\n")
	reg.WriteString("[HKEY_CLASSES_ROOT\\exefile\\shell\\runas2]\n")
	reg.WriteString("@=\"管理员取得所有权\"\n")
	reg.WriteString("\"Icon\"=\"C:\\\\Windows\\\\System32\\\\imageres.dll,73\"\n")
	reg.WriteString("\"NoWorkingDirectory\"=\"\"\n")
	reg.WriteString("[HKEY_CLASSES_ROOT\\exefile\\shell\\runas2\\command]\n")
	reg.WriteString("@=\"cmd.exe /c takeown /f \\\"%1\\\" && icacls \\\"%1\\\" /grant administrators:F\"\n")
	reg.WriteString("\"IsolatedCommand\"=\"cmd.exe /c takeown /f \\\"%1\\\" && icacls \\\"%1\\\" /grant administrators:F\"\n")
	reg.WriteString("[HKEY_CLASSES_ROOT\\Directory\\shell\\runas]\n")
	reg.WriteString("@=\"管理员取得所有权\"\n")
	reg.WriteString("\"Icon\"=\"C:\\\\Windows\\\\System32\\\\imageres.dll,73\"\n")
	reg.WriteString("\"NoWorkingDirectory\"=\"\"\n")
	reg.WriteString("[HKEY_CLASSES_ROOT\\Directory\\shell\\runas\\command]\n")
	reg.WriteString("@=\"cmd.exe /c takeown /f \\\"%1\\\" /r /d y && icacls \\\"%1\\\" /grant administrators:F /t\"\n")
	reg.WriteString("\"IsolatedCommand\"=\"cmd.exe /c takeown /f \\\"%1\\\" /r /d y && icacls \\\"%1\\\" /grant administrators:F /t\"\n")

	// 转成字符串格式
	regStr := reg.String()
	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("Config\\AcquisitionOfOwnership.reg", regStr)
}
