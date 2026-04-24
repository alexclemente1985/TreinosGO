const constants = {
    api: {
        API: "http://localhost:8080/api/v1",
        CUSTOMERS_API: "/customers",
        ACCOUNTS_API: "/accounts",
        AUTH_API: "/auth"
    },    
    http: {
        POST: "POST",
        GET: "GET",
        PUT: "PUT",
        PATCH: "PATCH",
        DELETE: "DELETE",
        status: {
            OK: 200,
            CREATED: 201,
            BAD_REQUEST: 400,
            UNAUTHORIZED: 401,
            NOT_FOUND: 404,
            INTERNAL_SERVER_ERROR: 500
        }
    },
    headers: {
        CONTENT_TYPE: "Content-Type",
        contentTypes: {
            JSON: "application/json"
        },
        authorization: {
            BEARER_TOKEN: "Bearer %s"
        }
    },
    localStorage: {
        ACCESS_TOKEN: "access_token",
        REFRESH_TOKEN: "refresh_token"
    }
}

export default constants;