/* 消息弹窗 */
var $$ = mdui.JQ;
$$('#signal-1').on('click', function () {
    mdui.snackbar({
        message: '正在自动裁剪，请稍等...',
        timeout: 1000,
        position: 'bottom',
    });
});
$$('#signal-2').on('click', function () {
    mdui.snackbar({
        message: '正在重建新文档，请稍等...',
        timeout: 1000,
        position: 'bottom',
    });
});
$$('#signal-3').on('click', function () {
    mdui.snackbar({
        message: '正在深度清理PSD，请稍等...',
        timeout: 1000,
        position: 'bottom',
    });
});
$$('#signal-6').on('click', function () {
    mdui.snackbar({
        message: '正在快速清理PSD，请稍等...',
        timeout: 1000,
        position: 'bottom',
    });
});
$$('#signal-7').on('click', function () {
    mdui.snackbar({
        message: '正在为当前文档加黑边，请稍等...',
        timeout: 1000,
        position: 'bottom',
    });
});
$$('#signal-9').on('click', function () {
    mdui.snackbar({
        message: '正在打开切图历史，请稍等...',
        timeout: 1000,
        position: 'bottom',
    });
});
$$('#signal-10').on('click', function () {
    mdui.snackbar({
        message: '正在优化另存中，请稍等...',
        timeout: 1000,
        position: 'bottom',
    });
});
$$('#signal-98').on('click', function () {
    mdui.snackbar({
        message: '正在导出为 Web 格式的 JPG，请稍等...',
        timeout: 1000,
        position: 'bottom',
    });
});





/* 鼠标点击文字特效 */
var f_idx = 0;
var c_idx = 0;
jQuery(document).ready(function ($) {
    $("body").click(function (e) {
        var font = new Array("富强", "民主", "文明", "和谐", "自由", "平等", "公正", "法治", "爱国", "敬业", "诚信", "友善");
        var color = new Array('#ff0000', '#eb4310', '#f6941d', '#fbb417', '#ffff00', '#cdd541', '#99cc33', '#3f9337', '#219167', '#239676', '#24998d', '#1f9baa', '#0080ff', '#3366cc', '#333399', '#003366', '#800080', '#a1488e', '#c71585', '#bd2158');
        var $i = $("<span />").text(font[f_idx]);
        f_idx = (f_idx + 1) % font.length;
        c_idx = (c_idx + 1) % color.length;
        var x = e.pageX,
            y = e.pageY;
        $i.css({
            "z-index": 99999999999999999999999999999999999999999999999999999999999999999999999999,
            "top": y - 20,
            "left": x,
            "position": "absolute",
            "font-weight": "bold",
            "color": color[c_idx],
            "font-size": "12px" /* 字体大小 */
        });
        $("body").append($i);
        $i.animate({
                "top": y - 180,
                "opacity": 0
            },
            1500,
            function () {
                $i.remove();
            });
    });
});
























