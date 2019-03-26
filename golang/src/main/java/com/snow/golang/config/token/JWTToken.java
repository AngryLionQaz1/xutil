package com.snow.golang.config.token;


import com.snow.golang.common.bean.Config;
import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.Date;

@Component
public class JWTToken {

    private static final int secondIn1day = 1000 * 60 * 60 * 24;

    @Autowired
    private Config config;


    //创建Token
    public  String createToken(String userId){
        long now = (new Date()).getTime();              //获取当前时间戳
        Date validity = new Date(now + secondIn1day*config.getJwtTokenValidity());
        return Jwts.builder()                                   //创建Token令牌
                .setSubject(userId)                             //设置面向用户
                .claim(config.getJwtKey(),userId)                  //添加权限属性
                .setExpiration(validity)                        //设置失效时间
                .signWith(SignatureAlgorithm.HS512,config.getJwtSecretKey())   //生成签名
                .compact();
    }


    //获取用户id
    public  String getUserId(String token){
        Claims claims = Jwts.parser()                           //解析Token的payload
                .setSigningKey(config.getJwtSecretKey())
                .parseClaimsJws(token)
                .getBody();

        return  claims.get(config.getJwtKey()).toString();
    }


    //验证Token是否正确
    public  boolean validateToken(String token){
        try {
            Jwts.parser().setSigningKey(config.getJwtSecretKey()).parseClaimsJws(token);   //通过密钥验证Token
            return true;
        } catch (Exception e) {

        }
        return false;
    }


}
