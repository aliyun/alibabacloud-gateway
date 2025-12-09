# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.

from Tea.model import TeaModel
from typing import Dict, Union, List, Any, get_type_hints
from xml.etree import ElementTree
from collections import defaultdict
from typing_extensions import get_origin, get_args
from Tea.exceptions import RequiredArgumentException
from alibabacloud_oss20190517 import models as oss_models
from alibabacloud_hcs_mgw20240626 import models as hcs_mgw_models

basic_instance = {}
basic_instance[str] = ''
basic_instance[int] = 0
basic_instance[bool] = False
basic_instance[Dict[str, str]] = {'':''}
basic_instance[Any] = {}

def get_origin_type(tp):
    if get_origin(tp) is Union:
        args = [t for t in get_args(tp) if t is not type(None)]
        return args[0] if args else None
    return get_origin(tp)

def build_instance_from_model(model):
    if model in basic_instance:
        return basic_instance[model]
    params = {}
    type_hints = get_type_hints(model.__init__, globalns=globals(), localns=locals())
    for para_name, param in type_hints.items():
        if para_name == "self":
            continue
        origin_type = get_origin_type(param)

        # for Python >= 3.10
        if origin_type is None:
            params[para_name] = build_instance_from_model(param)
        elif isinstance(origin_type, type) and issubclass(origin_type, list):
            params[para_name] = [build_instance_from_model(get_args(param)[0])]

        # for Python <= 3.9
        elif get_origin(origin_type) is list:
            params[para_name] = [build_instance_from_model(get_args(origin_type)[0])]
        else:
            params[para_name] = build_instance_from_model(origin_type)
    return model(**params)


instanceRegistry = {}
# for oss
main_models = oss_models
instanceRegistry["CompleteMultipartUpload"] = build_instance_from_model(main_models.CompleteMultipartUploadResponseBody)
instanceRegistry["CopyObject"] = build_instance_from_model(main_models.CopyObjectResponseBody)
instanceRegistry["CopyObjects"] = build_instance_from_model(main_models.CopyObjectsResponseBody)
instanceRegistry["CreateAccessPoint"] = build_instance_from_model(main_models.CreateAccessPointResponseBody)
instanceRegistry["CreateAccessPointForObjectProcess"] = build_instance_from_model(main_models.CreateAccessPointForObjectProcessResponseBody)
instanceRegistry["CreateBucketDataRedundancyTransition"] = build_instance_from_model(main_models.CreateBucketDataRedundancyTransitionResponseBody)
instanceRegistry["CreateCnameToken"] = build_instance_from_model(main_models.CreateCnameTokenResponseBody)
instanceRegistry["CreateJob"] = build_instance_from_model(main_models.CreateJobResponseBody)
instanceRegistry["CreateReservedCapacity"] = build_instance_from_model(main_models.CreateReservedCapacityResponseBody)
instanceRegistry["DeleteMultipleObjects"] = build_instance_from_model(main_models.DeleteMultipleObjectsResponseBody)
instanceRegistry["DescribeJob"] = build_instance_from_model(main_models.DescribeJobResponseBody)
instanceRegistry["DescribeRegions"] = build_instance_from_model(main_models.DescribeRegionsResponseBody)
instanceRegistry["DoMetaQuery"] = build_instance_from_model(main_models.DoMetaQueryResponseBody)
instanceRegistry["GetAccessPoint"] = build_instance_from_model(main_models.GetAccessPointResponseBody)
instanceRegistry["GetAccessPointConfigForObjectProcess"] = build_instance_from_model(main_models.GetAccessPointConfigForObjectProcessResponseBody)
instanceRegistry["GetAccessPointForObjectProcess"] = build_instance_from_model(main_models.GetAccessPointForObjectProcessResponseBody)
instanceRegistry["GetAccessPointPublicAccessBlock"] = build_instance_from_model(main_models.GetAccessPointPublicAccessBlockResponseBody)
instanceRegistry["GetAsyncFetchTask"] = build_instance_from_model(main_models.GetAsyncFetchTaskResponseBody)
instanceRegistry["GetBucketAccessMonitor"] = build_instance_from_model(main_models.GetBucketAccessMonitorResponseBody)
instanceRegistry["GetBucketAcl"] = build_instance_from_model(main_models.GetBucketAclResponseBody)
instanceRegistry["GetBucketArchiveDirectRead"] = build_instance_from_model(main_models.GetBucketArchiveDirectReadResponseBody)
instanceRegistry["GetBucketCacheConfiguration"] = build_instance_from_model(main_models.GetBucketCacheConfigurationResponseBody)
instanceRegistry["GetBucketCallbackPolicy"] = build_instance_from_model(main_models.GetBucketCallbackPolicyResponseBody)
instanceRegistry["GetBucketCommonHeader"] = build_instance_from_model(main_models.GetBucketCommonHeaderResponseBody)
instanceRegistry["GetBucketCors"] = build_instance_from_model(main_models.GetBucketCorsResponseBody)
instanceRegistry["GetBucketDataAccelerator"] = build_instance_from_model(main_models.GetBucketDataAcceleratorResponseBody)
instanceRegistry["GetBucketDataRedundancyTransition"] = build_instance_from_model(main_models.GetBucketDataRedundancyTransitionResponseBody)
instanceRegistry["GetBucketEncryption"] = build_instance_from_model(main_models.GetBucketEncryptionResponseBody)
instanceRegistry["GetBucketEventNotification"] = build_instance_from_model(main_models.GetBucketEventNotificationResponseBody)
instanceRegistry["GetBucketHash"] = build_instance_from_model(main_models.GetBucketHashResponseBody)
instanceRegistry["GetBucketHttpsConfig"] = build_instance_from_model(main_models.GetBucketHttpsConfigResponseBody)
instanceRegistry["GetBucketInfo"] = build_instance_from_model(main_models.GetBucketInfoResponseBody)
instanceRegistry["GetBucketInventory"] = build_instance_from_model(main_models.GetBucketInventoryResponseBody)
instanceRegistry["GetBucketLifecycle"] = build_instance_from_model(main_models.GetBucketLifecycleResponseBody)
instanceRegistry["GetBucketLocation"] = build_instance_from_model(main_models.GetBucketLocationResponseBody)
instanceRegistry["GetBucketLogging"] = build_instance_from_model(main_models.GetBucketLoggingResponseBody)
instanceRegistry["GetBucketNotification"] = build_instance_from_model(main_models.GetBucketNotificationResponseBody)
instanceRegistry["GetBucketOverwriteConfig"] = build_instance_from_model(main_models.GetBucketOverwriteConfigResponseBody)
instanceRegistry["GetBucketPolicyStatus"] = build_instance_from_model(main_models.GetBucketPolicyStatusResponseBody)
instanceRegistry["GetBucketPublicAccessBlock"] = build_instance_from_model(main_models.GetBucketPublicAccessBlockResponseBody)
instanceRegistry["GetBucketQoSInfo"] = build_instance_from_model(main_models.GetBucketQoSInfoResponseBody)
instanceRegistry["GetBucketReferer"] = build_instance_from_model(main_models.GetBucketRefererResponseBody)
instanceRegistry["GetBucketReplication"] = build_instance_from_model(main_models.GetBucketReplicationResponseBody)
instanceRegistry["GetBucketReplicationLocation"] = build_instance_from_model(main_models.GetBucketReplicationLocationResponseBody)
instanceRegistry["GetBucketReplicationProgress"] = build_instance_from_model(main_models.GetBucketReplicationProgressResponseBody)
instanceRegistry["GetBucketRequestPayment"] = build_instance_from_model(main_models.GetBucketRequestPaymentResponseBody)
instanceRegistry["GetBucketRequesterQoSInfo"] = build_instance_from_model(main_models.GetBucketRequesterQoSInfoResponseBody)
instanceRegistry["GetBucketResourceGroup"] = build_instance_from_model(main_models.GetBucketResourceGroupResponseBody)
instanceRegistry["GetBucketResponseHeader"] = build_instance_from_model(main_models.GetBucketResponseHeaderResponseBody)
instanceRegistry["GetBucketStat"] = build_instance_from_model(main_models.GetBucketStatResponseBody)
instanceRegistry["GetBucketTags"] = build_instance_from_model(main_models.GetBucketTagsResponseBody)
instanceRegistry["GetBucketTransferAcceleration"] = build_instance_from_model(main_models.GetBucketTransferAccelerationResponseBody)
instanceRegistry["GetBucketVersioning"] = build_instance_from_model(main_models.GetBucketVersioningResponseBody)
instanceRegistry["GetBucketWebsite"] = build_instance_from_model(main_models.GetBucketWebsiteResponseBody)
instanceRegistry["GetBucketWorm"] = build_instance_from_model(main_models.GetBucketWormResponseBody)
instanceRegistry["GetCache"] = build_instance_from_model(main_models.GetCacheResponseBody)
instanceRegistry["GetChannel"] = build_instance_from_model(main_models.GetChannelResponseBody)
instanceRegistry["GetCnameToken"] = build_instance_from_model(main_models.GetCnameTokenResponseBody)
instanceRegistry["GetDataLakeCachePrefetchJob"] = build_instance_from_model(main_models.GetDataLakeCachePrefetchJobResponseBody)
instanceRegistry["GetDataLakeStorageTransferJob"] = build_instance_from_model(main_models.GetDataLakeStorageTransferJobResponseBody)
instanceRegistry["GetLiveChannelHistory"] = build_instance_from_model(main_models.GetLiveChannelHistoryResponseBody)
instanceRegistry["GetLiveChannelInfo"] = build_instance_from_model(main_models.GetLiveChannelInfoResponseBody)
instanceRegistry["GetLiveChannelStat"] = build_instance_from_model(main_models.GetLiveChannelStatResponseBody)
instanceRegistry["GetMetaQueryStatus"] = build_instance_from_model(main_models.GetMetaQueryStatusResponseBody)
instanceRegistry["GetObjectAcl"] = build_instance_from_model(main_models.GetObjectAclResponseBody)
instanceRegistry["GetObjectGroupIndex"] = build_instance_from_model(main_models.GetObjectGroupIndexResponseBody)
instanceRegistry["GetObjectInfo"] = build_instance_from_model(main_models.GetObjectInfoResponseBody)
instanceRegistry["GetObjectLink"] = build_instance_from_model(main_models.GetObjectLinkResponseBody)
instanceRegistry["GetObjectTagging"] = build_instance_from_model(main_models.GetObjectTaggingResponseBody)
instanceRegistry["GetProcessConfiguration"] = build_instance_from_model(main_models.GetProcessConfigurationResponseBody)
instanceRegistry["GetPublicAccessBlock"] = build_instance_from_model(main_models.GetPublicAccessBlockResponseBody)
instanceRegistry["GetReservedCapacity"] = build_instance_from_model(main_models.GetReservedCapacityResponseBody)
instanceRegistry["GetResourcePoolBucketGroupQoSInfo"] = build_instance_from_model(main_models.GetResourcePoolBucketGroupQoSInfoResponseBody)
instanceRegistry["GetResourcePoolInfo"] = build_instance_from_model(main_models.GetResourcePoolInfoResponseBody)
instanceRegistry["GetResourcePoolRequesterQoSInfo"] = build_instance_from_model(main_models.GetResourcePoolRequesterQoSInfoResponseBody)
instanceRegistry["GetStyle"] = build_instance_from_model(main_models.GetStyleResponseBody)
instanceRegistry["GetUserAntiDDosInfo"] = build_instance_from_model(main_models.GetUserAntiDDosInfoResponseBody)
instanceRegistry["GetUserDefinedLogFieldsConfig"] = build_instance_from_model(main_models.GetUserDefinedLogFieldsConfigResponseBody)
instanceRegistry["GetUserQoSInfo"] = build_instance_from_model(main_models.GetUserQoSInfoResponseBody)
instanceRegistry["GetVectorBucket"] = build_instance_from_model(main_models.GetVectorBucketResponseBody)
instanceRegistry["GetVectorIndex"] = build_instance_from_model(main_models.GetVectorIndexResponseBody)
instanceRegistry["GetVectors"] = build_instance_from_model(main_models.GetVectorsResponseBody)
instanceRegistry["GetVirtualBucket"] = build_instance_from_model(main_models.GetVirtualBucketResponseBody)
instanceRegistry["InitiateMultipartUpload"] = build_instance_from_model(main_models.InitiateMultipartUploadResponseBody)
instanceRegistry["ListAccessPoints"] = build_instance_from_model(main_models.ListAccessPointsResponseBody)
instanceRegistry["ListAccessPointsForObjectProcess"] = build_instance_from_model(main_models.ListAccessPointsForObjectProcessResponseBody)
instanceRegistry["ListBucketAntiDDosInfo"] = build_instance_from_model(main_models.ListBucketAntiDDosInfoResponseBody)
instanceRegistry["ListBucketDataRedundancyTransition"] = build_instance_from_model(main_models.ListBucketDataRedundancyTransitionResponseBody)
instanceRegistry["ListBucketInventory"] = build_instance_from_model(main_models.ListBucketInventoryResponseBody)
instanceRegistry["ListBucketRequesterQoSInfos"] = build_instance_from_model(main_models.ListBucketRequesterQoSInfosResponseBody)
instanceRegistry["ListBuckets"] = build_instance_from_model(main_models.ListBucketsResponseBody)
instanceRegistry["ListCache"] = build_instance_from_model(main_models.ListCacheResponseBody)
instanceRegistry["ListCname"] = build_instance_from_model(main_models.ListCnameResponseBody)
instanceRegistry["ListDataLakeCachePrefetchJob"] = build_instance_from_model(main_models.ListDataLakeCachePrefetchJobResponseBody)
instanceRegistry["ListDataLakeCachePrefetchJobHistory"] = build_instance_from_model(main_models.ListDataLakeCachePrefetchJobHistoryResponseBody)
instanceRegistry["ListDataLakeStorageTransferJob"] = build_instance_from_model(main_models.ListDataLakeStorageTransferJobResponseBody)
instanceRegistry["ListDataLakeStorageTransferJobHistory"] = build_instance_from_model(main_models.ListDataLakeStorageTransferJobHistoryResponseBody)
instanceRegistry["ListJobs"] = build_instance_from_model(main_models.ListJobsResponseBody)
instanceRegistry["ListLiveChannel"] = build_instance_from_model(main_models.ListLiveChannelResponseBody)
instanceRegistry["ListMultipartUploads"] = build_instance_from_model(main_models.ListMultipartUploadsResponseBody)
instanceRegistry["ListObjectVersions"] = build_instance_from_model(main_models.ListObjectVersionsResponseBody)
instanceRegistry["ListObjects"] = build_instance_from_model(main_models.ListObjectsResponseBody)
instanceRegistry["ListObjectsV2"] = build_instance_from_model(main_models.ListObjectsV2ResponseBody)
instanceRegistry["ListParts"] = build_instance_from_model(main_models.ListPartsResponseBody)
instanceRegistry["ListReservedCapacity"] = build_instance_from_model(main_models.ListReservedCapacityResponseBody)
instanceRegistry["ListResourcePoolBucketGroupQoSInfos"] = build_instance_from_model(main_models.ListResourcePoolBucketGroupQoSInfosResponseBody)
instanceRegistry["ListResourcePoolBucketGroups"] = build_instance_from_model(main_models.ListResourcePoolBucketGroupsResponseBody)
instanceRegistry["ListResourcePoolBuckets"] = build_instance_from_model(main_models.ListResourcePoolBucketsResponseBody)
instanceRegistry["ListResourcePoolRequesterQoSInfos"] = build_instance_from_model(main_models.ListResourcePoolRequesterQoSInfosResponseBody)
instanceRegistry["ListResourcePools"] = build_instance_from_model(main_models.ListResourcePoolsResponseBody)
instanceRegistry["ListStyle"] = build_instance_from_model(main_models.ListStyleResponseBody)
instanceRegistry["ListUserDataRedundancyTransition"] = build_instance_from_model(main_models.ListUserDataRedundancyTransitionResponseBody)
instanceRegistry["ListUserRegions"] = build_instance_from_model(main_models.ListUserRegionsResponseBody)
instanceRegistry["ListVectorBuckets"] = build_instance_from_model(main_models.ListVectorBucketsResponseBody)
instanceRegistry["ListVectorIndexes"] = build_instance_from_model(main_models.ListVectorIndexesResponseBody)
instanceRegistry["ListVectors"] = build_instance_from_model(main_models.ListVectorsResponseBody)
instanceRegistry["ListVirtualBucket"] = build_instance_from_model(main_models.ListVirtualBucketResponseBody)
instanceRegistry["PostAsyncFetchTask"] = build_instance_from_model(main_models.PostAsyncFetchTaskResponseBody)
instanceRegistry["PostObjectGroup"] = build_instance_from_model(main_models.PostObjectGroupResponseBody)
instanceRegistry["PutBucketReplicationBandwidth"] = build_instance_from_model(main_models.PutBucketReplicationBandwidthResponseBody)
instanceRegistry["PutDataLakeCachePrefetchJob"] = build_instance_from_model(main_models.PutDataLakeCachePrefetchJobResponseBody)
instanceRegistry["PutDataLakeStorageTransferJob"] = build_instance_from_model(main_models.PutDataLakeStorageTransferJobResponseBody)
instanceRegistry["PutLiveChannel"] = build_instance_from_model(main_models.PutLiveChannelResponseBody)
instanceRegistry["PutObjectLink"] = build_instance_from_model(main_models.PutObjectLinkResponseBody)
instanceRegistry["QueryVectors"] = build_instance_from_model(main_models.QueryVectorsResponseBody)
instanceRegistry["StartDataLakeStorageTransferJob"] = build_instance_from_model(main_models.StartDataLakeStorageTransferJobResponseBody)
instanceRegistry["StartPartUpload"] = build_instance_from_model(main_models.StartPartUploadResponseBody)
instanceRegistry["UpdateJobPriority"] = build_instance_from_model(main_models.UpdateJobPriorityResponseBody)
instanceRegistry["UpdateJobStatus"] = build_instance_from_model(main_models.UpdateJobStatusResponseBody)
instanceRegistry["UploadPartCopy"] = build_instance_from_model(main_models.UploadPartCopyResponseBody)


# for hcs-mgw
main_models = hcs_mgw_models
instanceRegistry["GetAddress"] = build_instance_from_model(main_models.GetAddressResponseBody)
instanceRegistry["GetAgent"] = build_instance_from_model(main_models.GetAgentResponseBody)
instanceRegistry["GetAgentStatus"] = build_instance_from_model(main_models.GetAgentStatusResponseBody)
instanceRegistry["GetJob"] = build_instance_from_model(main_models.GetJobResponseBody)
instanceRegistry["GetJobResult"] = build_instance_from_model(main_models.GetJobResultResponseBody)
instanceRegistry["GetReport"] = build_instance_from_model(main_models.GetReportResponseBody)
instanceRegistry["GetTunnel"] = build_instance_from_model(main_models.GetTunnelResponseBody)
instanceRegistry["ListAddress"] = build_instance_from_model(main_models.ListAddressResponseBody)
instanceRegistry["ListAgent"] = build_instance_from_model(main_models.ListAgentResponseBody)
instanceRegistry["ListJob"] = build_instance_from_model(main_models.ListJobResponseBody)
instanceRegistry["ListJobHistory"] = build_instance_from_model(main_models.ListJobHistoryResponseBody)
instanceRegistry["ListTunnel"] = build_instance_from_model(main_models.ListTunnelResponseBody)
instanceRegistry["VerifyAddress"] = build_instance_from_model(main_models.VerifyAddressResponseBody)


class Client:
    @staticmethod
    def __parse_xml_impl(t, m):
        if t.tag not in m:
            return {}
        d = {t.tag: {} if t.attrib else None}
        children = list(t)
        if children:
            m = m[t.tag]
            mc = m[0] if isinstance(m, list) else m
            dd = defaultdict(list)
            for dc in [Client.__parse_xml_impl(c, mc) for c in children]:
                for k, v in dc.items():
                    dd[k].append(v)
            d = {t.tag: {k: v[0] if len(v) == 1 and not isinstance(mc[k], list) else v for k, v in dd.items() if k in mc}}

        if t.attrib:
            d[t.tag].update(('@' + k, v) for k, v in t.attrib.items())

        if t.text:
            text = t.text.strip()
            if children or t.attrib:
                if text:
                    d[t.tag]['#text'] = text
            else:
                type_ = type(m[t.tag][0]) if isinstance(m[t.tag], list) else type(m[t.tag])
                if type_ == bool:
                    d[t.tag] = text == "true"
                else:
                    d[t.tag] = type_(text)
        return d

    @staticmethod
    def parse_xml(bodyStr: str, apiName: str):
        return Client.__parse_xml_impl(ElementTree.fromstring(bodyStr), instanceRegistry[apiName].to_map())
    
    @staticmethod
    def __get_xml_factory(elem, val, parent_element=None):
        if val is None:
            return

        if isinstance(val, dict):
            Client.__get_xml_by_dict(elem, val)
        elif isinstance(val, (list, tuple, set)):
            if parent_element is None:
                raise RequiredArgumentException("Missing root tag")
            Client.__get_xml_by_list(elem, val, parent_element)
        elif isinstance(val, bool):
            elem.text = str(val).lower()
        else:
            elem.text = str(val)

    @staticmethod
    def __get_xml_by_dict(elem, val):
        for k in val:
            sub_elem = ElementTree.SubElement(elem, k)
            Client.__get_xml_factory(sub_elem, val[k], elem)

    @staticmethod
    def __get_xml_by_list(elem, val, parent_element):
        i = 0
        tag_name = elem.tag
        if val.__len__() > 0:
            Client.__get_xml_factory(elem, val[0], parent_element)

        for item in val:
            if i > 0:
                sub_elem = ElementTree.SubElement(parent_element, tag_name)
                Client.__get_xml_factory(sub_elem, item, parent_element)
            i = i + 1
    
    @staticmethod
    def to_xml(body):
        if body is None:
            return

        dic = {}
        if isinstance(body, TeaModel):
            dic = body.to_map()
        elif isinstance(body, dict):
            dic = body

        if dic.__len__() == 0:
            return ""
        else:
            result_xml = '<?xml version="1.0" encoding="utf-8"?>'
            for k in dic:
                elem = ElementTree.Element(k)
                Client.__get_xml_factory(elem, dic[k])
                result_xml += bytes.decode(ElementTree.tostring(elem), encoding="utf-8")
            return result_xml

