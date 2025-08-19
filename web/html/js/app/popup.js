
//popup user setup page
function popupUserSetup() {
  //setup delay func for ajax page
  var delayFunc = function() {
    //open the pop div
    $('#popup-setup-div').magnificPopup({
      type: 'inline',
      preloader: false,
      alignTop:true,
      //focus: '#name',
    }).magnificPopup('open');
  }

  //request ajax page
  var targetDivId = "popSetupDiv";
  var finalPageUrl = formatFinalPage("/page/setup");

  //send ajax page
  sendAjaxPageReq(finalPageUrl, targetDivId, null, delayFunc);
}

//gen feedback pop form
function popupFeedback() {
  //init popup div
  var popDiv = `
  <div id="pop-feedback-div" class="mfp-hide white-popup-block">
  <p>
  `+popupFBLang.get("info")+`
  </p>
  <p>
  `+popupFBLang.get("kind")+`: <select name="kind" id="kind">
  <option value="">`+popupFBLang.get("select")+`</option>
  <option value=1>`+popupFBLang.get("question")+`</option>
  <option value=2>`+popupFBLang.get("suggest")+`</option>
  <option value=3>`+popupFBLang.get("feedback")+`</option>
  </select>
  </p>
  <p>
  `+popupFBLang.get("introduce")+`:<br/>
  <textarea name="info" id="info" cols=50 rows=6 placeholder="`+popupFBLang.get("detail")+`"></textarea>
  </p>
  <input type="button" class="postButtonLink" value="`+popupFBLang.get("submit")+`" onclick="javascript:sendFeedback();">
  <input type="reset" class="postButtonLink" value="`+popupFBLang.get("reset")+`">
  </div>`;
  $("#popFeedbackDiv").html(popDiv);

  //setup delay func for ajax page
  var delayFunc = function() {
    //open the pop div
    $('#popup-feedback-div').magnificPopup({
      type: 'inline',
      preloader: false,
      alignTop:true,
      //focus: '#name',
    }).magnificPopup('open');
  }
  delayFunc();
}

//message tip
function messageTip(message, errCode) {
  //check
  var info = "";
  if(typeof(errCode) != "undefined" && errCode != errCodeOfSucceed) {
    //failed
    info = baseLang.get("errCode") + errCode + "," + baseLang.get("tip") + errCodeLang.get(errCode) + " " + message
  }else{
    //succeed
    info = message
  }

  var popDiv = `<div id="pop-msg-div" class="mfp-hide white-popup-block"><p>`+baseLang.get("tip")+`</p><p>`+info+`</p></div>`

 //update div html info
 $("#popMsgDiv").html(popDiv);

  //open the pop div
 $('#popup-message-div').magnificPopup({
    type: 'inline',
    preloader: false,
    alignTop:true,
    //focus: '#name',
  }).magnificPopup('open');
}