const assert = require('assert');
const Client = require('../dist/client').default;
const cases = require('../../testdata/signature-version.json');

// Also runnable directly with node after npm run build.
async function testSignatureVersion() {
  const client = new Client();
  for (const test of cases) {
    const context = {
      request: { signatureVersion: test.version },
      configuration: { regionId: test.region, endpoint: test.endpoint }
    };
    assert.strictEqual(await client.getSignatureVersion(context), test.want, test.name);
    assert.strictEqual(context.configuration.regionId, test.resolved, test.name);
    assert.strictEqual(context.request.signatureVersion, test.version, test.name);
  }
}

testSignatureVersion().then(() => console.log(`${cases.length} signature version cases passed`)).catch(error => {
  console.error(error);
  process.exitCode = 1;
});
