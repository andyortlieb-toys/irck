showChatIdentity = function(name){
    $(".irck-chat-identity").hide();
    $("#irck-chat-identity-"+name).show();
  }

showChatStream = function(idid, streamid){
    console.log("#irck-chat-stream-"+idid+"-"+streamid)
    $(".irck-chat-stream").hide();
    $("#irck-chat-stream-"+idid+"-"+streamid).show();
  }

function main_controller($scope) {
  $scope.usernameText = '';

  $scope.auth = {
    username: null
  };
  $scope.HistoryIdx = 0;
  $scope.identities = [];
  $scope.interruptListen = false;

  function getIdentity(iidx){
    if (!$scope.identities.lookup) { 
      $scope.identities.lookup = {};
    }
    
    // First try cache
    if ($scope.identities.lookup[iidx])
        return $scope.identities.lookup[iidx]
    
    for (var i=0; i<$scope.identities.length; ++i){
      if ($scope.identities[i].IdentityIdx == iidx){
        // Cache it
        $scope.identities.lookup[iidx] = $scope.identities[i];
        // Return it
        return $scope.identities.lookup[iidx]
      }
    }
  }

  function genId(){
    return ++genId.next;
  }
  genId.next = 0;

  function processMessage(msg, forceApply){
    if (forceApply){
      return $scope.$apply( function(){ processMessage(msg); } );
    }

    var streamname;
    var identity = getIdentity(msg.IdentityIdx);

    if (!identity){
      return console.error(" No identity found for this message: ", msg);
    }

    identity.streams = identity.streams || {};

    //Is it a channel or a privmsg?
    if (msg.Recipient == identity.Nick){
      streamname = msg.Originator;
    } else {
      streamname = msg.Recipient;
    }

    identity.streams[streamname] = identity.streams[streamname] || { messages: [], streamid: genId() }
    identity.streams[streamname].messages.push(msg)


  }

  function getHistory(historyIdx, startListening){
    var HistoryIdx = HistoryIdx || 0;

   $.post(
      '/history/', 
      JSON.stringify({
        Auth:{
          Username: $scope.auth.username,
        }
      }),
      function(data){
        $scope.auth.authenticated = true;
        $scope.$apply(function(){
          $scope.identities = data.Identities;
          $scope.HistoryIdx = data.HistoryIdx;

          // Find the history of messages.
          for (var i=0;i<$scope.identities.length;++i){
            if (!$scope.identities[i].History){ continue; }
            for (var h=0; h<$scope.identities[i].History.length;++h){
              processMessage($scope.identities[i].History[h], false);
            }
          }

          console.log("History:", data)
          if (startListening) { startListen(); }

        });
      },
      'json'
    );
  }

  function startListen(recurrence){
    if (!recurrence) { $scope.interruptListen = false; }
    if (!$scope.interruptListen){

    if ($scope.watcher){ $scope.watcher.abort(); }
    
    $scope.watcher = $.post(
      '/watch/all/', 
      JSON.stringify({
        Auth:{
          Username: $scope.auth.username
        }
      }),
      function(data){
        $scope.$apply(function(){
          // Process incoming messages
          for (var i=0; i<data.length; ++i){
            processMessage(data[i], false);

          }

          console.log("Got new messages: ", data)
          startListen(true);
        });
      },
      'json'
    );

    }
  }

  $scope.login = function() {
    console.log("bink")
    $scope.auth.username = $scope.usernameText
    console.log($scope.auth.username, dingo2=$scope)
    getHistory(-1, true);
  };

  $scope.logout = function() {
    console.log("plink")
    $scope.watcher.abort();
    $scope.auth = {};
  };  

  $scope.msg = function(identity, streamname){
    console.log("dink!", arguments)

    $.post(
      '/msg/',
      JSON.stringify({
        Auth: {
          Username: $scope.auth.username
        },
        "Message":{
          "Servername":identity.Servername,
          "Nick":identity.Nick,
          "Recipient":streamname,
          "Message":identity.textinput
        }
      })
    );

    identity.textinput = ''


  }
  console.log(dingo=$scope);
}