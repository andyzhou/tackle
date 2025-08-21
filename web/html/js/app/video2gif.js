/*
* convert video to gif face
*/

//global variables
var video2gifReqUrl = "/video2gif";
var video2gifListUrl = video2gifReqUrl + "/list";

var video2gifApiUrl = "/api/video2gif";
var video2gifApiUpload = video2gifApiUrl + "/upload";
var video2gifApiDownload = video2gifApiUrl + "/download";

var video2gifListPage = 1;
var video2gifListMoreDiv = true;


//set loaded video2gif images
function setLoadedVideo2Gif() {
    //image box interactive
    $(".imgBox").on("mouseenter", function(){
        var gif = $(this).data("gif");
        $(this).find("img").attr("src", gif);
        $(this).find(".play-btn").hide();
        // 显示操作面板
        $(this).find('.action-panel').fadeIn(200);
    });

   $(".imgBox").on("mouseleave", function(){
     var staticImg = $(this).data("static");
     $(this).find("img").attr("src", staticImg);
     $(this).find(".play-btn").show();
     // 隐藏操作面板
     $(this).find('.action-panel').fadeOut(200);
    });

   //actiion panel interactive
   $('.action-panel button').click(function(e) {
      e.stopPropagation(); // 阻止事件冒泡到 .imgBox
      var shortUrl = $(this).closest('.action-panel').attr('shortUrl');
      var gifUrl = $(this).closest('.action-panel').attr('gif');
      console.log('shortUrl:', shortUrl);
      console.log('gifUrl:', gifUrl);

      var opt = $(this).attr('opt'); // 获取 opt 属性
      var action = $(this).text(); // 或使用 data-action
      //console.log('点击了按钮：', action);
      console.log('opt', opt);

      // 根据按钮执行操作，例如：
      if(opt == "download") {
        //download gif file
        downloadFile(`/file/video2gif/`+shortUrl+`?uri=`+gifUrl+`&download=true`);
        //window.open(`/file/video2gif/`+shortUrl+`?uri=`+gifUrl+`&download=true`, "_blank");
      }
      // if(action === '操作1') {
      //   alert('执行操作1逻辑');
      // } else if(action === '操作2') {
      //   alert('执行操作2逻辑');
      // }
    });
}

//like gif

//download gif

//load more gif files
function loadMoreVideo2Gif(targetDivId, resetPage) {
    //check para
    if(typeof(targetDivId) == "undefined" || targetDivId == "") {
        return;
    }
    if(typeof(resetPage) != "undefined" && resetPage == true) {
        video2gifListPage = 1;
        video2gifListMoreDiv = true;
    }
    if(video2gifListMoreDiv == false) {
        return;
    }

    //detect page url
    var pageUrl = video2gifListUrl + "?pageNo=" + video2gifListPage;
    var finalPageUrl = formatFinalPage(pageUrl);
    var cbFunc = function() {
        // 动态添加图片后
       if(typeof(macyInstance) != "undefined" && macyInstance != null) {
         //console.log("loadMoreVideo2Gif.cbFunc.recalculate");
         var delayFunc = function() {
            macyInstance.recalculate(true);
            video2gifListPage++;
         }
         delayRun(delayLoadMSeconds, delayFunc);
       }
    }

    //send ajax page load
    //console.log("finalPageUrl:"+finalPageUrl);
    sendAjaxPageReq(finalPageUrl, targetDivId, null, cbFunc, true);
}

//upload origin video
function video2gifUploadVideo(paraMap) {
  //check para
  if(typeof(paraMap) == "undefined" || paraMap == null){
    return
  }

  //get key data
  var fileId = paraMap["fileId"];
  var startTime = paraMap["startTime"];
  var submitBtn = paraMap["submitBtn"];

  //check key data
  if(typeof(fileId) == "undefined" || fileId == "") {
    return
  }

  //disable submit button
  $("#"+submitBtn).prop("disabled", true);

  //format request data para
  var dataPara = {
    fileId:fileId,
    startTime:startTime,
    act:"save",
  }

  //send ajax request
  //need upload file
  $.ajaxFileUpload ({
      type: "Post",
      url: video2gifApiUpload,
      secureuri: false,
      fileElementId: fileId,//file id
      dataType: 'json',
      data: dataPara,
      success: function (data, status)
      {
        $("#"+submitBtn).prop("disabled", false);
        //check data
        if(typeof(data) == undefined || data == null) {
          return
        }

        //get resp of json
        var errCode = data.errCode
        var errMsg = data.errMsg
        var val = data.val; //new short url
        if(errCode != errCodeOfSucceed) {
          floatTipMessage(errMsg, "error");
          return
        }

        //setup url
        //var shortUrl = val.toString();
        //var detailViewUrl = video2gifReqUrl + shortUrl ;

        //jump to home page
        pageJump(video2gifReqUrl);
      },
      error: function (data, status, e)
      {
          //error tips
          floatTipMessage(e, "error");

          //enable submit button
          $("#"+submitBtn).prop("disabled", false);
      }
  });

}
