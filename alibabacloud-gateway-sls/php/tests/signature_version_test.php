<?php
require __DIR__ . '/../vendor/autoload.php';

use Darabonba\GatewaySls\Client;
use Darabonba\GatewaySpi\Models\InterceptorContext;
use Darabonba\GatewaySpi\Models\InterceptorContext\request;
use Darabonba\GatewaySpi\Models\InterceptorContext\configuration;

$cases = json_decode(file_get_contents(__DIR__ . '/../../testdata/signature-version.json'), true);
$client = new Client();
foreach ($cases as $case) {
    $context = new InterceptorContext([
        'request' => new request(['signatureVersion' => $case['version']]),
        'configuration' => new configuration(['regionId' => $case['region'], 'endpoint' => $case['endpoint']])
    ]);
    $version = $client->getSignatureVersion($context);
    if ($version !== $case['want'] || $context->configuration->regionId !== $case['resolved'] || $context->request->signatureVersion !== $case['version']) {
        throw new \RuntimeException('Failed signature version case: ' . $case['name']);
    }
}
echo count($cases) . " signature version cases passed\n";
