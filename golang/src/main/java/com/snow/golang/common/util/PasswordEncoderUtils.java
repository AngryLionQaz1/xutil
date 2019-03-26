package com.snow.golang.common.util;

import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;

public class PasswordEncoderUtils {


    private static BCryptPasswordEncoder passwordEncoder = new BCryptPasswordEncoder();


    public static String encode(String password){
        return passwordEncoder.encode(password);
    }

    public static Boolean decode(String password,String encodePassword){
        if (password==null||encodePassword==null)return false;
        return passwordEncoder.matches(password,encodePassword);
    }


   public static void main(String[]args){

        System.out.println(encode("123456"));
        System.out.println(decode("1234565",encode("123456")));

   }




}
