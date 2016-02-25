module Application.Services {
    export class GuestService implements Application.Interfaces.IUserService {
        httpService: ng.IHttpService
        static $inject = ["$http"];
        constructor($http: ng.IHttpService) {
            this.httpService = $http;
        }

        getUser = () => {

            return null;
        }
    }

    export class HttpHandlerService {
        httpService: ng.IHttpService
        handlerUrl: string;

        constructor($http: ng.IHttpService) {
            this.httpService = $http
        }

        public GetHanlder(params: any): ng.IPromise<any> {
            var result: ng.IPromise<any> = this.httpService.get(this.handlerUrl, params)
                .then((response: any): ng.IPromise<any> => this.handlerResponded(response, params));

            return result;
        }

        private handlerResponded(response: any, params: any): any {
            response.data.requestParams = params;
            return response.data;
        }
    }

    angular.module("Application").service("Application.Services.GuestService", GuestService);
}