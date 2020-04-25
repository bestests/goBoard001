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
        console.log("method : " + method);
        console.log("type : " + type);

        if(window.XMLHttpRequest) {
            xmlHttp = new XMLHttpRequest();
        } else {
            xmlHttp = new ActiveXObject("Microsoft.XMLHTTP");
        }

        let formData = new FormData();

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
        } else {
            if(conf.data) {
                let entries = Object.entries(conf.data);
                console.log(entries)
                if(entries.length > 0) {
                    let i = 0;
                    for(let [key, value] of entries) {
                        console.log(key, value)
                        formData.append(key, value);
                    }
                }
            }
        }
        console.log("url : " + conf.url);
        xmlHttp.open(conf.method, conf.url, true)

        if(type === "JSON") {
           //xmlHttp.setRequestHeader("Content-type", "application/json; charset=utf-8;")
        }

        xmlHttp.onreadystatechange = function () {
            console.log("state : " + xmlHttp.readyState)
            console.log("code : " + xmlHttp.status)
            if(xmlHttp.readyState === 4 && xmlHttp.status === 200) {
                if(conf.success) {
                    conf.success.call(xmlHttp, xmlHttp.responseText)
                }
            } else {
                if(xmlHttp.readyState === 4 && xmlHttp.status >= 400) {
                    if(conf.error) {
                        conf.error.call(xmlHttp, xmlHttp.status)
                    }
                }
            }
        }
        console.dir(conf.data)
        if(method === "POST" || method === "PUT") {
            formData.forEach(function (value, key) {
                console.log("key", key, "value", value)
            })
            xmlHttp.send(formData)
        } else {
            xmlHttp.send(null)
        }
    }
}
