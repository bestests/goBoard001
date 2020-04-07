const $ = {
    ajax : function (conf) {
        var xmlHttp;

        if(window.XMLHttpRequest) {
            xmlHttp = new XMLHttpRequest();
        } else {
            xmlHttp = new ActiveXObject("Microsoft.XMLHTTP");
        }

        xmlHttp.onreadystatechange = function () {
            console.log(xmlHttp.readyState)
            console.log(xmlHttp.status)
            console.log("================================");
            if(xmlHttp.readyState === 4 && xmlHttp.status === 200) {
                conf.success.bind(xmlHttp)
            }
        }

        xmlHttp.open(conf.method, conf.url, true)
        xmlHttp.send()
    }
}
