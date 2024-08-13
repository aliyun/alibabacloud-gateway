import Foundation
import Tea
import AlibabacloudGatewaySPI

open class Client : AlibabacloudGatewaySPI.Client {
    public var _sha256: String?

    public var _sm3: String?

    public override init() throws {
        try super.init()
    }

    @available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
    public func modifyConfiguration(_ context: AlibabacloudGatewaySPI.InterceptorContext, _ attributeMap: AlibabacloudGatewaySPI.AttributeMap) async throws -> Void {
    }

    @available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
    public func modifyRequest(_ context: AlibabacloudGatewaySPI.InterceptorContext, _ attributeMap: AlibabacloudGatewaySPI.AttributeMap) async throws -> Void {
        
    }

    @available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
    public func modifyResponse(_ context: AlibabacloudGatewaySPI.InterceptorContext, _ attributeMap: AlibabacloudGatewaySPI.AttributeMap) async throws -> Void {
        
    }
}
