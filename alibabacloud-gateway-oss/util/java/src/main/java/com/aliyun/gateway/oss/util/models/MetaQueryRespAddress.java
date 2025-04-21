// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class MetaQueryRespAddress extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>中国浙江省杭州市余杭区文一西路969号</p>
     */
    @NameInMap("AddressLine")
    public String addressLine;

    /**
     * <strong>example:</strong>
     * <p>杭州市</p>
     */
    @NameInMap("City")
    public String city;

    /**
     * <strong>example:</strong>
     * <p>余杭区</p>
     */
    @NameInMap("District")
    public String district;

    /**
     * <strong>example:</strong>
     * <p>zh-Hans</p>
     */
    @NameInMap("Language")
    public String language;

    /**
     * <strong>example:</strong>
     * <p>浙江省</p>
     */
    @NameInMap("Province")
    public String province;

    /**
     * <strong>example:</strong>
     * <p>文一西路</p>
     */
    @NameInMap("Township")
    public String township;

    public static MetaQueryRespAddress build(java.util.Map<String, ?> map) throws Exception {
        MetaQueryRespAddress self = new MetaQueryRespAddress();
        return TeaModel.build(map, self);
    }

    public MetaQueryRespAddress setAddressLine(String addressLine) {
        this.addressLine = addressLine;
        return this;
    }
    public String getAddressLine() {
        return this.addressLine;
    }

    public MetaQueryRespAddress setCity(String city) {
        this.city = city;
        return this;
    }
    public String getCity() {
        return this.city;
    }

    public MetaQueryRespAddress setDistrict(String district) {
        this.district = district;
        return this;
    }
    public String getDistrict() {
        return this.district;
    }

    public MetaQueryRespAddress setLanguage(String language) {
        this.language = language;
        return this;
    }
    public String getLanguage() {
        return this.language;
    }

    public MetaQueryRespAddress setProvince(String province) {
        this.province = province;
        return this;
    }
    public String getProvince() {
        return this.province;
    }

    public MetaQueryRespAddress setTownship(String township) {
        this.township = township;
        return this;
    }
    public String getTownship() {
        return this.township;
    }

}
