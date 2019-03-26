/**
 * Copyright (C), 2015-2019, 重庆了赢科技有限公司
 * FileName: UserService
 * Author:   萧毅
 * Date:     2019/3/26 16:32
 * Description:
 */
package com.snow.golang.module.login;

import com.snow.golang.common.bean.Result;
import com.snow.golang.common.pojo.User;
import com.snow.golang.common.repository.UserRepository;
import com.snow.golang.common.util.PasswordEncoderUtils;
import com.snow.golang.config.token.JWTToken;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.Optional;

@Service
public class LoginService {

    @Autowired
    private UserRepository userRepository;
    @Autowired
    private JWTToken jwtToken;

    public Result login(String name, String password) {
        Optional<User> byUsername = userRepository.findByUsername(name);
        if (!byUsername.isPresent())return Result.fail();
        User user=byUsername.get();
        if (!PasswordEncoderUtils.decode(password,user.getPassword()))return Result.fail();
        user.setToken(jwtToken.createToken(String.valueOf(user.getId())));
        return Result.success(user);
    }
}