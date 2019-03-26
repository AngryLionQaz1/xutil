package com.snow.golang.config.advice;

import com.alibaba.fastjson.JSONObject;
import com.snow.golang.common.bean.Config;
import com.snow.golang.common.util.AESUtils;
import com.snow.golang.config.annotation.Encrypt;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.MethodParameter;
import org.springframework.http.MediaType;
import org.springframework.http.converter.HttpMessageConverter;
import org.springframework.http.server.ServerHttpRequest;
import org.springframework.http.server.ServerHttpResponse;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.servlet.mvc.method.annotation.ResponseBodyAdvice;

/**
 * 请求响应处理类<br>
 * 对加了@Encrypt的方法的数据进行加密操作
 */
@ControllerAdvice
@Slf4j
public class EncryptResponseBodyAdvice implements ResponseBodyAdvice<Object> {

    @Autowired
    private Config config;
    private static ThreadLocal<Boolean> encryptLocal = new ThreadLocal<Boolean>();
    private JSONObject objectMapper = new JSONObject();

    @Override
    public boolean supports(MethodParameter methodParameter, Class<? extends HttpMessageConverter<?>> aClass) {
        return true;
    }

    @Override
    public Object beforeBodyWrite(Object o, MethodParameter methodParameter, MediaType mediaType, Class<? extends HttpMessageConverter<?>> aClass, ServerHttpRequest serverHttpRequest, ServerHttpResponse serverHttpResponse) {
        Boolean status = encryptLocal.get();
        if (status != null && status == false) {
            encryptLocal.remove();
            return o;
        }
        long startTime = System.currentTimeMillis();
        boolean encrypt = false;
        if (methodParameter.getMethod().isAnnotationPresent(Encrypt.class) && !config.isAesDebug()) {
            encrypt = true;
        }
        if (encrypt) {
            try {
                String content=objectMapper.toJSONString(o);
//                String result =  AesEncryptUtils.aesEncrypt(content, config.getAesKey());
                String result =  AESUtils.aesEncrypt(content, config.getAesKey());
                long endTime = System.currentTimeMillis();
                log.debug("Encrypt Time:" + (endTime - startTime));
                return result;
            } catch (Exception e) {
                log.error("加密数据异常", e);
            }
        }

        return o;

    }
}
