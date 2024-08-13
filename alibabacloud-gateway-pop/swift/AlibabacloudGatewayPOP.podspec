Pod::Spec.new do |spec|

  spec.name         = "AlibabacloudGatewayPOP"
  spec.version      = "0.0.1"
  spec.license      = "Apache 2.0"
  spec.summary      = "Alibaba Cloud Gateway SPI SDK Library for Swift"
  spec.homepage     = "https://github.com/alibabacloud-sdk-swift/alibabacloud-gateway-pop" 
  spec.author       = { "Alibaba Cloud SDK" => "sdk-team@alibabacloud.com" }

  spec.source       = { :git => spec.homepage + '.git', :tag => spec.version }
  spec.source_files = 'Sources/**/*.swift'

  spec.ios.framework   = 'Foundation'

  spec.ios.deployment_target     = '13.0'
  spec.osx.deployment_target     = '10.15'
  spec.watchos.deployment_target = '6.0'
  spec.tvos.deployment_target    = '13.0'

  spec.dependency 'Tea',  '~> 1.0.2'
  spec.dependency 'AlibabacloudGatewaySPI',  '~> 0.0.1'

  spec.swift_version='5.6'
end
