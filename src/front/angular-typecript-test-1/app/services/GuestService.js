var Application;
(function (Application) {
    var Services;
    (function (Services) {
        var GuestService = (function () {
            function GuestService($http) {
                this.getUser = function () {
                    return null;
                };
                this.httpService = $http;
            }
            GuestService.$inject = ["$http"];
            return GuestService;
        })();
        Services.GuestService = GuestService;
        var HttpHandlerService = (function () {
            function HttpHandlerService($http) {
                this.httpService = $http;
            }
            HttpHandlerService.prototype.GetHanlder = function (params) {
                var _this = this;
                var result = this.httpService.get(this.handlerUrl, params)
                    .then(function (response) { return _this.handlerResponded(response, params); });
                return result;
            };
            HttpHandlerService.prototype.handlerResponded = function (response, params) {
                response.data.requestParams = params;
                return response.data;
            };
            return HttpHandlerService;
        })();
        Services.HttpHandlerService = HttpHandlerService;
        angular.module("Application").service("Application.Services.GuestService", GuestService);
    })(Services = Application.Services || (Application.Services = {}));
})(Application || (Application = {}));
