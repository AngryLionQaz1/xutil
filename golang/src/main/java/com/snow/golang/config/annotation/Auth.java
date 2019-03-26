/**
 * Copyright (C), 2015-2019, 重庆了赢科技有限公司
 * FileName: Auth
 * Author:   萧毅
 * Date:     2019/3/4 10:12
 * Description: 生成权限,作用于controller 中的 方法上
 */
package com.snow.golang.config.annotation;

import java.lang.annotation.*;

@Target({ElementType.METHOD})
@Retention(RetentionPolicy.RUNTIME)
@Documented
public @interface Auth {

    String value() default "";



}
