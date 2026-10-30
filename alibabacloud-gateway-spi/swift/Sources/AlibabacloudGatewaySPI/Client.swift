import Foundation
import Tea
import AlibabaCloudCredentials

open class Client {
    public init() throws {
    }

    @available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
    open func modifyConfiguration(_ context: InterceptorContext, _ attributeMap: AttributeMap) async throws -> Void {
        throw Tea.TeaError("Un-implemented")
    }

    @available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
    open func modifyRequest(_ context: InterceptorContext, _ attributeMap: AttributeMap) async throws -> Void {
        throw Tea.TeaError("Un-implemented")
    }

    @available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
    open func modifyResponse(_ context: InterceptorContext, _ attributeMap: AttributeMap) async throws -> Void {
        throw Tea.TeaError("Un-implemented")
    }
}
