/*
* convert video to gif face
*/

//global variables
var video2gifReqUrl = "/video2gif/";
var video2gifApiUrl = "/api/video2gif";
var video2gifApiUpload = video2gifApiUrl + "/upload";


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
  if(typeof(startTime) == "undefined" || startTime == "") {
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

        var shortUrl = val.toString();
        var detailViewUrl = video2gifReqUrl + shortUrl ;

        //jump to view page
        pageJump(detailViewUrl);
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
