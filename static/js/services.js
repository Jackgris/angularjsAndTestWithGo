'use strict';

/* Services */

var phonecatServices = angular.module('phonecatServices', ['ngResource']);

// '/static/js/phones.json'  '/static/js/json/' + $routeParams.phoneId + '.json'
phonecatServices.factory('Phone', ['$resource', 
  function($resource){
    return $resource('static/js/json/:phoneId.json', {}, {
      query: {method:'GET', params:{phoneId:'phones'}, isArray:true}
    });
}]);
