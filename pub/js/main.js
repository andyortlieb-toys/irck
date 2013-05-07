function main_controller($scope) {
  $scope.usernameText = '';

  $scope.auth = {
    username: null
  };
  $scope.HistoryIdx = 0;
 
  function getHistory(HistoryIdx){
    var HistoryIdx = HistoryIdx || 0;
    $.post(
      '/history/', 
      JSON.stringify({
        Auth:{
          Username: $scope.auth.username,
        }
      }),
      function(data){
        console.log(data)
      },
      'json'
    );
  }

  $scope.login = function() {
    console.log("bink")
    $scope.auth.username = $scope.usernameText
    console.log($scope.auth.username, dingo2=$scope)
    getHistory();
  };

  console.log(dingo=$scope);
}