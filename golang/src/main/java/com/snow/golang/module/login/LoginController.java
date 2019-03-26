/**
 * Copyright (C), 2015-2019, 重庆了赢科技有限公司
 * FileName: LongController
 * Author:   萧毅
 * Date:     2019/3/26 16:27
 * Description:
 */
package com.snow.golang.module.login;

import com.snow.golang.common.bean.Result;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@Api(tags = "登录")
@RestController
public class LoginController {

    @Autowired
    private LoginService loginService;

    @PostMapping("login")
    @ApiOperation(value = "登录")
    public Result login(@RequestParam String name,
                        @RequestParam String password){
        return loginService.login(name,password);
    }







}