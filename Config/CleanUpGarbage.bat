:: 临时文件(如:*.tmp?*._mp)日志文件(*.log)?临时帮助文件(*.gid)?磁盘检查文件(*.chk)?临时备份文件(如: *.old?*.bak) 
:: %systemdrive%是系统盘，%windir%是系统文件夹，%userprofile%当前用户文件夹,recent是最近浏览过的文档 

:: del /f /s /q %windir%\prefetch\*.*  删除应用程序临时文件
:: 上面这一行删除的是“系统预先装载”文件，一般不要删 

:: rd /s /q %windir%\temp & md %windir%\temp  删除系统维护等操作产生的临时文件（删除目录后再创建目录）
:: 上一行的操作会改变文件夹的属性，这个属性会影响asp+access程序,所以给注释掉了，换成下边的了,缺点是不能删除这个文件夹下的文件夹，但文件都删除了 

:: 去掉了,下边这一行，今天在控制面板里装DNS时提示要sp2 光盘，插上我制作的光盘不行。好像必须到下面这个文件夹里找，虽然一般人都不会用到，但还是不要删除这个为好。 
:: del /f /s /q %windir%\ServicePackFiles\*.* 

:: 去掉了，下边这一行，虽然传说中没事，但没有事，不代表不会有事,系统经常出错了，可以把这个文件夹下的东西删除，系统会重建。 
:: del /f /s /q %windir%\SoftwareDistribution\Download\*.* 

:: 去掉了下边这一行，因为：有些安装信息会放到这里面，删除的话卸载软件时会有很大的麻烦，还有，如果您是用在服务器方面，这些记录是绝对不可以删除的，日志文件很重要的。 
:: del /f /s /q %systemdrive%\*.log 

:: 1>nul 意思是不显示命令运行的正确提示
:: 2>nul 是不显示错误提示
:: 一起就是 正确错误的都不显示


:: 关闭回显
@echo off
:: 换行
echo.
echo 正在清除系统垃圾文件，请稍等...... 

:: 删除系统盘目录下临时文件
del /f /s /q %systemdrive%\*.tmp 
del /f /s /q %systemdrive%\*._mp 

:: 删除系统盘目录下GID文件（属于临时文件，具体作用不详）
del /f /s /q %systemdrive%\*.gid 

:: 删除系统目录下scandisk（磁盘扫描）留下的无用文件
del /f /s /q %systemdrive%\*.chk 

:: 删除系统目录下old文件
del /f /s /q %systemdrive%\*.old 

:: 删除回收站的无用文件
del /f /s /q %systemdrive%\recycled\*.* 

:: 新增加的清除C:根目录下的MSN"用户体验改善计划"生成的临时文件 
del /f /a /q %systemdrive%\*.sqm 

:: 删除系统目录下备份文件
del /f /s /q %windir%\*.bak 

:: 删除系统维护等操作产生的临时文件
del /f /s /q %windir%\temp\*.* 

:: 可以把这个文件夹下的东西删除，系统会重建
del /f /s /q %windir%\SoftwareDistribution\Download\*.* 

:: 删除当前用户的COOKIE（IE）
del /f /s /q %userprofile%\cookies\*.* 

:: 删除internet临时文件
del /f /s /q "%userprofile%\Local Settings\Temporary Internet Files\*.*" 

:: 删除当前用户日常操作临时文件
del /f /s /q "%userprofile%\Local Settings\Temp\*.*" 

:: 删除访问记录（开始菜单中的文档里面的东西）
del /f /s /q "%userprofile%\recent\*.*" 

echo 清除系统LJ完成！
echo. & pause