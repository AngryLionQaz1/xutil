package com.snow.golang.config.annotation;

import java.lang.annotation.*;

@Target({ElementType.PARAMETER, ElementType.METHOD,ElementType.TYPE})
@Retention(RetentionPolicy.RUNTIME)
@Documented
public @interface Security {


    boolean flag() default false;//是否是菜单
    int order() default 0;//菜单排序
    int menu() default 0;//第几级菜单
    String names() default "";//菜单名称 : 一级菜单,二级菜单，三级菜单........
    int sign() default 0;//菜单标记，用于区分同级菜单
    String value()  default "";//哪些不行拦截的uri




}

