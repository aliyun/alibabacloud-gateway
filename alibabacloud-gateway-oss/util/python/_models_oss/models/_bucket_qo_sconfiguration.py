# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class BucketQoSConfiguration(DaraModel):
    def __init__(
        self,
        exclusive: bool = None,
        extranet_download_bandwidth: int = None,
        extranet_qps: int = None,
        extranet_upload_bandwidth: int = None,
        intranet_download_bandwidth: int = None,
        intranet_qps: int = None,
        intranet_upload_bandwidth: int = None,
        total_download_bandwidth: int = None,
        total_qps: int = None,
        total_upload_bandwidth: int = None,
    ):
        self.exclusive = exclusive
        self.extranet_download_bandwidth = extranet_download_bandwidth
        self.extranet_qps = extranet_qps
        self.extranet_upload_bandwidth = extranet_upload_bandwidth
        self.intranet_download_bandwidth = intranet_download_bandwidth
        self.intranet_qps = intranet_qps
        self.intranet_upload_bandwidth = intranet_upload_bandwidth
        self.total_download_bandwidth = total_download_bandwidth
        self.total_qps = total_qps
        self.total_upload_bandwidth = total_upload_bandwidth

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.exclusive is not None:
            result['Exclusive'] = self.exclusive

        if self.extranet_download_bandwidth is not None:
            result['ExtranetDownloadBandwidth'] = self.extranet_download_bandwidth

        if self.extranet_qps is not None:
            result['ExtranetQps'] = self.extranet_qps

        if self.extranet_upload_bandwidth is not None:
            result['ExtranetUploadBandwidth'] = self.extranet_upload_bandwidth

        if self.intranet_download_bandwidth is not None:
            result['IntranetDownloadBandwidth'] = self.intranet_download_bandwidth

        if self.intranet_qps is not None:
            result['IntranetQps'] = self.intranet_qps

        if self.intranet_upload_bandwidth is not None:
            result['IntranetUploadBandwidth'] = self.intranet_upload_bandwidth

        if self.total_download_bandwidth is not None:
            result['TotalDownloadBandwidth'] = self.total_download_bandwidth

        if self.total_qps is not None:
            result['TotalQps'] = self.total_qps

        if self.total_upload_bandwidth is not None:
            result['TotalUploadBandwidth'] = self.total_upload_bandwidth

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Exclusive') is not None:
            self.exclusive = m.get('Exclusive')

        if m.get('ExtranetDownloadBandwidth') is not None:
            self.extranet_download_bandwidth = m.get('ExtranetDownloadBandwidth')

        if m.get('ExtranetQps') is not None:
            self.extranet_qps = m.get('ExtranetQps')

        if m.get('ExtranetUploadBandwidth') is not None:
            self.extranet_upload_bandwidth = m.get('ExtranetUploadBandwidth')

        if m.get('IntranetDownloadBandwidth') is not None:
            self.intranet_download_bandwidth = m.get('IntranetDownloadBandwidth')

        if m.get('IntranetQps') is not None:
            self.intranet_qps = m.get('IntranetQps')

        if m.get('IntranetUploadBandwidth') is not None:
            self.intranet_upload_bandwidth = m.get('IntranetUploadBandwidth')

        if m.get('TotalDownloadBandwidth') is not None:
            self.total_download_bandwidth = m.get('TotalDownloadBandwidth')

        if m.get('TotalQps') is not None:
            self.total_qps = m.get('TotalQps')

        if m.get('TotalUploadBandwidth') is not None:
            self.total_upload_bandwidth = m.get('TotalUploadBandwidth')

        return self

