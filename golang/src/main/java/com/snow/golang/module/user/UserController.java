/**
 * Copyright (C), 2015-2019, 重庆了赢科技有限公司
 * FileName: UserController
 * Author:   萧毅
 * Date:     2019/3/26 16:28
 * Description:
 */
package com.snow.golang.module.user;

import com.snow.golang.common.bean.Result;
import com.snow.golang.config.annotation.Security;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@Api(tags = "用户")
@RestController
@RequestMapping("user")
@Security(value = "/security/sx",order = 1,names = "用户")
public class UserController {


    @Autowired
    private UserService userService;


    @GetMapping("userInfo")
    @ApiOperation(value = "获取用户基本信息")
    public Result userInfo(){
        return userService.userInfo();
    }


    @GetMapping("s2")
    @ApiOperation(value = "用户——二级权限")
    public Result s2(){

        return Result.success();
    }


    @GetMapping("s3")
    @ApiOperation(value = "用户——三级权限")
    @Security(menu = 3,sign = 9,names = "用户——权限二级级菜单")
    public Result s3(){


        return Result.success();
    }
    @GetMapping("s3_2")
    @ApiOperation(value = "用户——三级权限2")
    @Security(menu = 3,sign = 8,names = "用户——权限二级级菜单2")
    public Result s3_2(){


        return Result.success();
    }

    @GetMapping("sx")
    @ApiOperation(value = "权限排除项")
    public Result sx(){

        return Result.success();
    }


}