package com.snow.golang.config.json;

import com.alibaba.fastjson.serializer.SerializerFeature;
import com.alibaba.fastjson.support.config.FastJsonConfig;
import com.alibaba.fastjson.support.spring.FastJsonHttpMessageConverter;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.MediaType;
import org.springframework.http.converter.HttpMessageConverter;

import java.util.ArrayList;
import java.util.List;

@Configuration
public class JsonConfig {
    /**
     * 定制HTTP消息转换器
     * @return
     */
    @Bean
    public HttpMessageConverter fastJsonConverters(){
        // 1.定义一个convert 转换消息的对象
        FastJsonHttpMessageConverter fastConverter = new FastJsonHttpMessageConverter();
        // 2 添加fastjson 的配置信息 比如 是否要格式化 返回的json数据
        FastJsonConfig fastJsonConfig = new FastJsonConfig();
        fastJsonConfig.setSerializerFeatures(
                //格式化
                SerializerFeature.PrettyFormat,
                //是否输出值为null的字段,默认为false
                SerializerFeature.WriteMapNullValue,
                //null属性显示为""
                SerializerFeature.WriteNullStringAsEmpty,
                //消除对同一对象循环引用的问题，默认为false，去掉$监测
                SerializerFeature.DisableCircularReferenceDetect,
                //如果为null,输出为[],而非null
                SerializerFeature.WriteNullListAsEmpty
                );
        fastConverter.setFastJsonConfig(fastJsonConfig);
        // 解决乱码的问题
        List<MediaType> fastMediaTypes = new ArrayList<MediaType>();
        fastMediaTypes.add(MediaType.APPLICATION_JSON_UTF8);
        fastConverter.setSupportedMediaTypes(fastMediaTypes);
        HttpMessageConverter<?> converter = fastConverter;
        return converter;
    }
}
