package model

import (
	"github.com/yesilin/go-cutting/tools"
	"os/exec"
)

// SaveAndCloseAllDocumentsInit 保存并关闭全部文档的初始化 暗号-12的实现
func SaveAndCloseAllDocumentsInit() {
	const script = `// 保存并关闭所有文件
function saveAndClose() {
    // 得到要保存关闭的文档的数量，主要是documents关闭后会直接刷新
    const count = documents.length
 
    // 循环保存所有
    for (var i = 0; i < count; i++) {

        try {
            // 最后保存并关闭
            app.activeDocument.close(SaveOptions.SAVECHANGES);
        } catch (error) {
            // 忽略错误
        }
    }

}

// 添加进度条
function progressBar() {
    app.doForcedProgress("正在保存并关闭全部文件... ", "saveAndClose()")
}

// 主函数
function main() {
    if (!documents.length) {
        //alert("没有打开的文档，请打开一个文档来运行此脚本！");
        return;
    }

    // 运行进度条
    progressBar()
}

// 运行
main()`

	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("config/jsx/saveAndCloseAllDocuments.jsx", script)
}

// SaveAndCloseAllDocuments 保存并关闭全部文档的调用
func SaveAndCloseAllDocuments(){
	// 创建一个协程使用cmd启动外部程序
	dataPath := "config/jsx/saveAndCloseAllDocuments.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}