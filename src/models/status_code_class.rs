/*
 * Svix API
 *
 * Welcome to the Svix API documentation!  Useful links: [Homepage](https://www.svix.com) | [Support email](mailto:support+docs@svix.com) | [Blog](https://www.svix.com/blog/) | [Slack Community](https://www.svix.com/slack/)  # Introduction  This is the reference documentation and schemas for the [Svix webhook service](https://www.svix.com) API. For tutorials and other documentation please refer to [the documentation](https://docs.svix.com).  ## Main concepts  In Svix you have four important entities you will be interacting with:  - `messages`: these are the webhooks being sent. They can have contents and a few other properties. - `application`: this is where `messages` are sent to. Usually you want to create one application for each user on your platform. - `endpoint`: endpoints are the URLs messages will be sent to. Each application can have multiple `endpoints` and each message sent to that application will be sent to all of them (unless they are not subscribed to the sent event type). - `event-type`: event types are identifiers denoting the type of the message being sent. Event types are primarily used to decide which events are sent to which endpoint.   ## Authentication  Get your authentication token (`AUTH_TOKEN`) from the [Svix dashboard](https://dashboard.svix.com) and use it as part of the `Authorization` header as such: `Authorization: Bearer ${AUTH_TOKEN}`.  <SecurityDefinitions />   ## Code samples  The code samples assume you already have the respective libraries installed and you know how to use them. For the latest information on how to do that, please refer to [the documentation](https://docs.svix.com/).   ## Idempotency  Svix supports [idempotency](https://en.wikipedia.org/wiki/Idempotence) for safely retrying requests without accidentally performing the same operation twice. This is useful when an API call is disrupted in transit and you do not receive a response.  To perform an idempotent request, pass the idempotency key in the `Idempotency-Key` header to the request. The idempotency key should be a unique value generated by the client. You can create the key in however way you like, though we suggest using UUID v4, or any other string with enough entropy to avoid collisions.  Svix's idempotency works by saving the resulting status code and body of the first request made for any given idempotency key for any successful request. Subsequent requests with the same key return the same result.  Please note that idempotency is only supported for `POST` requests.   ## Cross-Origin Resource Sharing  This API features Cross-Origin Resource Sharing (CORS) implemented in compliance with [W3C spec](https://www.w3.org/TR/cors/). And that allows cross-domain communication from the browser. All responses have a wildcard same-origin which makes them completely public and accessible to everyone, including any code on any site. 
 *
 * The version of the OpenAPI document: 1.4
 * 
 * Generated by: https://openapi-generator.tech
 */

/// StatusCodeClass : The different classes of HTTP status codes: - CodeNone = 0 - Code1xx = 100 - Code2xx = 200 - Code3xx = 300 - Code4xx = 400 - Code5xx = 500

/// The different classes of HTTP status codes: - CodeNone = 0 - Code1xx = 100 - Code2xx = 200 - Code3xx = 300 - Code4xx = 400 - Code5xx = 500
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum StatusCodeClass {
    #[serde(rename = "0")]
    CodeNone,
    #[serde(rename = "100")]
    Code1xx,
    #[serde(rename = "200")]
    Code2xx,
    #[serde(rename = "300")]
    Code3xx,
    #[serde(rename = "400")]
    Code4xx,
    #[serde(rename = "500")]
    Code5xx,

}

impl ToString for StatusCodeClass {
    fn to_string(&self) -> String {
        match self {
            Self::CodeNone => String::from("0"),
            Self::Code1xx => String::from("100"),
            Self::Code2xx => String::from("200"),
            Self::Code3xx => String::from("300"),
            Self::Code4xx => String::from("400"),
            Self::Code5xx => String::from("500"),
        }
    }
}

impl Default for StatusCodeClass {
    fn default() -> StatusCodeClass {
        Self::CodeNone
    }
}




