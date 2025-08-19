//pve websocket functions
var wsAddr = "localhost:5500/ws";
var ws 	  = null; //websocket object

//core variables
var $messageContainer;

//add message into container
function addMessage(sender, content, type) {
	const time = new Date().toLocaleTimeString();
	const messageClass = type === 'self' ? 'message-self' : 'message-other';
	const html = `<div class="message ${messageClass}">
                        <div><strong>${sender}</strong> [${time}]</div>
                        <div>${content}</div>
                    </div>`;
	$messageContainer.append(html);
    //scroll to the div bottom
	$messageContainer.scrollTop($messageContainer[0].scrollHeight);
}

//player login the battle server
function playerLogin() {
	//check
	if(typeof(ws) == "undefined" || ws == null) {
		return
	}

	//disable soldier and skill button
	$('#soldier-btn').prop("disabled", true);
	$('#skill-btn').prop("disabled", true);

    //tip
	addMessage('系统', '玩家登录服务器..', 'system');

    //create login json object
	var loginObj = new Object();
	loginObj.playerId = parseInt(playerId);

    //gen opt json object
	var genOpt = new Object();
	genOpt.kind = "login"
	genOpt.jsonObj = loginObj

    //enc json object
	var jsonStr = JSON.stringify(genOpt);

    //send login object to server
	ws.send(jsonStr);
}

//player input command to the battle server
//magicType -> magicTypeOfSolider, magicTypeOfSkill
function playerInputCommand(magicType) {
	//check
	if(typeof(ws) == "undefined" || ws == null) {
		return
	}
	if(typeof(magicType) == "undefined" || magicType < 0) {
		return
	}

	//disable soldier button
	//$('#soldier-btn').prop("disabled", true);

	//create input command json object
	var commandObj = new Object();
	commandObj.magicType = magicType;

	switch(magicType) {
	case magicTypeOfSolider:
		{
			//disable soldier button
			$('#soldier-btn').prop("disabled", true);
			commandObj.targetType = "SwordMan";
			break
		}
	case magicTypeOfSkill:
		{
			//disable skill button
			$('#skill-btn').prop("disabled", true);
			commandObj.targetType = "FireBall";
		}
	}

    //gen opt json object
	var genOpt = new Object();
	genOpt.kind = "command";
	genOpt.jsonObj = commandObj;

    //enc json object
	var jsonStr = JSON.stringify(genOpt);
	//console.log("playerInputCommand, jsonStr:"+jsonStr);

    //send login object to server
	ws.send(jsonStr);
}

//received server message
function recvServerMessage(dataObj) {
    //check
	if(typeof(dataObj) == "undefined" || dataObj == null) {
		return;
	}

    //get opt info
	var kind = dataObj.kind
	var jsonObj = dataObj.jsonObj;
	var error = dataObj.error;

    //check error
	if(typeof(error) != "undefined" && error != "") {
		addMessage('系统', '错误发生:' + error, 'system');
		return;
	}

    //do diff opt for opt kind
	switch(kind) {
	case "login":
		{
			//player login
			addMessage('系统', '玩家登录成功', 'system');
			$('#join-btn').prop("disabled", true);
			$('#soldier-btn').prop("disabled", false);
			$('#skill-btn').prop("disabled", false);
			break;
		}
	case "gameStart":
		{
			//game start notify
			addMessage('系统', '游戏开始..', 'system');

			//setup opt button
			$('#create-btn').prop("disabled", true);
			$('#join-btn').prop("disabled", true);
			$('#soldier-btn').prop("disabled", false);
			$('#skill-btn').prop("disabled", false);
			break;
		}
	case "gameWinner":
		{
			//game winner
			var playerId = parseInt(jsonObj);
			addMessage('系统', '胜者:'+playerId, 'system');
			break;
		}	
	case "gameEnd":
		{
			//game end notify
			addMessage('系统', '游戏结束..', 'system');
			$('#create-btn').prop("disabled", false);
			$('#join-btn').prop("disabled", true);
			$('#soldier-btn').prop("disabled", true);
			$('#skill-btn').prop("disabled", true);
			break;
		}
	case "magicActive":
		{
			//player magic active notify
			var magicObj = JSON.stringify(jsonObj);
			var magicKind = jsonObj.kind;
			var magicCount = jsonObj.count;
			//console.log("magicActive, magicObj:"+magicObj+", kind:"+magicKind+", count:"+magicCount);

			switch(magicKind) {
			case magicTypeOfSolider:
				{
					//soldier
					$('#soldier-btn').prop("disabled", false);
					break
				}
			case magicTypeOfSkill:
				{
					//skill
					$('#skill-btn').prop("disabled", false);
					break
				}	
			}

			//addMessage('玩家', '玩家魔法通知..'+magicObj, 'self');
			break;
		}	
	case "command":
		{   
			//player input command
			//include soldier and skill
			var comandObj = JSON.stringify(jsonObj);
			console.log("comandObj:"+comandObj);

			//get core data
			var playerId = jsonObj.playerId;
			var path = jsonObj.path;
			var magicType = jsonObj.magicType;
			var targetType = jsonObj.targetType;
			var startPos = jsonObj.startPos;

			console.log("command, playerId:"+playerId);
			var tipInfo = 'magicType:'+magicType+', targetType:'+targetType;

			if(playerId == pvePlayerId) {
				//pve player
				addMessage('对手', '对手指令, '+tipInfo, 'message-other');
			}else{
				//current player
				addMessage('玩家', '玩家指令, '+tipInfo, 'self');
			}
			break;
		}
	case "frame":
		{
			//frame data broadcast
			//kind include `PutSoldier, PutSkill, Snap`
			var frameObj = JSON.stringify(jsonObj);

			//get core data
			var kind = jsonObj.kind;
			var frameData = jsonObj.frameData;
			//addMessage('系统', '定时帧数据'+frameObj, 'system');
			break
		}	
	}

    //enable send button
	$('#send-btn').prop("disabled", false);
}

//connect game server
function connectGameServer(gameName, groupId) {
	//check
	if(typeof(gameName) == "undefined" || gameName == "") {
		return
	}

	//setup real ws address
	const wsUrl = 'ws:' + wsAddr + '/' + gameName + '/1';

	//connect websocket obj
    ws = new WebSocket(wsUrl);

    //connected
    ws.onopen = function() {
    	addMessage('系统', '已连接到服务器'+gameName, 'system');

	    //delay login opt for force order!!!
    	//delayRun(playerLogin, delayLoadMSeconds);
    };

	//received message
    ws.onmessage = function(event) {
    	//check
    	if(typeof(event) == "undefined" || event == null) {
    		return
    	}

    	addMessage('系统', '接受到数据', 'system');

    	var dataObj = JSON.parse(event.data);
	    //console.log("ws data:" + dataObj);

	    //cb for recv server message
    	//recvServerMessage(dataObj);
    };

	//error process
    ws.onerror = function(error) {
    	addMessage('系统', '连接发生错误: ' + error.message, 'system');
    };

	//connect closed
    ws.onclose = function() {
    	addMessage('系统', '连接已关闭', 'system');
    };
}