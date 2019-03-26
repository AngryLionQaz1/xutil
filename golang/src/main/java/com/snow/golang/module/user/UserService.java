/**
 * Copyright (C), 2015-2019, 重庆了赢科技有限公司
 * FileName: UserService
 * Author:   萧毅
 * Date:     2019/3/26 16:40
 * Description:
 */
package com.snow.golang.module.user;

import com.snow.golang.common.bean.Result;
import com.snow.golang.config.security.SecurityContextHolder;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class UserService {


    @Autowired
    private SecurityContextHolder securityContextHolder;


    public Result userInfo() {
        return Result.success(securityContextHolder.getUser());
    }
}