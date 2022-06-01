## README

### Installation

+ install with `pip` tool.

```bash
# install alibabacloud_fc_open20210406
pip install -U alibabacloud_fc_open20210406
```

### Usage

+ Invoke HTTP Trigger

```python
# -*- coding: utf-8 -*-
import os

from alibabacloud_fc_open20210406.client import Client
from alibabacloud_tea_openapi import models as open_api_models

ak = os.getenv('ak')
sk = os.getenv('sk')

client = Client(config=open_api_models.Config(access_key_id=ak,
                                              access_key_secret=sk,
                                              region_id='cn-hangzhou'))

resp = client.invoke_httptrigger(url="https://xxx.fcapp.run/action?key=value",
                                 method="GET", 
                                 body="anything".encode(encoding='utf-8'),
                                 headers={"k1": "v1", "k2": "v2"})
```

+ Invoke Anonymous HTTP Trigger

```python
import os

from alibabacloud_fc_open20210406.client import Client
from alibabacloud_tea_openapi import models as open_api_models

ak = os.getenv('ak')
sk = os.getenv('sk')

client = Client(config=open_api_models.Config(access_key_id=ak,
                                              access_key_secret=sk,
                                              region_id='cn-hangzhou'))

resp = client.invoke_anonymous_httptrigger(url="https://xxx.fcapp.run/action?key=value",
                                 method="GET", 
                                 body="anything".encode(encoding='utf-8'),
                                 headers={"k1": "v1", "k2": "v2"})


```


+ Integration with your own http_client

```python

import requests
import os

from alibabacloud_fc_open20210406.client import Client
from alibabacloud_tea_openapi import models as open_api_models

ak = os.getenv('ak')
sk = os.getenv('sk')
client = Client(config=open_api_models.Config(access_key_id=ak,
                                              access_key_secret=sk,
                                              region_id='cn-hangzhou'))

# build your own request
req = requests.Request(
    url='https://xxx.fcapp.run/action?key=value',
    method='GET'
)
req = client.sign_request(req)
with requests.Session() as s:
    prep=s.prepare_request(req)
    resp = s.send(prep)

```




