# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class DataLakeStorageTransferJobHistory(DaraModel):
    def __init__(
        self,
        detail_info: main_models.DataLakeStorageTransferJobHistoryDetailInfo = None,
        end_time: int = None,
        id: str = None,
        job_id: str = None,
        start_time: int = None,
        status: str = None,
        succeed_count: int = None,
        total_count: int = None,
    ):
        self.detail_info = detail_info
        self.end_time = end_time
        self.id = id
        self.job_id = job_id
        self.start_time = start_time
        self.status = status
        self.succeed_count = succeed_count
        self.total_count = total_count

    def validate(self):
        if self.detail_info:
            self.detail_info.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.detail_info is not None:
            result['DetailInfo'] = self.detail_info.to_map()

        if self.end_time is not None:
            result['EndTime'] = self.end_time

        if self.id is not None:
            result['Id'] = self.id

        if self.job_id is not None:
            result['JobId'] = self.job_id

        if self.start_time is not None:
            result['StartTime'] = self.start_time

        if self.status is not None:
            result['Status'] = self.status

        if self.succeed_count is not None:
            result['SucceedCount'] = self.succeed_count

        if self.total_count is not None:
            result['TotalCount'] = self.total_count

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DetailInfo') is not None:
            temp_model = main_models.DataLakeStorageTransferJobHistoryDetailInfo()
            self.detail_info = temp_model.from_map(m.get('DetailInfo'))

        if m.get('EndTime') is not None:
            self.end_time = m.get('EndTime')

        if m.get('Id') is not None:
            self.id = m.get('Id')

        if m.get('JobId') is not None:
            self.job_id = m.get('JobId')

        if m.get('StartTime') is not None:
            self.start_time = m.get('StartTime')

        if m.get('Status') is not None:
            self.status = m.get('Status')

        if m.get('SucceedCount') is not None:
            self.succeed_count = m.get('SucceedCount')

        if m.get('TotalCount') is not None:
            self.total_count = m.get('TotalCount')

        return self

class DataLakeStorageTransferJobHistoryDetailInfo(DaraModel):
    def __init__(
        self,
        error_msg: str = None,
        hdfsfailed_count: int = None,
        hdfstransfer_data_dir: str = None,
        hdfstransfer_err_info_dir: str = None,
        hdfstransfer_import_meta_dir: str = None,
        hdfstransfer_job_id: str = None,
        log_base_dir: str = None,
        verify_err_info_dir: str = None,
        verify_status: str = None,
        verify_total_count: int = None,
    ):
        self.error_msg = error_msg
        self.hdfsfailed_count = hdfsfailed_count
        self.hdfstransfer_data_dir = hdfstransfer_data_dir
        self.hdfstransfer_err_info_dir = hdfstransfer_err_info_dir
        self.hdfstransfer_import_meta_dir = hdfstransfer_import_meta_dir
        self.hdfstransfer_job_id = hdfstransfer_job_id
        self.log_base_dir = log_base_dir
        self.verify_err_info_dir = verify_err_info_dir
        self.verify_status = verify_status
        self.verify_total_count = verify_total_count

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.error_msg is not None:
            result['ErrorMsg'] = self.error_msg

        if self.hdfsfailed_count is not None:
            result['HDFSFailedCount'] = self.hdfsfailed_count

        if self.hdfstransfer_data_dir is not None:
            result['HDFSTransferDataDir'] = self.hdfstransfer_data_dir

        if self.hdfstransfer_err_info_dir is not None:
            result['HDFSTransferErrInfoDir'] = self.hdfstransfer_err_info_dir

        if self.hdfstransfer_import_meta_dir is not None:
            result['HDFSTransferImportMetaDir'] = self.hdfstransfer_import_meta_dir

        if self.hdfstransfer_job_id is not None:
            result['HDFSTransferJobId'] = self.hdfstransfer_job_id

        if self.log_base_dir is not None:
            result['LogBaseDir'] = self.log_base_dir

        if self.verify_err_info_dir is not None:
            result['VerifyErrInfoDir'] = self.verify_err_info_dir

        if self.verify_status is not None:
            result['VerifyStatus'] = self.verify_status

        if self.verify_total_count is not None:
            result['VerifyTotalCount'] = self.verify_total_count

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ErrorMsg') is not None:
            self.error_msg = m.get('ErrorMsg')

        if m.get('HDFSFailedCount') is not None:
            self.hdfsfailed_count = m.get('HDFSFailedCount')

        if m.get('HDFSTransferDataDir') is not None:
            self.hdfstransfer_data_dir = m.get('HDFSTransferDataDir')

        if m.get('HDFSTransferErrInfoDir') is not None:
            self.hdfstransfer_err_info_dir = m.get('HDFSTransferErrInfoDir')

        if m.get('HDFSTransferImportMetaDir') is not None:
            self.hdfstransfer_import_meta_dir = m.get('HDFSTransferImportMetaDir')

        if m.get('HDFSTransferJobId') is not None:
            self.hdfstransfer_job_id = m.get('HDFSTransferJobId')

        if m.get('LogBaseDir') is not None:
            self.log_base_dir = m.get('LogBaseDir')

        if m.get('VerifyErrInfoDir') is not None:
            self.verify_err_info_dir = m.get('VerifyErrInfoDir')

        if m.get('VerifyStatus') is not None:
            self.verify_status = m.get('VerifyStatus')

        if m.get('VerifyTotalCount') is not None:
            self.verify_total_count = m.get('VerifyTotalCount')

        return self

