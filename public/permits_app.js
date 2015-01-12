var app = angular.module('permitsApp', []);

app.controller('permitsCtrl', ['$http', function($http) {
  var vm = this;

  activate();

  function activate() { // TODO: move me to a service
    $http.get("permits").
      success(function(data) {
        vm.permits = data;
      }).
      error(function() {
        console.log("Request broke");
      });
  }
}]);
