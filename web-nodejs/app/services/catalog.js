'use strict';

angular.module("app")

.factory('catalog', ['$http', '$q', 'COOLSTORE_CONFIG', 'Auth', '$location', function($http, $q, COOLSTORE_CONFIG, $auth, $location) {
	var factory = {}, products, baseUrl;

	if ($location.protocol() === 'https') {
		baseUrl = (COOLSTORE_CONFIG.SECURE_API_ENDPOINT.startsWith("https://") ? COOLSTORE_CONFIG.SECURE_API_ENDPOINT : "https://" + $location.host().replace(COOLSTORE_CONFIG.MY_SERVICE, COOLSTORE_CONFIG.SECURE_API_ENDPOINT)) + '/api/products';
	} else {
		baseUrl = (COOLSTORE_CONFIG.API_ENDPOINT.startsWith("http://") ? COOLSTORE_CONFIG.API_ENDPOINT : "http://" + $location.host().replace(COOLSTORE_CONFIG.MY_SERVICE, COOLSTORE_CONFIG.API_ENDPOINT)) + '/api/products';
	}


//    console.log("baseUrl: " + baseUrl);

	factory.getProducts = function() {

        function get(fsuccess, ferror)
        {
            $http(
                {
                method: 'GET',
                url: baseUrl
                }
                ).then(fsuccess, 
                    function(err) 
                    {
                    // force a retry once, mainly due to Chrome timeouts on web requests
                    $http({
                        method: 'GET',
                        url: baseUrl
                    }).then(fsuccess, function(err) 
                    {
                        ferror;
                    });
            });
        }

		var deferred = $q.defer();
        if (products) {
            deferred.resolve(products);
        } 
        else 
        {
            get(function(resp) 
                    {
                    products = resp.data;
                    deferred.resolve(resp.data);
                    }, 
                function(err) {
                    deferred.reject(err);
                });
        }
	   return deferred.promise;
	};

	return factory;
}]);
