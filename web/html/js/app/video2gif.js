/*
* convert video to gif face
*/

//global variables
var video2gifReqUrl = "/video2gif";
var video2gifListUrl = video2gifReqUrl + "/list";

var video2gifApiUrl = "/api/video2gif";
var video2gifApiUpload = video2gifApiUrl + "/upload";
var video2gifApiDownload = video2gifApiUrl + "/download";
var video2gifApiDelete = video2gifApiUrl + "/delete";

var viedo2gifMaxSeconds = 10;
var video2gifListPage = 1;
var video2gifListMoreDiv = true;


//show video2gif upload form page
function showVideo2gifUploadForm(paraMap) {
    if (!paraMap) return;

    var videoElement = paraMap["videoElement"];
    var duration = paraMap["duration"];
    var start = paraMap["start"] || 0;
    var end = paraMap["end"] || 0;
    var maxSeconds = paraMap["maxSeconds"] || 10;

    let isPlaying = false;

    // 上传并预览视频
    $("#upload").on("change", function(e) {
        const file = e.target.files[0];
        if (!file) return;

        const preview = videoElement;
        const url = URL.createObjectURL(file);
        preview.src = url;
        preview.load();

        // 避免 iOS 全屏
        preview.setAttribute('playsinline', '');
        preview.setAttribute('webkit-playsinline', '');

        preview.onloadedmetadata = function() {
            duration = preview.duration;
            start = 0;
            end = Math.floor(duration);

            $("#duration").text("视频时长: " + duration.toFixed(2) + " 秒");
            $("#startSlider, #endSlider").prop("disabled", false)
                .attr("max", Math.floor(duration));
            $("#startSlider").val(0);
            $("#endSlider").val(Math.floor(duration));
            $("#startTime").text(0);
            $("#endTime").text(Math.floor(duration));
            $("#submitBtn").prop("disabled", false);
        };
    });

    // 播放状态监听
    videoElement.addEventListener("play", () => {
        isPlaying = true;
        videoElement.currentTime = start; // 播放前跳到选区起点
    });
    videoElement.addEventListener("pause", () => { isPlaying = false; });

    // 滑块拖动：只更新显示，不修改 video.currentTime
    $("#startSlider").on("input", function() {
        start = parseInt($(this).val());
        if (start + maxSeconds <= duration) end = start + maxSeconds;
        else end = duration;
        if (end - start > maxSeconds) start = end - maxSeconds;

        $("#startTime").text(start);
        $("#endTime").text(end);
        $("#startSlider").val(start);
        $("#endSlider").val(end);
    });

    $("#endSlider").on("input", function() {
        end = parseInt($(this).val());
        if (end - maxSeconds <= start) start = Math.max(0, end - maxSeconds);
        if (end - start > maxSeconds) start = end - maxSeconds;

        $("#startTime").text(start);
        $("#endTime").text(end);
        $("#startSlider").val(start);
        $("#endSlider").val(end);
    });

    // 滑块松开时同步 video.currentTime
    $("#startSlider, #endSlider").on("change", function() {
        videoElement.currentTime = start;
    });

    // 播放控制：只播放选定片段
    videoElement.addEventListener("timeupdate", function() {
        if (!isPlaying) return;
        if (videoElement.currentTime > end - 0.05) { // 0.05s 缓冲
            videoElement.pause();
        }
    });

    // 点击提交
    $("#submitBtn").on("click", function(e) {
        e.preventDefault();
        let length = end - start;
        if (length > maxSeconds) end = start + maxSeconds;

        var paraMap = {
            fileId: "upload",
            startTime: start,
            submitBtn: "submitBtn"
        };

        video2gifUploadVideo(paraMap);
    });
}

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

      // 根据按钮执行操作
      switch(opt){
      case 'download':
        {
            //download gif file
            downloadFile(`/file/video2gif/`+shortUrl+`?uri=`+gifUrl+`&download=true`);
            break;
        }
    case 'delete':
        {
            //delete gif file
            deleteVideo2Gif(shortUrl);
            break;
        }
      }
    });
}

//delete gif file
function deleteVideo2Gif(shortUrl) {
    //check
    if(typeof(shortUrl) == "undefined" || shortUrl == "") {
        return;
    }

    //set ext para for cb func
    var extCbPara = {};
    extCbPara["shortUrl"] = shortUrl;

    //send ajax request
    var data = {
        'uri':shortUrl,
    }
    sendAjaxReqWithCB(video2gifApiDelete, data, cbForDeleteVideoGif, extCbPara);
}

//cb for delete video gif
function cbForDeleteVideoGif(dataVal, paraMap, errCode, errMsg) {
   if(errCode != errCodeOfSucceed) {
        //message tip
        floatTipMessage("删除失败,"+errMsg, "error");
        return;
    }

    //float tips
    floatTipMessage("删除成功", "success");

    //remove deleted gif file
    var shortUrl = paraMap["shortUrl"];
    var gridDiv = `grid_` + shortUrl;
    $(`#`+gridDiv).remove();

    //trigger macy image flows
    if(typeof(macyInstance) != "undefined" && macyInstance != null) {
        macyInstance.recalculate(true);
    }
}

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
