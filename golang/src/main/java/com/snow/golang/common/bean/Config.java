package com.snow.golang.common.bean;

import lombok.Data;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Component;

@Data
@Component
@ConfigurationProperties(prefix ="config" )
public class Config {
    /**
     * 请求头
     */
    private String authorization = "authorization";
    /**
     * 存储当前登录token
     */
    private String token = "authorization";
    /**
     * JWT字段名
     */
    private String jwtKey = "AUTHORITIES_KEY";
    /**
     * JWT签名密钥
     */
    private String jwtSecretKey = "secretKey";
    /**
     * JWT有效期
     */
    private Long jwtTokenValidity = 7L;

    /**拦截uri*/
    private String addPath="/token/**";
    /**不拦截uri*/
    private String excludePath="/test/**";

   /**端口号*/
    private Integer filePort=8080;
    /**本地文件地址*/
    private String  filePath;
    /**IP地址*/
    private String  fileHost;
    /**请求地址*/
    private String  fileUrl="/upload/**";
    /**设置不能上传的文件类型*/
    private String  fileType="php,java,jsp";

    /**权限管理的 超级管理员角色*/
    private String authorityAdmin="admin";
    /**权限管理 是否初始化 权限*/
    private Boolean authorityInit=false;
    /**AES加密KEY*/
    private String aesKey="QAZWSXEDCR123456";
    /**字符集*/
    private String aesCharset="UTF-8";
    /**
     * 开启调试模式，调试模式下不进行加解密操作，用于像Swagger这种在线API测试场景
     */
    private boolean aesDebug = false;





}


