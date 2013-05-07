function main_controller($scope) {
  $scope.usernameText = '';

  $scope.auth = {
    username: null
  };
  $scope.HistoryIdx = 0;
  $scope.identities = [];
  $scope.interruptListen = false;
 
  function getHistory(historyIdx, startListen){
    var HistoryIdx = HistoryIdx || 0;
    $.post(
      '/history/', 
      JSON.stringify({
        Auth:{
          Username: $scope.auth.username,
        }
      }),
      function(data){
        $scope.$apply(function(){
          $scope.identities = data.Identities;
          $scope.HistoryIdx = data.HistoryIdx;

          if (startListen) { startListen(); }

        });
      },
      'json'
    );
  }

  function startListen(recurrence){
    if (!recurrence) { $scope.interruptListen = false; }
    if (!$scope.interruptListen){

    $.post(
      '/watch/all/', 
      JSON.stringify({
        Auth:{
          Username: $scope.auth.username,
        }
      }),
      function(data){
        $scope.$apply(function(){
          $scope.identities = data.Identities;
          $scope.HistoryIdx = data.HistoryIdx;

          if (startListen) { startListen(); }

        });
      },
      'json'
    );

    }
  }

  $scope.showChatIdentity = function(name){
    console.log(name);
    $(".irck-chat-identity").hide();
    $("#irck-chat-identity-"+name).show();
  }

  $scope.login = function() {
    console.log("bink")
    $scope.auth.username = $scope.usernameText
    console.log($scope.auth.username, dingo2=$scope)
    getHistory();
  };

  console.log(dingo=$scope);
}