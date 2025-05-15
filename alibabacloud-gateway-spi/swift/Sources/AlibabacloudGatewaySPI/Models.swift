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

        public override func fromMap(_ dict: [String: Any?]?) -> Void {
            guard let dict else { return }
            if let value = dict["headers"] as? [String: String] {
                self.headers = value
            }
            if let value = dict["query"] as? [String: String] {
                self.query = value
            }
            if let value = dict["body"] as? Any {
                self.body = value
            }
            if let value = dict["stream"] as? InputStream {
                self.stream = value
            }
            if let value = dict["hostMap"] as? [String: String] {
                self.hostMap = value
            }
            if let value = dict["pathname"] as? String {
                self.pathname = value
            }
            if let value = dict["productId"] as? String {
                self.productId = value
            }
            if let value = dict["action"] as? String {
                self.action = value
            }
            if let value = dict["version"] as? String {
                self.version = value
            }
            if let value = dict["protocol"] as? String {
                self.protocol_ = value
            }
            if let value = dict["method"] as? String {
                self.method = value
            }
            if let value = dict["authType"] as? String {
                self.authType = value
            }
            if let value = dict["bodyType"] as? String {
                self.bodyType = value
            }
            if let value = dict["reqBodyType"] as? String {
                self.reqBodyType = value
            }
            if let value = dict["style"] as? String {
                self.style = value
            }
            if let value = dict["credential"] as? AlibabaCloudCredentials.Client {
                self.credential = value
            }
            if let value = dict["signatureVersion"] as? String {
                self.signatureVersion = value
            }
            if let value = dict["signatureAlgorithm"] as? String {
                self.signatureAlgorithm = value
            }
            if let value = dict["userAgent"] as? String {
                self.userAgent = value
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

        public override func fromMap(_ dict: [String: Any?]?) -> Void {
            guard let dict else { return }
            if let value = dict["regionId"] as? String {
                self.regionId = value
            }
            if let value = dict["endpoint"] as? String {
                self.endpoint = value
            }
            if let value = dict["endpointRule"] as? String {
                self.endpointRule = value
            }
            if let value = dict["endpointMap"] as? [String: String] {
                self.endpointMap = value
            }
            if let value = dict["endpointType"] as? String {
                self.endpointType = value
            }
            if let value = dict["network"] as? String {
                self.network = value
            }
            if let value = dict["suffix"] as? String {
                self.suffix = value
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

        public override func fromMap(_ dict: [String: Any?]?) -> Void {
            guard let dict else { return }
            if let value = dict["statusCode"] as? Int {
                self.statusCode = value
            }
            if let value = dict["headers"] as? [String: String] {
                self.headers = value
            }
            if let value = dict["body"] as? InputStream {
                self.body = value
            }
            if let value = dict["deserializedBody"] as? Any {
                self.deserializedBody = value
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

    public override func fromMap(_ dict: [String: Any?]?) -> Void {
        guard let dict else { return }
        if let value = dict["request"] as? [String: Any?] {
            var model = InterceptorContext.Request()
            model.fromMap(value)
            self.request = model
        }
        if let value = dict["configuration"] as? [String: Any?] {
            var model = InterceptorContext.Configuration()
            model.fromMap(value)
            self.configuration = model
        }
        if let value = dict["response"] as? [String: Any?] {
            var model = InterceptorContext.Response()
            model.fromMap(value)
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

    public override func fromMap(_ dict: [String: Any?]?) -> Void {
        guard let dict else { return }
        if let value = dict["attributes"] as? [String: Any] {
            self.attributes = value
        }
        if let value = dict["key"] as? [String: String] {
            self.key = value
        }
    }
}
