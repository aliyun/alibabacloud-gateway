# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
# -*- coding: utf-8 -*-
# import alibabacloud_tea_xml.client

from Tea.model import TeaModel
from typing import Dict
from xml.etree import ElementTree
from collections import defaultdict
import inspect
from typing_extensions import get_origin, get_args
from Tea.exceptions import RequiredArgumentException
from .structs import *


basic_instance = {}
basic_instance[str] = ''
basic_instance[int] = 0
basic_instance[bool] = False
basic_instance[Dict[str, str]] = {'':''}

def build_instance_from_model(model):
    if model in basic_instance:
        return basic_instance[model]
    sig = inspect.signature(model.__init__)
    params = {}
    for para_name, param in sig.parameters.items():
        if para_name == "self":
            continue
        origin_type = get_origin(param.annotation)
        if origin_type is not None and issubclass(origin_type, list):
            params[para_name] = [build_instance_from_model(get_args(param.annotation)[0])]
        else:
            params[para_name] = build_instance_from_model(param.annotation)
    return model(**params)


instanceRegistry = {}
# for oss
instanceRegistry["CompleteMultipartUpload"] = build_instance_from_model(CompleteMultipartUploadResponseBody)
instanceRegistry["CopyObject"] = build_instance_from_model(CopyObjectResponseBody)
instanceRegistry["CopyObjects"] = build_instance_from_model(CopyObjectsResponseBody)
instanceRegistry["CreateAccessPoint"] = build_instance_from_model(CreateAccessPointResponseBody)
instanceRegistry["CreateAccessPointForObjectProcess"] = build_instance_from_model(CreateAccessPointForObjectProcessResponseBody)
instanceRegistry["CreateBucketDataRedundancyTransition"] = build_instance_from_model(CreateBucketDataRedundancyTransitionResponseBody)
instanceRegistry["CreateCnameToken"] = build_instance_from_model(CreateCnameTokenResponseBody)
instanceRegistry["CreateReservedCapacity"] = build_instance_from_model(CreateReservedCapacityResponseBody)
instanceRegistry["DeleteMultipleObjects"] = build_instance_from_model(DeleteMultipleObjectsResponseBody)
instanceRegistry["DescribeRegions"] = build_instance_from_model(DescribeRegionsResponseBody)
instanceRegistry["DoMetaQuery"] = build_instance_from_model(DoMetaQueryResponseBody)
instanceRegistry["GetAccessPoint"] = build_instance_from_model(GetAccessPointResponseBody)
instanceRegistry["GetAccessPointConfigForObjectProcess"] = build_instance_from_model(GetAccessPointConfigForObjectProcessResponseBody)
instanceRegistry["GetAccessPointForObjectProcess"] = build_instance_from_model(GetAccessPointForObjectProcessResponseBody)
instanceRegistry["GetAccessPointPublicAccessBlock"] = build_instance_from_model(GetAccessPointPublicAccessBlockResponseBody)
instanceRegistry["GetAsyncFetchTask"] = build_instance_from_model(GetAsyncFetchTaskResponseBody)
instanceRegistry["GetBucketAccessMonitor"] = build_instance_from_model(GetBucketAccessMonitorResponseBody)
instanceRegistry["GetBucketAcl"] = build_instance_from_model(GetBucketAclResponseBody)
instanceRegistry["GetBucketArchiveDirectRead"] = build_instance_from_model(GetBucketArchiveDirectReadResponseBody)
instanceRegistry["GetBucketCacheConfiguration"] = build_instance_from_model(GetBucketCacheConfigurationResponseBody)
instanceRegistry["GetBucketCallbackPolicy"] = build_instance_from_model(GetBucketCallbackPolicyResponseBody)
instanceRegistry["GetBucketCommonHeader"] = build_instance_from_model(GetBucketCommonHeaderResponseBody)
instanceRegistry["GetBucketCors"] = build_instance_from_model(GetBucketCorsResponseBody)
instanceRegistry["GetBucketDataAccelerator"] = build_instance_from_model(GetBucketDataAcceleratorResponseBody)
instanceRegistry["GetBucketDataRedundancyTransition"] = build_instance_from_model(GetBucketDataRedundancyTransitionResponseBody)
instanceRegistry["GetBucketEncryption"] = build_instance_from_model(GetBucketEncryptionResponseBody)
instanceRegistry["GetBucketEventNotification"] = build_instance_from_model(GetBucketEventNotificationResponseBody)
instanceRegistry["GetBucketHash"] = build_instance_from_model(GetBucketHashResponseBody)
instanceRegistry["GetBucketHttpsConfig"] = build_instance_from_model(GetBucketHttpsConfigResponseBody)
instanceRegistry["GetBucketInfo"] = build_instance_from_model(GetBucketInfoResponseBody)
instanceRegistry["GetBucketInventory"] = build_instance_from_model(GetBucketInventoryResponseBody)
instanceRegistry["GetBucketLifecycle"] = build_instance_from_model(GetBucketLifecycleResponseBody)
instanceRegistry["GetBucketLocation"] = build_instance_from_model(GetBucketLocationResponseBody)
instanceRegistry["GetBucketLogging"] = build_instance_from_model(GetBucketLoggingResponseBody)
instanceRegistry["GetBucketNotification"] = build_instance_from_model(GetBucketNotificationResponseBody)
instanceRegistry["GetBucketPolicyStatus"] = build_instance_from_model(GetBucketPolicyStatusResponseBody)
instanceRegistry["GetBucketPublicAccessBlock"] = build_instance_from_model(GetBucketPublicAccessBlockResponseBody)
instanceRegistry["GetBucketQoSInfo"] = build_instance_from_model(GetBucketQoSInfoResponseBody)
instanceRegistry["GetBucketReferer"] = build_instance_from_model(GetBucketRefererResponseBody)
instanceRegistry["GetBucketReplication"] = build_instance_from_model(GetBucketReplicationResponseBody)
instanceRegistry["GetBucketReplicationLocation"] = build_instance_from_model(GetBucketReplicationLocationResponseBody)
instanceRegistry["GetBucketReplicationProgress"] = build_instance_from_model(GetBucketReplicationProgressResponseBody)
instanceRegistry["GetBucketRequestPayment"] = build_instance_from_model(GetBucketRequestPaymentResponseBody)
instanceRegistry["GetBucketRequesterQoSInfo"] = build_instance_from_model(GetBucketRequesterQoSInfoResponseBody)
instanceRegistry["GetBucketResourceGroup"] = build_instance_from_model(GetBucketResourceGroupResponseBody)
instanceRegistry["GetBucketResponseHeader"] = build_instance_from_model(GetBucketResponseHeaderResponseBody)
instanceRegistry["GetBucketStat"] = build_instance_from_model(GetBucketStatResponseBody)
instanceRegistry["GetBucketTags"] = build_instance_from_model(GetBucketTagsResponseBody)
instanceRegistry["GetBucketTransferAcceleration"] = build_instance_from_model(GetBucketTransferAccelerationResponseBody)
instanceRegistry["GetBucketVersioning"] = build_instance_from_model(GetBucketVersioningResponseBody)
instanceRegistry["GetBucketWebsite"] = build_instance_from_model(GetBucketWebsiteResponseBody)
instanceRegistry["GetBucketWorm"] = build_instance_from_model(GetBucketWormResponseBody)
instanceRegistry["GetCache"] = build_instance_from_model(GetCacheResponseBody)
instanceRegistry["GetChannel"] = build_instance_from_model(GetChannelResponseBody)
instanceRegistry["GetCnameToken"] = build_instance_from_model(GetCnameTokenResponseBody)
instanceRegistry["GetDataLakeCachePrefetchJob"] = build_instance_from_model(GetDataLakeCachePrefetchJobResponseBody)
instanceRegistry["GetDataLakeStorageTransferJob"] = build_instance_from_model(GetDataLakeStorageTransferJobResponseBody)
instanceRegistry["GetLiveChannelHistory"] = build_instance_from_model(GetLiveChannelHistoryResponseBody)
instanceRegistry["GetLiveChannelInfo"] = build_instance_from_model(GetLiveChannelInfoResponseBody)
instanceRegistry["GetLiveChannelStat"] = build_instance_from_model(GetLiveChannelStatResponseBody)
instanceRegistry["GetMetaQueryStatus"] = build_instance_from_model(GetMetaQueryStatusResponseBody)
instanceRegistry["GetObjectAcl"] = build_instance_from_model(GetObjectAclResponseBody)
instanceRegistry["GetObjectGroupIndex"] = build_instance_from_model(GetObjectGroupIndexResponseBody)
instanceRegistry["GetObjectInfo"] = build_instance_from_model(GetObjectInfoResponseBody)
instanceRegistry["GetObjectLink"] = build_instance_from_model(GetObjectLinkResponseBody)
instanceRegistry["GetObjectTagging"] = build_instance_from_model(GetObjectTaggingResponseBody)
instanceRegistry["GetProcessConfiguration"] = build_instance_from_model(GetProcessConfigurationResponseBody)
instanceRegistry["GetPublicAccessBlock"] = build_instance_from_model(GetPublicAccessBlockResponseBody)
instanceRegistry["GetReservedCapacity"] = build_instance_from_model(GetReservedCapacityResponseBody)
instanceRegistry["GetResourcePoolInfo"] = build_instance_from_model(GetResourcePoolInfoResponseBody)
instanceRegistry["GetResourcePoolRequesterQoSInfo"] = build_instance_from_model(GetResourcePoolRequesterQoSInfoResponseBody)
instanceRegistry["GetStyle"] = build_instance_from_model(GetStyleResponseBody)
instanceRegistry["GetUserAntiDDosInfo"] = build_instance_from_model(GetUserAntiDDosInfoResponseBody)
instanceRegistry["GetUserDefinedLogFieldsConfig"] = build_instance_from_model(GetUserDefinedLogFieldsConfigResponseBody)
instanceRegistry["GetUserQoSInfo"] = build_instance_from_model(GetUserQoSInfoResponseBody)
instanceRegistry["GetVirtualBucket"] = build_instance_from_model(GetVirtualBucketResponseBody)
instanceRegistry["InitiateMultipartUpload"] = build_instance_from_model(InitiateMultipartUploadResponseBody)
instanceRegistry["ListAccessPoints"] = build_instance_from_model(ListAccessPointsResponseBody)
instanceRegistry["ListAccessPointsForObjectProcess"] = build_instance_from_model(ListAccessPointsForObjectProcessResponseBody)
instanceRegistry["ListBucketAntiDDosInfo"] = build_instance_from_model(ListBucketAntiDDosInfoResponseBody)
instanceRegistry["ListBucketDataRedundancyTransition"] = build_instance_from_model(ListBucketDataRedundancyTransitionResponseBody)
instanceRegistry["ListBucketInventory"] = build_instance_from_model(ListBucketInventoryResponseBody)
instanceRegistry["ListBucketRequesterQoSInfos"] = build_instance_from_model(ListBucketRequesterQoSInfosResponseBody)
instanceRegistry["ListBuckets"] = build_instance_from_model(ListBucketsResponseBody)
instanceRegistry["ListCache"] = build_instance_from_model(ListCacheResponseBody)
instanceRegistry["ListCname"] = build_instance_from_model(ListCnameResponseBody)
instanceRegistry["ListDataLakeCachePrefetchJob"] = build_instance_from_model(ListDataLakeCachePrefetchJobResponseBody)
instanceRegistry["ListDataLakeCachePrefetchJobHistory"] = build_instance_from_model(ListDataLakeCachePrefetchJobHistoryResponseBody)
instanceRegistry["ListDataLakeStorageTransferJob"] = build_instance_from_model(ListDataLakeStorageTransferJobResponseBody)
instanceRegistry["ListDataLakeStorageTransferJobHistory"] = build_instance_from_model(ListDataLakeStorageTransferJobHistoryResponseBody)
instanceRegistry["ListLiveChannel"] = build_instance_from_model(ListLiveChannelResponseBody)
instanceRegistry["ListMultipartUploads"] = build_instance_from_model(ListMultipartUploadsResponseBody)
instanceRegistry["ListObjectVersions"] = build_instance_from_model(ListObjectVersionsResponseBody)
instanceRegistry["ListObjects"] = build_instance_from_model(ListObjectsResponseBody)
instanceRegistry["ListObjectsV2"] = build_instance_from_model(ListObjectsV2ResponseBody)
instanceRegistry["ListParts"] = build_instance_from_model(ListPartsResponseBody)
instanceRegistry["ListReservedCapacity"] = build_instance_from_model(ListReservedCapacityResponseBody)
instanceRegistry["ListResourcePoolBuckets"] = build_instance_from_model(ListResourcePoolBucketsResponseBody)
instanceRegistry["ListResourcePoolRequesterQoSInfos"] = build_instance_from_model(ListResourcePoolRequesterQoSInfosResponseBody)
instanceRegistry["ListResourcePools"] = build_instance_from_model(ListResourcePoolsResponseBody)
instanceRegistry["ListStyle"] = build_instance_from_model(ListStyleResponseBody)
instanceRegistry["ListUserDataRedundancyTransition"] = build_instance_from_model(ListUserDataRedundancyTransitionResponseBody)
instanceRegistry["ListUserRegions"] = build_instance_from_model(ListUserRegionsResponseBody)
instanceRegistry["ListVirtualBucket"] = build_instance_from_model(ListVirtualBucketResponseBody)
instanceRegistry["PostAsyncFetchTask"] = build_instance_from_model(PostAsyncFetchTaskResponseBody)
instanceRegistry["PostObjectGroup"] = build_instance_from_model(PostObjectGroupResponseBody)
instanceRegistry["PutDataLakeCachePrefetchJob"] = build_instance_from_model(PutDataLakeCachePrefetchJobResponseBody)
instanceRegistry["PutDataLakeStorageTransferJob"] = build_instance_from_model(PutDataLakeStorageTransferJobResponseBody)
instanceRegistry["PutLiveChannel"] = build_instance_from_model(PutLiveChannelResponseBody)
instanceRegistry["PutObjectLink"] = build_instance_from_model(PutObjectLinkResponseBody)
instanceRegistry["StartDataLakeStorageTransferJob"] = build_instance_from_model(StartDataLakeStorageTransferJobResponseBody)
instanceRegistry["StartPartUpload"] = build_instance_from_model(StartPartUploadResponseBody)
instanceRegistry["UploadPartCopy"] = build_instance_from_model(UploadPartCopyResponseBody)


# for hcs-mgw
instanceRegistry["GetAddress"] = build_instance_from_model(GetAddressResponseBody)
instanceRegistry["GetAgent"] = build_instance_from_model(GetAgentResponseBody)
instanceRegistry["GetAgentStatus"] = build_instance_from_model(GetAgentStatusResponseBody)
instanceRegistry["GetJob"] = build_instance_from_model(GetJobResponseBody)
instanceRegistry["GetJobResult"] = build_instance_from_model(GetJobResultResponseBody)
instanceRegistry["GetReport"] = build_instance_from_model(GetReportResponseBody)
instanceRegistry["GetTunnel"] = build_instance_from_model(GetTunnelResponseBody)
instanceRegistry["ListAddress"] = build_instance_from_model(ListAddressResponseBody)
instanceRegistry["ListAgent"] = build_instance_from_model(ListAgentResponseBody)
instanceRegistry["ListJob"] = build_instance_from_model(ListJobResponseBody)
instanceRegistry["ListJobHistory"] = build_instance_from_model(ListJobHistoryResponseBody)
instanceRegistry["ListTunnel"] = build_instance_from_model(ListTunnelResponseBody)
instanceRegistry["VerifyAddress"] = build_instance_from_model(VerifyAddressResponseBody)


class Client:
    @staticmethod
    def __parse_xml_impl(t, m):
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
                d[t.tag] = text
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

