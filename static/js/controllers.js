var phonecatControllers = angular.module('phonecatControllers', []);

phonecatControllers.controller('PhoneListCtrl', ['$scope', '$http',
  function($scope, $http){
      $http.get('/static/js/phones.json').success(function(data){
          $scope.phones = data;
      });

      $scope.orderProp = 'age';
  }]);

phonecatControllers.controller('PhoneDetailCtrl', ['$scope', '$routeParams', '$http',
  function($scope, $routeParams, $http){
      $http.get('/static/js/json/' + $routeParams.phoneId + '.json').success(function(data){
          console.log('/static/js/json/' + $routeParams.phoneId + '.json')
          $scope.phone = data;
          $scope.mainImageUrl = data.images[0];
      });

      $scope.setImage = function(imageUrl){
          $scope.mainImageUrl = imageUrl;
      };
  }]);
