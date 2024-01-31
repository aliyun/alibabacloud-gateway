import Foundation
import Tea
import AlibabaCloudCredentials

public class InterceptorContext : Tea.TeaModel {
    public class Request : Tea.TeaModel {
        public var headers: [String: String]?

        public var query: [String: String]?

        public var body: Any?

        public var stream: InputStream?

        public var hostMap: [String: String]?

        public var pathname: String?

        public var productId: String?

        public var action: String?

        public var version: String?

        public var protocol_: String?

        public var method: String?

        public var authType: String?

        public var bodyType: String?

        public var reqBodyType: String?

        public var style: String?

        public var credential: AlibabaCloudCredentials.Client?

        public var signatureVersion: String?

        public var signatureAlgorithm: String?

        public var userAgent: String?

        public override init() {
            super.init()
        }

        public init(_ dict: [String: Any]) {
            super.init()
            self.fromMap(dict)
        }

        public override func validate() throws -> Void {
            try self.validateRequired(self.pathname, "pathname")
            try self.validateRequired(self.productId, "productId")
            try self.validateRequired(self.action, "action")
            try self.validateRequired(self.version, "version")
            try self.validateRequired(self.protocol_, "protocol_")
            try self.validateRequired(self.method, "method")
            try self.validateRequired(self.authType, "authType")
            try self.validateRequired(self.bodyType, "bodyType")
            try self.validateRequired(self.reqBodyType, "reqBodyType")
            try self.validateRequired(self.credential, "credential")
            try self.validateRequired(self.userAgent, "userAgent")
        }

        public override func toMap() -> [String : Any] {
            var map = super.toMap()
            if self.headers != nil {
                map["headers"] = self.headers!
            }
            if self.query != nil {
                map["query"] = self.query!
            }
            if self.body != nil {
                map["body"] = self.body!
            }
            if self.stream != nil {
                map["stream"] = self.stream!
            }
            if self.hostMap != nil {
                map["hostMap"] = self.hostMap!
            }
            if self.pathname != nil {
                map["pathname"] = self.pathname!
            }
            if self.productId != nil {
                map["productId"] = self.productId!
            }
            if self.action != nil {
                map["action"] = self.action!
            }
            if self.version != nil {
                map["version"] = self.version!
            }
            if self.protocol_ != nil {
                map["protocol"] = self.protocol_!
            }
            if self.method != nil {
                map["method"] = self.method!
            }
            if self.authType != nil {
                map["authType"] = self.authType!
            }
            if self.bodyType != nil {
                map["bodyType"] = self.bodyType!
            }
            if self.reqBodyType != nil {
                map["reqBodyType"] = self.reqBodyType!
            }
            if self.style != nil {
                map["style"] = self.style!
            }
            if self.credential != nil {
                map["credential"] = self.credential!
            }
            if self.signatureVersion != nil {
                map["signatureVersion"] = self.signatureVersion!
            }
            if self.signatureAlgorithm != nil {
                map["signatureAlgorithm"] = self.signatureAlgorithm!
            }
            if self.userAgent != nil {
                map["userAgent"] = self.userAgent!
            }
            return map
        }

        public override func fromMap(_ dict: [String: Any]) -> Void {
            if dict.keys.contains("headers") && dict["headers"] != nil {
                self.headers = dict["headers"] as! [String: String]
            }
            if dict.keys.contains("query") && dict["query"] != nil {
                self.query = dict["query"] as! [String: String]
            }
            if dict.keys.contains("body") && dict["body"] != nil {
                self.body = dict["body"] as! Any
            }
            if dict.keys.contains("stream") && dict["stream"] != nil {
                self.stream = dict["stream"] as! InputStream
            }
            if dict.keys.contains("hostMap") && dict["hostMap"] != nil {
                self.hostMap = dict["hostMap"] as! [String: String]
            }
            if dict.keys.contains("pathname") && dict["pathname"] != nil {
                self.pathname = dict["pathname"] as! String
            }
            if dict.keys.contains("productId") && dict["productId"] != nil {
                self.productId = dict["productId"] as! String
            }
            if dict.keys.contains("action") && dict["action"] != nil {
                self.action = dict["action"] as! String
            }
            if dict.keys.contains("version") && dict["version"] != nil {
                self.version = dict["version"] as! String
            }
            if dict.keys.contains("protocol") && dict["protocol"] != nil {
                self.protocol_ = dict["protocol"] as! String
            }
            if dict.keys.contains("method") && dict["method"] != nil {
                self.method = dict["method"] as! String
            }
            if dict.keys.contains("authType") && dict["authType"] != nil {
                self.authType = dict["authType"] as! String
            }
            if dict.keys.contains("bodyType") && dict["bodyType"] != nil {
                self.bodyType = dict["bodyType"] as! String
            }
            if dict.keys.contains("reqBodyType") && dict["reqBodyType"] != nil {
                self.reqBodyType = dict["reqBodyType"] as! String
            }
            if dict.keys.contains("style") && dict["style"] != nil {
                self.style = dict["style"] as! String
            }
            if dict.keys.contains("credential") && dict["credential"] != nil {
                self.credential = dict["credential"] as! AlibabaCloudCredentials.Client
            }
            if dict.keys.contains("signatureVersion") && dict["signatureVersion"] != nil {
                self.signatureVersion = dict["signatureVersion"] as! String
            }
            if dict.keys.contains("signatureAlgorithm") && dict["signatureAlgorithm"] != nil {
                self.signatureAlgorithm = dict["signatureAlgorithm"] as! String
            }
            if dict.keys.contains("userAgent") && dict["userAgent"] != nil {
                self.userAgent = dict["userAgent"] as! String
            }
        }
    }
    public class Configuration : Tea.TeaModel {
        public var regionId: String?

        public var endpoint: String?

        public var endpointRule: String?

        public var endpointMap: [String: String]?

        public var endpointType: String?

        public var network: String?

        public var suffix: String?

        public override init() {
            super.init()
        }

        public init(_ dict: [String: Any]) {
            super.init()
            self.fromMap(dict)
        }

        public override func validate() throws -> Void {
            try self.validateRequired(self.regionId, "regionId")
        }

        public override func toMap() -> [String : Any] {
            var map = super.toMap()
            if self.regionId != nil {
                map["regionId"] = self.regionId!
            }
            if self.endpoint != nil {
                map["endpoint"] = self.endpoint!
            }
            if self.endpointRule != nil {
                map["endpointRule"] = self.endpointRule!
            }
            if self.endpointMap != nil {
                map["endpointMap"] = self.endpointMap!
            }
            if self.endpointType != nil {
                map["endpointType"] = self.endpointType!
            }
            if self.network != nil {
                map["network"] = self.network!
            }
            if self.suffix != nil {
                map["suffix"] = self.suffix!
            }
            return map
        }

        public override func fromMap(_ dict: [String: Any]) -> Void {
            if dict.keys.contains("regionId") && dict["regionId"] != nil {
                self.regionId = dict["regionId"] as! String
            }
            if dict.keys.contains("endpoint") && dict["endpoint"] != nil {
                self.endpoint = dict["endpoint"] as! String
            }
            if dict.keys.contains("endpointRule") && dict["endpointRule"] != nil {
                self.endpointRule = dict["endpointRule"] as! String
            }
            if dict.keys.contains("endpointMap") && dict["endpointMap"] != nil {
                self.endpointMap = dict["endpointMap"] as! [String: String]
            }
            if dict.keys.contains("endpointType") && dict["endpointType"] != nil {
                self.endpointType = dict["endpointType"] as! String
            }
            if dict.keys.contains("network") && dict["network"] != nil {
                self.network = dict["network"] as! String
            }
            if dict.keys.contains("suffix") && dict["suffix"] != nil {
                self.suffix = dict["suffix"] as! String
            }
        }
    }
    public class Response : Tea.TeaModel {
        public var statusCode: Int?

        public var headers: [String: String]?

        public var body: InputStream?

        public var deserializedBody: Any?

        public override init() {
            super.init()
        }

        public init(_ dict: [String: Any]) {
            super.init()
            self.fromMap(dict)
        }

        public override func validate() throws -> Void {
        }

        public override func toMap() -> [String : Any] {
            var map = super.toMap()
            if self.statusCode != nil {
                map["statusCode"] = self.statusCode!
            }
            if self.headers != nil {
                map["headers"] = self.headers!
            }
            if self.body != nil {
                map["body"] = self.body!
            }
            if self.deserializedBody != nil {
                map["deserializedBody"] = self.deserializedBody!
            }
            return map
        }

        public override func fromMap(_ dict: [String: Any]) -> Void {
            if dict.keys.contains("statusCode") && dict["statusCode"] != nil {
                self.statusCode = dict["statusCode"] as! Int
            }
            if dict.keys.contains("headers") && dict["headers"] != nil {
                self.headers = dict["headers"] as! [String: String]
            }
            if dict.keys.contains("body") && dict["body"] != nil {
                self.body = dict["body"] as! InputStream
            }
            if dict.keys.contains("deserializedBody") && dict["deserializedBody"] != nil {
                self.deserializedBody = dict["deserializedBody"] as! Any
            }
        }
    }
    public var request: InterceptorContext.Request?

    public var configuration: InterceptorContext.Configuration?

    public var response: InterceptorContext.Response?

    public override init() {
        super.init()
    }

    public init(_ dict: [String: Any]) {
        super.init()
        self.fromMap(dict)
    }

    public override func validate() throws -> Void {
        try self.validateRequired(self.request, "request")
        try self.request?.validate()
        try self.validateRequired(self.configuration, "configuration")
        try self.configuration?.validate()
        try self.validateRequired(self.response, "response")
        try self.response?.validate()
    }

    public override func toMap() -> [String : Any] {
        var map = super.toMap()
        if self.request != nil {
            map["request"] = self.request?.toMap()
        }
        if self.configuration != nil {
            map["configuration"] = self.configuration?.toMap()
        }
        if self.response != nil {
            map["response"] = self.response?.toMap()
        }
        return map
    }

    public override func fromMap(_ dict: [String: Any]) -> Void {
        if dict.keys.contains("request") && dict["request"] != nil {
            var model = InterceptorContext.Request()
            model.fromMap(dict["request"] as! [String: Any])
            self.request = model
        }
        if dict.keys.contains("configuration") && dict["configuration"] != nil {
            var model = InterceptorContext.Configuration()
            model.fromMap(dict["configuration"] as! [String: Any])
            self.configuration = model
        }
        if dict.keys.contains("response") && dict["response"] != nil {
            var model = InterceptorContext.Response()
            model.fromMap(dict["response"] as! [String: Any])
            self.response = model
        }
    }
}

public class AttributeMap : Tea.TeaModel {
    public var attributes: [String: Any]?

    public var key: [String: String]?

    public override init() {
        super.init()
    }

    public init(_ dict: [String: Any]) {
        super.init()
        self.fromMap(dict)
    }

    public override func validate() throws -> Void {
        try self.validateRequired(self.attributes, "attributes")
        try self.validateRequired(self.key, "key")
    }

    public override func toMap() -> [String : Any] {
        var map = super.toMap()
        if self.attributes != nil {
            map["attributes"] = self.attributes!
        }
        if self.key != nil {
            map["key"] = self.key!
        }
        return map
    }

    public override func fromMap(_ dict: [String: Any]) -> Void {
        if dict.keys.contains("attributes") && dict["attributes"] != nil {
            self.attributes = dict["attributes"] as! [String: Any]
        }
        if dict.keys.contains("key") && dict["key"] != nil {
            self.key = dict["key"] as! [String: String]
        }
    }
}
