//base property and function

//global variable
var rootAppReqUrl = "";
var rootApiReqUrl = "/api";
var ssoLoginUrl = rootAppReqUrl+"/sso";
var pageMainDiv = `mainDiv`;

//delay load micro seconds
var delayLoadMSeconds = 500; //0.5 seconds, general
var delayLoadLongMSeconds = 1000; //1 seconds
var delayLoadSlowMSeconds = 3000; //3 seconds

//special error code
var errCodeOfSucceed = 1000;
var errCodeOfNeedLogin = 1010;
var errCodeOfHasBadwords = 1011;

//for tip message
var tipMessageTimeOut = 5000; //xx micro seconds


//general download file
// 通用下载文件函数（兼容 PC + iOS）
async function downloadFile(url) {
  try {
    const response = await fetch(url, {
      method: 'GET',
      // 如果需要携带 Cookie，取消下面注释
      // credentials: 'include'
    });

    if (!response.ok) throw new Error('下载失败');

    // 获取文件名
    let filename = 'download';
    const disposition = response.headers.get('Content-Disposition');
    if (disposition && disposition.includes('filename=')) {
      const match = disposition.match(/filename\*?=(?:UTF-8'')?["']?([^;"']+)/i);
      if (match && match[1]) filename = decodeURIComponent(match[1]);
    }

    const blob = await response.blob();

    // iOS Safari / WebView 检测
    const isIOS = /iP(ad|hone|od)/.test(navigator.userAgent);

    if (isIOS) {
      // 使用 FileReader 转 Data URL 打开新标签页
      const reader = new FileReader();
      reader.onload = function(e) {
        const dataUrl = e.target.result;
        const a = document.createElement('a');
        a.href = dataUrl;
        a.target = '_blank'; // iOS 打开新窗口
        a.click();
      };
      reader.readAsDataURL(blob);
    } else {
      // PC 浏览器直接使用 Blob URL + download
      const blobUrl = window.URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = blobUrl;
      a.download = filename;
      document.body.appendChild(a);
      a.click();
      a.remove();

      // 延迟释放 URL，保证浏览器完成下载
      setTimeout(() => window.URL.revokeObjectURL(blobUrl), 1000);
    }

  } catch (err) {
    console.error('下载失败:', err);
  }
}



//count down trigger
function countDownTrigger(countDiv, maxSeconds) {
    //check
    if(countDiv == "" || maxSeconds <= 0) {
        return
    }
    let countdownInterval;
    let remainingSeconds = maxSeconds;

    function updateDisplay() {
        $('#'+countDiv).text(remainingSeconds);
    }

    function countDown() {
        clearInterval(countdownInterval);
        updateDisplay();

        countdownInterval = setInterval(function() {
            remainingSeconds--;
            updateDisplay();

            if (remainingSeconds <= 0) {
                clearInterval(countdownInterval);
            }
        }, 1000);
    }

    //begin count down
    countDown();
}

//get string length
function getStrLen(input) {
    return new TextEncoder().encode(input).length;
}

//scroll to top
function scrollToTop() {
  $(window).scrollTop(0);
}


//sub page url jump
function pageJump(pageUrl) {
    //check
    if(typeof(pageUrl) == "undefined" || pageUrl == null) {
        return
    }

    //detect page url
    var finalPageUrl = formatFinalPage(pageUrl);

    //console.log("finalPageUrl:"+finalPageUrl);

    //change browser url
    //use original page url
    changeBrowserUrl(pageUrl);

    //send ajax page load
    sendAjaxPageReq(finalPageUrl, pageMainDiv, null);
}

//format final page
function formatFinalPage(pageUrl) {
    var finalPageUrl = pageUrl;
    if (pageUrl.indexOf('?') == -1) {
        finalPageUrl = pageUrl + "?page=true";
    }else{
        finalPageUrl = pageUrl + "&page=true";
    }
    return finalPageUrl;
}

//go back page
function goBack() {
    var url = window.location.href;
    //isGoBack = true;
    history.back();
}

//check input is numeric
function isNumeric(input) {
    return $.isNumeric(input);
}

//get current href
function getCurHref() {
    return window.location.href;
}

//gen new object
function genNewObject() {
    return new Object();
}

//decode json string into obj
function decJsonStr(str) {
    return JSON.parse(str);
}

//encode obj into json string
function encJsonObj(obj) {
    return JSON.stringify(obj);
}

function getCurTimeStamp() {
    var currentDate = new Date();
    var timestamp = currentDate.getTime();
    return timestamp;
}

//format timestamp
function formatTimeStamp(unixtime) {
    if(typeof(unixtime) == "undefined" || unixtime <= 0) {
        return "";
    }
    var newDate = new Date();
    newDate.setTime(unixtime*1000);
    dateString = newDate.toUTCString();
    return dateString;
}

//gen random num
function genRandomNum(max) {
    return Math.floor((Math.random()*max)+1);
}

//delay a while and run function
function delayRun(mSeconds, func) {
    //check
    if(typeof(mSeconds) == "undefined" || mSeconds <= 0) {
        return
    }
    if(typeof(mSeconds) == "undefined" || func == null) {
        return
    }
    setTimeout(func, mSeconds);
}


//json string check
function isJsonStr(str) {
    if (typeof str == 'string') {
        try {
            var obj=JSON.parse(str);
            if(typeof obj == 'object' && obj ){
                return true;
            }else{
                return false;
            }
        } catch(e) {
            return false;
        }
    }else if (typeof str == 'object'  && str) {
        return true;
    }
}

//copy div content to clipboard
function copyToClipboard(divId) {
    //check
    if(typeof(divId) == "undefined" || divId == "") {
        return
    }

    var text = $("#"+divId).val();
    var $temp = $("<input>");
    $("body").append($temp);
    $temp.val(text).select();
    document.execCommand("copy");
    $temp.remove();

    //console.log("copyToClipboard, divId:"+divId+", text:"+text);

    //fload tips
    var tips = baseLang.get("copied") + " " + text;
    floatTipMessage(tips, "info");
}

//dynamic load page div by ajax request way
//fill the response page info into div
function registeDynamicPageLoadLink(linkId, pageUrl, targetDivId, paraMap, isFullUrl) {
    //console.log("registeDynamicPageLoadLink, linkId:"+linkId)
    //input para check
    if(typeof(linkId) == "undefined" || linkId == "") {
        return
    }
    if(typeof(pageUrl) == "undefined" || pageUrl == "") {
        return
    }
    if(typeof(targetDivId) == "undefined" || targetDivId == "") {
        return
    }

    //setup link id click call back
    $("#"+linkId).on('click', function() {
        //console.log("registeDynamicPageLoadLink, linkId:"+ linkId +", clicked...")
        //change browser url
        changeBrowserUrl(pageUrl);

        //setup final page url
        var finalPageUrl = pageUrl + "?page=true";
        if(typeof(isFullUrl) != "undefined" && isFullUrl == true) {
            finalPageUrl = pageUrl + "&page=true";
        }

        //send ajax page load
        sendAjaxPageReq(finalPageUrl, targetDivId, paraMap);
    })
}

//direct jump to dynamic ajax page
//fill the response page info into div
function jumpToDynamicPageLink(pageUrl, targetDivId, isFullUrl, skipChangeBrowser, skipStopScroll) {
    //check
    if(typeof(pageUrl) == "undefined" || pageUrl == "") {
        return
    }
    if(typeof(targetDivId) == "undefined" || targetDivId == "") {
        return
    }
    if(typeof(skipChangeBrowser) == "undefined" || skipChangeBrowser == null) {
        skipChangeBrowser = false;
    }

    if(skipChangeBrowser == false) {
        //change browser url
        changeBrowserUrl(pageUrl);
    }

    //setup final page url
    var finalPageUrl = pageUrl + "?page=true";
    if(typeof(isFullUrl) != "undefined" && isFullUrl == true) {
        finalPageUrl = pageUrl + "&page=true";
    }

    //console.log("jumpToDynamicPageLink:"+finalPageUrl)
    if(typeof(skipStopScroll) == "undefined" || skipStopScroll == false) {
        //unbind scoll, this is important!!!
        $(window).unbind("scroll");
    }

    //send ajax page load
    sendAjaxPageReq(finalPageUrl, targetDivId);
}

//dynamic setup browser url
function changeBrowserUrl(url, page) {
    //check
    if(typeof(url) == "undefined" || url == "") {
        return
    }

    //setup url
    if (typeof(history.pushState) != "undefined") {
        var obj = { Page: page, Url: url };
        history.pushState(obj, obj.Page, obj.Url);
    } else {
        console.log("Browser does not support HTML5.");
    }
}


//float tip message
//kind: info, warning, error, success, loading
//base on tips/message.min.js
function floatTipMessage(message, kind, autoClose) {
    //check
    if(typeof(message) == "undefined" || message == "") {
        return
    }
    if(typeof(kind) == "undefined" || kind == "") {
        kind = "info";
    }

    //setup config
    var config = {
        showClose:true,
        html:true,
    };
    if(typeof(autoClose) != "undefined" && autoClose == null) {
        config[autoClose] = autoClose;
        if(autoClose == true) {
            config[timeout] = tipMessageTimeOut;
        }
    }

    //show diff message by kind
    switch (kind) {
        case "info":
            Qmsg.info(message, config);
            break
        case "warning":
            Qmsg.warning(message, config);
            break
        case "error":
            Qmsg.error(message, config);
            break
        case "success":
            Qmsg.success(message, config);
            break
        case "loading":
            Qmsg.loading(message, config);
            break
        default:
            Qmsg.info(message, config);
            break
    }
}


/////////////////////////////////
//gen dyanmic div with close opt
/////////////////////////////////

var dynamicDivPrefix = "dynamicDiv_";
var dynamicDivIdMap = {};
var dynamicDivClosedCBMap = {}; //divId -> closedCB

function resetDynamicClosebleDiv() {
    dynamicDivIdMap = {};
    dynamicDivClosedCBMap = {};
}

function genDynamicClosebleDiv(divId, info, closedCB) {
    //check
    if(typeof(divId) == "undefined" || divId == "") {
        return "";
    }
    if(typeof(info) == "undefined" || info == "") {
        return "";
    }
    var divIdOld = dynamicDivIdMap[divId];
    if(typeof(divIdOld) != "undefined" && divIdOld != false){
        return "";
    }

    //begin gen dynamic div
    var realDivId = dynamicDivPrefix + divId;
    var div = `<div id="`+realDivId+`" class="dynamicDiv">
  <p>`+info+`&nbsp;<a class="close" href="javascript:closeDynamicDiv('`+divId+`');">x</a></p>
  </div>
  `;
    dynamicDivIdMap[divId] = true;

    //check and register closed cb
    if(typeof(closedCB) != "undefined" && closedCB != null){
        dynamicDivClosedCBMap[divId] = closedCB
    }
    return div
}

function closeDynamicDiv(divId) {
    //check
    if(typeof(divId) == "undefined" || divId == "") {
        return
    }
    //get real div id
    var realDivId = dynamicDivPrefix + divId;

    //remove div and running element
    $('#'+realDivId).remove();
    delete dynamicDivIdMap[divId];

    //check and call cb
    var closedCB = dynamicDivClosedCBMap[divId]
    if(typeof(closedCB) != "undefined" && closedCB != null){
        closedCB(divId);
        delete dynamicDivClosedCBMap[divId];
    }
}


//////////////////////////
//send ajax request
//////////////////////////

//general ajax request
function sendAjaxReq(reqUrl, successUrl, data) {
    if(typeof(reqUrl) == "undefined" || reqUrl == "") {
        return
    }
    //send ajax request
    $.ajax({
        type: "Post",
        url: reqUrl,
        data: data,
        async : true,
        dataType : "json",
        success: function(data){
            if(typeof(data) == undefined) {
                floatTipMessage("sendAjaxReq, invalid response data", "error");
                //console.log('sendAjaxReq, reqUrl:' + reqUrl + ', invalid response data');
                return false;
            }
            //get resp of json
            var errCode = data.errCode
            var errMsg = data.errMsg
            if(errCode != errCodeOfSucceed) {
                console.log('sendAjaxReq, reqUrl:' + reqUrl
                    + ', errCode:' + errCode
                    + ', errMsg:' + errMsg);
                //if need login
                if(errCode == errCodeOfNeedLogin) {
                    //jump to login page
                    jumpToLoginPage();
                }
                return false;
            }
            if(typeof(successUrl) != "undefined" && successUrl != null && successUrl != "") {
                //jump to succeed url
                window.location = successUrl;
            }
        },
        error: function(err) {
            floatTipMessage("sendAjaxReq, "+err, "error");
            console.log('sendAjaxReq, reqUrl:' + reqUrl + ', err:' + err);
        }
    });
}

//ajax request with cb func
function sendAjaxReqWithCB(reqUrl, data, cbFunc, cbPara) {
    //input check
    if(typeof(reqUrl) == "undefined" || reqUrl == "") {
        return
    }
    //send ajax request
    $.ajax({
        type: "POST",
        url: reqUrl,
        data: data,
        async : true,
        dataType : "json",
        success: function(data){
            if(typeof(data) == undefined) {
                floatTipMessage("sendAjaxReqWithCB, invalid response data", "error");
                //console.log('sendAjaxReqWithCB, reqUrl:' + reqUrl + ', invalid response data');
                return
            }
            //get resp of json
            var errCode = data.errCode;
            var errMsg = data.errMsg;
            var dataVal = data.val;
            if(errCode != errCodeOfSucceed) {
                // console.log('sendAjaxReqWithCB, reqUrl:' + reqUrl
                //      + ', errCode:' + errCode
                //      + ', errMsg:' + errMsg);
                //if need login
                if(errCode == errCodeOfNeedLogin) {
                    //jump to login page
                    jumpToLoginPage();
                    return
                }
            }
            cbFunc(dataVal, cbPara, errCode, errMsg);
        },
        error: function(err) {
            floatTipMessage("sendAjaxReqWithCB, "+err, "error");
            console.log('sendAjaxReqWithCB, reqUrl:' + reqUrl + ', err:' + err);
        }
    });
}

//ajax page content and fill target div
//cbFunc called after got page response
//if isGoBack, maybe need hit cache for user expression
function sendAjaxPageReq(reqUrl, fillDivId, paraMap, cbFunc, isAppend) {
    //input check
    if(typeof(reqUrl) == "undefined" || reqUrl == "") {
        return
    }
    if(typeof(fillDivId) == "undefined" || fillDivId == "") {
        return
    }
    if(typeof(isAppend) == "undefined" || isAppend == null) {
        isAppend = false;
    }

    //unbind scoll, this is important!!!
    //$(window).unbind("scroll");
    console.log("sendAjaxPageReq, reqUrl:"+reqUrl);

    //send ajax request to fetch page
    $.ajax({
        type: "POST",
        url: reqUrl,
        data: paraMap,
        async : true,
        cache : true,
        dataType: "html",
        beforeSend: function () {
            return true;
        },
        success: function(data){
            if(typeof(data) == undefined) {
                floatTipMessage("sendAjaxPageReq, invalid response data", "error");
                console.log('sendAjaxPageReq, reqUrl:' + reqUrl + ', invalid response data');
                return
            }
            //get resp of page content
            var page = data.toString();

            //console.log('sendAjaxPageReq, reqUrl:'+reqUrl+', fillDivId:'+fillDivId);

            //fill page content to target div
            if(isAppend == true) {
                $("#"+fillDivId).append(page);
            }else{
                $("#"+fillDivId).html(page);
            }

            //check and call cb func
            if(page != "" && typeof(cbFunc) != "undefined" && cbFunc != null) {
                cbFunc();
            }
        },
        error: function(err) {
            floatTipMessage("sendAjaxPageReq, "+err, "error");
            //console.log(''sendAjaxPageReq, reqUrl:' + reqUrl + ', err:' + err);
        }
    });
}
