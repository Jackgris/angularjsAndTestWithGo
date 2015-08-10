<!doctype html>
<html lang="en" ng-app="phonecatApp">
  <head>
    <meta charset="utf-8">
    <title>My HTML File</title>
    <link rel="stylesheet" href="static/css/bootstrap.css">
    <link rel="stylesheet" href="static/css/app.css">
    <script src="static/js/angular/angular.js"></script>
    <script src="static/js/controllers.js"></script>
  </head>
  <body ng-controller="PhoneListCtrl">
    <div class="container-fluid">
      <div class="row">
        <div class="col-md-2">
          <!--Sidebar content-->

          Search: <input ng-model="query">

        </div>
        <div class="col-md-10">
          <!--Body content-->

          <ul class="phones">
            <li ng-repeat="phone in phones | filter:query">
              {{phone.name}}
              <p>{{phone.snippet}}</p>
            </li>
          </ul>

        </div>
      </div>
    </div>
  </body>
</html>
