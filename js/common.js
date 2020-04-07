const $ = {
    ajax : function (conf) {
        let xmlHttp;
        
        let method = conf.method;
        let type   = conf.type;

        if(method) {
            method = method.toUpperCase();
        }

        if(type) {
            type = type.toUpperCase();
        }

        if(window.XMLHttpRequest) {
            xmlHttp = new XMLHttpRequest();
        } else {
            xmlHttp = new ActiveXObject("Microsoft.XMLHTTP");
        }

        if(method === "GET" || method === "DELETE") {
            let data = conf.data

            if(data) {
                if(typeof data === "string") {
                    conf.url = conf.url + "/" + data
                } else if (typeof data === "object") {
                    let entries = Object.entries(data);
                    let dataStr = "";
                    if(entries.length > 0) {
                        let i = 0;
                        for(let [key, value] of entries) {
                            if(i === 0) dataStr += "?"
                            else if (i !== 0 && i == (entries.length - 1)) dataStr += "&"
                            dataStr += key + "=" + value
                            i++;
                        }
                        conf.url += dataStr;
                    }
                }
            }
        }

        xmlHttp.open(conf.method, conf.url, true)

        if(type === "JSON") {
            xmlHttp.setRequestHeader("Content-type", "application/json; charset=utf-8;")
        }

        xmlHttp.onreadystatechange = function () {
            if(xmlHttp.readyState === 4 && xmlHttp.status === 200) {
                if(conf.success) {
                    conf.success.call(xmlHttp, xmlHttp.responseText)
                }
            } else {
                if(conf.error) {
                    conf.error.call(xmlHttp)
                }
            }
        }

        if(method === "POST" || method === "PUT") {
            xmlHttp.send(JSON.stringify(conf.data))
        } else {
            xmlHttp.send(null)
        }
    }
}
